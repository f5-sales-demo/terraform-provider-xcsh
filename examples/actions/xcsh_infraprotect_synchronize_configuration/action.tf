# InfraprotectSynchronizeConfiguration Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_infraprotect_synchronize_configuration" "example" {
  config {
    namespace = "example-value"
  }
}
