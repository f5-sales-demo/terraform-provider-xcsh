---
page_title: "peers.routing_policies.route_policy"
subcategory: ""
description: "peers.routing_policies.route_policy for xcsh_bgp."
xcsh_docs: {"aliases": [], "body_bytes": 3452, "body_sha256": "sha256:60f2f1101bb4b5cba62671b4c15cf7e3f9074312377877454cf73d836c50cbaf", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:inbound", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:node_name", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:object_refs", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:outbound"], "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies", "path": "documentation/data-sources/bgp/properties/peers/routing_policies/route_policy/index.md", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/routing_policies/route_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "peers.routing_policies.route_policy for xcsh_bgp.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/)
- peers.routing_policies.route_policy

<a id="section"></a>

Type: `"list"`. Computed.

Policy configuration for this feature.

Upstream description:

Route policy to be applied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

## Direct properties

- [all_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/all_nodes/): complete subsection reference.

- [inbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/inbound/): complete subsection reference.

- [node_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/node_name/): complete subsection reference.

- [object_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/object_refs/): complete subsection reference.

- [outbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/outbound/): complete subsection reference.

## Next pages

- [peers.routing_policies.route_policy.all_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/all_nodes/)
- [peers.routing_policies.route_policy.inbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/inbound/)
- [peers.routing_policies.route_policy.node_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/node_name/)
- [peers.routing_policies.route_policy.object_refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/object_refs/)
- [peers.routing_policies.route_policy.outbound](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/outbound/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
