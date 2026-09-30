---
page_title: "type.data_plane_mesh"
subcategory: ""
description: "type.data_plane_mesh for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": [], "body_bytes": 867, "body_sha256": "sha256:16eae049f37f1272d9cdde2f75b86671f53073bc22c2acf3318a648d868e5ab9", "canonical_id": "xcsh-docs:resources:dc_cluster_group:properties:type:data_plane_mesh", "child_ids": [], "collection_id": "xcsh-docs:resources:dc_cluster_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:dc_cluster_group:properties:type:data_plane_mesh", "parent_id": "xcsh-docs:resources:dc_cluster_group:properties:type", "path": "docs/guides/resources--dc_cluster_group--properties--type--data_plane_mesh.md", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["type", "data_plane_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dc_cluster_group/properties/type/data_plane_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "type.data_plane_mesh for xcsh_dc_cluster_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# type.data_plane_mesh

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md)
- [Property reference](resources--dc_cluster_group--reference.md)
- [type](resources--dc_cluster_group--properties--type.md)
- type.data_plane_mesh

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
data_plane_mesh = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [type](resources--dc_cluster_group--properties--type.md)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md)
