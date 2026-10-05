---
page_title: "re_acl"
subcategory: ""
description: "Fast ACL definition for RE."
xcsh_docs: {"aliases": ["re acl"], "body_bytes": 3119, "body_sha256": "sha256:340c4a84ef95ffd222537349194670392a591748dfefdeb988bd0026c5660407", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:re_acl", "parent_id": "xcsh-docs:resources:fast_acl:reference", "path": "documentation/resources/fast_acl/properties/re_acl/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1231022222302112-1010301103333012-2213133232333203-2331003233132011-2120200223000031-2211020102213001-0013011133103333-3102033223202221", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:all_public_vips,default_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:all_public_vips,selected_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:all_public_vips,default_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:default_tenant_vip,selected_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:all_public_vips,selected_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl:ConflictingObjectAttributes:default_tenant_vip,selected_tenant_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["re_acl"], "schema_version": 1, "sections": [{"aliases": ["re acl all public vips"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:all_public_vips", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "all_public_vips"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl default tenant vip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:default_tenant_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["re_acl", "default_tenant_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["re acl fast acl rules"], "anchor": "section", "description": "Fast ACL rules to match.", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:ip_prefix_set", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.fast_acl_rules:ConflictingListObjectAttributes:ip_prefix_set,prefix", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:fast_acl_rules:prefix", "type": "conflicts"}], "schema_path": ["re_acl", "fast_acl_rules"], "syntax": "block", "type": "object"}, {"aliases": ["re acl selected tenant vip"], "anchor": "section", "description": "Select various tenant public VIP(s)", "document_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "re_acl.selected_tenant_vip:RequiredObjectAttributes:public_ip_refs", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fast_acl:properties:re_acl:selected_tenant_vip:public_ip_refs", "type": "requires"}], "schema_path": ["re_acl", "selected_tenant_vip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/re_acl/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Fast ACL definition for RE.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
EnumExtractionComplete: false
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
