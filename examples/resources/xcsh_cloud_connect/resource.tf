# CloudConnect Resource Example
# Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud provider networks.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudConnect configuration
resource "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}
