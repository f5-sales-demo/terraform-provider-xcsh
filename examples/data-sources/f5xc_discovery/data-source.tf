# Discovery Data Source Example
# Retrieves information about an existing Discovery

# Look up an existing Discovery by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_discovery" "example" {
  name      = "example-discovery"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "discovery_id" {
#   value = data.f5xc_discovery.example.id
# }
