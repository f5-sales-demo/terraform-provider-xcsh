# DetectionSettings — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_app_firewall" "test" {
  name      = "example"
  namespace = "system"

  blocking = {}


  detection_settings {
    default_violation_settings = {}

    enable_suppression      = {}
    enable_threat_campaigns = {}
    signature_selection_setting {
      high_medium_accuracy_signatures = {}
      signature_settings_by_accuracy {
        high_accuracy_action   = "SIG_BLOCK"
        medium_accuracy_action = "SIG_BLOCK"
        low_accuracy_action    = "SIG_BLOCK"
      }
      default_attack_type_settings = {}
    }
  }
}
