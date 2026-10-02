---
page_title: "api_group_matcher"
subcategory: ""
description: "A matcher specifies a list of values for matching an input string. The match is considered successful if the input value is present in the list. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["api group matcher", "login success", "succeeded", "success", "successful"], "body_bytes": 2949, "body_sha256": "sha256:f33fee17492a9ffae788c409e10ef3821e17eb0819c83234c04da8430295a3c2", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:api_group_matcher", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "documentation/data-sources/service_policy_rule/properties/api_group_matcher/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2323220103100030-2300221222133301-0312321332223112-2330231103123033-2032200011132020-3100000303330210-1222301113122130-2312000321110100", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_group_matcher"], "schema_version": 1, "sections": [{"aliases": ["invert matcher"], "anchor": "schema-api_group_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:api_group_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_group_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["match"], "anchor": "schema-api_group_matcher--match", "description": "A list of exact values to match the input against.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:api_group_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_group_matcher", "match"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/api_group_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A matcher specifies a list of values for matching an input string. The match is considered successful if the input value is present in the list. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_group_matcher

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- api_group_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Matcher specifies a list of values for matching an input string. The match is considered successful
if the input value is present in the list. The result of the match is inverted if invert\_matcher is
true.

Upstream description:

A matcher specifies a list of values for matching an input string. The match is considered
successful if the input value is present in the list. The result of the match is inverted if
invert\_matcher is true.

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

## Direct properties

<a id="schema-api_group_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

Invert String Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="schema-api_group_matcher--match"></a>

### match property

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
