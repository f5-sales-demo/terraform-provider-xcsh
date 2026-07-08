# VoltstackSite Resource Example
# Manages a Voltstack Site resource in F5 Distributed Cloud for deploying Volterra stack sites for edge computing.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VoltstackSite configuration
resource "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
  worker_nodes          = ["example-value"]
  address               = "example-value"
}
