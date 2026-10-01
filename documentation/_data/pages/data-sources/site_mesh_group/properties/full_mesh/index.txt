---
page_title: "full_mesh"
subcategory: "Infrastructure"
description: "full_mesh for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 2319, "body_sha256": "sha256:763bde2ad1369fd85c45b498d345385a536549b5ab3f9291bca2bbfa7cdd00db", "child_ids": ["xcsh-docs:data-sources:site_mesh_group:properties:full_mesh:control_and_data_plane_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh:data_plane_mesh"], "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh", "parent_id": "xcsh-docs:data-sources:site_mesh_group:reference", "path": "documentation/data-sources/site_mesh_group/properties/full_mesh/index.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["full_mesh"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/full_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "full_mesh for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# full_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/)
- full_mesh

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: full\_mesh, hub\_mesh, spoke\_mesh\] Full Mesh. Details of Full Mesh Group Type.

Upstream description:

Details of Full Mesh Group Type.

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

- [full_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/#section)
- [hub_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/hub_mesh/#section)
- [spoke_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/spoke_mesh/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/control_and_data_plane_mesh/): complete subsection reference.

- [data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/data_plane_mesh/): complete subsection reference.

## Next pages

- [full_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/control_and_data_plane_mesh/)
- [full_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/full_mesh/data_plane_mesh/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/properties/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
