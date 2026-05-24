# Cminstance Data Source Example
# Retrieves information about an existing Cminstance

# Look up an existing Cminstance by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "cminstance_id" {
#   value = data.f5xc_cminstance.example.id
# }
