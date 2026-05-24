# Trusted Ca List Data Source Example
# Retrieves information about an existing Trusted Ca List

# Look up an existing Trusted Ca List by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "trusted_ca_list_id" {
#   value = data.f5xc_trusted_ca_list.example.id
# }
