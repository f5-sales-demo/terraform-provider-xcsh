# Nfv Service Data Source Example
# Retrieves information about an existing Nfv Service

# Look up an existing Nfv Service by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "nfv_service_id" {
#   value = data.f5xc_nfv_service.example.id
# }
