# Certificate Chain Data Source Example
# Retrieves information about an existing Certificate Chain

# Look up an existing Certificate Chain by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "certificate_chain_id" {
#   value = data.f5xc_certificate_chain.example.id
# }
