---
page_title: "kafka_receiver.use_tls.mtls_enable"
subcategory: ""
description: "kafka_receiver.use_tls.mtls_enable for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2528, "body_sha256": "sha256:823a9855ed201c9192f31adba054ac0d4517f2edefbe19c377d9570408d6242d", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable:key_url"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls", "path": "docs/guides/data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kafka_receiver", "use_tls", "mtls_enable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/kafka_receiver/use_tls/mtls_enable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kafka_receiver.use_tls.mtls_enable for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver.use_tls.mtls_enable

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [kafka_receiver](data-sources--global_log_receiver--properties--kafka_receiver.md)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--properties--kafka_receiver--use_tls.md)
- kafka_receiver.use_tls.mtls_enable

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

<a id="schema-kafka_receiver--use_tls--mtls_enable--certificate"></a>

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

- [key_url](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url.md): complete subsection reference.

## Next pages

- [kafka_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--properties--kafka_receiver--use_tls--mtls_enable--key_url.md)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--properties--kafka_receiver--use_tls.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
