---
page_title: "xcsh_site_bgp_status"
subcategory: ""
description: "Polls authoritative F5 XC BGP and route observations until they agree with MAC-bound AWS expectations."
xcsh_docs: {"aliases": ["site bgp status"], "body_bytes": 2267, "body_sha256": "sha256:35fae35b271d4a143c2757fc2046e2a93d62d858c129b410e9b8a56ca9e38f09", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_bgp_status:reference", "xcsh-docs:data-sources:site_bgp_status:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_bgp_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_bgp_status:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/site_bgp_status/index.md", "product": "distributed-cloud", "provider_name": "site_bgp_status", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3202131020311001-0122003032002021-0022320312033010-1011001132303230-3130221113000032-1203011221022212-2212133031122301-1222231101030231", "registry_path": "docs/data-sources/site_bgp_status.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_bgp_status/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Polls authoritative F5 XC BGP and route observations until they agree with MAC-bound AWS expectations.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_bgp_status

Breadcrumbs:

- xcsh_site_bgp_status

Polls authoritative F5 XC BGP and route observations until they agree with MAC-bound AWS
expectations.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `expected_exported_routes`, `expected_peers`, `namespace`, `site`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/examples/)
