---
page_title: "qradar_receiver.compression"
subcategory: ""
description: "qradar_receiver.compression for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1928, "body_sha256": "sha256:3c72863a0ca27afa08ebaed6d18598915262214f60835cfb3b96a4af40015d3b", "canonical_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:compression", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:compression:compression_default", "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:compression:compression_gzip", "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:compression:compression_none"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver:compression", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:qradar_receiver", "path": "docs/guides/data-sources--global_log_receiver--properties--qradar_receiver--compression.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["qradar_receiver", "compression"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/qradar_receiver/compression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "qradar_receiver.compression for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
- [Property reference](data-sources--global_log_receiver--reference.md)
- [qradar_receiver](data-sources--global_log_receiver--properties--qradar_receiver.md)
- qradar_receiver.compression

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

- [compression_default](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_default.md): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_gzip.md): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_none.md): complete subsection reference.

## Next pages

- [qradar_receiver.compression.compression_default](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_default.md)
- [qradar_receiver.compression.compression_gzip](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_gzip.md)
- [qradar_receiver.compression.compression_none](data-sources--global_log_receiver--properties--qradar_receiver--compression--compression_none.md)
- [qradar_receiver](data-sources--global_log_receiver--properties--qradar_receiver.md)
- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md)
