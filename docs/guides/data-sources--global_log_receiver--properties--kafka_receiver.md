---
page_title: "kafka_receiver"
subcategory: ""
description: "kafka_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 4214, "body_sha256": "sha256:a4e6adce61a398048212ba299ba0938ebbfa9cbc12c04aa9e309cfe969526270", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:no_tls", "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:use_tls"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "docs/guides/data-sources--global_log_receiver--properties--kafka_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kafka_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/kafka_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kafka_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# kafka_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- kafka_receiver

<a id="section"></a>

Type: `"single"`. Computed.

Kafka Configuration for Global Log Receiver.

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

## Direct properties

- [batch](data-sources--global_log_receiver--properties--kafka_receiver--batch.md): complete subsection reference.

<a id="schema-kafka_receiver--bootstrap_servers"></a>

### bootstrap_servers property

Type: `["list", "string"]`. Computed.

List of host:port pairs of the Kafka brokers.

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

- [compression](data-sources--global_log_receiver--properties--kafka_receiver--compression.md): complete subsection reference.

<a id="schema-kafka_receiver--kafka_topic"></a>

### kafka_topic property

Type: `"string"`. Computed.

The Kafka topic name to write events to.

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

- [no_tls](data-sources--global_log_receiver--properties--kafka_receiver--no_tls.md): complete subsection reference.

- [use_tls](data-sources--global_log_receiver--properties--kafka_receiver--use_tls.md): complete subsection reference.

## Next pages

- [kafka_receiver.batch](data-sources--global_log_receiver--properties--kafka_receiver--batch.md)
- [kafka_receiver.compression](data-sources--global_log_receiver--properties--kafka_receiver--compression.md)
- [kafka_receiver.no_tls](data-sources--global_log_receiver--properties--kafka_receiver--no_tls.md)
- [kafka_receiver.use_tls](data-sources--global_log_receiver--properties--kafka_receiver--use_tls.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
