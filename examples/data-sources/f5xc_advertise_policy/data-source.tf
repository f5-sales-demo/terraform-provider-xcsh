# Advertise Policy Data Source Example
# Retrieves information about an existing Advertise Policy

# Look up an existing Advertise Policy by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "advertise_policy_id" {
#   value = data.f5xc_advertise_policy.example.id
# }
