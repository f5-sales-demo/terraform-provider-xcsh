# Wait for MAC-correlated BGP peers and both BGP and simplified route views to
# converge. Peer addresses and expected routes come from authoritative AWS
# TGW Connect resource attributes in a real configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_site_bgp_status" "site" {
  namespace = "system"
  site      = "example-smsv2-site"

  expected_peers = {
    node_0_slo = {
      node            = "node-0"
      role            = "slo"
      mac             = "02:00:00:00:00:10"
      peer_address    = "169.254.100.1"
      expected_routes = ["10.20.0.0/16"]
    }
    node_0_sli = {
      node            = "node-0"
      role            = "sli"
      mac             = "02:00:00:00:00:11"
      peer_address    = "169.254.101.1"
      expected_routes = ["10.30.0.0/16"]
    }
  }

  timeout_seconds             = 300
  poll_interval_seconds       = 10
  max_observation_age_seconds = 120
}

output "bgp_converged" {
  value = data.xcsh_site_bgp_status.site.converged
}

output "bgp_peers" {
  value = data.xcsh_site_bgp_status.site.peers
}
