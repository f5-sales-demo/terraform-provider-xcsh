---
page_title: "site_local_inside_network"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["site local inside network"], "body_bytes": 1244, "body_sha256": "sha256:f01113b35474ba86f09fe32968e295499e230b8b75255b29776761e258c99dea", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:site_local_inside_network", "parent_id": "xcsh-docs:data-sources:proxy:reference", "path": "documentation/data-sources/proxy/properties/site_local_inside_network/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-1101012223103130-2201131330221131-1331022102330232-3330000112313011-2223113210330001-3012032001012200-0303100022020121-1102012131121322", "registry_path": "docs/guides/data-sources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_local_inside_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/site_local_inside_network/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["proxyCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_local_inside_network

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- site_local_inside_network

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: site\_local\_inside\_network, site\_local\_network\] Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_local_inside_network/#section)
- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_local_network/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
