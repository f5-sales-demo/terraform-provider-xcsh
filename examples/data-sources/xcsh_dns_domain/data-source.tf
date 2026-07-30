# DNSDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSDomain by name
data "xcsh_dns_domain" "example" {
  name      = "example-dns-domain"
  namespace = "staging"
}

output "dns_domain_id" {
  value = data.xcsh_dns_domain.example.id
}
