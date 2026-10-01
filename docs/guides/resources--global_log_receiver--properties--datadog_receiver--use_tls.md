---
page_title: "datadog_receiver.use_tls"
subcategory: ""
description: "datadog_receiver.use_tls for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 5209, "body_sha256": "sha256:e821377f2433e7f6da82c1d18ca27c30ccffd8033ed558f732373afdf8d70aeb", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls:disable_verify_certificate", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls:disable_verify_hostname", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls:enable_verify_certificate", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls:enable_verify_hostname", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls:mtls_disabled", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls:mtls_enable", "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls:no_ca"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:use_tls", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver", "path": "docs/guides/resources--global_log_receiver--properties--datadog_receiver--use_tls.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["datadog_receiver", "use_tls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/datadog_receiver/use_tls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "datadog_receiver.use_tls for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# datadog_receiver.use_tls

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [datadog_receiver](resources--global_log_receiver--properties--datadog_receiver.md)
- datadog_receiver.use_tls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for client connection to the endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_verify_certificate",
    "enable_verify_certificate"),
  validators.ConflictingObjectAttributes("disable_verify_hostname",
    "enable_verify_hostname"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("no_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_verify_certificate](resources--global_log_receiver--properties--datadog_receiver--use_tls--disable_verify_certificate.md): complete subsection reference.

- [disable_verify_hostname](resources--global_log_receiver--properties--datadog_receiver--use_tls--disable_verify_hostname.md): complete subsection reference.

- [enable_verify_certificate](resources--global_log_receiver--properties--datadog_receiver--use_tls--enable_verify_certificate.md): complete subsection reference.

- [enable_verify_hostname](resources--global_log_receiver--properties--datadog_receiver--use_tls--enable_verify_hostname.md): complete subsection reference.

- [mtls_disabled](resources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_disabled.md): complete subsection reference.

- [mtls_enable](resources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable.md): complete subsection reference.

- [no_ca](resources--global_log_receiver--properties--datadog_receiver--use_tls--no_ca.md): complete subsection reference.

<a id="schema-datadog_receiver--use_tls--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Optional.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
}
```

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

- [datadog_receiver.use_tls.disable_verify_certificate](resources--global_log_receiver--properties--datadog_receiver--use_tls--disable_verify_certificate.md)
- [datadog_receiver.use_tls.disable_verify_hostname](resources--global_log_receiver--properties--datadog_receiver--use_tls--disable_verify_hostname.md)
- [datadog_receiver.use_tls.enable_verify_certificate](resources--global_log_receiver--properties--datadog_receiver--use_tls--enable_verify_certificate.md)
- [datadog_receiver.use_tls.enable_verify_hostname](resources--global_log_receiver--properties--datadog_receiver--use_tls--enable_verify_hostname.md)
- [datadog_receiver.use_tls.mtls_disabled](resources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_disabled.md)
- [datadog_receiver.use_tls.mtls_enable](resources--global_log_receiver--properties--datadog_receiver--use_tls--mtls_enable.md)
- [datadog_receiver.use_tls.no_ca](resources--global_log_receiver--properties--datadog_receiver--use_tls--no_ca.md)
- [datadog_receiver](resources--global_log_receiver--properties--datadog_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
