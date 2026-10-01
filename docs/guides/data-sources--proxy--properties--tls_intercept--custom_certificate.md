---
page_title: "tls_intercept.custom_certificate"
subcategory: ""
description: "tls_intercept.custom_certificate for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 3719, "body_sha256": "sha256:d602b3d268991c112e017c03c713614801d4f8edf8dbe3d49f061a18abb5c045", "canonical_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate", "child_ids": ["xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:custom_hash_algorithms", "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:disable_ocsp_stapling", "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:private_key", "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate:use_system_defaults"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:custom_certificate", "parent_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept", "path": "docs/guides/data-sources--proxy--properties--tls_intercept--custom_certificate.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "custom_certificate"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/tls_intercept/custom_certificate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.custom_certificate for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.custom_certificate

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- [Property reference](data-sources--proxy--reference.md)
- [tls_intercept](data-sources--proxy--properties--tls_intercept.md)
- tls_intercept.custom_certificate

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for custom certificate.

Upstream description:

Handle to fetch certificate and key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

## Direct properties

<a id="schema-tls_intercept--custom_certificate--certificate_url"></a>

### certificate_url property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--proxy--properties--tls_intercept--custom_certificate--custom_hash_algorithms.md): complete subsection reference.

<a id="schema-tls_intercept--custom_certificate--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--proxy--properties--tls_intercept--custom_certificate--disable_ocsp_stapling.md): complete subsection reference.

- [private_key](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key.md): complete subsection reference.

- [use_system_defaults](data-sources--proxy--properties--tls_intercept--custom_certificate--use_system_defaults.md): complete subsection reference.

## Next pages

- [tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--proxy--properties--tls_intercept--custom_certificate--custom_hash_algorithms.md)
- [tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--proxy--properties--tls_intercept--custom_certificate--disable_ocsp_stapling.md)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key.md)
- [tls_intercept.custom_certificate.use_system_defaults](data-sources--proxy--properties--tls_intercept--custom_certificate--use_system_defaults.md)
- [tls_intercept](data-sources--proxy--properties--tls_intercept.md)
- [xcsh_proxy](../data-sources/proxy.md)
