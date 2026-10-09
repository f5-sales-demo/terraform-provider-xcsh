---
page_title: "full_mesh"
subcategory: "Infrastructure"
description: "Details of Full Mesh Group Type."
xcsh_docs: {"aliases": ["full mesh"], "body_bytes": 1681, "body_sha256": "sha256:6e567d7aa69d0c45d04c541c8a979e9cd5bdcdb87d7e0a9a6add304d98427a19", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:site_mesh_group:properties:full_mesh:control_and_data_plane_mesh", "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh:data_plane_mesh"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh", "parent_id": "xcsh-docs:data-sources:site_mesh_group:reference", "path": "documentation/data-sources/site_mesh_group/properties/full_mesh/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2011311010303311-3101333312312330-0321213121133131-0132222021121301-0211101131120320-2330312212002010-0320002102331223-2100120012300212", "registry_path": "docs/guides/data-sources--site_mesh_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["full_mesh"], "schema_version": 1, "sections": [{"aliases": ["full mesh control and data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh:control_and_data_plane_mesh", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["full_mesh", "control_and_data_plane_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["full mesh data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:site_mesh_group:properties:full_mesh:data_plane_mesh", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["full_mesh", "data_plane_mesh"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/properties/full_mesh/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Details of Full Mesh Group Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
