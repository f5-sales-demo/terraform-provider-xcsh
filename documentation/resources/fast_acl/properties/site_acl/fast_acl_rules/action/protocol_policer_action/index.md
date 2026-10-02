---
page_title: "site_acl.fast_acl_rules.action.protocol_policer_action"
subcategory: ""
description: "Reference to policer object."
xcsh_docs: {"aliases": ["site acl fast acl rules action protocol policer action"], "body_bytes": 1992, "body_sha256": "sha256:398fc73413895b141b76d1d51adb39978591e60e629be2d5fcbde49c86c8fa12", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action:ref"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action", "parent_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action", "path": "documentation/resources/fast_acl/properties/site_acl/fast_acl_rules/action/protocol_policer_action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3223200333033130-3333302210113232-3022221201003133-1200121200002033-0303020320120312-0210322323011100-1110022122213113-2313000130133000", "registry_path": "docs/guides/resources--fast_acl--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "action", "protocol_policer_action"], "schema_version": 1, "sections": [{"aliases": ["ref"], "anchor": "section", "description": "Reference to protocol policer object.", "document_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action:ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["site_acl", "fast_acl_rules", "action", "protocol_policer_action", "ref"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/fast_acl_rules/action/protocol_policer_action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Reference to policer object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules.action.protocol_policer_action

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/)
- [site_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/)
- [site_acl.fast_acl_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/)
- [site_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/)
- site_acl.fast_acl_rules.action.protocol_policer_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Protocol Policer Reference. Reference to policer object.

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
protocol_policer_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/protocol_policer_action/ref/): complete subsection reference.

## Next pages

- [site_acl.fast_acl_rules.action.protocol_policer_action.ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/protocol_policer_action/ref/)
- [site_acl.fast_acl_rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/properties/site_acl/fast_acl_rules/action/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fast_acl/)
