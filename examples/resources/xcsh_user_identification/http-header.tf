# HttpHeader — Verified Configuration Example
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
  name      = "example"
  namespace = "system"

  rules {
    http_header_name = "X-Forwarded-For"
  }
}
