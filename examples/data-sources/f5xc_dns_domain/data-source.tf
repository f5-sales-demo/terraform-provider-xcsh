# DNS Domain Data Source Example
# Retrieves information about an existing DNS Domain

# Look up an existing DNS Domain by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_dns_domain" "example" {
  name      = "example-dns-domain"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "dns_domain_id" {
#   value = data.f5xc_dns_domain.example.id
# }
