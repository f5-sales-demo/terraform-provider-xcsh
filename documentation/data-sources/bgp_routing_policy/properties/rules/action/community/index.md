---
page_title: "rules.action.community"
subcategory: ""
description: "List of BGP communities."
xcsh_docs: {"aliases": ["rules action community"], "body_bytes": 2337, "body_sha256": "sha256:c2c9b174e7e8971577c381a43c642693775a45933363092c6a310dd6488853d0", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp_routing_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:community", "parent_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action", "path": "documentation/data-sources/bgp_routing_policy/properties/rules/action/community/index.md", "product": "distributed-cloud", "provider_name": "bgp_routing_policy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0033110321000302-2132303301212021-1332331010232320-3101332112131311-0322321331333201-1321010001213133-3222300233100233-1220322110221031", "registry_path": "docs/guides/data-sources--bgp_routing_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "action", "community"], "schema_version": 1, "sections": [{"aliases": ["rules action community community"], "anchor": "schema-rules--action--community--community", "description": "An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits being value.", "document_id": "xcsh-docs:data-sources:bgp_routing_policy:properties:rules:action:community", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "action", "community", "community"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_routing_policy/properties/rules/action/community/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "List of BGP communities.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["bgp_routing_policyCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.community

Breadcrumbs:

- [xcsh_bgp_routing_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_routing_policy/properties/rules/action/)
- rules.action.community

<a id="section"></a>

Type: `"single"`. Computed.

BGP Community list. List of BGP communities.

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

<a id="schema-rules--action--community--community"></a>

### community property

Type: `["list", "string"]`. Computed.

An unordered set of RFC 1997 defined 4-byte community, first 16 bits being ASN and lower 16 bits
being value.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-15T10:49:30+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9]{1,5}:[0-9]{1,5}$",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
