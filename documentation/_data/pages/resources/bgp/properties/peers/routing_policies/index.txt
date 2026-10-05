---
page_title: "peers.routing_policies"
subcategory: ""
description: "List of rules which can be applied on all or particular nodes."
xcsh_docs: {"aliases": ["peers routing policies"], "body_bytes": 1419, "body_sha256": "sha256:cf03102559ca45504b525da74391ff67dacac13906dad7de717757da7af2c09a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies", "parent_id": "xcsh-docs:resources:bgp:properties:peers", "path": "documentation/resources/bgp/properties/peers/routing_policies/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies"], "schema_version": 1, "sections": [{"aliases": ["peers routing policies route policy"], "anchor": "section", "description": "Route policy to be applied.", "document_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:all_nodes,node_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:inbound,outbound", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:inbound", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:all_nodes,node_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:inbound,outbound", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:outbound", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:RequiredListObjectAttributes:object_refs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:object_refs", "type": "requires"}], "schema_path": ["peers", "routing_policies", "route_policy"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of rules which can be applied on all or particular nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.routing_policies

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- peers.routing_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
routing_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/): complete subsection reference.

## Next pages

- [peers.routing_policies.route_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/routing_policies/route_policy/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
