# VirtualHost Resource Example
# Manages virtual host in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualHost configuration
resource "xcsh_virtual_host" "example" {
  name      = "example-virtual-host"
  namespace = "staging"

  domains                     = ["example-value"]
  request_cookies_to_remove   = ["example-value"]
  request_headers_to_remove   = ["example-value"]
  response_cookies_to_remove  = ["example-value"]
  response_headers_to_remove  = ["example-value"]
  add_location                = true
  connection_idle_timeout     = 1
  disable_default_error_pages = true
  disable_dns_resolve         = true
  idle_timeout                = 1
  max_request_header_size     = 1
  proxy                       = "UDP_PROXY"
}
