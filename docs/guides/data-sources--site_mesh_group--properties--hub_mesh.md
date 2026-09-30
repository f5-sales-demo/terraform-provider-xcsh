---
page_title: "hub_mesh"
subcategory: "Infrastructure"
description: "hub_mesh for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1278, "body_sha256": "sha256:1e5a4c45fc1ef21ff2e70e7998c37a7b6bcf71c223094fc0a845ee47bd7ea428", "canonical_id": "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh", "child_ids": ["xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh:control_and_data_plane_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh:data_plane_mesh"], "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:properties:hub_mesh", "parent_id": "xcsh-docs:data-sources:site_mesh_group:reference", "path": "docs/guides/data-sources--site_mesh_group--properties--hub_mesh.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["hub_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/hub_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "hub_mesh for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# hub_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
- [Property reference](data-sources--site_mesh_group--reference.md)
- hub_mesh

<a id="section"></a>

Type: `"single"`. Computed.

Hub Full Mesh. Details of Hub Full Mesh Group Type.

Upstream description:

Details of Hub Full Mesh Group Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-hub_full_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

## Direct properties

- [control_and_data_plane_mesh](data-sources--site_mesh_group--properties--hub_mesh--control_and_data_plane_mesh.md): complete subsection reference.

- [data_plane_mesh](data-sources--site_mesh_group--properties--hub_mesh--data_plane_mesh.md): complete subsection reference.

## Next pages

- [hub_mesh.control_and_data_plane_mesh](data-sources--site_mesh_group--properties--hub_mesh--control_and_data_plane_mesh.md)
- [hub_mesh.data_plane_mesh](data-sources--site_mesh_group--properties--hub_mesh--data_plane_mesh.md)
- [Property reference](data-sources--site_mesh_group--reference.md)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
