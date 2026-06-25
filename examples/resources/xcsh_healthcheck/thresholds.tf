terraform {
  required_providers {
    xcsh = {
      source  = "f5xc-salesdemos/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Thresholds — Verified Configuration Example
# This configuration is extracted from acceptance tests
# and verified against the live F5 XC API.

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 443
  unhealthy_threshold = 3
  timeout             = 5
  interval            = 15

  tcp_health_check {}
}
