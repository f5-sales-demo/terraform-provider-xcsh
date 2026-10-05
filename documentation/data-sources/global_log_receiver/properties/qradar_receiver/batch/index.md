---
page_title: "qradar_receiver.batch"
subcategory: ""
description: "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint."
xcsh_docs: {"aliases": ["qradar receiver batch"], "body_bytes": 5768, "body_sha256": "sha256:42a2d43417e69b872aa110d1bf7a9bc6e3fdcf69ea3ffdf8bc20b6fab412581e", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch:max_bytes_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch:max_events_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch:timeout_seconds_default"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver", "path": "documentation/data-sources/global_log_receiver/properties/qradar_receiver/batch/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3322102130323103-1103332202002221-0010102101231230-1122230223022300-2103102121313033-1201330301311321-1311300322122211-2310302202321201", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["qradar_receiver", "batch"], "schema_version": 1, "sections": [{"aliases": ["qradar receiver batch max bytes"], "anchor": "schema-qradar_receiver--batch--max_bytes", "description": "Exclusive with Send batch to endpoint after the batch is equal to or larger than this many bytes.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "batch", "max_bytes"], "syntax": "attribute", "type": "number"}, {"aliases": ["qradar receiver batch max bytes disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch:max_bytes_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "batch", "max_bytes_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["qradar receiver batch max events"], "anchor": "schema-qradar_receiver--batch--max_events", "description": "Exclusive with Send batch to endpoint after this many log messages are in the batch.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "batch", "max_events"], "syntax": "attribute", "type": "number"}, {"aliases": ["qradar receiver batch max events disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch:max_events_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "batch", "max_events_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "qradar receiver batch timeout seconds"], "anchor": "schema-qradar_receiver--batch--timeout_seconds", "description": "Exclusive with Send batch to the endpoint after this many seconds.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "batch", "timeout_seconds"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "qradar receiver batch timeout seconds default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch:timeout_seconds_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["qradar_receiver", "batch", "timeout_seconds_default"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/qradar_receiver/batch/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver.batch

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/)
- qradar_receiver.batch

<a id="section"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

## Direct properties

<a id="schema-qradar_receiver--batch--max_bytes"></a>

### max_bytes property

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Upstream description:

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/batch/max_bytes_disabled/): complete subsection reference.

<a id="schema-qradar_receiver--batch--max_events"></a>

### max_events property

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Upstream description:

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/batch/max_events_disabled/): complete subsection reference.

<a id="schema-qradar_receiver--batch--timeout_seconds"></a>

### timeout_seconds property

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Upstream description:

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/batch/timeout_seconds_default/): complete subsection reference.

## Next pages

- [qradar_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/batch/max_bytes_disabled/)
- [qradar_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/batch/max_events_disabled/)
- [qradar_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/batch/timeout_seconds_default/)
- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
