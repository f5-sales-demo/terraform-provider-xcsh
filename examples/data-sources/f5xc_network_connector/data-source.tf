# Network Connector Data Source Example
# Retrieves information about an existing Network Connector

# Look up an existing Network Connector by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "network_connector_id" {
#   value = data.f5xc_network_connector.example.id
# }
