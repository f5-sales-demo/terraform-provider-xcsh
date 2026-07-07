# Site Resource Example
# Manages virtual site object in given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Site configuration
resource "xcsh_site" "example" {
  name      = "example-site"
  namespace = "system"

  site_type = "INVALID"
}
