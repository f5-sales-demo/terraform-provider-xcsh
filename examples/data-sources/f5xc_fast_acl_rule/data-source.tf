# Fast ACL Rule Data Source Example
# Retrieves information about an existing Fast ACL Rule

# Look up an existing Fast ACL Rule by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "fast_acl_rule_id" {
#   value = data.f5xc_fast_acl_rule.example.id
# }
