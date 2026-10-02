---
page_title: "bot_action"
subcategory: ""
description: "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing."
xcsh_docs: {"aliases": ["bot action"], "body_bytes": 2006, "body_sha256": "sha256:e6fa871f38ffdcd82485b1b857a2a3edaeee9570d3df5f4314425d62493e442a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:bot_action:bot_skip_processing", "xcsh-docs:resources:service_policy_rule:properties:bot_action:none"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:bot_action", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/bot_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2103123012221321-1103321133011201-1313232020222001-2012201222102131-0300222221031123-2133203022221201-3200132333302332-1223010122321102", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_action:ConflictingObjectAttributes:bot_skip_processing,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:bot_skip_processing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_action:ConflictingObjectAttributes:bot_skip_processing,none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:none", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_action"], "schema_version": 1, "sections": [{"aliases": ["bot skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:bot_skip_processing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_action", "bot_skip_processing"], "syntax": "attribute", "type": "object"}, {"aliases": ["none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:bot_action:none", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_action", "none"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/bot_action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Modify Bot protection behavior for a matching request. The modification could be to entirely skip Bot processing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_action

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- bot_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot protection behavior for a matching request. The modification could be to entirely skip
Bot processing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_skip_processing",
    "none")}
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
  "x-ves-oneof-field-action_type": "[\"bot_skip_processing\",\"none\"]"
}
```

Terraform syntax:

```terraform
bot_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bot_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/bot_skip_processing/): complete subsection reference.

- [none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/none/): complete subsection reference.

## Next pages

- [bot_action.bot_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/bot_skip_processing/)
- [bot_action.none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/bot_action/none/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
