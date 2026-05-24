# Network Policy Data Source Example
# Retrieves information about an existing Network Policy

# Look up an existing Network Policy by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "network_policy_id" {
#   value = data.f5xc_network_policy.example.id
# }
