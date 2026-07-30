# TCPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TCPLoadBalancer by name
data "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}

output "tcp_loadbalancer_id" {
  value = data.xcsh_tcp_loadbalancer.example.id
}
