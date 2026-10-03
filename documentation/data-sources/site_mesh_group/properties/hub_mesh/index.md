---
page_title: "hub_mesh"
subcategory: "Infrastructure"
description: "Details of Hub Full Mesh Group Type."
xcsh_docs: {"aliases": ["hub mesh"], "body_bytes": 1785, "body_sha256": "sha256:136839bf764d337eb5872f372adde56d192d8b8d3dc66b80a01ad1c10c3883d4", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh:control_and_data_plane_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh:data_plane_mesh"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh", "parent_id": "xcsh-docs:data-sources:site_mesh_group:reference", "path": "documentation/data-sources/site_mesh_group/properties/hub_mesh/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3310111110031031-0031312012323212-1332110322233232-3121231211210320-1320111110002130-2310110100121303-2100213103112331-3310121200102323", "registry_path": "docs/guides/data-sources--site_mesh_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["hub_mesh"], "schema_version": 1, "sections": [{"aliases": ["hub mesh control and data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh:control_and_data_plane_mesh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["hub_mesh", "control_and_data_plane_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["hub mesh data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh:data_plane_mesh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["hub_mesh", "data_plane_mesh"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/hub_mesh/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Details of Hub Full Mesh Group Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hub_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/)
- hub_mesh

<a id="section"></a>

Type: `"single"`. Computed.

Hub Full Mesh. Details of Hub Full Mesh Group Type.

Upstream description:

Details of Hub Full Mesh Group Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-hub_full_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

## Direct properties

- [control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/control_and_data_plane_mesh/): complete subsection reference.

- [data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/data_plane_mesh/): complete subsection reference.

## Next pages

- [hub_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/control_and_data_plane_mesh/)
- [hub_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/data_plane_mesh/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
