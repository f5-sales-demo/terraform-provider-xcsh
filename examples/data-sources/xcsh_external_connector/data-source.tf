# ExternalConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ExternalConnector by name
data "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}

output "external_connector_id" {
  value = data.xcsh_external_connector.example.id
}
