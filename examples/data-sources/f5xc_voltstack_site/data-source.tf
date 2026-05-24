# Voltstack Site Data Source Example
# Retrieves information about an existing Voltstack Site

# Look up an existing Voltstack Site by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "voltstack_site_id" {
#   value = data.f5xc_voltstack_site.example.id
# }
