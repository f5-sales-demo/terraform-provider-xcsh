# Domain Resource Example
# Manages allowed domain.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Domain configuration
resource "xcsh_domain" "example" {
  name      = "example-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
