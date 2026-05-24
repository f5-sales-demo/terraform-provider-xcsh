# Network Firewall Data Source Example
# Retrieves information about an existing Network Firewall

# Look up an existing Network Firewall by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_network_firewall" "example" {
  name      = "example-network-firewall"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "network_firewall_id" {
#   value = data.f5xc_network_firewall.example.id
# }
