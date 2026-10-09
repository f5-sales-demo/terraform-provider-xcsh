---
page_title: "client_name_matcher"
subcategory: ""
description: "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions."
xcsh_docs: {"aliases": ["client name matcher", "succeeded", "success", "successful"], "body_bytes": 3859, "body_sha256": "sha256:13f39c6ebc7a9908d326bf3a27972f65f8a2740a17487f8475f18bc0314aef21", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:client_name_matcher", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/client_name_matcher/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0332123330301021-1100223121211022-3121010103210131-3033010222100300-1123002321301311-1112213102202113-1311100312010010-0101133331030013", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["client_name_matcher"], "schema_version": 1, "sections": [{"aliases": ["client name matcher exact values"], "anchor": "schema-client_name_matcher--exact_values", "description": "A list of exact values to match the input against.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:client_name_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_name_matcher", "exact_values"], "syntax": "attribute", "type": "list"}, {"aliases": ["client name matcher regex values"], "anchor": "schema-client_name_matcher--regex_values", "description": "A list of regular expressions to match the input against.", "document_id": "xcsh-docs:resources:service_policy_rule:properties:client_name_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_name_matcher", "regex_values"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/client_name_matcher/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# client_name_matcher

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- client_name_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
client_name_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-client_name_matcher--exact_values"></a>

### exact_values property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-client_name_matcher--regex_values"></a>

### regex_values property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
