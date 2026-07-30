# AllowedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AllowedDomain by name
data "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"
}

output "allowed_domain_id" {
  value = data.xcsh_allowed_domain.example.id
}
