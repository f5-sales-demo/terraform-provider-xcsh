---
page_title: "type"
subcategory: ""
description: "type for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": [], "body_bytes": 1788, "body_sha256": "sha256:4792e4913d7f908540d2c55bb4dc21e2642a8641fed84837d83671642661bfe4", "child_ids": ["xcsh-docs:data-sources:dc_cluster_group:properties:type:control_and_data_plane_mesh", "xcsh-docs:data-sources:dc_cluster_group:properties:type:data_plane_mesh"], "collection_id": "xcsh-docs:data-sources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dc_cluster_group:properties:type", "parent_id": "xcsh-docs:data-sources:dc_cluster_group:reference", "path": "documentation/data-sources/dc_cluster_group/properties/type/index.md", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dc_cluster_group/properties/type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "type for xcsh_dc_cluster_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
