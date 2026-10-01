# DenyList — Acceptance-test-derived Configuration
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

resource "xcsh_service_policy" "test" {
  name      = "example"
  namespace = "system"

  deny_list {
    prefix_list {
      prefixes = ["172.16.0.0/12"]
    }
    default_action_allow = {}
  }

  any_server = {}
}
