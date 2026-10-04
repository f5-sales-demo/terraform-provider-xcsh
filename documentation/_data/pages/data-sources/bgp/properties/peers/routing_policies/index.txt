---
page_title: "peers.routing_policies"
subcategory: ""
description: "List of rules which can be applied on all or particular nodes."
xcsh_docs: {"aliases": ["peers routing policies"], "body_bytes": 1315, "body_sha256": "sha256:a7920669770db88a0d9e8e07b587f5123bbd6f08a751bb020fa65a599126bebe", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers", "path": "documentation/data-sources/bgp/properties/peers/routing_policies/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies"], "schema_version": 1, "sections": [{"aliases": ["peers routing policies route policy"], "anchor": "section", "description": "Route policy to be applied.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/routing_policies/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of rules which can be applied on all or particular nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["bgpCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- peers.routing_policies

<a id="section"></a>

Type: `"single"`. Computed.

List of rules which can be applied on all or particular nodes.

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

## Direct properties

- [route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/): complete subsection reference.

## Next pages

- [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
