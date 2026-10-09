---
page_title: "peers.external.family_inet.enable"
subcategory: ""
description: "IPv4 Unicast."
xcsh_docs: {"aliases": ["peers external family inet enable"], "body_bytes": 1207, "body_sha256": "sha256:467f521c659ce488a04fa5ea52ab7380616d46131fcb3e90300159ea604e13bb", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet", "path": "documentation/data-sources/bgp/properties/peers/external/family_inet/enable/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-1202221221321332-1012011021203233-0322011222210001-1211000120000333-0221213030222220-0313313120000111-2231112031231303-3222212323102132", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable"], "schema_version": 1, "sections": [{"aliases": ["peers external family inet enable aggregation"], "anchor": "section", "description": "BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing table and applies to outbound advertisements.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:external:family_inet:enable:aggregation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers", "external", "family_inet", "enable", "aggregation"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/family_inet/enable/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "IPv4 Unicast.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["bgpCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- [peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/)
- peers.external.family_inet.enable

<a id="section"></a>

Type: `"single"`. Computed.

Unicast IPv4. IPv4 Unicast.

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

- [aggregation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/family_inet/enable/aggregation/): complete subsection reference.
