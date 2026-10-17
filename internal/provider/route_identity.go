package provider

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"regexp"
)

// This read evidence is paired with the same optimistic-concurrency token as
// the saved plan. It never comes from a GET performed during an update.
const routeIdentitiesPrivateSlot = "f5xcRouteIdentities"

var routeUUIDPattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
var routeSelectorPattern = regexp.MustCompile(`^[0-9a-f]{64}$`)

type routeIdentity struct {
	Kind     string `json:"kind"`
	Selector string `json:"selector"`
	UUID     string `json:"uuid"`
}
type routeIdentityEvidence struct {
	Version   int             `json:"version"`
	Namespace string          `json:"namespace"`
	Name      string          `json:"name"`
	Token     string          `json:"token"`
	Routes    []routeIdentity `json:"routes"`
}

func cloneRouteIdentitySpec(spec map[string]interface{}) (map[string]interface{}, error) {
	raw, err := json.Marshal(spec)
	if err != nil {
		return nil, errors.New("route configuration cannot be encoded")
	}
	var result map[string]interface{}
	if json.Unmarshal(raw, &result) != nil || result == nil {
		return nil, errors.New("route configuration is malformed")
	}
	return result, nil
}

// Empty API defaults do not distinguish selectors. Header match values
// remain part of the selector; policy, query rewriting, pools and order do not.
func normalizeRouteSelector(value interface{}) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		result := map[string]interface{}{}
		for key, child := range v {
			// The header matcher defaults invert_match to false. Unlike presence,
			// omitting it does not change the request match.
			if key == "invert_match" && child == false {
				continue
			}
			if normalized := normalizeRouteSelector(child); normalized != nil {
				result[key] = normalized
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case []interface{}:
		if len(v) == 0 {
			return nil
		}
		result := make([]interface{}, len(v))
		for i, child := range v {
			result[i] = normalizeRouteSelector(child)
		}
		return result
	default:
		return value
	}
}

func inlineRouteIdentity(route interface{}) (string, string, map[string]interface{}, error) {
	outer, ok := route.(map[string]interface{})
	if !ok {
		return "", "", nil, errors.New("route entry is malformed")
	}
	var kind string
	var body map[string]interface{}
	for _, candidate := range []string{"simple_route", "redirect_route", "direct_response_route"} {
		if value, present := outer[candidate]; present && value != nil {
			if kind != "" {
				return "", "", nil, errors.New("route has ambiguous inline alternatives")
			}
			body, ok = value.(map[string]interface{})
			if !ok {
				return "", "", nil, errors.New("inline route is malformed")
			}
			kind = candidate
		}
	}
	if kind == "" {
		if ref, ok := outer["custom_route_object"].(map[string]interface{}); ok && len(ref) > 0 {
			return "", "", nil, nil
		}
		return "", "", nil, errors.New("route identity kind is unsupported")
	}
	// Live HTTP LB readback assigns UUIDs only to simple routes. Redirect and
	// direct-response routes have no identity field and remain untouched.
	if kind != "simple_route" {
		if _, supplied := body["uuid"]; supplied {
			return "", "", nil, errors.New("route kind does not support a UUID")
		}
		return "", "", nil, nil
	}
	path, ok := body["path"].(map[string]interface{})
	if !ok {
		return "", "", nil, errors.New("route selector path is missing")
	}
	count := 0
	for _, key := range []string{"prefix", "path", "regex"} {
		if value, present := path[key]; present && value != nil {
			if _, ok := value.(string); !ok {
				return "", "", nil, errors.New("route selector path is malformed")
			}
			count++
		}
	}
	if count != 1 {
		return "", "", nil, errors.New("route selector path is ambiguous")
	}
	selector := map[string]interface{}{"kind": kind, "path": normalizeRouteSelector(path), "http_method": "ANY"}
	for _, key := range []string{"http_method", "incoming_port", "headers"} {
		if normalized := normalizeRouteSelector(body[key]); normalized != nil {
			selector[key] = normalized
		}
	}
	if method, ok := selector["http_method"].(string); !ok || method == "" {
		return "", "", nil, errors.New("route selector method is malformed")
	}
	raw, err := json.Marshal(selector)
	if err != nil {
		return "", "", nil, errors.New("route selector cannot be encoded")
	}
	digest := sha256.Sum256(raw)
	return kind, hex.EncodeToString(digest[:]), body, nil
}

func routeIdentityRows(spec map[string]interface{}) ([]interface{}, error) {
	value, present := spec["routes"]
	if !present || value == nil {
		return nil, nil
	}
	routes, ok := value.([]interface{})
	if !ok {
		return nil, errors.New("route collection is malformed")
	}
	return routes, nil
}

func encodeRouteIdentities(spec map[string]interface{}, namespace, name, token string) ([]byte, error) {
	if namespace == "" || name == "" || token == "" {
		return nil, errors.New("route read provenance is incomplete")
	}
	cloned, err := cloneRouteIdentitySpec(spec)
	if err != nil {
		return nil, err
	}
	routes, err := routeIdentityRows(cloned)
	if err != nil {
		return nil, err
	}
	evidence := routeIdentityEvidence{Version: 1, Namespace: namespace, Name: name, Token: token, Routes: []routeIdentity{}}
	selectors := map[string]bool{}
	identities := map[string]bool{}
	for _, route := range routes {
		kind, selector, body, err := inlineRouteIdentity(route)
		if err != nil {
			return nil, err
		}
		if body == nil {
			continue
		}
		identity, ok := body["uuid"].(string)
		if !ok || !routeUUIDPattern.MatchString(identity) {
			return nil, errors.New("server route UUID is missing or malformed")
		}
		if selectors[selector] || identities[identity] {
			return nil, errors.New("server route identities are ambiguous")
		}
		selectors[selector] = true
		identities[identity] = true
		evidence.Routes = append(evidence.Routes, routeIdentity{Kind: kind, Selector: selector, UUID: identity})
	}
	return json.Marshal(evidence)
}

func preserveRouteIdentities(raw []byte, spec map[string]interface{}, namespace, name, token string) (map[string]interface{}, error) {
	if len(raw) == 0 {
		return nil, errors.New("route identity read evidence is missing; refresh and replan")
	}
	var evidence routeIdentityEvidence
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&evidence) != nil {
		return nil, errors.New("route identity read evidence is malformed")
	}
	if err := decoder.Decode(new(interface{})); err != io.EOF {
		return nil, errors.New("route identity read evidence has trailing data")
	}
	if evidence.Version != 1 || namespace == "" || name == "" || token == "" || evidence.Namespace != namespace || evidence.Name != name || evidence.Token != token || evidence.Routes == nil {
		return nil, errors.New("route identity read provenance does not match this plan; refresh and replan")
	}
	identities := map[string]routeIdentity{}
	uuids := map[string]bool{}
	for _, row := range evidence.Routes {
		if !routeSelectorPattern.MatchString(row.Selector) || !routeUUIDPattern.MatchString(row.UUID) || row.Kind != "simple_route" {
			return nil, errors.New("route identity evidence entry is malformed")
		}
		if _, duplicate := identities[row.Selector]; duplicate || uuids[row.UUID] {
			return nil, errors.New("route identity read evidence is ambiguous")
		}
		identities[row.Selector] = row
		uuids[row.UUID] = true
	}
	result, err := cloneRouteIdentitySpec(spec)
	if err != nil {
		return nil, err
	}
	routes, err := routeIdentityRows(result)
	if err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	for _, route := range routes {
		kind, selector, body, err := inlineRouteIdentity(route)
		if err != nil {
			return nil, err
		}
		if body == nil {
			continue
		}
		if seen[selector] {
			return nil, errors.New("desired route selectors are ambiguous")
		}
		seen[selector] = true
		if _, supplied := body["uuid"]; supplied {
			return nil, errors.New("route UUID must come from verified read evidence")
		}
		if row, found := identities[selector]; found {
			if row.Kind != kind {
				return nil, errors.New("route identity kind mismatch")
			}
			body["uuid"] = row.UUID
		}
	}
	return result, nil
}

func verifyPreservedRouteIdentities(desired, observed map[string]interface{}) error {
	expected, err := cloneRouteIdentitySpec(desired)
	if err != nil {
		return err
	}
	actual, err := cloneRouteIdentitySpec(observed)
	if err != nil {
		return err
	}
	rows, err := routeIdentityRows(actual)
	if err != nil {
		return err
	}
	ids := map[string]string{}
	for _, route := range rows {
		_, selector, body, err := inlineRouteIdentity(route)
		if err != nil {
			return err
		}
		if body == nil {
			continue
		}
		identity, ok := body["uuid"].(string)
		if !ok || !routeUUIDPattern.MatchString(identity) {
			return errors.New("updated route UUID is missing or malformed")
		}
		if _, duplicate := ids[selector]; duplicate {
			return errors.New("updated route selectors are ambiguous")
		}
		ids[selector] = identity
	}
	rows, err = routeIdentityRows(expected)
	if err != nil {
		return err
	}
	for _, route := range rows {
		_, selector, body, err := inlineRouteIdentity(route)
		if err != nil {
			return err
		}
		if body == nil {
			continue
		}
		if identity, ok := body["uuid"].(string); ok && ids[selector] != identity {
			return errors.New("server did not preserve the reviewed route UUID")
		}
	}
	return nil
}
