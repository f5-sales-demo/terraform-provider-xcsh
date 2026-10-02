---
page_title: "rule_list.rules.spec.mum_action"
subcategory: "Security"
description: "Modify behavior for a matching request. The modification could be to entirely skip processing."
xcsh_docs: {"aliases": ["rule list rules spec mum action"], "body_bytes": 2510, "body_sha256": "sha256:61585d63bf2abba885c98ecc773d91d2a689d8eb5dc7c2e8f1d7f9d11fecaa44", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:default", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:skip_processing"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/mum_action/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1031203103211101-1320130203110223-1203222000022323-0200122233332223-3330121000123133-2233010210331321-1120111112011131-2031122333013131", "registry_path": "docs/guides/resources--service_policy--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.mum_action:ConflictingObjectAttributes:default,skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rule_list.rules.spec.mum_action:ConflictingObjectAttributes:default,skip_processing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:skip_processing", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "mum_action"], "schema_version": 1, "sections": [{"aliases": ["default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:default", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "mum_action", "default"], "syntax": "attribute", "type": "object"}, {"aliases": ["skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:skip_processing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "mum_action", "skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/mum_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Modify behavior for a matching request. The modification could be to entirely skip processing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.mum_action

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- rule_list.rules.spec.mum_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default",
    "skip_processing")}
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
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

Terraform syntax:

```terraform
mum_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/default/): complete subsection reference.

- [skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/skip_processing/): complete subsection reference.

## Next pages

- [rule_list.rules.spec.mum_action.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/default/)
- [rule_list.rules.spec.mum_action.skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/mum_action/skip_processing/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
