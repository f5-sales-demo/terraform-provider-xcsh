---
page_title: "syslog"
subcategory: "Monitoring"
description: "syslog for xcsh_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 3041, "body_sha256": "sha256:3cd0b941bb936e7dbbb2d50ff2f4d8f8cd6bc5d8de8b7725ed50ab4b3232a08e", "canonical_id": "xcsh-docs:resources:log_receiver:properties:syslog", "child_ids": ["xcsh-docs:resources:log_receiver:properties:syslog:tcp_server", "xcsh-docs:resources:log_receiver:properties:syslog:tls_server", "xcsh-docs:resources:log_receiver:properties:syslog:udp_server"], "collection_id": "xcsh-docs:resources:log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:log_receiver:properties:syslog", "parent_id": "xcsh-docs:resources:log_receiver:reference", "path": "docs/guides/resources--log_receiver--properties--syslog.md", "provider_name": "log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["syslog"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/log_receiver/properties/syslog/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "syslog for xcsh_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# syslog

Breadcrumbs:

- [xcsh_log_receiver](../resources/log_receiver.md)
- [Property reference](resources--log_receiver--reference.md)
- syslog

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Syslog Server Configuration. Configuration for syslog server.

Upstream description:

Configuration for syslog server.

Provider validators and defaults (from schema source):

```go
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

Upstream description:

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [tcp_server](resources--log_receiver--properties--syslog--tcp_server.md): complete subsection reference.

- [tls_server](resources--log_receiver--properties--syslog--tls_server.md): complete subsection reference.

- [udp_server](resources--log_receiver--properties--syslog--udp_server.md): complete subsection reference.

## Next pages

- [syslog.tcp_server](resources--log_receiver--properties--syslog--tcp_server.md)
- [syslog.tls_server](resources--log_receiver--properties--syslog--tls_server.md)
- [syslog.udp_server](resources--log_receiver--properties--syslog--udp_server.md)
- [Property reference](resources--log_receiver--reference.md)
- [xcsh_log_receiver](../resources/log_receiver.md)
