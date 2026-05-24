# Endpoint Data Source Example
# Retrieves information about an existing Endpoint

# Look up an existing Endpoint by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "endpoint_id" {
#   value = data.f5xc_endpoint.example.id
# }
