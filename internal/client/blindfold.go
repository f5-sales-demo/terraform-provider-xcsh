package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

// GetReplaceForm retains the outer concurrency token when XC wraps complete
// private-key configuration in replace_form. A missing token stays missing.
func (c *Client) GetReplaceForm(ctx context.Context, path string, result any) error {
	u, err := url.Parse(path)
	if err != nil {
		return errors.New("invalid certificate read path")
	}
	q := u.Query()
	q.Set("response_format", "2")
	u.RawQuery = q.Encode()
	var body json.RawMessage
	if err := c.Get(ctx, u.String(), &body); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if json.Unmarshal(body, &fields) != nil {
		return errors.New("malformed certificate response")
	}
	if form, ok := fields["replace_form"]; ok {
		var replacement map[string]json.RawMessage
		if json.Unmarshal(form, &replacement) != nil {
			return errors.New("malformed certificate replace form")
		}
		if token, ok := fields["resource_version"]; ok {
			replacement["resource_version"] = token
		}
		body, err = json.Marshal(replacement)
		if err != nil {
			return errors.New("malformed certificate replace form")
		}
	}
	if err := json.Unmarshal(body, result); err != nil {
		return errors.New("malformed certificate configuration")
	}
	return nil
}

// Reconciles native Blindfold writes only. Legacy client behavior is unchanged.
func (c *Client) blindfoldWrite(ctx context.Context, method, path string, data, result any) error {
	encoded, err := json.Marshal(data)
	if err != nil {
		return errors.New("cannot encode certificate configuration")
	}
	var desired map[string]any
	if json.Unmarshal(encoded, &desired) != nil {
		return errors.New("cannot encode certificate configuration")
	}
	meta, _ := desired["metadata"].(map[string]any)
	annotations, _ := meta["annotations"].(map[string]any)
	if _, native := annotations["f5-sales-demo.com/blindfold"]; !native {
		return errors.New("native provenance missing")
	}
	_, writeErr := c.doRequestWithRetry(ctx, method, path, data, false)
	if HasHTTPStatus(writeErr, 409) {
		return errors.New("Blindfold concurrency conflict; refresh and replan")
	}
	readPath := path
	if method == "POST" {
		name, ok := meta["name"].(string)
		if !ok || name == "" || strings.Contains(name, "/") {
			return errors.New("invalid certificate resource identity")
		}
		readPath += "/" + name
	}
	var observed map[string]any
	err = c.GetReplaceForm(ctx, readPath, &observed)
	if err != nil || !nativeMaterialEqual(desired["spec"], observed["spec"]) {
		return errors.New("Blindfold write outcome unresolved; inspect named resource before retrying")
	}
	observedMeta, _ := observed["metadata"].(map[string]any)
	observedAnnotations, _ := observedMeta["annotations"].(map[string]any)
	if observedAnnotations["f5-sales-demo.com/blindfold"] != annotations["f5-sales-demo.com/blindfold"] {
		return errors.New("Blindfold provenance readback failed")
	}
	if result != nil {
		body, _ := json.Marshal(observed)
		if json.Unmarshal(body, result) != nil {
			return errors.New("cannot decode Blindfold readback")
		}
	}
	return nil
}
func containsNativeMaterial(value any) bool {
	switch v := value.(type) {
	case map[string]any:
		if _, ok := v["certificate_url"]; ok {
			return true
		}
		for _, child := range v {
			if containsNativeMaterial(child) {
				return true
			}
		}
	case []any:
		for _, child := range v {
			if containsNativeMaterial(child) {
				return true
			}
		}
	}
	return false
}
func nativeMaterialEqual(desired, actual any) bool {
	if !containsNativeMaterial(desired) {
		return true
	}
	switch d := desired.(type) {
	case map[string]any:
		a, ok := actual.(map[string]any)
		if !ok {
			return false
		}
		if cert, ok := d["certificate_url"]; ok {
			if a["certificate_url"] != cert {
				return false
			}
			dk, _ := d["private_key"].(map[string]any)
			ak, _ := a["private_key"].(map[string]any)
			db, _ := dk["blindfold_secret_info"].(map[string]any)
			ab, _ := ak["blindfold_secret_info"].(map[string]any)
			if db["location"] != ab["location"] {
				return false
			}
		}
		for k, v := range d {
			if containsNativeMaterial(v) && !nativeMaterialEqual(v, a[k]) {
				return false
			}
		}
		return true
	case []any:
		a, ok := actual.([]any)
		if !ok || len(a) != len(d) {
			return false
		}
		for i, v := range d {
			if !nativeMaterialEqual(v, a[i]) {
				return false
			}
		}
		return true
	}
	return true
}
func isNativeBlindfold(data any) bool {
	b, err := json.Marshal(data)
	if err != nil {
		return false
	}
	var body map[string]any
	if json.Unmarshal(b, &body) != nil {
		return false
	}
	meta, _ := body["metadata"].(map[string]any)
	a, _ := meta["annotations"].(map[string]any)
	_, ok := a["f5-sales-demo.com/blindfold"]
	return ok
}
