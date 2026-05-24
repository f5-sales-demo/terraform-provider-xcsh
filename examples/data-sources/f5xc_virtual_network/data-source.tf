# Virtual Network Data Source Example
# Retrieves information about an existing Virtual Network

# Look up an existing Virtual Network by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "virtual_network_id" {
#   value = data.f5xc_virtual_network.example.id
# }
