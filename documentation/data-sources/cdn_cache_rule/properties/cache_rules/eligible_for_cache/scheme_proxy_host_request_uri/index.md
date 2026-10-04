---
page_title: "cache_rules.eligible_for_cache.scheme_proxy_host_request_uri"
subcategory: ""
description: "Cache TTL Enable Values."
xcsh_docs: {"aliases": ["cache rules eligible for cache scheme proxy host request uri"], "body_bytes": 3524, "body_sha256": "sha256:524fed0b191ef7a2e7fc7b5f787c358e7fbc2c4976cdc7711af143436fb08732", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "parent_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "path": "documentation/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1133121110333221-2102301131222222-1101333101231332-1000333213102100-0000231312133030-1300101132033322-1010301231001112-3331032133211302", "registry_path": "docs/guides/data-sources--cdn_cache_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_request_uri"], "schema_version": 1, "sections": [{"aliases": ["cache rules eligible for cache scheme proxy host request uri cache override"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_override", "description": "Honour Cache Override.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_request_uri", "cache_override"], "syntax": "attribute", "type": "bool"}, {"aliases": ["cache rules eligible for cache scheme proxy host request uri cache ttl"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_ttl", "description": "Cache TTL value is used to cache the resource/content for the specified amount of time Format: , where s - seconds, m - minutes, h - hours, d - days.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_request_uri", "cache_ttl"], "syntax": "attribute", "type": "string"}, {"aliases": ["cache rules eligible for cache scheme proxy host request uri ignore response cookie"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--ignore_response_cookie", "description": "By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present.", "document_id": "xcsh-docs:data-sources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_request_uri", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_request_uri", "ignore_response_cookie"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_request_uri/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Cache TTL Enable Values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.eligible_for_cache.scheme_proxy_host_request_uri

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/)
- [cache_rules.eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/)
- cache_rules.eligible_for_cache.scheme_proxy_host_request_uri

<a id="section"></a>

Type: `"single"`. Computed.

Cache TTL Enable Props. Cache TTL Enable Values.

Upstream description:

Cache TTL Enable Values.

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

<a id="schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_override"></a>

### cache_override property

Type: `"bool"`. Computed.

Cache Override. Honour Cache Override.

Upstream description:

Honour Cache Override.

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

<a id="schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--cache_ttl"></a>

### cache_ttl property

Type: `"string"`. Computed.

Cache TTL value is used to cache the resource/content for the specified amount of time Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days.

Upstream description:

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

<a id="schema-cache_rules--eligible_for_cache--scheme_proxy_host_request_uri--ignore_response_cookie"></a>

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

## Next pages

- [cache_rules.eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_cache_rule/)
