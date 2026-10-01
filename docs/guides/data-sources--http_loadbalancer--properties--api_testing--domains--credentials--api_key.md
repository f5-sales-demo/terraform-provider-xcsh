---
page_title: "api_testing.domains.credentials.api_key"
subcategory: "Load Balancing"
description: "api_testing.domains.credentials.api_key for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2299, "body_sha256": "sha256:d77e1b5d0af0f6ccf0609719a009e4d806737d2c1a979a626103b275684c7e97", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key:value"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials:api_key", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains:credentials", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_testing", "domains", "credentials", "api_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_testing/domains/credentials/api_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_testing.domains.credentials.api_key for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing.domains.credentials.api_key

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_testing](data-sources--http_loadbalancer--properties--api_testing.md)
- [api_testing.domains](data-sources--http_loadbalancer--properties--api_testing--domains.md)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--properties--api_testing--domains--credentials.md)
- api_testing.domains.credentials.api_key

<a id="section"></a>

Type: `"single"`. Computed.

API Key

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

<a id="schema-api_testing--domains--credentials--api_key--key"></a>

### key property

Type: `"string"`. Computed.

Key. Cryptographic key material

Upstream description:

Cryptographic key material

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key--value.md): complete subsection reference.

## Next pages

- [api_testing.domains.credentials.api_key.value](data-sources--http_loadbalancer--properties--api_testing--domains--credentials--api_key--value.md)
- [api_testing.domains.credentials](data-sources--http_loadbalancer--properties--api_testing--domains--credentials.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
