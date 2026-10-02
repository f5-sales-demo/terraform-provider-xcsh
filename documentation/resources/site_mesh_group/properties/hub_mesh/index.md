---
page_title: "hub_mesh"
subcategory: "Infrastructure"
description: "Details of Hub Full Mesh Group Type."
xcsh_docs: {"aliases": ["hub mesh"], "body_bytes": 2070, "body_sha256": "sha256:d8bb4deb3bc034a1383c6f2aa67dc1d718478108d6474081def735c9bb21d110", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:site_mesh_group:properties:hub_mesh:control_and_data_plane_mesh", "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:data_plane_mesh"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "documentation/resources/site_mesh_group/properties/hub_mesh/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1011133001312013-2312132233110312-1332302222102230-1122130331221011-0121321303130030-0111233333133330-2112011323002320-2210320323220112", "registry_path": "docs/guides/resources--site_mesh_group--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "hub_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:control_and_data_plane_mesh", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "hub_mesh:ConflictingObjectAttributes:control_and_data_plane_mesh,data_plane_mesh", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:data_plane_mesh", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["hub_mesh"], "schema_version": 1, "sections": [{"aliases": ["control and data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:control_and_data_plane_mesh", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["hub_mesh", "control_and_data_plane_mesh"], "syntax": "attribute", "type": "object"}, {"aliases": ["data plane mesh"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:site_mesh_group:properties:hub_mesh:data_plane_mesh", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["hub_mesh", "data_plane_mesh"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/hub_mesh/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Details of Hub Full Mesh Group Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hub_mesh

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
- hub_mesh

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Hub Full Mesh. Details of Hub Full Mesh Group Type.

Upstream description:

Details of Hub Full Mesh Group Type.

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
  "x-ves-oneof-field-hub_full_mesh_choice": "[\"control_and_data_plane_mesh\",\"data_plane_mesh\"]"
}
```

Terraform syntax:

```terraform
hub_mesh {
  # Configure direct properties listed below.
}
```

## Direct properties

- [control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/control_and_data_plane_mesh/): complete subsection reference.

- [data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/data_plane_mesh/): complete subsection reference.

## Next pages

- [hub_mesh.control_and_data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/control_and_data_plane_mesh/)
- [hub_mesh.data_plane_mesh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/hub_mesh/data_plane_mesh/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
