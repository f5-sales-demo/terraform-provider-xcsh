---
page_title: "qradar_receiver.use_tls.mtls_enable"
subcategory: ""
description: "MTLS Client config allows configuration of mTLS client OPTIONS."
xcsh_docs: {"aliases": ["qradar receiver use tls mtls enable"], "body_bytes": 2942, "body_sha256": "sha256:a117ac4dded89f3bf4e70ca8c7402014ea599a6bd1b78aa23d002c6b940f215c", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable:key_url"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls", "path": "documentation/data-sources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["qradar_receiver", "use_tls", "mtls_enable"], "schema_version": 1, "sections": [{"aliases": ["qradar receiver use tls mtls enable certificate"], "anchor": "schema-qradar_receiver--use_tls--mtls_enable--certificate", "description": "Client certificate is PEM-encoded certificate or certificate-chain.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "mtls_enable", "certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["qradar receiver use tls mtls enable key url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls:mtls_enable:key_url", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["qradar_receiver", "use_tls", "mtls_enable", "key_url"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "MTLS Client config allows configuration of mTLS client OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver.use_tls.mtls_enable

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/)
- [qradar_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/use_tls/)
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/): complete subsection reference.

## Next pages

- [qradar_receiver.use_tls.mtls_enable.key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/use_tls/mtls_enable/key_url/)
- [qradar_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/use_tls/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
