---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bgp_routing_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1091, "body_sha256": "sha256:f9496f866512ca8fe5710a763c5490abf9be52cd283710f952397f81efdb038e", "canonical_id": "xcsh-docs:resources:bgp_routing_policy:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp_routing_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5114b3e2f71f278003899287dd9fce44709b024db2b4ac16ae2e23054bd84151", "source_path": "examples/resources/xcsh_bgp_routing_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bgp_routing_policy:example:resource", "parent_id": "xcsh-docs:resources:bgp_routing_policy:examples", "path": "docs/guides/resources--bgp_routing_policy--example--resource.md", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp_routing_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_bgp_routing_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
- [Examples](resources--bgp_routing_policy--examples.md)
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

## Next pages

- [Examples](resources--bgp_routing_policy--examples.md)
- [xcsh_bgp_routing_policy](../resources/bgp_routing_policy.md)
