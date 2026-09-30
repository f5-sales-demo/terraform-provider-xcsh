---
page_title: "xcsh_bgp_routing_policy"
subcategory: ""
description: "xcsh_bgp_routing_policy for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1557, "body_sha256": "sha256:7a487b7898ffcf55cbefe6aa4a87c98159bb8b99608069d898eec3645d92087b", "canonical_id": "xcsh-docs:resources:bgp_routing_policy:fundamentals", "child_ids": ["xcsh-docs:resources:bgp_routing_policy:reference", "xcsh-docs:resources:bgp_routing_policy:examples", "xcsh-docs:resources:bgp_routing_policy:import", "xcsh-docs:resources:bgp_routing_policy:timeouts"], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp_routing_policy:fundamentals", "parent_id": null, "path": "docs/resources/bgp_routing_policy.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bgp_routing_policy for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](../guides/resources--bgp_routing_policy--reference.md)
- [Examples](../guides/resources--bgp_routing_policy--examples.md)
- [Import](../guides/resources--bgp_routing_policy--import.md)
- [Timeouts](../guides/resources--bgp_routing_policy--timeouts.md)
