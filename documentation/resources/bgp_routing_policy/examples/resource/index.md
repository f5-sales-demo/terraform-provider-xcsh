---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1156, "body_sha256": "sha256:1f0a8eba577b8736dddb2d94f380c7668550a3eba0dcdf2f0ae99ecfbd8ab90c", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5114b3e2f71f278003899287dd9fce44709b024db2b4ac16ae2e23054bd84151", "source_path": "examples/resources/xcsh_bgp_routing_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bgp_routing_policy:example:resource", "parent_id": "xcsh-docs:resources:bgp_routing_policy:examples", "path": "documentation/resources/bgp_routing_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3331212200303011-3320120231232031-1121221232110113-2110313120103300-3000112321010011-0322122111210001-2322011202313211-0103301030332021", "registry_path": "docs/guides/resources--bgp_routing_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Resource for xcsh_bgp_routing_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bgp_routing_policy/resource.tf`; digest `sha256:5114b3e2f71f278003899287dd9fce44709b024db2b4ac16ae2e23054bd84151`.

```terraform
# BGPRoutingPolicy Resource Example
# Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of rules containing match criteria and action to be applied.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGPRoutingPolicy configuration
resource "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}
```
