---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_bgp_status."
xcsh_docs: {"aliases": [], "body_bytes": 2160, "body_sha256": "sha256:9598a7c740d77e6b63c3fcca24f102822dc60b9781c82753e7055391da836e62", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_bgp_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:78479126aafed9697f5378abaad404d355901539446a65e0ee8620a492c5b9cd", "source_path": "examples/data-sources/xcsh_site_bgp_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_bgp_status:example:data-source", "parent_id": "xcsh-docs:data-sources:site_bgp_status:examples", "path": "documentation/data-sources/site_bgp_status/examples/data-source/index.md", "provider_name": "site_bgp_status", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_bgp_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_bgp_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_bgp_status/data-source.tf`; digest `sha256:78479126aafed9697f5378abaad404d355901539446a65e0ee8620a492c5b9cd`.

```terraform
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

  expected_exported_routes = ["10.40.0.10/32"]

  expected_peers = {
    node_0_slo = {
      node                     = "node-0"
      role                     = "slo"
      mac                      = "02:00:00:00:00:10"
      peer_address             = "169.254.100.1"
      expected_imported_routes = ["10.20.0.0/16"]
    }
    node_0_sli = {
      node                     = "node-0"
      role                     = "sli"
      mac                      = "02:00:00:00:00:11"
      peer_address             = "169.254.101.1"
      expected_imported_routes = ["10.30.0.0/16"]
    }
  }

  timeout_seconds       = 300
  poll_interval_seconds = 10
}

output "bgp_converged" {
  value = data.xcsh_site_bgp_status.site.converged
}

output "bgp_peers" {
  value = data.xcsh_site_bgp_status.site.peers
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/examples/)
- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
