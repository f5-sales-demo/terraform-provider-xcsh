# WorkloadFlavor Resource Example
# Manages workload_flavor.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic WorkloadFlavor configuration
resource "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "staging"

  ephemeral_storage = "example-value"
  memory            = "example-value"
  vcpus             = 1
}
