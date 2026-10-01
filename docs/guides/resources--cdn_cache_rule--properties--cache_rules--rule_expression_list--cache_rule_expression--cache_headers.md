---
page_title: "cache_rules.rule_expression_list.cache_rule_expression.cache_headers"
subcategory: ""
description: "cache_rules.rule_expression_list.cache_rule_expression.cache_headers for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 3756, "body_sha256": "sha256:7102245081d4581d9bda4a0ec1d70eb31c53a768932843f9d75ebe52f3a09c10", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cache_headers", "child_ids": ["xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cache_headers:operator"], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression:cache_headers", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:rule_expression_list:cache_rule_expression", "path": "docs/guides/resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cache_rules", "rule_expression_list", "cache_rule_expression", "cache_headers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/rule_expression_list/cache_rule_expression/cache_headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.rule_expression_list.cache_rule_expression.cache_headers for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.rule_expression_list.cache_rule_expression.cache_headers

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
- [Property reference](resources--cdn_cache_rule--reference.md)
- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- [cache_rules.rule_expression_list](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list.md)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md)
- cache_rules.rule_expression_list.cache_rule_expression.cache_headers

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Configure cache rule headers to match the criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
cache_headers {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cache_rules--rule_expression_list--cache_rule_expression--cache_headers--name"></a>

### name property

Type: `"string"`. Optional.

\[Enum: PROXY\_HOST|REFERER|SCHEME|USER\_AGENT\] - PROXY\_HOST: Proxy Host Name of the proxied
server - REFERER: Referer This is the address of the previous web page from which a link to the
currently requested page was followed - SCHEME: Scheme The HTTP scheme used: HTTP or HTTPS -
USER\_AGENT: User Agent The user agent string of the user agent. Possible values are
\`PROXY\_HOST\`, \`REFERER\`, \`SCHEME\`, \`USER\_AGENT\`. Defaults to \`PROXY\_HOST\`.

Upstream description:

&#8203;- PROXY\_HOST: Proxy Host

Name of the proxied server &#8203;- REFERER: Referer

This is the address of the previous web page from which a link to the currently requested page was
followed &#8203;- SCHEME: Scheme

The HTTP scheme used: HTTP or HTTPS &#8203;- USER\_AGENT: User Agent

The user agent string of the user agent.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROXY_HOST",
    "REFERER",
    "SCHEME",
    "USER_AGENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROXY_HOST",
  "enum": [
    "PROXY_HOST",
    "REFERER",
    "SCHEME",
    "USER_AGENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [operator](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md): complete subsection reference.

## Next pages

- [cache_rules.rule_expression_list.cache_rule_expression.cache_headers.operator](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression--cache_headers--operator.md)
- [cache_rules.rule_expression_list.cache_rule_expression](resources--cdn_cache_rule--properties--cache_rules--rule_expression_list--cache_rule_expression.md)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
