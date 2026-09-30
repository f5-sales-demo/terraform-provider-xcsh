---
page_title: "kafka_receiver"
subcategory: ""
description: "kafka_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4819, "body_sha256": "sha256:2e69089d5004f2700b46d51c9e90172c0cae5b94aab4c244e5239d8b7e0be475", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:no_tls", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--kafka_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kafka_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/kafka_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kafka_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# kafka_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- kafka_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Kafka Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("bootstrap_servers",
    "kafka_topic"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
kafka_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [batch](resources--global_log_receiver--properties--kafka_receiver--batch.md): complete subsection reference.

<a id="schema-kafka_receiver--bootstrap_servers"></a>

### bootstrap_servers property

Type: `["list", "string"]`. Optional.

List of host:port pairs of the Kafka brokers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostport": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [compression](resources--global_log_receiver--properties--kafka_receiver--compression.md): complete subsection reference.

<a id="schema-kafka_receiver--kafka_topic"></a>

### kafka_topic property

Type: `"string"`. Optional.

The Kafka topic name to write events to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-zA-Z0-9\\\\._\\\\-]+$"
  }
}
```

- [no_tls](resources--global_log_receiver--properties--kafka_receiver--no_tls.md): complete subsection reference.

- [use_tls](resources--global_log_receiver--properties--kafka_receiver--use_tls.md): complete subsection reference.

## Next pages

- [kafka_receiver.batch](resources--global_log_receiver--properties--kafka_receiver--batch.md)
- [kafka_receiver.compression](resources--global_log_receiver--properties--kafka_receiver--compression.md)
- [kafka_receiver.no_tls](resources--global_log_receiver--properties--kafka_receiver--no_tls.md)
- [kafka_receiver.use_tls](resources--global_log_receiver--properties--kafka_receiver--use_tls.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
