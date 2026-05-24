# Usb Policy Data Source Example
# Retrieves information about an existing Usb Policy

# Look up an existing Usb Policy by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "usb_policy_id" {
#   value = data.f5xc_usb_policy.example.id
# }
