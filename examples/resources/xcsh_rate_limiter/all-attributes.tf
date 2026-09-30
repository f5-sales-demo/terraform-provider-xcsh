# AllAttributes — Acceptance-test-derived Configuration
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

resource "xcsh_rate_limiter" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test rate limiter with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "testing"
  }
}
