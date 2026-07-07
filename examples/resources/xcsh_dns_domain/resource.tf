# DNSDomain Resource Example
# Manages DNS Domain in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSDomain configuration
resource "xcsh_dns_domain" "example" {
  name      = "example-dns-domain"
  namespace = "system"

  dnssec_mode = "DNSSEC_DISABLE"
}
