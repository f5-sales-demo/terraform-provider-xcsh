# Alert Policy Data Source Example
# Retrieves information about an existing Alert Policy

# Look up an existing Alert Policy by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "alert_policy_id" {
#   value = data.f5xc_alert_policy.example.id
# }
