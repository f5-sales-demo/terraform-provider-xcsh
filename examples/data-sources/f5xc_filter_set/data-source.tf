# Filter Set Data Source Example
# Retrieves information about an existing Filter Set

# Look up an existing Filter Set by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "filter_set_id" {
#   value = data.f5xc_filter_set.example.id
# }
