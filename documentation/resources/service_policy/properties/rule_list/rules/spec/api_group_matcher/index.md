---
page_title: "rule_list.rules.spec.api_group_matcher"
subcategory: "Security"
description: "A matcher specifies a list of values for matching an input string. The match is considered successful if the input value is present in the list. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["rule list rules spec api group matcher", "succeeded", "success", "successful"], "body_bytes": 3253, "body_sha256": "sha256:794215f9b2c40c039841661433cd89d7ec097de9b9f5b69d33dc9f76f0725148", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1200123013332212-3203303010013021-3203002133013210-1003313101132211-0113302233322111-2222033012113113-0111001320232211-3203213221111311", "registry_path": "docs/guides/resources--service_policy--reference--group-001.md", "relationships": [{"anchor": "schema-rule_list--rules--spec--api_group_matcher--match", "enforcement": "provider-schema", "group": "rule_list.rules.spec.api_group_matcher:RequiredObjectAttributes:match", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "api_group_matcher"], "schema_version": 1, "sections": [{"aliases": ["rule list rules spec api group matcher invert matcher"], "anchor": "schema-rule_list--rules--spec--api_group_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "api_group_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["rule list rules spec api group matcher match"], "anchor": "schema-rule_list--rules--spec--api_group_matcher--match", "description": "A list of exact values to match the input against.", "document_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:api_group_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "spec", "api_group_matcher", "match"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/api_group_matcher/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "A matcher specifies a list of values for matching an input string. The match is considered successful if the input value is present in the list. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["service_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.api_group_matcher

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- rule_list.rules.spec.api_group_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies a list of values for matching an input string. The match is considered
successful if the input value is present in the list. The result of the match is inverted if
invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("match")}
```

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
api_group_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--spec--api_group_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert String Matcher. Invert the match result.

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

<a id="schema-rule_list--rules--spec--api_group_matcher--match"></a>

### match property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "63",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
