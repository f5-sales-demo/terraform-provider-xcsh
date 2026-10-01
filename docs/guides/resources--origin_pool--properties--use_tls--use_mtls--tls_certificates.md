---
page_title: "use_tls.use_mtls.tls_certificates"
subcategory: "Load Balancing"
description: "use_tls.use_mtls.tls_certificates for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 5144, "body_sha256": "sha256:8c24fdef698f7ffc6edf79a0776e83289f706f34242587efd3893bea36935c71", "canonical_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "child_ids": ["xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:private_key", "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:use_system_defaults"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls", "path": "docs/guides/resources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "use_mtls", "tls_certificates"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.use_mtls.tls_certificates for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.use_mtls.tls_certificates

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [use_tls](resources--origin_pool--properties--use_tls.md)
- [use_tls.use_mtls](resources--origin_pool--properties--use_tls--use_mtls.md)
- use_tls.use_mtls.tls_certificates

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-use_tls--use_mtls--tls_certificates--certificate_url"></a>

### certificate_url property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [custom_hash_algorithms](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--custom_hash_algorithms.md): complete subsection reference.

<a id="schema-use_tls--use_mtls--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--disable_ocsp_stapling.md): complete subsection reference.

- [private_key](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key.md): complete subsection reference.

- [use_system_defaults](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--use_system_defaults.md): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--custom_hash_algorithms.md)
- [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--disable_ocsp_stapling.md)
- [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key.md)
- [use_tls.use_mtls.tls_certificates.use_system_defaults](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--use_system_defaults.md)
- [use_tls.use_mtls](resources--origin_pool--properties--use_tls--use_mtls.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
