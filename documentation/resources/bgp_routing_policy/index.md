---
page_title: "xcsh_bgp_routing_policy"
subcategory: ""
description: "Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of rules containing match criteria and action to be applied. these rules help control routes which are imported or exported to bgp peers. configuration."
xcsh_docs: {"aliases": ["bgp routing policy"], "body_bytes": 1845, "body_sha256": "sha256:e9b269e64b947af1c36b52560a14679cf3a1851a4038eac636b8452daa589221", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:reference", "xcsh-docs:resources:bgp_routing_policy:examples", "xcsh-docs:resources:bgp_routing_policy:import", "xcsh-docs:resources:bgp_routing_policy:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/bgp_routing_policy/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2211022302323020-0302320013313033-3100300331020023-2001230222330132-3112010301102323-1231210203133211-2221203232121211-3030023302033111", "registry_path": "docs/resources/bgp_routing_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of rules containing match criteria and action to be applied. these rules help control routes which are imported or exported to bgp peers. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bgp_routing_policy

Breadcrumbs:

- xcsh_bgp_routing_policy

Manages a BGP Routing Policy resource in F5 Distributed Cloud for bgp routing policy is a list of
rules containing match criteria and action to be applied. these rules help control routes which are
imported or exported to bgp peers. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp_routing_policy/lifecycle/timeouts/)
