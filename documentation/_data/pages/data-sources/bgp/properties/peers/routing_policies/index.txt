---
page_title: "peers.routing_policies"
subcategory: ""
description: "List of rules which can be applied on all or particular nodes."
xcsh_docs: {"aliases": ["peers routing policies"], "body_bytes": 947, "body_sha256": "sha256:71965674e52f6b949c090e76d853c9af4f02c58893ee4dba3525261ecc1edf28", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers", "path": "documentation/data-sources/bgp/properties/peers/routing_policies/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-1023013200012011-2323323211323211-0121003131323302-1321231123212311-0203130231120011-0030213231023130-3313220021112010-1101312101303321", "registry_path": "docs/guides/data-sources--bgp--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "routing_policies"], "schema_version": 1, "sections": [{"aliases": ["peers routing policies route policy"], "anchor": "section", "description": "Route policy to be applied.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:routing_policies:route_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers", "routing_policies", "route_policy"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/routing_policies/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "List of rules which can be applied on all or particular nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["bgpCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
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
