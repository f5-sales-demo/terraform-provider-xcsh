# Address Allocator Data Source Example
# Retrieves information about an existing Address Allocator

# Look up an existing Address Allocator by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_address_allocator" "example" {
  name      = "example-address-allocator"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "address_allocator_id" {
#   value = data.f5xc_address_allocator.example.id
# }
