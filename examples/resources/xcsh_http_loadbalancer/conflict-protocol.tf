# ConflictProtocol — Negative Configuration Example
# Acceptance-test-derived conflict fixture; not a successful configuration.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"
  domains   = ["test.example.com"]

  http {
    port = 80
  }

  https_auto_cert {
    add_hsts                 = false
    no_mtls                  = {}
    default_header           = {}
    enable_path_normalize    = {}
    non_default_loadbalancer = {}
  }

  advertise_on_public_default_vip = {}
}
