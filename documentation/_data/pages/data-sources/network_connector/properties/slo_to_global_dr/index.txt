---
page_title: "slo_to_global_dr"
subcategory: "Networking"
description: "Global network reference for direct connection."
xcsh_docs: {"aliases": ["slo to global dr"], "body_bytes": 1275, "body_sha256": "sha256:2ed9c50e8c071c2d9b9931e1a9406f5d5e9e39ccd290d3ceb4cdaec17e222412", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_connector:properties:slo_to_global_dr:global_vn"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:properties:slo_to_global_dr", "parent_id": "xcsh-docs:data-sources:network_connector:reference", "path": "documentation/data-sources/network_connector/properties/slo_to_global_dr/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3202010033321000-2321111032303003-3110021002030311-2331130332311233-1122123100201221-1020301003032212-3023213002320313-1232123203201023", "registry_path": "docs/guides/data-sources--network_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["slo_to_global_dr"], "schema_version": 1, "sections": [{"aliases": ["slo to global dr global vn"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:network_connector:properties:slo_to_global_dr:global_vn", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["slo_to_global_dr", "global_vn"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/properties/slo_to_global_dr/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Global network reference for direct connection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slo_to_global_dr

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- slo_to_global_dr

<a id="section"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

- [global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/slo_to_global_dr/global_vn/): complete subsection reference.

## Next pages

- [slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/slo_to_global_dr/global_vn/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/)
