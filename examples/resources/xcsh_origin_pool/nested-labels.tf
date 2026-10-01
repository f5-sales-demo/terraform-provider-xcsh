# NestedLabels — Acceptance-test-derived Configuration
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

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 8080

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
    labels = {
      "env" = "test"
      "app" = "demo"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
