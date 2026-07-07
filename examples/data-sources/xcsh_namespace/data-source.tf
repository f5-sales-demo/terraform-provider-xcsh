# Namespace Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name      = "example-namespace"
  namespace = "staging"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
