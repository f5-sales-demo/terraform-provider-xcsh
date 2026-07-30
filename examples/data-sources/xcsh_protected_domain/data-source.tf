# ProtectedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedDomain by name
data "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"
}

output "protected_domain_id" {
  value = data.xcsh_protected_domain.example.id
}
