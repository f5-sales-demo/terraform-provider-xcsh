---
page_title: "syslog.tls_server"
subcategory: "Monitoring"
description: "TLS config for client of discovery service."
xcsh_docs: {"aliases": ["syslog tls server"], "body_bytes": 6550, "body_sha256": "sha256:ea64defd576f08b051155e9c920304b9174a6156a9f712f5c202607e7a700375", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:default_https_port", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_disabled", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_enable", "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:volterra_ca"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "parent_id": "xcsh-docs:data-sources:log_receiver:properties:syslog", "path": "documentation/data-sources/log_receiver/properties/syslog/tls_server/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032", "registry_path": "docs/guides/data-sources--log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["syslog", "tls_server"], "schema_version": 1, "sections": [{"aliases": ["default https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:default_https_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "default_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["default syslog tls port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "default_syslog_tls_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtls disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "mtls_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtls enable"], "anchor": "section", "description": "TLS config for client.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:mtls_enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["syslog", "tls_server", "mtls_enable"], "syntax": "attribute", "type": "object"}, {"aliases": ["port"], "anchor": "schema-syslog--tls_server--port", "description": "Exclusive with Custom port number used for communication.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["cert", "certificate", "existing certificates", "server name", "tls certificates"], "anchor": "schema-syslog--tls_server--server_name", "description": "ServerName is passed to the server for SNI and is used in the client to check server certificates against.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "trusted ca url"], "anchor": "schema-syslog--tls_server--trusted_ca_url", "description": "Exclusive with The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format including the PEM headers.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["volterra ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:log_receiver:properties:syslog:tls_server:volterra_ca", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "volterra_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/log_receiver/properties/syslog/tls_server/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "TLS config for client of discovery service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/)
- [syslog](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/)
- syslog.tls_server

<a id="section"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"trusted_ca_url\",\"volterra_ca\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"default_syslog_tls_port\",\"port\"]"
}
```

## Direct properties

- [default_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/default_https_port/): complete subsection reference.

- [default_syslog_tls_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/default_syslog_tls_port/): complete subsection reference.

- [mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/mtls_disabled/): complete subsection reference.

- [mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/mtls_enable/): complete subsection reference.

<a id="schema-syslog--tls_server--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Upstream description:

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-syslog--tls_server--server_name"></a>

### server_name property

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-syslog--tls_server--trusted_ca_url"></a>

### trusted_ca_url property

Type: `"string"`. Computed.

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/volterra_ca/): complete subsection reference.

## Next pages

- [syslog.tls_server.default_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/default_https_port/)
- [syslog.tls_server.default_syslog_tls_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/default_syslog_tls_port/)
- [syslog.tls_server.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/mtls_disabled/)
- [syslog.tls_server.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/mtls_enable/)
- [syslog.tls_server.volterra_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/tls_server/volterra_ca/)
- [syslog](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/properties/syslog/)
- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/log_receiver/)
