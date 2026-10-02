---
page_title: "type"
subcategory: ""
description: "Details of DC Cluster Group Mesh Type."
xcsh_docs: {"aliases": ["type"], "body_bytes": 1788, "body_sha256": "sha256:4792e4913d7f908540d2c55bb4dc21e2642a8641fed84837d83671642661bfe4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:dc_cluster_group:properties:type:control_and_data_plane_mesh", "xcsh-docs:data-sources:dc_cluster_group:properties:type:data_plane_mesh"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dc_cluster_group:properties:type", "parent_id": "xcsh-docs:data-sources:dc_cluster_group:reference", "path": "documentation/data-sources/dc_cluster_group/properties/type/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2121331020332121-1223023231322301-0313101210122203-0120122211011202-2000020120123123-1221321122023132-0133202212221032-0322033003200322", "registry_path": "docs/guides/data-sources--dc_cluster_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["type"], "schema_version": 1, "sections": [{"aliases": ["control and data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dc_cluster_group:properties:type:control_and_data_plane_mesh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["type", "control_and_data_plane_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dc_cluster_group:properties:type:data_plane_mesh", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["type", "data_plane_mesh"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dc_cluster_group/properties/type/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Details of DC Cluster Group Mesh Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# type

Breadcrumbs:

- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/properties/)
- type

<a id="section"></a>

Type: `"single"`. Computed.

DC Cluster Group Mesh Type. Details of DC Cluster Group Mesh Type.

Upstream description:

Details of DC Cluster Group Mesh Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

## Direct properties

- [control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/properties/type/control_and_data_plane_mesh/): complete subsection reference.

- [data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/properties/type/data_plane_mesh/): complete subsection reference.

## Next pages

- [type.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/properties/type/control_and_data_plane_mesh/)
- [type.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/properties/type/data_plane_mesh/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/properties/)
- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/)
