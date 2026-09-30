---
page_title: "use_tls.use_mtls.tls_certificates.custom_hash_algorithms"
subcategory: "Load Balancing"
description: "use_tls.use_mtls.tls_certificates.custom_hash_algorithms for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 2420, "body_sha256": "sha256:4b6420fe34d06b0c8820025a737931efc47682337a89fe37e90db65fbd5513a3", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "path": "docs/guides/data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates--custom_hash_algorithms.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "use_mtls", "tls_certificates", "custom_hash_algorithms"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/use_mtls/tls_certificates/custom_hash_algorithms/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.use_mtls.tls_certificates.custom_hash_algorithms for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# use_tls.use_mtls.tls_certificates.custom_hash_algorithms

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [use_tls](data-sources--origin_pool--properties--use_tls.md)
- [use_tls.use_mtls](data-sources--origin_pool--properties--use_tls--use_mtls.md)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md)
- use_tls.use_mtls.tls_certificates.custom_hash_algorithms

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

<a id="schema-use_tls--use_mtls--tls_certificates--custom_hash_algorithms--hash_algorithms"></a>

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

- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
