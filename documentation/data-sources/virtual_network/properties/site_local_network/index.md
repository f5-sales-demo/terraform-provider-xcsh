---
page_title: "site_local_network"
subcategory: "Networking"
description: "Select a site-local virtual network when connectivity must remain within one site."
xcsh_docs: {"aliases": ["site local network"], "body_bytes": 777, "body_sha256": "sha256:6f91e1ca287aadcbab56951c0b99783a52bab95ddd56ae946e87de05b153b90f", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_network:properties:site_local_network", "parent_id": "xcsh-docs:data-sources:virtual_network:reference", "path": "documentation/data-sources/virtual_network/properties/site_local_network/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0021131112022130-2303120033303331-1112013031030100-1222221101001130-3221130212300323-1020002201031323-1003130311113210-0211332332033212", "registry_path": "docs/guides/data-sources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_local_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_network/properties/site_local_network/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Select a site-local virtual network when connectivity must remain within one site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_local_network

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/)
- site_local_network

<a id="section"></a>

Type: `["object", {}]`. Computed.

Select a site-local virtual network when connectivity must remain within one site.

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

This is an empty object or choice marker. It has no direct properties.
