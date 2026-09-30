---
page_title: "spoke_mesh"
subcategory: "Infrastructure"
description: "spoke_mesh for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1792, "body_sha256": "sha256:965396175d0c36576bf56ae4f157888eb0d06e0412227601540d547cb75dbf68", "canonical_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh", "child_ids": ["xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:control_and_data_plane_mesh", "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:data_plane_mesh", "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:hub_mesh_group"], "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "docs/guides/resources--site_mesh_group--properties--spoke_mesh.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["spoke_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/spoke_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "spoke_mesh for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# spoke_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
- [Property reference](resources--site_mesh_group--reference.md)
- spoke_mesh

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Spoke. Details of Spoke Mesh Group Type.

Upstream description:

Details of Spoke Mesh Group Type.

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
  "x-ves-oneof-field-spoke_hub_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

Terraform syntax:

```terraform
spoke_mesh {
  # Configure direct properties listed below.
}
```

## Direct properties

- [control_and_data_plane_mesh](resources--site_mesh_group--properties--spoke_mesh--control_and_data_plane_mesh.md): complete subsection reference.

- [data_plane_mesh](resources--site_mesh_group--properties--spoke_mesh--data_plane_mesh.md): complete subsection reference.

- [hub_mesh_group](resources--site_mesh_group--properties--spoke_mesh--hub_mesh_group.md): complete subsection reference.

## Next pages

- [spoke_mesh.control_and_data_plane_mesh](resources--site_mesh_group--properties--spoke_mesh--control_and_data_plane_mesh.md)
- [spoke_mesh.data_plane_mesh](resources--site_mesh_group--properties--spoke_mesh--data_plane_mesh.md)
- [spoke_mesh.hub_mesh_group](resources--site_mesh_group--properties--spoke_mesh--hub_mesh_group.md)
- [Property reference](resources--site_mesh_group--reference.md)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
