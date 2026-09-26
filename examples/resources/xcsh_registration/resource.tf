# Registration Resource Example
# Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Registration configuration
resource "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  token = "example-value"
}
