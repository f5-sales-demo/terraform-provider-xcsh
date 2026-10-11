---
page_title: "peers.routing_policies"
subcategory: ""
description: "List of rules which can be applied on all or particular nodes."
xcsh_docs: {"aliases": ["peers routing policies"], "body_bytes": 1060, "body_sha256": "sha256:971026b51c33bbb9f15419da05b3f7623a60f901eeae9b8f35e774ce45eac422", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:routing_policies", "parent_id": "xcsh-docs:resources:bgp:properties:peers", "path": "documentation/resources/bgp/properties/peers/routing_policies/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313", "registry_path": "docs/guides/resources--bgp--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies"], "schema_version": 1, "sections": [{"aliases": ["peers routing policies route policy"], "anchor": "section", "description": "Route policy to be applied.", "document_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:all_nodes,node_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:inbound,outbound", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:inbound", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:all_nodes,node_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:node_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:ConflictingListObjectAttributes:inbound,outbound", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:outbound", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.routing_policies.route_policy:RequiredListObjectAttributes:object_refs", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:routing_policies:route_policy:object_refs", "type": "requires"}], "schema_path": ["peers", "routing_policies", "route_policy"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/routing_policies/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of rules which can be applied on all or particular nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bgpCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
