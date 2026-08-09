# Cluster Resource Example
# Manages cluster will create the object in the storage backend for namespace metadata.namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cluster configuration
resource "xcsh_cluster" "example" {
  name      = "example-cluster"
  namespace = "staging"

  connection_timeout     = 1
  endpoint_selection     = "DISTRIBUTED"
  fallback_policy        = "NO_FALLBACK"
  http_idle_timeout      = 1
  loadbalancer_algorithm = "ROUND_ROBIN"
}
