# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
