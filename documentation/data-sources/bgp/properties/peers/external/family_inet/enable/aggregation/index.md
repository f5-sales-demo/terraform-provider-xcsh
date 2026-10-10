---
page_title: "peers.external.family_inet.enable.aggregation"
subcategory: ""
description: "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements."
xcsh_docs: {"aliases": ["peers external family inet enable aggregation"], "body_bytes": 2817, "body_sha256": "sha256:d7aa08fd0de61f94028ecf37e8ad2d2fde18ecdd7e9af48d0eb615758917e5da", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation:options"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable", "path": "documentation/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0302321100202233-0231213322311003-3030221203212302-3100323132212200-2120102003121311-2133210022223132-2310332021103113-2202223020221302", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable", "aggregation"], "schema_version": 1, "sections": [{"aliases": ["peers external family inet enable aggregation ip prefix"], "anchor": "schema-peers--external--family_inet--enable--aggregation--ip_prefix", "description": "Specify IPv4 subnet for aggregation.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "family_inet", "enable", "aggregation", "ip_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["peers external family inet enable aggregation options"], "anchor": "section", "description": "Configuration parameter for options", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation:options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers", "external", "family_inet", "enable", "aggregation", "options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bgpCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable.aggregation

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- [peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/)
- [peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/)
- peers.external.family_inet.enable.aggregation

<a id="section"></a>

Type: `"list"`. Computed.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-peers--external--family_inet--enable--aggregation--ip_prefix"></a>

### ip_prefix property

Type: `"string"`. Computed.

IP Prefix. Specify IPv4 subnet for aggregation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

- [options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/options/): complete subsection reference.
