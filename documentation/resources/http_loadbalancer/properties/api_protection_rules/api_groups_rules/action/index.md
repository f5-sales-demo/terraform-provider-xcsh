---
page_title: "api_protection_rules.api_groups_rules.action"
subcategory: "Load Balancing"
description: "The action to take if the input request matches the rule."
xcsh_docs: {"aliases": ["api protection rules api groups rules action"], "body_bytes": 2486, "body_sha256": "sha256:41642ab94f839d8704b2cf8354ea01adb1c5612b03e470ef1a552b0dcc2512b8", "capabilities": ["load-balancing", "security.api-protection"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:deny"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules", "path": "documentation/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.action:ConflictingObjectAttributes:allow,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_protection_rules.api_groups_rules.action:ConflictingObjectAttributes:allow,deny", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:deny", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_protection_rules", "api_groups_rules", "action"], "schema_version": 1, "sections": [{"aliases": ["allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:allow", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "action", "allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["deny"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules:action:deny", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_protection_rules", "api_groups_rules", "action", "deny"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "The action to take if the input request matches the rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules.api_groups_rules.action

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_protection_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- api_protection_rules.api_groups_rules.action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

The action to take if the input request matches the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "deny")}
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
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/allow/): complete subsection reference.

- [deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/deny/): complete subsection reference.

## Next pages

- [api_protection_rules.api_groups_rules.action.allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/allow/)
- [api_protection_rules.api_groups_rules.action.deny](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/action/deny/)
- [api_protection_rules.api_groups_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_protection_rules/api_groups_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
