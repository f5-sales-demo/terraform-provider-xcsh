# WithRules — Verified Configuration Example
# This configuration is extracted from acceptance tests
# and verified against the live F5 XC API.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_user_identification" "test" {
  name        = "example"
  namespace   = "system"
  description = "User identification with identification rules"

  rules {
    client_ip {}
  }
}
