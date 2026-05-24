# Cloud Link Data Source Example
# Retrieves information about an existing Cloud Link

# Look up an existing Cloud Link by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cloud_link_id" {
#   value = data.f5xc_cloud_link.example.id
# }
