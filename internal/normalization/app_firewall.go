// Package normalization reconciles verified equivalent API representations.
package normalization

var defaultDisabledViolations = [...]string{
	"VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
	"VIOL_COOKIE_MALFORMED",
	"VIOL_ENCODING",
	"VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
	"VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
	"VIOL_EVASION_BAD_UNESCAPE",
}

// AppFirewallDetectionSettings recognizes only the exact six-entry default
// selection materialized by XC on PUT. A changed, extra, duplicate, or malformed
// value remains API-derived drift. Explicit list configurations stay explicit;
// configured default markers and imports use the equivalent default marker.
// The caller's API evidence is never modified.
func AppFirewallDetectionSettings(settings map[string]interface{}, useDefault bool) map[string]interface{} {
	if !useDefault {
		return settings
	}
	violations, ok := settings["violation_settings"].(map[string]interface{})
	if !ok || len(violations) != 1 {
		return settings
	}
	values, ok := violations["disabled_violation_types"].([]interface{})
	if !ok || len(values) != len(defaultDisabledViolations) {
		return settings
	}
	want := make(map[string]bool, len(defaultDisabledViolations))
	for _, value := range defaultDisabledViolations {
		want[value] = true
	}
	for _, value := range values {
		text, ok := value.(string)
		if !ok || !want[text] {
			return settings
		}
		delete(want, text)
	}
	if len(want) != 0 {
		return settings
	}
	result := make(map[string]interface{}, len(settings))
	for key, value := range settings {
		result[key] = value
	}
	delete(result, "violation_settings")
	result["default_violation_settings"] = map[string]interface{}{}
	return result
}
