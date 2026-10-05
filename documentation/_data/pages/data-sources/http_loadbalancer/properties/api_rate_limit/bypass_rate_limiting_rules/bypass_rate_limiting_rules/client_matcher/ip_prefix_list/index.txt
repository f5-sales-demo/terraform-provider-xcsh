---
page_title: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list"
subcategory: "Load Balancing"
description: "List of IP Prefix strings to match against."
xcsh_docs: {"aliases": ["api rate limit bypass rate limiting rules bypass rate limiting rules client matcher ip prefix list"], "body_bytes": 3884, "body_sha256": "sha256:6b4886a8a9b84207f2a2d467a80565acca2595c51b9989a437086fc508bc2022", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:ip_prefix_list", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher", "path": "documentation/data-sources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/ip_prefix_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3102132222212321-1112001130122301-3120001322031323-0201001213103202-0013311222212313-1233032022030211-3221110233031122-1012212110002223", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "client_matcher", "ip_prefix_list"], "schema_version": 1, "sections": [{"aliases": ["api rate limit bypass rate limiting rules bypass rate limiting rules client matcher ip prefix list invert match"], "anchor": "schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher--ip_prefix_list--invert_match", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:ip_prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "client_matcher", "ip_prefix_list", "invert_match"], "syntax": "attribute", "type": "bool"}, {"aliases": ["api rate limit bypass rate limiting rules bypass rate limiting rules client matcher ip prefix list ip prefixes"], "anchor": "schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher--ip_prefix_list--ip_prefixes", "description": "List of IPv4 prefix strings.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:ip_prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "client_matcher", "ip_prefix_list", "ip_prefixes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/ip_prefix_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of IP Prefix strings to match against.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="section"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher--ip_prefix_list--invert_match"></a>

### invert_match property

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="schema-api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules--client_matcher--ip_prefix_list--ip_prefixes"></a>

### ip_prefixes property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
