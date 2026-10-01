---
page_title: "qradar_receiver.compression"
subcategory: ""
description: "qradar_receiver.compression for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2375, "body_sha256": "sha256:b953315576d86ba74407fc9be241afe5edbf849a1558f9b559cb63ccc4af00a8", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_default", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_gzip", "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression:compression_none"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver:compression", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:qradar_receiver", "path": "docs/guides/resources--global_log_receiver--properties--qradar_receiver--compression.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["qradar_receiver", "compression"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/qradar_receiver/compression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "qradar_receiver.compression for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# qradar_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [qradar_receiver](resources--global_log_receiver--properties--qradar_receiver.md)
- qradar_receiver.compression

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for compression.

Upstream description:

Compression Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("compression_default",
    "compression_gzip"),
  validators.ConflictingObjectAttributes("compression_default",
    "compression_none"),
  validators.ConflictingObjectAttributes("compression_gzip",
    "compression_none")}
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
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

Terraform syntax:

```terraform
compression {
  # Configure direct properties listed below.
}
```

## Direct properties

- [compression_default](resources--global_log_receiver--properties--qradar_receiver--compression--compression_default.md): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--properties--qradar_receiver--compression--compression_gzip.md): complete subsection reference.

- [compression_none](resources--global_log_receiver--properties--qradar_receiver--compression--compression_none.md): complete subsection reference.

## Next pages

- [qradar_receiver.compression.compression_default](resources--global_log_receiver--properties--qradar_receiver--compression--compression_default.md)
- [qradar_receiver.compression.compression_gzip](resources--global_log_receiver--properties--qradar_receiver--compression--compression_gzip.md)
- [qradar_receiver.compression.compression_none](resources--global_log_receiver--properties--qradar_receiver--compression--compression_none.md)
- [qradar_receiver](resources--global_log_receiver--properties--qradar_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
