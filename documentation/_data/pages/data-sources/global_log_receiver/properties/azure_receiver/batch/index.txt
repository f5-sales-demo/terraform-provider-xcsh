---
page_title: "azure_receiver.batch"
subcategory: ""
description: "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint."
xcsh_docs: {"aliases": ["azure receiver batch"], "body_bytes": 5750, "body_sha256": "sha256:e11413833c32983bfb98f016266ea804956e2aabda872cf4aa8f8aabf384455c", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch:max_bytes_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch:max_events_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch:timeout_seconds_default"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver", "path": "documentation/data-sources/global_log_receiver/properties/azure_receiver/batch/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2110231131123022-3210021331120313-0230130302332321-2122330301322100-0032000123010111-3000120033100202-2313222223213221-0231002202020032", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_receiver", "batch"], "schema_version": 1, "sections": [{"aliases": ["max bytes"], "anchor": "schema-azure_receiver--batch--max_bytes", "description": "Exclusive with Send batch to endpoint after the batch is equal to or larger than this many bytes.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "batch", "max_bytes"], "syntax": "attribute", "type": "number"}, {"aliases": ["max bytes disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch:max_bytes_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "batch", "max_bytes_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["max events"], "anchor": "schema-azure_receiver--batch--max_events", "description": "Exclusive with Send batch to endpoint after this many log messages are in the batch.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "batch", "max_events"], "syntax": "attribute", "type": "number"}, {"aliases": ["max events disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch:max_events_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "batch", "max_events_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeout seconds"], "anchor": "schema-azure_receiver--batch--timeout_seconds", "description": "Exclusive with Send batch to the endpoint after this many seconds.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "batch", "timeout_seconds"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "operation timeout", "timeout seconds default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:azure_receiver:batch:timeout_seconds_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_receiver", "batch", "timeout_seconds_default"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/azure_receiver/batch/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.batch

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/)
- azure_receiver.batch

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

<a id="schema-azure_receiver--batch--max_bytes"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/batch/max_bytes_disabled/): complete subsection reference.

<a id="schema-azure_receiver--batch--max_events"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/batch/max_events_disabled/): complete subsection reference.

<a id="schema-azure_receiver--batch--timeout_seconds"></a>

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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/batch/timeout_seconds_default/): complete subsection reference.

## Next pages

- [azure_receiver.batch.max_bytes_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/batch/max_bytes_disabled/)
- [azure_receiver.batch.max_events_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/batch/max_events_disabled/)
- [azure_receiver.batch.timeout_seconds_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/batch/timeout_seconds_default/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
