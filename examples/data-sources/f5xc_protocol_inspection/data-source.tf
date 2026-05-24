# Protocol Inspection Data Source Example
# Retrieves information about an existing Protocol Inspection

# Look up an existing Protocol Inspection by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "protocol_inspection_id" {
#   value = data.f5xc_protocol_inspection.example.id
# }
