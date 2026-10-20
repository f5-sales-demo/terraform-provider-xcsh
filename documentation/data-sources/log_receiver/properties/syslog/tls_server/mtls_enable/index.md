---
page_title: "syslog.tls_server.mtls_enable"
subcategory: "Monitoring"
description: "TLS config for client."
xcsh_docs: {"aliases": ["syslog tls server mtls enable"], "body_bytes": 2393, "body_sha256": "sha256:277443511e750aad3e301e695bd76eec4a60695025a1f03cc1e6c897f1a6b978", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_enable", "parent_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "path": "documentation/data-sources/log_receiver/properties/syslog/tls_server/mtls_enable/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0021031100311221-0200310322310210-1003132312021011-2300020313013311-1212233202221013-0321223012230213-2200330310310103-2212020332100213", "registry_path": "docs/guides/data-sources--log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["syslog", "tls_server", "mtls_enable"], "schema_version": 1, "sections": [{"aliases": ["syslog tls server mtls enable certificate"], "anchor": "schema-syslog--tls_server--mtls_enable--certificate", "description": "Client certificate is PEM-encoded certificate or certificate-chain.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "mtls_enable", "certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["syslog tls server mtls enable key url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_enable:key_url", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["syslog", "tls_server", "mtls_enable", "key_url"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/log_receiver/properties/syslog/tls_server/mtls_enable/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "TLS config for client.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["log_receiverCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server.mtls_enable

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/)
- [syslog](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/)
- [syslog.tls_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/)
- syslog.tls_server.mtls_enable

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for mtls enable.

Additional upstream details:

TLS config for client.

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

<a id="schema-syslog--tls_server--mtls_enable--certificate"></a>

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
      "validatedAt": "2026-10-20T00:10:01+00:00"
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

- [key_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/mtls_enable/key_url/): complete subsection reference.
