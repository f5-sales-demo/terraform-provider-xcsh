---
page_title: "disable_api_discovery"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable api discovery"], "body_bytes": 1386, "body_sha256": "sha256:ead8d86bcabf9a425906846f7ba6d6f8f8731cf303b1c62f199c7df49c40e929", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:disable_api_discovery", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/disable_api_discovery/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-1302330231202101-0223003210100011-1103001000322103-2130013013301113-0031323333122311-1013313202011020-2010303010122300-2321321122232201", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_api_discovery"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/disable_api_discovery/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_api_discovery

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- disable_api_discovery

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

- [disable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/disable_api_discovery/#section)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/enable_api_discovery/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
