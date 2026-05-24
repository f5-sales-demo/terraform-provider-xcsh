# Cloud Connect Data Source Example
# Retrieves information about an existing Cloud Connect

# Look up an existing Cloud Connect by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cloud_connect_id" {
#   value = data.f5xc_cloud_connect.example.id
# }
