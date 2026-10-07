---
page_title: "cache_rules"
subcategory: ""
description: "This defines a CDN Cache Rule."
xcsh_docs: {"aliases": ["cache rules"], "body_bytes": 2151, "body_sha256": "sha256:5422628be474883cea75ad07b6e5cf269c6d987bfb490dfc62c0d85b17a4bd51", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:cache_bypass", "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:reference", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3232300021003032-2311103103303303-2212230332022111-3002010311233233-3301102211203232-1013221222030211-1013001102220021-3103203212031313", "registry_path": "docs/guides/data-sources--cdn_cache_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules"], "schema_version": 1, "sections": [{"aliases": ["cache rules cache bypass"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:cache_bypass", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "cache_bypass"], "syntax": "attribute", "type": "object"}, {"aliases": ["cache rules eligible for cache"], "anchor": "section", "description": "List of OPTIONS for Cache Action.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache"], "syntax": "attribute", "type": "object"}, {"aliases": ["cache rules rule expression list"], "anchor": "section", "description": "Expressions are evaluated in the order in which they are specified. The evaluation stops when the first rule match occurs..", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:rule_expression_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cache_rules", "rule_expression_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["cache rules rule name"], "anchor": "schema-cache_rules--rule_name", "description": "Name of the Cache Rule.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "rule_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "This defines a CDN Cache Rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/)
- cache_rules

<a id="section"></a>

Type: `"single"`. Computed.

Cache Rule. This defines a CDN Cache Rule.

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

## Direct properties

- [cache_bypass](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/cache_bypass/): complete subsection reference.

- [eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/): complete subsection reference.

- [rule_expression_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/rule_expression_list/): complete subsection reference.

<a id="schema-cache_rules--rule_name"></a>

### rule_name property

Type: `"string"`. Computed.

Rule Name. Name of the Cache Rule.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
