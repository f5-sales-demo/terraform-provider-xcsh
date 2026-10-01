---
page_title: "type"
subcategory: ""
description: "type for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": [], "body_bytes": 1661, "body_sha256": "sha256:7604c3845bfe286fcf919806936fe8047b441e2644848a746c898820a996b9f9", "canonical_id": "xcsh-docs:resources:dc_cluster_group:properties:type", "child_ids": ["xcsh-docs:resources:dc_cluster_group:properties:type:control_and_data_plane_mesh", "xcsh-docs:resources:dc_cluster_group:properties:type:data_plane_mesh"], "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:dc_cluster_group:properties:type", "parent_id": "xcsh-docs:resources:dc_cluster_group:reference", "path": "docs/guides/resources--dc_cluster_group--properties--type.md", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/properties/type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "type for xcsh_dc_cluster_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# type

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md)
- [Property reference](resources--dc_cluster_group--reference.md)
- type

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

DC Cluster Group Mesh Type. Details of DC Cluster Group Mesh Type.

Upstream description:

Details of DC Cluster Group Mesh Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("control_and_data_plane_mesh",
    "data_plane_mesh")}
```

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

Terraform syntax:

```terraform
type {
  # Configure direct properties listed below.
}
```

## Direct properties

- [control_and_data_plane_mesh](resources--dc_cluster_group--properties--type--control_and_data_plane_mesh.md): complete subsection reference.

- [data_plane_mesh](resources--dc_cluster_group--properties--type--data_plane_mesh.md): complete subsection reference.

## Next pages

- [type.control_and_data_plane_mesh](resources--dc_cluster_group--properties--type--control_and_data_plane_mesh.md)
- [type.data_plane_mesh](resources--dc_cluster_group--properties--type--data_plane_mesh.md)
- [Property reference](resources--dc_cluster_group--reference.md)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md)
