# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"

  health_check_port = 1
  port              = 1
  protocol          = "TCP"
}
