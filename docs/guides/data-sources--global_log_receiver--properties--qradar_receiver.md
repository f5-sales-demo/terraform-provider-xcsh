---
page_title: "qradar_receiver"
subcategory: ""
description: "qradar_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2686, "body_sha256": "sha256:d3c53099d24cedc1771e810849023ce676501dae6cb8cf1018183af023bb56f0", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:no_tls", "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:use_tls"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "docs/guides/data-sources--global_log_receiver--properties--qradar_receiver.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["qradar_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/qradar_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "qradar_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- qradar_receiver

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for qradar receiver.

Upstream description:

Configuration for IBM QRadar endpoint.

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

- [batch](data-sources--global_log_receiver--properties--qradar_receiver--batch.md): complete subsection reference.

- [compression](data-sources--global_log_receiver--properties--qradar_receiver--compression.md): complete subsection reference.

- [no_tls](data-sources--global_log_receiver--properties--qradar_receiver--no_tls.md): complete subsection reference.

<a id="schema-qradar_receiver--uri"></a>

### uri property

Type: `"string"`. Computed.

Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to,.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](data-sources--global_log_receiver--properties--qradar_receiver--use_tls.md): complete subsection reference.

## Next pages

- [qradar_receiver.batch](data-sources--global_log_receiver--properties--qradar_receiver--batch.md)
- [qradar_receiver.compression](data-sources--global_log_receiver--properties--qradar_receiver--compression.md)
- [qradar_receiver.no_tls](data-sources--global_log_receiver--properties--qradar_receiver--no_tls.md)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--properties--qradar_receiver--use_tls.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
