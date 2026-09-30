---
page_title: "full_mesh"
subcategory: "Infrastructure"
description: "full_mesh for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1936, "body_sha256": "sha256:d5ba8565773f81bb29b0dd60be4931e16e3c6941dde666fb97a4c2ead473060a", "canonical_id": "xcsh-docs:resources:site_mesh_group:properties:full_mesh", "child_ids": ["xcsh-docs:resources:site_mesh_group:properties:full_mesh:control_and_data_plane_mesh", "xcsh-docs:resources:site_mesh_group:properties:full_mesh:data_plane_mesh"], "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:full_mesh", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "docs/guides/resources--site_mesh_group--properties--full_mesh.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["full_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/full_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "full_mesh for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# full_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
- [Property reference](resources--site_mesh_group--reference.md)
- full_mesh

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: full\_mesh, hub\_mesh, spoke\_mesh\] Full Mesh. Details of Full Mesh Group Type.

Upstream description:

Details of Full Mesh Group Type.

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
  "x-ves-oneof-field-full_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

OneOf alternatives in this subsection:

- [full_mesh](resources--site_mesh_group--properties--full_mesh.md#section)
- [hub_mesh](resources--site_mesh_group--properties--hub_mesh.md#section)
- [spoke_mesh](resources--site_mesh_group--properties--spoke_mesh.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
full_mesh {
  # Configure direct properties listed below.
}
```

## Direct properties

- [control_and_data_plane_mesh](resources--site_mesh_group--properties--full_mesh--control_and_data_plane_mesh.md): complete subsection reference.

- [data_plane_mesh](resources--site_mesh_group--properties--full_mesh--data_plane_mesh.md): complete subsection reference.

## Next pages

- [full_mesh.control_and_data_plane_mesh](resources--site_mesh_group--properties--full_mesh--control_and_data_plane_mesh.md)
- [full_mesh.data_plane_mesh](resources--site_mesh_group--properties--full_mesh--data_plane_mesh.md)
- [Property reference](resources--site_mesh_group--reference.md)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
