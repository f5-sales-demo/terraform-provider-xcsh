---
page_title: "cache_rules.eligible_for_cache.scheme_proxy_host_uri"
subcategory: ""
description: "cache_rules.eligible_for_cache.scheme_proxy_host_uri for xcsh_cdn_cache_rule."
xcsh_docs: {"aliases": [], "body_bytes": 3438, "body_sha256": "sha256:72abe5eccda7a7b079ae06c6624c536ea725f43ab0103d940047c6f769a10ae8", "canonical_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_cache_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache:scheme_proxy_host_uri", "parent_id": "xcsh-docs:resources:cdn_cache_rule:properties:cache_rules:eligible_for_cache", "path": "docs/guides/resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache--scheme_proxy_host_uri.md", "provider_name": "cdn_cache_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cache_rules", "eligible_for_cache", "scheme_proxy_host_uri"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_cache_rule/properties/cache_rules/eligible_for_cache/scheme_proxy_host_uri/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cache_rules.eligible_for_cache.scheme_proxy_host_uri for xcsh_cdn_cache_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_cache_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cache_rules.eligible_for_cache.scheme_proxy_host_uri

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
- [Property reference](resources--cdn_cache_rule--reference.md)
- [cache_rules](resources--cdn_cache_rule--properties--cache_rules.md)
- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache.md)
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

- [cache_rules.eligible_for_cache](resources--cdn_cache_rule--properties--cache_rules--eligible_for_cache.md)
- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md)
