---
page_title: "tls_intercept.custom_certificate"
subcategory: ""
description: "tls_intercept.custom_certificate for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 4392, "body_sha256": "sha256:64318f0ceb28db7dd1fd6a8f35746603f11d989b949fb55d77a3e0e8a3a3d576", "canonical_id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "child_ids": ["xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:custom_hash_algorithms", "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:disable_ocsp_stapling", "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:private_key", "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate:use_system_defaults"], "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:tls_intercept:custom_certificate", "parent_id": "xcsh-docs:resources:proxy:properties:tls_intercept", "path": "docs/guides/resources--proxy--properties--tls_intercept--custom_certificate.md", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_intercept", "custom_certificate"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/tls_intercept/custom_certificate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_intercept.custom_certificate for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.custom_certificate

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md)
- [Property reference](resources--proxy--reference.md)
- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- tls_intercept.custom_certificate

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for custom certificate.

Upstream description:

Handle to fetch certificate and key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

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

Terraform syntax:

```terraform
custom_certificate {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-tls_intercept--custom_certificate--certificate_url"></a>

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

- [custom_hash_algorithms](resources--proxy--properties--tls_intercept--custom_certificate--custom_hash_algorithms.md): complete subsection reference.

<a id="schema-tls_intercept--custom_certificate--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--proxy--properties--tls_intercept--custom_certificate--disable_ocsp_stapling.md): complete subsection reference.

- [private_key](resources--proxy--properties--tls_intercept--custom_certificate--private_key.md): complete subsection reference.

- [use_system_defaults](resources--proxy--properties--tls_intercept--custom_certificate--use_system_defaults.md): complete subsection reference.

## Next pages

- [tls_intercept.custom_certificate.custom_hash_algorithms](resources--proxy--properties--tls_intercept--custom_certificate--custom_hash_algorithms.md)
- [tls_intercept.custom_certificate.disable_ocsp_stapling](resources--proxy--properties--tls_intercept--custom_certificate--disable_ocsp_stapling.md)
- [tls_intercept.custom_certificate.private_key](resources--proxy--properties--tls_intercept--custom_certificate--private_key.md)
- [tls_intercept.custom_certificate.use_system_defaults](resources--proxy--properties--tls_intercept--custom_certificate--use_system_defaults.md)
- [tls_intercept](resources--proxy--properties--tls_intercept.md)
- [xcsh_proxy](../resources/proxy.md)
