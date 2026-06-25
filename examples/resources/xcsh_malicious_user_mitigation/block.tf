terraform {
  required_providers {
    xcsh = {
      source  = "f5xc-salesdemos/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Block — Verified Configuration Example
# This configuration is extracted from acceptance tests
# and verified against the live F5 XC API.

resource "xcsh_malicious_user_mitigation" "test" {
  name      = "example"
  namespace = "system"

  mitigation_type {
    rules {
      threat_level {
        high {}
      }
      mitigation_action {
        block_temporarily {}
      }
    }
  }
}
