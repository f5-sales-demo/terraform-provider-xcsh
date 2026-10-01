# CustomBlockingPage — Acceptance-test-derived Configuration
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

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  default_bot_setting        = {}
  default_anonymization      = {}

  blocking_page {
    blocking_page = "https://example.com/blocked.html"
    response_code = "Forbidden"
  }
}
