---
page_title: "s3_receiver.batch"
subcategory: ""
description: "s3_receiver.batch for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 5145, "body_sha256": "sha256:6f610656a1ad5bc8d0e6b1b2cf68849f9828a44b88f37e76f8f193c19ec87213", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:batch", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:batch:max_bytes_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:batch:max_events_disabled", "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:batch:timeout_seconds_default"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:batch", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver", "path": "docs/guides/data-sources--global_log_receiver--properties--s3_receiver--batch.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["s3_receiver", "batch"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/s3_receiver/batch/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "s3_receiver.batch for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# s3_receiver.batch

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [s3_receiver](data-sources--global_log_receiver--properties--s3_receiver.md)
- s3_receiver.batch

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

<a id="schema-s3_receiver--batch--max_bytes"></a>

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

- [max_bytes_disabled](data-sources--global_log_receiver--properties--s3_receiver--batch--max_bytes_disabled.md): complete subsection reference.

<a id="schema-s3_receiver--batch--max_events"></a>

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

- [max_events_disabled](data-sources--global_log_receiver--properties--s3_receiver--batch--max_events_disabled.md): complete subsection reference.

<a id="schema-s3_receiver--batch--timeout_seconds"></a>

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

- [timeout_seconds_default](data-sources--global_log_receiver--properties--s3_receiver--batch--timeout_seconds_default.md): complete subsection reference.

## Next pages

- [s3_receiver.batch.max_bytes_disabled](data-sources--global_log_receiver--properties--s3_receiver--batch--max_bytes_disabled.md)
- [s3_receiver.batch.max_events_disabled](data-sources--global_log_receiver--properties--s3_receiver--batch--max_events_disabled.md)
- [s3_receiver.batch.timeout_seconds_default](data-sources--global_log_receiver--properties--s3_receiver--batch--timeout_seconds_default.md)
- [s3_receiver](data-sources--global_log_receiver--properties--s3_receiver.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
