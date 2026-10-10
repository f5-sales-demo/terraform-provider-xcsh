---
page_title: "xcsh_bgp_routing_policy"
subcategory: ""
description: "Reads BGP routing rules with match criteria and actions that control imported or exported routes."
xcsh_docs: {"aliases": ["bgp routing policy"], "body_bytes": 1415, "body_sha256": "sha256:6f0b1ca83b49321e5fd21cb8eddbffbd6990bb5cee7cf18a21833660ee1966ad", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp_routing_policy:reference", "xcsh-docs:data-sources:bgp_routing_policy:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bgp_routing_policy/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011", "registry_path": "docs/data-sources/bgp_routing_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads BGP routing rules with match criteria and actions that control imported or exported routes.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bgp_routing_policy

Breadcrumbs:

- xcsh_bgp_routing_policy

Reads BGP routing rules with match criteria and actions that control imported or exported routes.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPRoutingPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPRoutingPolicy by name
data "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}

output "bgp_routing_policy_id" {
  value = data.xcsh_bgp_routing_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/examples/)
