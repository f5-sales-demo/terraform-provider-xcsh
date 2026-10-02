---
page_title: "cache_rules"
subcategory: ""
description: "This defines a CDN Cache Rule."
xcsh_docs: {"aliases": ["cache rules"], "body_bytes": 3469, "body_sha256": "sha256:367305d6abe9fd2612b4373cf2888201f2a5a8004c260900182729098353e97e", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:cache_bypass", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "parent_id": "xcsh-docs:resources:cdn_cache_rule:reference", "path": "documentation/resources/cdn_cache_rule/properties/cache_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2100032033322312-2310230223310322-3031201111321100-0121200020133021-2121321023123233-1100132303013031-3310112213122032-1032301311001032", "registry_path": "docs/guides/resources--cdn_cache_rule--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules:ConflictingObjectAttributes:cache_bypass,eligible_for_cache", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:cache_bypass", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules:ConflictingObjectAttributes:cache_bypass,eligible_for_cache", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "type": "conflicts"}, {"anchor": "schema-cache_rules--rule_name", "enforcement": "provider-schema", "group": "cache_rules:RequiredObjectAttributes:rule_expression_list,rule_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules:RequiredObjectAttributes:rule_expression_list,rule_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules"], "schema_version": 1, "sections": [{"aliases": ["cache bypass"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:cache_bypass", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "cache_bypass"], "syntax": "attribute", "type": "object"}, {"aliases": ["eligible for cache"], "anchor": "section", "description": "List of OPTIONS for Cache Action.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules.eligible_for_cache:ConflictingObjectAttributes:scheme_proxy_host_request_uri,scheme_proxy_host_uri", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules.eligible_for_cache:ConflictingObjectAttributes:scheme_proxy_host_request_uri,scheme_proxy_host_uri", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "type": "conflicts"}], "schema_path": ["cache_rules", "eligible_for_cache"], "syntax": "block", "type": "object"}, {"aliases": ["rule expression list"], "anchor": "section", "description": "Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs..", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cache_rules--rule_expression_list--expression_name", "enforcement": "provider-schema", "group": "cache_rules.rule_expression_list:RequiredListObjectAttributes:cache_rule_expression,expression_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cache_rules.rule_expression_list:RequiredListObjectAttributes:cache_rule_expression,expression_name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "type": "requires"}], "schema_path": ["cache_rules", "rule_expression_list"], "syntax": "block", "type": "object"}, {"aliases": ["rule name"], "anchor": "schema-cache_rules--rule_name", "description": "Name of the Cache Rule.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines a CDN Cache Rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- cache_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Cache Rule. This defines a CDN Cache Rule.

Upstream description:

This defines a CDN Cache Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rule_expression_list",
    "rule_name"),
  validators.ConflictingObjectAttributes("cache_bypass",
    "eligible_for_cache")}
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
  "x-ves-oneof-field-cache_actions": "[\"cache_bypass\",\"eligible_for_cache\"]"
}
```

Terraform syntax:

```terraform
cache_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cache_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/cache_bypass/): complete subsection reference.

- [eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/): complete subsection reference.

- [rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/): complete subsection reference.

<a id="schema-cache_rules--rule_name"></a>

### rule_name property

Type: `"string"`. Optional.

Rule Name. Name of the Cache Rule.

Upstream description:

Name of the Cache Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

## Next pages

- [cache_rules.cache_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/cache_bypass/)
- [cache_rules.eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/)
- [cache_rules.rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
