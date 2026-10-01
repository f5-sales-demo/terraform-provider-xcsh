---
page_title: "spoke_mesh"
subcategory: "Infrastructure"
description: "spoke_mesh for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 2118, "body_sha256": "sha256:e1f8d6e4ef5f29d784d6c9acfe2856d6b88ed70cb126920699565a3276653dfc", "child_ids": ["xcsh-docs:data-sources:site_mesh_group:properties:spoke_mesh:control_and_data_plane_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:spoke_mesh:data_plane_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:spoke_mesh:hub_mesh_group"], "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:properties:spoke_mesh", "parent_id": "xcsh-docs:data-sources:site_mesh_group:reference", "path": "documentation/data-sources/site_mesh_group/properties/spoke_mesh/index.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["spoke_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/spoke_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "spoke_mesh for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# spoke_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/)
- spoke_mesh

<a id="section"></a>

Type: `"single"`. Computed.

Spoke. Details of Spoke Mesh Group Type.

Upstream description:

Details of Spoke Mesh Group Type.

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

## Direct properties

- [control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/control_and_data_plane_mesh/): complete subsection reference.

- [data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/data_plane_mesh/): complete subsection reference.

- [hub_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/): complete subsection reference.

## Next pages

- [spoke_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/control_and_data_plane_mesh/)
- [spoke_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/data_plane_mesh/)
- [spoke_mesh.hub_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
