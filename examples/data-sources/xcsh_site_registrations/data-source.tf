# SiteRegistrations DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations" "example" {
  namespace = "example-value"
}

output "site_registrations_result" {
  value = data.xcsh_site_registrations.example
}
