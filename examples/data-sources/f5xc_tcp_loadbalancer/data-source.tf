# TCP Loadbalancer Data Source Example
# Retrieves information about an existing TCP Loadbalancer

# Look up an existing TCP Loadbalancer by name
terraform {
  required_version = ">= 1.0"

  required_providers {
    f5xc = {
      source  = "f5xc-salesdemos/f5xc"
      version = ">= 0.1.0"
    }
  }
}


data "f5xc_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "system"
}

# Example: Use the data source in another resource
# output "tcp_loadbalancer_id" {
#   value = data.f5xc_tcp_loadbalancer.example.id
# }
