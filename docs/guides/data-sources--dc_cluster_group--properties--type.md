---
page_title: "type"
subcategory: ""
description: "type for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": [], "body_bytes": 1281, "body_sha256": "sha256:61cf506aa2f89b536ab5d085a16935ba299ea4331f591eb7d38ddff832613872", "canonical_id": "xcsh-docs:data-sources:dc_cluster_group:properties:type", "child_ids": ["xcsh-docs:data-sources:dc_cluster_group:properties:type:control_and_data_plane_mesh", "xcsh-docs:data-sources:dc_cluster_group:properties:type:data_plane_mesh"], "collection_id": "xcsh-docs:data-sources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dc_cluster_group:properties:type", "parent_id": "xcsh-docs:data-sources:dc_cluster_group:reference", "path": "docs/guides/data-sources--dc_cluster_group--properties--type.md", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dc_cluster_group/properties/type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "type for xcsh_dc_cluster_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# type

Breadcrumbs:

- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md)
- [Property reference](data-sources--dc_cluster_group--reference.md)
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

- [control_and_data_plane_mesh](data-sources--dc_cluster_group--properties--type--control_and_data_plane_mesh.md): complete subsection reference.

- [data_plane_mesh](data-sources--dc_cluster_group--properties--type--data_plane_mesh.md): complete subsection reference.

## Next pages

- [type.control_and_data_plane_mesh](data-sources--dc_cluster_group--properties--type--control_and_data_plane_mesh.md)
- [type.data_plane_mesh](data-sources--dc_cluster_group--properties--type--data_plane_mesh.md)
- [Property reference](data-sources--dc_cluster_group--reference.md)
- [xcsh_dc_cluster_group](../data-sources/dc_cluster_group.md)
