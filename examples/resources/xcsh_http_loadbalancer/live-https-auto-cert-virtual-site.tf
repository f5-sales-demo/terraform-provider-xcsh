# LiveHTTPSAutoCertVirtualSite — Verified Configuration Example
# This configuration is extracted from acceptance tests
# and verified against the live F5 XC API.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "0.14.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example-value"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_virtual_site" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "example-description"
  namespace  = xcsh_namespace.test.name
  site_type  = "CUSTOMER_EDGE"

  site_selector {
    expressions = ["site_type=customer_edge"]
  }
}

resource "xcsh_http_loadbalancer" "test" {
  depends_on = [xcsh_virtual_site.test]
  name       = "example"
  namespace  = xcsh_namespace.test.name
  domains    = ["test.example.com"]

  https_auto_cert {}

  advertise_custom {
    advertise_where {
      virtual_site {
        network = "SITE_NETWORK_INSIDE_AND_OUTSIDE"
        virtual_site {
          name      = xcsh_virtual_site.test.name
          namespace = xcsh_namespace.test.name
        }
      }
      use_default_port {}
    }
  }
}
