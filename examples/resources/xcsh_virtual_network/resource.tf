# VirtualNetwork Resource Example
# Manages virtual network in given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualNetwork configuration
resource "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"

  legacy_type = "VIRTUAL_NETWORK_SITE_LOCAL"
}
