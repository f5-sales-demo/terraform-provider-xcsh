---
page_title: "syslog"
subcategory: "Monitoring"
description: "Configuration for syslog server."
xcsh_docs: {"aliases": ["syslog"], "body_bytes": 2201, "body_sha256": "sha256:dca2ca9051c54ba8f128dd863c5303316c8a34c5bdc005fa5fff9d9a33a3fa30", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:log_receiver:properties:syslog:tcp_server", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "xcsh-docs:data-sources:log_receiver:properties:syslog:udp_server"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:log_receiver:properties:syslog", "parent_id": "xcsh-docs:data-sources:log_receiver:reference", "path": "documentation/data-sources/log_receiver/properties/syslog/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330", "registry_path": "docs/guides/data-sources--log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["syslog"], "schema_version": 1, "sections": [{"aliases": ["syslog syslog rfc5424"], "anchor": "schema-syslog--syslog_rfc5424", "description": "Exclusive with Select RFC5424 syslog format and maximum message length.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "syslog_rfc5424"], "syntax": "attribute", "type": "number"}, {"aliases": ["syslog tcp server"], "anchor": "section", "description": "Name and port number for a TCP server.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tcp_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["syslog", "tcp_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["syslog tls server"], "anchor": "section", "description": "TLS config for client of discovery service.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["syslog", "tls_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["syslog udp server"], "anchor": "section", "description": "Name and port number for a UDP server.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:udp_server", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["syslog", "udp_server"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/log_receiver/properties/syslog/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration for syslog server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["log_receiverCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/)
- syslog

<a id="section"></a>

Type: `"single"`. Computed.

Syslog Server Configuration. Configuration for syslog server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-format_choice": "[\"syslog_rfc5424\"]",
  "x-ves-oneof-field-mode_choice": "[\"tcp_server\",\"tls_server\",\"udp_server\"]"
}
```

## Direct properties

<a id="schema-syslog--syslog_rfc5424"></a>

### syslog_rfc5424 property

Type: `"number"`. Computed.

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 268435456,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 408
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  }
}
```

- [tcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tcp_server/): complete subsection reference.

- [tls_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/): complete subsection reference.

- [udp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/udp_server/): complete subsection reference.
