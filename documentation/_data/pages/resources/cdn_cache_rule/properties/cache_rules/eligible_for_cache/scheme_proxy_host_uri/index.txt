---
page_title: "cache_rules.eligible_for_cache.scheme_proxy_host_uri"
subcategory: ""
description: "Cache TTL Enable Values."
xcsh_docs: {"aliases": ["cache rules eligible for cache scheme proxy host uri"], "body_bytes": 3744, "body_sha256": "sha256:24d7d2838eddaec6805cd991d5b6ba85fc056504f13f89eb1a9960e02e133805", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "path": "documentation/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/index.md", "product": "distributed-cloud", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3131020013020013-3111002330212120-3331323132023223-2111313013102020-2031102010222030-2121100131011321-3100111011102223-0101003333131212", "registry_path": "docs/guides/resources--cdn_cache_rule--reference--group-001.md", "relationships": [{"anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl", "enforcement": "provider-schema", "group": "cache_rules.eligible_for_cache.scheme_proxy_host_uri:RequiredObjectAttributes:cache_ttl", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri"], "schema_version": 1, "sections": [{"aliases": ["cache rules eligible for cache scheme proxy host uri cache override"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_override", "description": "Honour Cache Override.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri", "cache_override"], "syntax": "attribute", "type": "bool"}, {"aliases": ["cache rules eligible for cache scheme proxy host uri cache ttl"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl", "description": "Cache TTL value is used to cache the resource/content for the specified amount of time Format: , where s - seconds, m - minutes, h - hours, d - days.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri", "cache_ttl"], "syntax": "attribute", "type": "string"}, {"aliases": ["cache rules eligible for cache scheme proxy host uri ignore response cookie"], "anchor": "schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--ignore_response_cookie", "description": "By default, response will not be cached if set-cookie header is present. This option will override the behavior and cache response even with set-cookie header present.", "document_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri", "ignore_response_cookie"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Cache TTL Enable Values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.eligible_for_cache.scheme_proxy_host_uri

Breadcrumbs:

- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/)
- [cache_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/)
- [cache_rules.eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/)
- cache_rules.eligible_for_cache.scheme_proxy_host_uri

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Cache TTL Enable Props. Cache TTL Enable Values.

Upstream description:

Cache TTL Enable Values.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cache_ttl")}
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
scheme_proxy_host_uri {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_override"></a>

### cache_override property

Type: `"bool"`. Optional.

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

<a id="schema-cache_rules--eligible_for_cache--scheme_proxy_host_uri--cache_ttl"></a>

### cache_ttl property

Type: `"string"`. Optional.

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

Type: `"bool"`. Optional.

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

- [cache_rules.eligible_for_cache](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/)
- [xcsh_cdn_cache_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_cache_rule/)
