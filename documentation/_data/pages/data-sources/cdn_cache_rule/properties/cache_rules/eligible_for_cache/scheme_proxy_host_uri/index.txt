---
page_title: "cache_rules.eligible_for_cache.scheme_proxy_host_uri"
subcategory: ""
description: "Cache TTL Enable Values."
xcsh_docs: {"aliases": ["cache rules eligible for cache scheme proxy host uri"], "body_bytes": 2910, "body_sha256": "sha256:636df7db8e72c17b60786dae3d5e3f452bee17b5489f3c4072a012be5b55731d", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2322323230330032-1013101130203322-2002311100202112-3230333221330001-3011103200320000-1332113333221302-2013001123000010-2210131133101313", "registry_path": "docs/guides/data-sources--cdn_cache_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri"], "schema_version": 1, "sections": [{"aliases": ["cache rules eligible for cache scheme proxy host uri cache override"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_override", "description": "Honour Cache Override.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri", "cache_override"], "syntax": "attribute", "type": "bool"}, {"aliases": ["cache rules eligible for cache scheme proxy host uri cache ttl"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl", "description": "Cache TTL value is used to cache the resource/content for the specified amount of time Format: , where s - seconds, m - minutes, h - hours, d - days.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri", "cache_ttl"], "syntax": "attribute", "type": "string"}, {"aliases": ["cache rules eligible for cache scheme proxy host uri ignore response cookie"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--ignore_response_cookie", "description": "By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri", "ignore_response_cookie"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Cache TTL Enable Values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.eligible_for_cache.scheme_proxy_host_uri

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- [cache_rules.eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/)
- cache_rules.eligible_for_cache.scheme_proxy_host_uri

<a id="section"></a>

Type: `"single"`. Computed.

Cache TTL Enable Props. Cache TTL Enable Values.

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

<a id="schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_override"></a>

### cache_override property

Type: `"bool"`. Computed.

Cache Override. Honour Cache Override.

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

<a id="schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl"></a>

### cache_ttl property

Type: `"string"`. Computed.

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--ignore_response_cookie"></a>

### ignore_response_cookie property

Type: `"bool"`. Computed.

By default, response will not be cached if set-cookie header is present. This option will override
the behavior and cache response even with set-cookie header present.

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
