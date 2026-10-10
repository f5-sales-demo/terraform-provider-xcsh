---
page_title: "action.policer_action"
subcategory: ""
description: "Reference to policer object."
xcsh_docs: {"aliases": ["action policer action"], "body_bytes": 1074, "body_sha256": "sha256:6f2fd0382b892b0850295664e7ad678a7a2ea8424f94e7d04257714dba0b22ee", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl_rule:properties:action:policer_action:ref"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "parent_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "path": "documentation/resources/fast_acl_rule/properties/action/policer_action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1220210013232110-1103010033322231-2200331031121032-1010021122312012-0112233031130010-2331101033211202-1022320330321321-0221101310331003", "registry_path": "docs/guides/resources--fast_acl_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["action", "policer_action"], "schema_version": 1, "sections": [{"aliases": ["action policer action ref"], "anchor": "section", "description": "A policer direct reference.", "document_id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action:ref", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["action", "policer_action", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/properties/action/policer_action/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reference to policer object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
