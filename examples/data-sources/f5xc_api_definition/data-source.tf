# API Definition Data Source Example
# Retrieves information about an existing API Definition

# Look up an existing API Definition by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "api_definition_id" {
#   value = data.f5xc_api_definition.example.id
# }
