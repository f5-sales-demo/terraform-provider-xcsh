package codegen

// These owned choices have observable API readback. Preserve no configured
// marker over the live selection; only the measured equivalent default list is
// normalized separately (#2573).
func appFirewallAuthoritativeMarker(resource, wireName string) bool {
	if resource != "AppFirewall" {
		return false
	}
	switch wireName {
	case "mitigate_high_risk_action", "mitigate_high_medium_risk_action", "default_violation_settings",
		"enable_suppression", "disable_suppression", "enable_threat_campaigns", "disable_threat_campaigns":
		return true
	}
	return false
}
