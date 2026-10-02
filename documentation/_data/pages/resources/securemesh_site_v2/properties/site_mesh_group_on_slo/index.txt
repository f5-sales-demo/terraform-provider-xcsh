---
page_title: "site_mesh_group_on_slo"
subcategory: ""
description: "Select how the site mesh group will be connected. By default, public IPs of the control nodes of the site will be used."
xcsh_docs: {"aliases": ["site mesh group on slo"], "body_bytes": 3177, "body_sha256": "sha256:a9d5e0b04f347131783c75ba9d8c43f6383ab312ab105c857f28171a640f585d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "site_mesh_group_on_slo:ConflictingObjectAttributes:no_site_mesh_group,site_mesh_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_mesh_group_on_slo:ConflictingObjectAttributes:no_site_mesh_group,site_mesh_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_mesh_group_on_slo:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_mesh_group_on_slo:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["site_mesh_group_on_slo"], "schema_version": 1, "sections": [{"aliases": ["no site mesh group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:no_site_mesh_group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "no_site_mesh_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["site mesh group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-site_mesh_group_on_slo--site_mesh_group--name", "enforcement": "provider-schema", "group": "site_mesh_group_on_slo.site_mesh_group:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:site_mesh_group", "type": "requires"}], "schema_path": ["site_mesh_group_on_slo", "site_mesh_group"], "syntax": "block", "type": "object"}, {"aliases": ["sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_pvt_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_mesh_group_on_slo", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Select how the site mesh group will be connected. By default, public IPs of the control nodes of the site will be used.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_mesh_group_on_slo

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- site_mesh_group_on_slo

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_site_mesh_group",
    "site_mesh_group"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
site_mesh_group_on_slo {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/no_site_mesh_group/): complete subsection reference.

- [site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/site_mesh_group/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_pvt_ip/): complete subsection reference.

## Next pages

- [site_mesh_group_on_slo.no_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/no_site_mesh_group/)
- [site_mesh_group_on_slo.site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/site_mesh_group/)
- [site_mesh_group_on_slo.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_public_ip/)
- [site_mesh_group_on_slo.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
