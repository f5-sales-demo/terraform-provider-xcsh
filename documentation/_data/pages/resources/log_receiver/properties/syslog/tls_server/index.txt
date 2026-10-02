---
page_title: "syslog.tls_server"
subcategory: "Monitoring"
description: "TLS config for client of discovery service."
xcsh_docs: {"aliases": ["syslog tls server"], "body_bytes": 7598, "body_sha256": "sha256:4531bd44536283dc00aef25b7c1eb0291277ee843af6a5885b67ec8169f09b19", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_https_port", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_disabled", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:volterra_ca"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "parent_id": "xcsh-docs:resources:log_receiver:properties:syslog", "path": "documentation/resources/log_receiver/properties/syslog/tls_server/index.md", "product": "distributed-cloud", "provider_name": "log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3200100101112223-3031012131332101-0233231200012200-3300210220323323-3010221310003023-2131333213030031-2302320323113222-2213232001330132", "registry_path": "docs/guides/resources--log_receiver--reference--group-001.md", "relationships": [{"anchor": "schema-syslog--tls_server--port", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_https_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "conflicts"}, {"anchor": "schema-syslog--tls_server--port", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_syslog_tls_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "conflicts"}, {"anchor": "schema-syslog--tls_server--trusted_ca_url", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:trusted_ca_url,volterra_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_https_port,default_syslog_tls_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_https_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_https_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_https_port,default_syslog_tls_port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:default_syslog_tls_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "syslog.tls_server:ConflictingObjectAttributes:trusted_ca_url,volterra_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:volterra_ca", "type": "conflicts"}, {"anchor": "schema-syslog--tls_server--server_name", "enforcement": "provider-schema", "group": "syslog.tls_server:RequiredObjectAttributes:server_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["syslog", "tls_server"], "schema_version": 1, "sections": [{"aliases": ["default https port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_https_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "default_https_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["default syslog tls port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:default_syslog_tls_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "default_syslog_tls_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtls disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "mtls_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtls enable"], "anchor": "section", "description": "TLS config for client.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:mtls_enable", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["syslog", "tls_server", "mtls_enable"], "syntax": "block", "type": "object"}, {"aliases": ["port"], "anchor": "schema-syslog--tls_server--port", "description": "Exclusive with Custom port number used for communication.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["cert", "certificate", "existing certificates", "server name", "tls certificates"], "anchor": "schema-syslog--tls_server--server_name", "description": "ServerName is passed to the server for SNI and is used in the client to check server certificates against.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "trusted ca url"], "anchor": "schema-syslog--tls_server--trusted_ca_url", "description": "Exclusive with The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format including the PEM headers.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "trusted_ca_url"], "syntax": "attribute", "type": "string"}, {"aliases": ["volterra ca"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:log_receiver:properties:syslog:tls_server:volterra_ca", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["syslog", "tls_server", "volterra_ca"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/tls_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "TLS config for client of discovery service.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog.tls_server

Breadcrumbs:

- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/)
- [syslog](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/)
- syslog.tls_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

TLS config for client of discovery service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server_name"),
  validators.ConflictingObjectAttributes("default_https_port",
    "default_syslog_tls_port"),
  validators.ConflictingObjectAttributes("default_https_port",
    "port"),
  validators.ConflictingObjectAttributes("default_syslog_tls_port",
    "port"),
  validators.ConflictingObjectAttributes("mtls_disabled",
    "mtls_enable"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_ca")}
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
  "x-ves-oneof-field-ca_choice": "[\"trusted_ca_url\",\"volterra_ca\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"default_syslog_tls_port\",\"port\"]"
}
```

Terraform syntax:

```terraform
tls_server {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/default_https_port/): complete subsection reference.

- [default_syslog_tls_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/default_syslog_tls_port/): complete subsection reference.

- [mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_disabled/): complete subsection reference.

- [mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_enable/): complete subsection reference.

<a id="schema-syslog--tls_server--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Upstream description:

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

Type: `"string"`. Optional.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

Type: `"string"`. Optional.

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

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

- [volterra_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/volterra_ca/): complete subsection reference.

## Next pages

- [syslog.tls_server.default_https_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/default_https_port/)
- [syslog.tls_server.default_syslog_tls_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/default_syslog_tls_port/)
- [syslog.tls_server.mtls_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_disabled/)
- [syslog.tls_server.mtls_enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/mtls_enable/)
- [syslog.tls_server.volterra_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/tls_server/volterra_ca/)
- [syslog](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/properties/syslog/)
- [xcsh_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/log_receiver/)
