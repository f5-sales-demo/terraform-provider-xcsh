---
page_title: "re_acl"
subcategory: ""
description: "Fast ACL definition for RE."
xcsh_docs: {"aliases": ["re acl"], "body_bytes": 3089, "body_sha256": "sha256:976ff01fc48e0018648df671a93a77241949cb4b1f4b3d0c8d92040927e1f885", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl", "parent_id": "xcsh-docs:resources:fast_acl:reference", "path": "documentation/resources/fast_acl/properties/re_acl/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:all_public_vips,default_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:all_public_vips,selected_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:all_public_vips,default_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:default_tenant_vip,selected_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:all_public_vips,selected_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:default_tenant_vip,selected_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl"], "schema_version": 1, "sections": [{"aliases": ["all public vips"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "all_public_vips"], "syntax": "attribute", "type": "object"}, {"aliases": ["default tenant vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "default_tenant_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["fast acl rules"], "anchor": "section", "description": "Fast ACL rules to match.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:prefix", "type": "conflicts"}], "schema_path": ["re_acl", "fast_acl_rules"], "syntax": "block", "type": "object"}, {"aliases": ["selected tenant vip"], "anchor": "section", "description": "Select various tenant public VIP(s)", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.selected_tenant_vip:RequiredObjectAttributes:public_ip_refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs", "type": "requires"}], "schema_path": ["re_acl", "selected_tenant_vip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Fast ACL definition for RE.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# re_acl

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- re_acl

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: re\_acl, site\_acl\] Fast ACL for RE. Fast ACL definition for RE.

Upstream description:

Fast ACL definition for RE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_public_vips",
    "default_tenant_vip"),
  validators.ConflictingObjectAttributes("all_public_vips",
    "selected_tenant_vip"),
  validators.ConflictingObjectAttributes("default_tenant_vip",
    "selected_tenant_vip")}
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
  "x-ves-oneof-field-vip_choice": "[\"all_public_vips\",\"default_tenant_vip\",\"selected_tenant_vip\"]"
}
```

OneOf alternatives in this subsection:

- [re_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/#section)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
re_acl {
  # Configure direct properties listed below.
}
```

## Direct properties

- [all_public_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/all_public_vips/): complete subsection reference.

- [default_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/default_tenant_vip/): complete subsection reference.

- [fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/): complete subsection reference.

- [selected_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/selected_tenant_vip/): complete subsection reference.

## Next pages

- [re_acl.all_public_vips](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/all_public_vips/)
- [re_acl.default_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/default_tenant_vip/)
- [re_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/fast_acl_rules/)
- [re_acl.selected_tenant_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/re_acl/selected_tenant_vip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
