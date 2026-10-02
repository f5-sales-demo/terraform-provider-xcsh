---
page_title: "kafka_receiver"
subcategory: ""
description: "Kafka Configuration for Global Log Receiver."
xcsh_docs: {"aliases": ["kafka receiver"], "body_bytes": 5526, "body_sha256": "sha256:1ab8ed2d1db4999a5f73abd29426c14ab7cb2f8c295c4565803d4103f09c98c6", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:no_tls", "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/kafka_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3232010031122222-1223223213331010-3301333221230101-2312232222210120-2110111021320310-1231013030323111-0302310312310312-0223030012322112", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:no_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver:ConflictingObjectAttributes:no_tls,use_tls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "type": "conflicts"}, {"anchor": "schema-kafka_receiver--bootstrap_servers", "enforcement": "provider-schema", "group": "kafka_receiver:RequiredObjectAttributes:bootstrap_servers,kafka_topic", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "type": "requires"}, {"anchor": "schema-kafka_receiver--kafka_topic", "enforcement": "provider-schema", "group": "kafka_receiver:RequiredObjectAttributes:bootstrap_servers,kafka_topic", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["kafka_receiver"], "schema_version": 1, "sections": [{"aliases": ["batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-kafka_receiver--batch--max_bytes", "enforcement": "provider-schema", "group": "kafka_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch", "type": "conflicts"}, {"anchor": "schema-kafka_receiver--batch--max_events", "enforcement": "provider-schema", "group": "kafka_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch", "type": "conflicts"}, {"anchor": "schema-kafka_receiver--batch--timeout_seconds", "enforcement": "provider-schema", "group": "kafka_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch:max_bytes_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch:max_events_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch:timeout_seconds_default", "type": "conflicts"}], "schema_path": ["kafka_receiver", "batch"], "syntax": "block", "type": "object"}, {"aliases": ["bootstrap servers"], "anchor": "schema-kafka_receiver--bootstrap_servers", "description": "List of host:port pairs of the Kafka brokers.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "bootstrap_servers"], "syntax": "attribute", "type": "list"}, {"aliases": ["compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_none", "type": "conflicts"}], "schema_path": ["kafka_receiver", "compression"], "syntax": "block", "type": "object"}, {"aliases": ["kafka topic"], "anchor": "schema-kafka_receiver--kafka_topic", "description": "The Kafka topic name to write events to.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "kafka_topic"], "syntax": "attribute", "type": "string"}, {"aliases": ["no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:no_tls", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kafka_receiver", "no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["use tls"], "anchor": "section", "description": "TLS Parameters for client connection to the endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-kafka_receiver--use_tls--trusted_ca_url", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:disable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:disable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:disable_verify_certificate,enable_verify_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:disable_verify_hostname,enable_verify_hostname", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:enable_verify_hostname", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:mtls_disabled,mtls_enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:mtls_enable", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kafka_receiver.use_tls:ConflictingObjectAttributes:no_ca,trusted_ca_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:use_tls:no_ca", "type": "conflicts"}], "schema_path": ["kafka_receiver", "use_tls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/kafka_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Kafka Configuration for Global Log Receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
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

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/): complete subsection reference.

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

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/): complete subsection reference.

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

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/no_tls/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/): complete subsection reference.

## Next pages

- [kafka_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/)
- [kafka_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/)
- [kafka_receiver.no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/no_tls/)
- [kafka_receiver.use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/use_tls/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
