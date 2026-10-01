---
page_title: "kafka_receiver.compression"
subcategory: ""
description: "kafka_receiver.compression for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1913, "body_sha256": "sha256:419feda9a1bad41aadd66cae27e8b031c5bbd7cf3bbd61b3699ad1505521a2f8", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:compression", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:compression:compression_default", "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:compression:compression_gzip", "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:compression:compression_none"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver:compression", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:kafka_receiver", "path": "docs/guides/data-sources--global_log_receiver--properties--kafka_receiver--compression.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["kafka_receiver", "compression"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/kafka_receiver/compression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "kafka_receiver.compression for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [kafka_receiver](data-sources--global_log_receiver--properties--kafka_receiver.md)
- kafka_receiver.compression

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Upstream description:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

## Direct properties

- [compression_default](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_default.md): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_gzip.md): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_none.md): complete subsection reference.

## Next pages

- [kafka_receiver.compression.compression_default](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_default.md)
- [kafka_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_gzip.md)
- [kafka_receiver.compression.compression_none](data-sources--global_log_receiver--properties--kafka_receiver--compression--compression_none.md)
- [kafka_receiver](data-sources--global_log_receiver--properties--kafka_receiver.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
