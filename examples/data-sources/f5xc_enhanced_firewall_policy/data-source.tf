# Enhanced Firewall Policy Data Source Example
# Retrieves information about an existing Enhanced Firewall Policy

# Look up an existing Enhanced Firewall Policy by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_enhanced_firewall_policy" "example" {
  name      = "example-enhanced-firewall-policy"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "enhanced_firewall_policy_id" {
#   value = data.f5xc_enhanced_firewall_policy.example.id
# }
