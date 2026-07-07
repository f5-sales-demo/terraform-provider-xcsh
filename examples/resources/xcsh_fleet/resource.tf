# Fleet Resource Example
# Manages fleet will create a fleet object in 'system' namespace of the user.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Fleet configuration
resource "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "system"

  fleet_label                          = "example-value"
  enable_default_fleet_config_download = true
  operating_system_version             = "example-value"
  volterra_software_version            = "example-value"
}
