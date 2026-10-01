---
page_title: "tls_parameters.common_params.tls_certificates"
subcategory: ""
description: "tls_parameters.common_params.tls_certificates for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 4719, "body_sha256": "sha256:99085cc11edc431916d9ae355143084c326dd4f99ce76d9b9f63d0bb6b6b5d0e", "canonical_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "child_ids": ["xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:custom_hash_algorithms", "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:disable_ocsp_stapling", "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:private_key", "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:use_system_defaults"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params", "path": "docs/guides/resources--virtual_host--properties--tls_parameters--common_params--tls_certificates.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.tls_certificates for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [tls_parameters](resources--virtual_host--properties--tls_parameters.md)
- [tls_parameters.common_params](resources--virtual_host--properties--tls_parameters--common_params.md)
- tls_parameters.common_params.tls_certificates

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

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
tls_certificates {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-tls_parameters--common_params--tls_certificates--certificate_url"></a>

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

- [custom_hash_algorithms](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--custom_hash_algorithms.md): complete subsection reference.

<a id="schema-tls_parameters--common_params--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--disable_ocsp_stapling.md): complete subsection reference.

- [private_key](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--private_key.md): complete subsection reference.

- [use_system_defaults](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--use_system_defaults.md): complete subsection reference.

## Next pages

- [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--custom_hash_algorithms.md)
- [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--disable_ocsp_stapling.md)
- [tls_parameters.common_params.tls_certificates.private_key](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--private_key.md)
- [tls_parameters.common_params.tls_certificates.use_system_defaults](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--use_system_defaults.md)
- [tls_parameters.common_params](resources--virtual_host--properties--tls_parameters--common_params.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
