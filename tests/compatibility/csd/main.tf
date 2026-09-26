terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = "= 99.0.0"
    }
  }
}

data "xcsh_addon_service_activation_status" "csd" {
  addon_service = "f5xc-client-side-defense-standard"
}

output "csd_activation_status" {
  value = data.xcsh_addon_service_activation_status.csd
}

data "xcsh_network_regional_edges" "origin_ingress" {}

resource "xcsh_namespace" "csd" {
  name = "provider-compatibility"
}

resource "xcsh_protected_domain" "csd" {
  name             = "client-side-defense"
  namespace        = xcsh_namespace.csd.name
  protected_domain = "f5-sales-demo.example"
}

resource "xcsh_origin_pool" "origin" {
  name      = "provider-compatibility"
  namespace = xcsh_namespace.csd.name
  port      = 80

  origin_servers {
    public_name {
      dns_name = "origin.f5-sales-demo.example"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_http_loadbalancer" "csd" {
  name      = "provider-compatibility"
  namespace = xcsh_namespace.csd.name
  domains   = ["app.f5-sales-demo.example"]

  https_auto_cert {
    http_redirect = true
  }

  advertise_on_public_default_vip = {}

  default_route_pools {
    pool {
      name      = xcsh_origin_pool.origin.name
      namespace = xcsh_namespace.csd.name
    }
    weight   = 1
    priority = 1
  }

  client_side_defense {
    policy {
      js_insert_all_pages = {}
    }
  }
}

output "regional_edge_source_entry_count" {
  value = length(data.xcsh_network_regional_edges.origin_ingress.source_entries)
}
