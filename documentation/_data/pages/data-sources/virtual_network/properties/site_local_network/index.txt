---
page_title: "site_local_network"
subcategory: "Networking"
description: "Select a site-local virtual network when connectivity must remain within one site."
xcsh_docs: {"aliases": ["site local network"], "body_bytes": 777, "body_sha256": "sha256:6f91e1ca287aadcbab56951c0b99783a52bab95ddd56ae946e87de05b153b90f", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_network:properties:site_local_network", "parent_id": "xcsh-docs:data-sources:virtual_network:reference", "path": "documentation/data-sources/virtual_network/properties/site_local_network/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0021131112022130-2303120033303331-1112013031030100-1222221101001130-3221130212300323-1020002201031323-1003130311113210-0211332332033212", "registry_path": "docs/guides/data-sources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_local_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_network/properties/site_local_network/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Select a site-local virtual network when connectivity must remain within one site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
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
