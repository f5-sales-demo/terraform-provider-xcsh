# Container Registry Data Source Example
# Retrieves information about an existing Container Registry

# Look up an existing Container Registry by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "container_registry_id" {
#   value = data.f5xc_container_registry.example.id
# }
