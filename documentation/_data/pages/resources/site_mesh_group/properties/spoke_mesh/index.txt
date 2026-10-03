---
page_title: "spoke_mesh"
subcategory: "Infrastructure"
description: "Details of Spoke Mesh Group Type."
xcsh_docs: {"aliases": ["spoke mesh"], "body_bytes": 2399, "body_sha256": "sha256:1de4e681c5db563b0b60789b29fe9d68f401cc18635e0ebbc7b85f2c17e66198", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:control_and_data_plane_mesh", "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:data_plane_mesh", "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:hub_mesh_group"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "documentation/resources/site_mesh_group/properties/spoke_mesh/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2020322201021122-0131011313222121-0101022011301302-3002331111101232-3230013331330023-0032312121112300-2322110320130302-2012302020131121", "registry_path": "docs/guides/resources--site_mesh_group--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "spoke_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:control_and_data_plane_mesh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "spoke_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:data_plane_mesh", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["spoke_mesh"], "schema_version": 1, "sections": [{"aliases": ["spoke mesh control and data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:control_and_data_plane_mesh", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["spoke_mesh", "control_and_data_plane_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["spoke mesh data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:data_plane_mesh", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["spoke_mesh", "data_plane_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["spoke mesh hub mesh group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:hub_mesh_group", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-spoke_mesh--hub_mesh_group--name", "enforcement": "provider-schema", "group": "spoke_mesh.hub_mesh_group:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:spoke_mesh:hub_mesh_group", "type": "requires"}], "schema_path": ["spoke_mesh", "hub_mesh_group"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/spoke_mesh/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Details of Spoke Mesh Group Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# spoke_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
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

- [control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/control_and_data_plane_mesh/): complete subsection reference.

- [data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/data_plane_mesh/): complete subsection reference.

- [hub_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/): complete subsection reference.

## Next pages

- [spoke_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/control_and_data_plane_mesh/)
- [spoke_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/data_plane_mesh/)
- [spoke_mesh.hub_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/spoke_mesh/hub_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
