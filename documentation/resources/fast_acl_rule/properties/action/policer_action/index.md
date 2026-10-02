---
page_title: "action.policer_action"
subcategory: ""
description: "Reference to policer object."
xcsh_docs: {"aliases": ["action policer action"], "body_bytes": 1508, "body_sha256": "sha256:3c05d664a0cc37ad52f1af5d0f4ab580ac1db403f71bb6cb9863f35378021830", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl_rule:properties:action:policer_action:ref"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "parent_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "path": "documentation/resources/fast_acl_rule/properties/action/policer_action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003", "registry_path": "docs/guides/resources--fast_acl_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["action", "policer_action"], "schema_version": 1, "sections": [{"aliases": ["ref"], "anchor": "section", "description": "A policer direct reference.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["action", "policer_action", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/properties/action/policer_action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Reference to policer object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# action.policer_action

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/)
- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/)
- action.policer_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
policer_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/ref/): complete subsection reference.

## Next pages

- [action.policer_action.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/policer_action/ref/)
- [action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/properties/action/)
- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl_rule/)
