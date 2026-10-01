---
page_title: "http_receiver.use_tls"
subcategory: ""
description: "http_receiver.use_tls for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4474, "body_sha256": "sha256:6104f233e22d0ae80e4aac1893ce4cab26cdce95dec5f2f33c6296af12d2d687", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_certificate", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:disable_verify_hostname", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:enable_verify_certificate", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:enable_verify_hostname", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:mtls_enable", "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls:no_ca"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver:use_tls", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:http_receiver", "path": "docs/guides/data-sources--global_log_receiver--properties--http_receiver--use_tls.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["http_receiver", "use_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/http_receiver/use_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_receiver.use_tls for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_receiver.use_tls

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [http_receiver](data-sources--global_log_receiver--properties--http_receiver.md)
- http_receiver.use_tls

<a id="section"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

## Direct properties

- [disable_verify_certificate](data-sources--global_log_receiver--properties--http_receiver--use_tls--disable_verify_certificate.md): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--properties--http_receiver--use_tls--disable_verify_hostname.md): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--properties--http_receiver--use_tls--enable_verify_certificate.md): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--properties--http_receiver--use_tls--enable_verify_hostname.md): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_disabled.md): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable.md): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--properties--http_receiver--use_tls--no_ca.md): complete subsection reference.

<a id="schema-http_receiver--use_tls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

## Next pages

- [http_receiver.use_tls.disable_verify_certificate](data-sources--global_log_receiver--properties--http_receiver--use_tls--disable_verify_certificate.md)
- [http_receiver.use_tls.disable_verify_hostname](data-sources--global_log_receiver--properties--http_receiver--use_tls--disable_verify_hostname.md)
- [http_receiver.use_tls.enable_verify_certificate](data-sources--global_log_receiver--properties--http_receiver--use_tls--enable_verify_certificate.md)
- [http_receiver.use_tls.enable_verify_hostname](data-sources--global_log_receiver--properties--http_receiver--use_tls--enable_verify_hostname.md)
- [http_receiver.use_tls.mtls_disabled](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_disabled.md)
- [http_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--properties--http_receiver--use_tls--mtls_enable.md)
- [http_receiver.use_tls.no_ca](data-sources--global_log_receiver--properties--http_receiver--use_tls--no_ca.md)
- [http_receiver](data-sources--global_log_receiver--properties--http_receiver.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
