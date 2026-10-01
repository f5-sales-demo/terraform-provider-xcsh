---
page_title: "other_settings.header_options"
subcategory: "Load Balancing"
description: "other_settings.header_options for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4216, "body_sha256": "sha256:83232a95f86db55323cee7c45f66be79d9403c18fd5f752bc03c4737aedd061f", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:request_headers_to_add", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options:response_headers_to_add"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--other_settings--header_options.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["other_settings", "header_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/header_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "other_settings.header_options for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings.header_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [other_settings](data-sources--cdn_loadbalancer--properties--other_settings.md)
- other_settings.header_options

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS related to request/response headers.

Upstream description:

This defines various OPTIONS related to request/response headers.

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

- [request_headers_to_add](data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add.md): complete subsection reference.

<a id="schema-other_settings--header_options--request_headers_to_remove"></a>

### request_headers_to_remove property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--cdn_loadbalancer--properties--other_settings--header_options--response_headers_to_add.md): complete subsection reference.

<a id="schema-other_settings--header_options--response_headers_to_remove"></a>

### response_headers_to_remove property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [other_settings.header_options.request_headers_to_add](data-sources--cdn_loadbalancer--properties--other_settings--header_options--request_headers_to_add.md)
- [other_settings.header_options.response_headers_to_add](data-sources--cdn_loadbalancer--properties--other_settings--header_options--response_headers_to_add.md)
- [other_settings](data-sources--cdn_loadbalancer--properties--other_settings.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
