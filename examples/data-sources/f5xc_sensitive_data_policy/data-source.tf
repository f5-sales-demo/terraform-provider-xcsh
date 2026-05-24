# Sensitive Data Policy Data Source Example
# Retrieves information about an existing Sensitive Data Policy

# Look up an existing Sensitive Data Policy by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "sensitive_data_policy_id" {
#   value = data.f5xc_sensitive_data_policy.example.id
# }
