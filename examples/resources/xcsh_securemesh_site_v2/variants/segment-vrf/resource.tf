terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2-segment-vrf"
  namespace = "system"

  baremetal {
    not_managed {}
  }

  disable_ha {}
  no_network_policy {}
  no_forward_proxy {}
  logs_streaming_disabled {}
  block_all_services {}

  segment_vrf {
    segment_network {
      name      = "example-segment"
      namespace = "system"
    }
  }
}
