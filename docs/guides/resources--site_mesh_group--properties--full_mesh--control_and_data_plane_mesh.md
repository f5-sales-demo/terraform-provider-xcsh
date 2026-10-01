---
page_title: "full_mesh.control_and_data_plane_mesh"
subcategory: "Infrastructure"
description: "full_mesh.control_and_data_plane_mesh for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1025, "body_sha256": "sha256:71c378634dbf51756f1508674997331b057a92e23f65fd2abd92fc4ab21cb2a4", "canonical_id": "xcsh-docs:resources:site_mesh_group:properties:full_mesh:control_and_data_plane_mesh", "child_ids": [], "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:full_mesh:control_and_data_plane_mesh", "parent_id": "xcsh-docs:resources:site_mesh_group:properties:full_mesh", "path": "docs/guides/resources--site_mesh_group--properties--full_mesh--control_and_data_plane_mesh.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["full_mesh", "control_and_data_plane_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/full_mesh/control_and_data_plane_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "full_mesh.control_and_data_plane_mesh for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# full_mesh.control_and_data_plane_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
- [Property reference](resources--site_mesh_group--reference.md)
- [full_mesh](resources--site_mesh_group--properties--full_mesh.md)
- full_mesh.control_and_data_plane_mesh

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
control_and_data_plane_mesh = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [full_mesh](resources--site_mesh_group--properties--full_mesh.md)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
