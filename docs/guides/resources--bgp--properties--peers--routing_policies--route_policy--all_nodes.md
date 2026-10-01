---
page_title: "peers.routing_policies.route_policy.all_nodes"
subcategory: ""
description: "peers.routing_policies.route_policy.all_nodes for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 1178, "body_sha256": "sha256:2f33a0fcac32c69a0ac5d3bc63bd426b1428c20c87133057d526bea5ef5c49f2", "canonical_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "child_ids": [], "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "parent_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "path": "docs/guides/resources--bgp--properties--peers--routing_policies--route_policy--all_nodes.md", "provider_name": "bgp", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy", "all_nodes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/route_policy/all_nodes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.routing_policies.route_policy.all_nodes for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy.all_nodes

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md)
- [Property reference](resources--bgp--reference.md)
- [peers](resources--bgp--properties--peers.md)
- [peers.routing_policies](resources--bgp--properties--peers--routing_policies.md)
- [peers.routing_policies.route_policy](resources--bgp--properties--peers--routing_policies--route_policy.md)
- peers.routing_policies.route_policy.all_nodes

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
all_nodes = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [peers.routing_policies.route_policy](resources--bgp--properties--peers--routing_policies--route_policy.md)
- [xcsh_bgp](../resources/bgp.md)
