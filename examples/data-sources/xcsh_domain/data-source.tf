# Domain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Domain by name
data "xcsh_domain" "example" {
  name      = "example-domain"
  namespace = "staging"
}

output "domain_id" {
  value = data.xcsh_domain.example.id
}
