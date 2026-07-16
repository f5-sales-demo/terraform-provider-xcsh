// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package suppress

import "reflect"

// ExtractDefaults compares a create request against the API response and returns
// the server-applied defaults keyed by their dotted spec path. A field present in
// the response but absent from the request is a server default; an empty {} block
// is a marker block (IsMarkerBlock). Both request and response are the full
// object maps ({"spec": {...}, ...}); only the spec subtree is compared.
//
// It recurses into nested maps AND list elements. The list-element recursion is
// what lets it discover defaults nested inside list elements — e.g.
// origin_servers[].labels {} and default_route_pools[].endpoint_subsets {} — which
// the previous map-only walker silently skipped, the root cause of issue #1103.
func ExtractDefaults(request, response map[string]interface{}) map[string]FieldDefault {
	defaults := make(map[string]FieldDefault)

	reqSpec, reqOk := request["spec"].(map[string]interface{})
	respSpec, respOk := response["spec"].(map[string]interface{})
	if !reqOk || !respOk {
		return defaults
	}

	compareObjects(reqSpec, respSpec, "spec", defaults)
	return defaults
}

// compareObjects records every response key that was absent from the request (a
// server default) and recurses into nested maps that both sides share.
func compareObjects(request, response map[string]interface{}, path string, defaults map[string]FieldDefault) {
	for key, respValue := range response {
		fullPath := path + "." + key

		reqValue, existsInReq := request[key]
		if !existsInReq {
			// Absent from the request → a server default.
			recordDefault(fullPath, respValue, defaults)
			continue
		}

		// Present in both. Recurse into maps AND lists so defaults nested at any depth
		// are discovered.
		switch rv := reqValue.(type) {
		case map[string]interface{}:
			respMap, ok := respValue.(map[string]interface{})
			if !ok {
				continue
			}
			if len(rv) == 0 {
				// Request sent an empty marker block; record whatever the response added.
				for subKey, subValue := range respMap {
					recordDefault(fullPath+"."+subKey, subValue, defaults)
				}
			} else {
				compareObjects(rv, respMap, fullPath, defaults)
			}
		case []interface{}:
			// #1103: recurse element-wise into list values. The API echoes empty-marker
			// sub-blocks (labels {}, endpoint_subsets {}) on every element of lists like
			// origin_servers / default_route_pools; without this branch those defaults
			// were invisible to discovery. The path carries no [i] index, so every
			// element collapses onto the same dotted path (deduped by leaf downstream).
			respList, ok := respValue.([]interface{})
			if !ok {
				continue
			}
			compareLists(rv, respList, fullPath, defaults)
		}
	}
}

// compareLists pairs response elements with their request counterpart by index and
// recurses into each element map. Response elements beyond the request's length (or
// whose request counterpart is not a map) are compared against an empty request, so
// every field they carry registers as a server default.
func compareLists(request, response []interface{}, path string, defaults map[string]FieldDefault) {
	for i, respElem := range response {
		respMap, ok := respElem.(map[string]interface{})
		if !ok {
			continue
		}
		reqMap := map[string]interface{}{}
		if i < len(request) {
			if m, ok := request[i].(map[string]interface{}); ok {
				reqMap = m
			}
		}
		compareObjects(reqMap, respMap, path, defaults)
	}
}

// recordDefault stores one discovered default, flagging empty {} objects as marker
// blocks (the shape origin_servers[].labels / default_route_pools[].endpoint_subsets
// take, and that isServerDefaultMember suppresses on import).
func recordDefault(fullPath string, respValue interface{}, defaults map[string]FieldDefault) {
	fd := FieldDefault{
		Path:         fullPath,
		DefaultValue: respValue,
		Type:         getValueType(respValue),
	}
	if m, ok := respValue.(map[string]interface{}); ok && len(m) == 0 {
		fd.IsMarkerBlock = true
	}
	defaults[fullPath] = fd
}

// getValueType classifies a JSON-decoded value for the defaults record.
func getValueType(v interface{}) string {
	if v == nil {
		return "null"
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int64, reflect.Float64:
		return "number"
	case reflect.Bool:
		return "bool"
	case reflect.Map:
		return "object"
	case reflect.Slice, reflect.Array:
		return "array"
	default:
		return "unknown"
	}
}
