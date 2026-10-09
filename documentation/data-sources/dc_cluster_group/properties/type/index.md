---
page_title: "type"
subcategory: ""
description: "Details of DC Cluster Group Mesh Type."
xcsh_docs: {"aliases": ["type"], "body_bytes": 1159, "body_sha256": "sha256:9b175f6c598264c7d6f7bcb3bd8e01ec564bc89d0ae32ff4ca135448585dc42c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:dc_cluster_group:properties:type:control_and_data_plane_mesh", "xcsh-docs:data-sources:dc_cluster_group:properties:type:data_plane_mesh"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dc_cluster_group:properties:type", "parent_id": "xcsh-docs:data-sources:dc_cluster_group:reference", "path": "documentation/data-sources/dc_cluster_group/properties/type/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2121331020332121-1223023231322301-0313101210122203-0120122211011202-2000020120123123-1221321122023132-0133202212221032-0322033003200322", "registry_path": "docs/guides/data-sources--dc_cluster_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["type"], "schema_version": 1, "sections": [{"aliases": ["type control and data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dc_cluster_group:properties:type:control_and_data_plane_mesh", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["type", "control_and_data_plane_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["type data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:dc_cluster_group:properties:type:data_plane_mesh", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["type", "data_plane_mesh"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dc_cluster_group/properties/type/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Details of DC Cluster Group Mesh Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
