---
page_title: "syslog"
subcategory: "Monitoring"
description: "Configuration for syslog server."
xcsh_docs: {"aliases": ["syslog"], "body_bytes": 2821, "body_sha256": "sha256:10c3a6dd9480d79611d39a6ad5dd48fa5c745bf80e3fcf8b2e2534d56548ec44", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:log_receiver:properties:syslog:tcp_server", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "xcsh-docs:resources:log_receiver:properties:syslog:udp_server"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog", "parent_id": "xcsh-docs:resources:log_receiver:reference", "path": "documentation/resources/log_receiver/properties/syslog/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0203202210332311-3021022100100323-3112010101230333-1102322323103120-0302200123210110-2102301220202230-0110212323333002-0311111012210231", "registry_path": "docs/guides/resources--log_receiver--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "syslog:ConflictingObjectAttributes:tcp_server,tls_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog:ConflictingObjectAttributes:tcp_server,udp_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tcp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog:ConflictingObjectAttributes:tcp_server,tls_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog:ConflictingObjectAttributes:tls_server,udp_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog:ConflictingObjectAttributes:tcp_server,udp_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:udp_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog:ConflictingObjectAttributes:tls_server,udp_server", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:udp_server", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["syslog"], "schema_version": 1, "sections": [{"aliases": ["syslog syslog rfc5424"], "anchor": "schema-syslog--syslog_rfc5424", "description": "Exclusive with Select RFC5424 syslog format and maximum message length.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "syslog_rfc5424"], "syntax": "attribute", "type": "number"}, {"aliases": ["syslog tcp server"], "anchor": "section", "description": "Name and port number for a TCP server.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tcp_server", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-syslog--tcp_server--port", "enforcement": "provider-schema", "group": "syslog.tcp_server:RequiredObjectAttributes:port,server_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tcp_server", "type": "requires"}, {"anchor": "schema-syslog--tcp_server--server_name", "enforcement": "provider-schema", "group": "syslog.tcp_server:RequiredObjectAttributes:port,server_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tcp_server", "type": "requires"}], "schema_path": ["syslog", "tcp_server"], "syntax": "block", "type": "object"}, {"aliases": ["syslog tls server"], "anchor": "section", "description": "TLS config for client of discovery service.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-syslog--tls_server--port", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_https_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "conflicts"}, {"anchor": "schema-syslog--tls_server--port", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_syslog_tls_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "conflicts"}, {"anchor": "schema-syslog--tls_server--trusted_ca_url", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:trusted_ca_url,volterra_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_https_port,default_syslog_tls_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_https_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_https_port,default_syslog_tls_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_syslog_tls_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:trusted_ca_url,volterra_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:volterra_ca", "type": "conflicts"}, {"anchor": "schema-syslog--tls_server--server_name", "enforcement": "provider-schema", "group": "syslog.tls_server:RequiredObjectAttributes:server_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "requires"}], "schema_path": ["syslog", "tls_server"], "syntax": "block", "type": "object"}, {"aliases": ["syslog udp server"], "anchor": "section", "description": "Name and port number for a UDP server.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:udp_server", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-syslog--udp_server--port", "enforcement": "provider-schema", "group": "syslog.udp_server:RequiredObjectAttributes:port,server_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:udp_server", "type": "requires"}, {"anchor": "schema-syslog--udp_server--server_name", "enforcement": "provider-schema", "group": "syslog.udp_server:RequiredObjectAttributes:port,server_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:udp_server", "type": "requires"}], "schema_path": ["syslog", "udp_server"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration for syslog server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["log_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/)
- syslog

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Syslog Server Configuration. Configuration for syslog server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("tcp_server",
    "tls_server"),
  validators.ConflictingObjectAttributes("tcp_server",
    "udp_server"),
  validators.ConflictingObjectAttributes("tls_server",
    "udp_server")}
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
  "x-ves-oneof-field-format_choice": "[\"syslog_rfc5424\"]",
  "x-ves-oneof-field-mode_choice": "[\"tcp_server\",\"tls_server\",\"udp_server\"]"
}
```

Terraform syntax:

```terraform
syslog {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-syslog--syslog_rfc5424"></a>

### syslog_rfc5424 property

Type: `"number"`. Optional.

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(408, 268435456),
}
```

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [tcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tcp_server/): complete subsection reference.

- [tls_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/): complete subsection reference.

- [udp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/udp_server/): complete subsection reference.
