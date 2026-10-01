---
page_title: "qradar_receiver.use_tls.mtls_enable"
subcategory: ""
description: "qradar_receiver.use_tls.mtls_enable for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2540, "body_sha256": "sha256:69cc00f059bab69b18def610d39e945fa6e2bac9c877073a08806c04cdd9c027", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable:key_url"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls", "path": "docs/guides/data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["qradar_receiver", "use_tls", "mtls_enable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "qradar_receiver.use_tls.mtls_enable for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver.use_tls.mtls_enable

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [qradar_receiver](data-sources--global_log_receiver--properties--qradar_receiver.md)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--properties--qradar_receiver--use_tls.md)
- qradar_receiver.use_tls.mtls_enable

<a id="section"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

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

<a id="schema-qradar_receiver--use_tls--mtls_enable--certificate"></a>

### certificate property

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "format": "uri",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url.md): complete subsection reference.

## Next pages

- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--qradar_receiver--use_tls--mtls_enable--key_url.md)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--properties--qradar_receiver--use_tls.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
