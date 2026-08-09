# MultipleOrigins — Verified Configuration Example
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

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 443

  origin_servers {
    public_name {
      dns_name = "backend1.example.com"
    }
  }

  origin_servers {
    public_name {
      dns_name = "backend2.example.com"
    }
  }

  no_tls {}
  same_as_endpoint_port {}
}
