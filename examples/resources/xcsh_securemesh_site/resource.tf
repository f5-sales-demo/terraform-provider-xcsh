# SecuremeshSite Resource Example
# Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with distributed security capabilities.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSite configuration
resource "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
  worker_nodes          = ["example-value"]
  address               = "example-value"
}
