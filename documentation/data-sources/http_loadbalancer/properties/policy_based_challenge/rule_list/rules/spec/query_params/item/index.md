---
page_title: "policy_based_challenge.rule_list.rules.spec.query_params.item"
subcategory: "Load Balancing"
description: "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions."
xcsh_docs: {"aliases": ["login success", "policy based challenge rule list rules spec query params item", "succeeded", "success", "successful"], "body_bytes": 6709, "body_sha256": "sha256:b5550ac59f5d7df6dc4a59d52fb8e5fd2d965c3b8779139f95bb704ac9adfa75", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "path": "documentation/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/item/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1302012230310122-2012322122203331-1330220123231131-0311210132332223-0100310120103122-2311031213033211-2321133023211223-2103103230011023", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "item"], "schema_version": 1, "sections": [{"aliases": ["exact values"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--query_params--item--exact_values", "description": "A list of exact values to match the input against.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "item", "exact_values"], "syntax": "attribute", "type": "list"}, {"aliases": ["regex values"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--query_params--item--regex_values", "description": "A list of regular expressions to match the input against.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "item", "regex_values"], "syntax": "attribute", "type": "list"}, {"aliases": ["transformers"], "anchor": "schema-policy_based_challenge--rule_list--rules--spec--query_params--item--transformers", "description": "An ordered list of transformers (starting from index 0) to be applied to the path before matching.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params:item", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec", "query_params", "item", "transformers"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/item/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A matcher specifies multiple criteria for matching an input string. The match is considered successful if any of the criteria are satisfied. The set of supported match criteria includes a list of exact values and a list of regular expressions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules.spec.query_params.item

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/)
- [policy_based_challenge.rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- [policy_based_challenge.rule_list.rules.spec.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="section"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

## Direct properties

<a id="schema-policy_based_challenge--rule_list--rules--spec--query_params--item--exact_values"></a>

### exact_values property

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

<a id="schema-policy_based_challenge--rule_list--rules--spec--query_params--item--regex_values"></a>

### regex_values property

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-policy_based_challenge--rule_list--rules--spec--query_params--item--transformers"></a>

### transformers property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [policy_based_challenge.rule_list.rules.spec.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/query_params/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
