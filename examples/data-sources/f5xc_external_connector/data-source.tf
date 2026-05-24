# External Connector Data Source Example
# Retrieves information about an existing External Connector

# Look up an existing External Connector by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "external_connector_id" {
#   value = data.f5xc_external_connector.example.id
# }
