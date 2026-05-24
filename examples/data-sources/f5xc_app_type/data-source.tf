# App Type Data Source Example
# Retrieves information about an existing App Type

# Look up an existing App Type by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_app_type" "example" {
  name      = "example-app-type"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "app_type_id" {
#   value = data.f5xc_app_type.example.id
# }
