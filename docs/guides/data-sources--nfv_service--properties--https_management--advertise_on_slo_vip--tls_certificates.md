---
page_title: "https_management.advertise_on_slo_vip.tls_certificates"
subcategory: ""
description: "https_management.advertise_on_slo_vip.tls_certificates for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 5147, "body_sha256": "sha256:9a742b2dd1fe863bc37187684705a30f7c83d839badd7a95f6f7cda7bc9b6c4a", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_certificates", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_certificates:custom_hash_algorithms", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_certificates:disable_ocsp_stapling", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_certificates:private_key", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_certificates:use_system_defaults"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_certificates", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip", "path": "docs/guides/data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_slo_vip", "tls_certificates"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_certificates/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_slo_vip.tls_certificates for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_vip.tls_certificates

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [https_management](data-sources--nfv_service--properties--https_management.md)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip.md)
- https_management.advertise_on_slo_vip.tls_certificates

<a id="section"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

## Direct properties

<a id="schema-https_management--advertise_on_slo_vip--tls_certificates--certificate_url"></a>

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

- [custom_hash_algorithms](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--custom_hash_algorithms.md): complete subsection reference.

<a id="schema-https_management--advertise_on_slo_vip--tls_certificates--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--disable_ocsp_stapling.md): complete subsection reference.

- [private_key](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key.md): complete subsection reference.

- [use_system_defaults](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--use_system_defaults.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_vip.tls_certificates.custom_hash_algorithms](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--custom_hash_algorithms.md)
- [https_management.advertise_on_slo_vip.tls_certificates.disable_ocsp_stapling](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--disable_ocsp_stapling.md)
- [https_management.advertise_on_slo_vip.tls_certificates.private_key](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--private_key.md)
- [https_management.advertise_on_slo_vip.tls_certificates.use_system_defaults](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip--tls_certificates--use_system_defaults.md)
- [https_management.advertise_on_slo_vip](data-sources--nfv_service--properties--https_management--advertise_on_slo_vip.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
