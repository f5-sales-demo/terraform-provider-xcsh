# APM Data Source Example
# Retrieves information about an existing APM

# Look up an existing APM by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_apm" "example" {
  name      = "example-apm"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "apm_id" {
#   value = data.f5xc_apm.example.id
# }
