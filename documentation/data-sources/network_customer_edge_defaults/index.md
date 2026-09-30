---
page_title: "xcsh_network_customer_edge_defaults"
subcategory: ""
description: "xcsh_network_customer_edge_defaults for xcsh_network_customer_edge_defaults."
xcsh_docs: {"aliases": [], "body_bytes": 1693, "body_sha256": "sha256:814c66c02e5d3ba2ffddb172c17f42b01f3be6d15afc71d2f991d9080b6c6bb3", "child_ids": ["xcsh-docs:data-sources:network_customer_edge_defaults:reference", "xcsh-docs:data-sources:network_customer_edge_defaults:examples"], "collection_id": "xcsh-docs:data-sources:network_customer_edge_defaults:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_customer_edge_defaults:fundamentals", "parent_id": null, "path": "documentation/data-sources/network_customer_edge_defaults/index.md", "provider_name": "network_customer_edge_defaults", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_defaults/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_customer_edge_defaults for xcsh_network_customer_edge_defaults.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_network_customer_edge_defaults

Breadcrumbs:

- xcsh_network_customer_edge_defaults

Default DNS and NTP destinations for Customer Edge firewall rules. Values are bundled from the
pinned OpenAPI release; this data source performs no network request. Ports and traffic direction
are not encoded in the manifest.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_customer_edge_defaults" "system_services" {}

output "customer_edge_default_egress" {
  value = {
    dns = {
      direction    = "egress"
      protocols    = ["udp", "tcp"]
      port         = 53
      destinations = data.xcsh_network_customer_edge_defaults.system_services.dns_servers
    }
    ntp = {
      direction    = "egress"
      protocols    = ["udp"]
      port         = 123
      destinations = data.xcsh_network_customer_edge_defaults.system_services.ntp_servers
    }
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/examples/)
