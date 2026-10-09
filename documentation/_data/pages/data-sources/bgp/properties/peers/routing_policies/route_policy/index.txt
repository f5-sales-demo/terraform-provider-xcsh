---
page_title: "peers.routing_policies.route_policy"
subcategory: ""
description: "Route policy to be applied."
xcsh_docs: {"aliases": ["peers routing policies route policy"], "body_bytes": 2311, "body_sha256": "sha256:fd6b930daf93195bcf8818eb3f0949e14a2d964a9e8072fb1f62a312d53fc4af", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:inbound", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:node_name", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:object_refs", "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:outbound"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies", "path": "documentation/data-sources/bgp/properties/peers/routing_policies/route_policy/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0131123322003121-2223310210101102-2130310123231230-0112020320212322-1103013103101323-3022013213101332-1033010112002013-1300030022030213", "registry_path": "docs/guides/data-sources--bgp--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies", "route_policy"], "schema_version": 1, "sections": [{"aliases": ["peers routing policies route policy all nodes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:all_nodes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "all_nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers routing policies route policy inbound"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:inbound", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "inbound"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers routing policies route policy node name"], "anchor": "section", "description": "List of nodes on which BGP routing policy has to be applied.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:node_name", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "node_name"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers routing policies route policy object refs"], "anchor": "section", "description": "Select route policy to apply.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:object_refs", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "object_refs"], "syntax": "attribute", "type": "object"}, {"aliases": ["peers routing policies route policy outbound"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy:outbound", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy", "outbound"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/routing_policies/route_policy/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Route policy to be applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["bgpCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
