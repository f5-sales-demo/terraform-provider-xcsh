# APM Resource Example
# Manages new APM as a service with configured parameters.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APM configuration
resource "xcsh_apm" "example" {
  name      = "example-apm"
  namespace = "staging"
}
