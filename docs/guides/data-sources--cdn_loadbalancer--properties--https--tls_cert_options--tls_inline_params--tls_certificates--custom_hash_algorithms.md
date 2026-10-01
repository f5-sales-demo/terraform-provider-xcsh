---
page_title: "https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms"
subcategory: "Load Balancing"
description: "https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2879, "body_sha256": "sha256:ac3c72ff4e0a977ecc43e0c790d18ce8957bebd962b6fb33235e944ac233cad7", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params:tls_certificates:custom_hash_algorithms", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params:tls_certificates:custom_hash_algorithms", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:https:tls_cert_options:tls_inline_params:tls_certificates", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_inline_params--tls_certificates--custom_hash_algorithms.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https", "tls_cert_options", "tls_inline_params", "tls_certificates", "custom_hash_algorithms"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/https/tls_cert_options/tls_inline_params/tls_certificates/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [https](data-sources--cdn_loadbalancer--properties--https.md)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--properties--https--tls_cert_options.md)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_inline_params.md)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_inline_params--tls_certificates.md)
- https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms

<a id="section"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="schema-https--tls_cert_options--tls_inline_params--tls_certificates--custom_hash_algorithms--hash_algorithms"></a>

### hash_algorithms property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--properties--https--tls_cert_options--tls_inline_params--tls_certificates.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
