---
page_title: "caching_policy.default_cache_action"
subcategory: "Load Balancing"
description: "This defines a Default Cache Action."
xcsh_docs: {"aliases": ["caching policy default cache action"], "body_bytes": 3736, "body_sha256": "sha256:826f9ed21a6cfec1ae3762ad50956c984e42ae3dad0dbc35bec8d700ff1996eb", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:default_cache_action:cache_disabled"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:default_cache_action", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy", "path": "documentation/data-sources/http_loadbalancer/properties/caching_policy/default_cache_action/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2030011133231130-2202111323312110-1300223120331333-0200132202210011-1330111202102310-2332303231202013-1322220332013231-3031001313200331", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-014.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["caching_policy", "default_cache_action"], "schema_version": 1, "sections": [{"aliases": ["caching policy default cache action cache disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:default_cache_action:cache_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["caching_policy", "default_cache_action", "cache_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["caching policy default cache action cache ttl default"], "anchor": "schema-caching_policy--default_cache_action--cache_ttl_default", "description": "Exclusive with Use Cache TTL Provided by Origin, and set a contigency TTL value in case one is not provided.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:default_cache_action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["caching_policy", "default_cache_action", "cache_ttl_default"], "syntax": "attribute", "type": "string"}, {"aliases": ["caching policy default cache action cache ttl override"], "anchor": "schema-caching_policy--default_cache_action--cache_ttl_override", "description": "Exclusive with Always override the Cache TTL provided by Origin.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:caching_policy:default_cache_action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["caching_policy", "default_cache_action", "cache_ttl_override"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/caching_policy/default_cache_action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a Default Cache Action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# caching_policy.default_cache_action

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/caching_policy/)
- caching_policy.default_cache_action

<a id="section"></a>

Type: `"single"`. Computed.

Default Cache Behaviour. This defines a Default Cache Action.

Upstream description:

This defines a Default Cache Action.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

## Direct properties

- [cache_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/caching_policy/default_cache_action/cache_disabled/): complete subsection reference.

<a id="schema-caching_policy--default_cache_action--cache_ttl_default"></a>

### cache_ttl_default property

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="schema-caching_policy--default_cache_action--cache_ttl_override"></a>

### cache_ttl_override property

Type: `"string"`. Computed.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

## Next pages

- [caching_policy.default_cache_action.cache_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/caching_policy/default_cache_action/cache_disabled/)
- [caching_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/caching_policy/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
