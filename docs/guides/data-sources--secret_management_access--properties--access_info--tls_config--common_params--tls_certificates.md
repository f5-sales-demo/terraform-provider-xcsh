---
page_title: "access_info.tls_config.common_params.tls_certificates"
subcategory: ""
description: "access_info.tls_config.common_params.tls_certificates for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 4502, "body_sha256": "sha256:868336c14e550aa2a59c8c22e33946d813af3f73ea2cb99999e4eb9fdca56013", "canonical_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "child_ids": ["xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:custom_hash_algorithms", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:disable_ocsp_stapling", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:private_key", "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates:use_system_defaults"], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params:tls_certificates", "parent_id": "xcsh-docs:data-sources:secret_management_access:properties:access_info:tls_config:common_params", "path": "docs/guides/data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "tls_config", "common_params", "tls_certificates"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/properties/access_info/tls_config/common_params/tls_certificates/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.tls_config.common_params.tls_certificates for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.common_params.tls_certificates

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
- [Property reference](data-sources--secret_management_access--reference.md)
- [access_info](data-sources--secret_management_access--properties--access_info.md)
- [access_info.tls_config](data-sources--secret_management_access--properties--access_info--tls_config.md)
- [access_info.tls_config.common_params](data-sources--secret_management_access--properties--access_info--tls_config--common_params.md)
- access_info.tls_config.common_params.tls_certificates

<a id="section"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

Upstream description:

Set of TLS certificates.

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

<a id="schema-access_info--tls_config--common_params--tls_certificates--certificate_url"></a>

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

- [custom_hash_algorithms](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms.md): complete subsection reference.

<a id="schema-access_info--tls_config--common_params--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--disable_ocsp_stapling.md): complete subsection reference.

- [private_key](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key.md): complete subsection reference.

- [use_system_defaults](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--use_system_defaults.md): complete subsection reference.

## Next pages

- [access_info.tls_config.common_params.tls_certificates.custom_hash_algorithms](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--custom_hash_algorithms.md)
- [access_info.tls_config.common_params.tls_certificates.disable_ocsp_stapling](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--disable_ocsp_stapling.md)
- [access_info.tls_config.common_params.tls_certificates.private_key](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--private_key.md)
- [access_info.tls_config.common_params.tls_certificates.use_system_defaults](data-sources--secret_management_access--properties--access_info--tls_config--common_params--tls_certificates--use_system_defaults.md)
- [access_info.tls_config.common_params](data-sources--secret_management_access--properties--access_info--tls_config--common_params.md)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md)
