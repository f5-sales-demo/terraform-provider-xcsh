---
page_title: "peers.routing_policies.route_policy.node_name"
subcategory: ""
description: "List of nodes on which BGP routing policy has to be applied."
xcsh_docs: {"aliases": ["peers routing policies route policy node name"], "body_bytes": 1745, "body_sha256": "sha256:3ad70e4a01060ed9fb34f424e2be33d04b2295335c832cbde0e87e026d17c6dc", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:node_name", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy", "path": "documentation/data-sources/bgp/properties/peers/routing_policies/route_policy/node_name/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3130312021232102-1102212310221310-0111003332211020-0112323103100110-0021103221220320-2202321201013013-1333301200000213-1303311003021121", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy", "node_name"], "schema_version": 1, "sections": [{"aliases": ["node"], "anchor": "schema-peers--routing_policies--route_policy--node_name--node", "description": "Select BGP Session on which policy will be applied.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:node_name", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "node_name", "node"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/routing_policies/route_policy/node_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of nodes on which BGP routing policy has to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies.route_policy.node_name

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [peers.routing_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/)
- [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/)
- peers.routing_policies.route_policy.node_name

<a id="section"></a>

Type: `"single"`. Computed.

List of nodes on which BGP routing policy has to be applied.

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

<a id="schema-peers--routing_policies--route_policy--node_name--node"></a>

### node property

Type: `["list", "string"]`. Computed.

Select BGP Session on which policy will be applied.

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

## Next pages

- [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/routing_policies/route_policy/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
