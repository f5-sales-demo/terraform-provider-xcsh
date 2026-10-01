---
page_title: "hub_mesh.data_plane_mesh"
subcategory: "Infrastructure"
description: "hub_mesh.data_plane_mesh for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 983, "body_sha256": "sha256:fe8dacef235b76c5a57c64747750ce5732a5b4d6a2424e7b82dfdefd72209b6c", "canonical_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:data_plane_mesh", "child_ids": [], "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:data_plane_mesh", "parent_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh", "path": "docs/guides/resources--site_mesh_group--properties--hub_mesh--data_plane_mesh.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["hub_mesh", "data_plane_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/hub_mesh/data_plane_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "hub_mesh.data_plane_mesh for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hub_mesh.data_plane_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
- [Property reference](resources--site_mesh_group--reference.md)
- [hub_mesh](resources--site_mesh_group--properties--hub_mesh.md)
- hub_mesh.data_plane_mesh

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

- [hub_mesh](resources--site_mesh_group--properties--hub_mesh.md)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
