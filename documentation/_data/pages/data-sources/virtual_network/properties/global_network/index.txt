---
page_title: "global_network"
subcategory: "Networking"
description: "Select the global virtual-network scope for connectivity across participating sites."
xcsh_docs: {"aliases": ["global network"], "body_bytes": 1786, "body_sha256": "sha256:943d1d3249766609977239f76eddf4031ab7a0f34b1d2e66077e4174a89b9af0", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_network:properties:global_network", "parent_id": "xcsh-docs:data-sources:virtual_network:reference", "path": "documentation/data-sources/virtual_network/properties/global_network/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1332013022320101-1021031123100220-2033102232003030-3201221313323023-2302312133113002-0113321030111130-3313002330222000-3033021020231011", "registry_path": "docs/guides/data-sources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["global_network"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_network/properties/global_network/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Select the global virtual-network scope for connectivity across participating sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# global_network

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/)
- global_network

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: global\_network, site\_local\_inside\_network, site\_local\_network\] Select the global
virtual-network scope for connectivity across participating sites.

Upstream description:

Select the global virtual-network scope for connectivity across participating sites.

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

- [global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/global_network/#section)
- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/site_local_inside_network/#section)
- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/site_local_network/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/)
- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
