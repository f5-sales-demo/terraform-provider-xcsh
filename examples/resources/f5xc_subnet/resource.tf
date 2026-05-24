# Subnet Resource Example
# Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod. it is created in user or shared namespace. configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}

# Basic Subnet configuration
resource "f5xc_subnet" "example" {
  name      = "example-subnet"
  namespace = "system"

  labels = {
    environment = "production"
    managed_by  = "terraform"
  }

  annotations = {
    "owner" = "platform-team"
  }

  # Resource-specific configuration
  # [OneOf: connect_to_layer2, connect_to_slo, isolated_nw] S...
  connect_to_layer2 {
    # Configure connect_to_layer2 settings
  }
  # Type establishes a direct reference from one object(the r...
  layer2_intf_ref {
    # Configure layer2_intf_ref settings
  }
  # Enable this option
  connect_to_slo {
    # Configure connect_to_slo settings
  }
}
