terraform {
  required_version = ">= 1.14"
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 12.3.0"
    }
  }
}

variable "allocated_public_ip_name" { type = string }
variable "allocated_public_ip" { type = string }

resource "xcsh_virtual_site" "canada" {
  name      = "example-canadian-edges"
  namespace = "system"
  site_type = "REGIONAL_EDGE"
  site_selector {
    expressions = ["ves.io/region in (ves-io-toronto, ves-io-montreal)"]
  }
}

resource "xcsh_public_ip_binding" "canada" {
  name                   = var.allocated_public_ip_name
  namespace              = "shared"
  expected_ip            = var.allocated_public_ip
  virtual_site           = xcsh_virtual_site.canada.name
  virtual_site_namespace = xcsh_virtual_site.canada.namespace
}
