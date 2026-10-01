---
page_title: "gcp_bucket_receiver.compression"
subcategory: ""
description: "gcp_bucket_receiver.compression for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2435, "body_sha256": "sha256:3de4a4b51aaeb7429e8623cfa2f6b7bb404ed3ff55b584af4fc29a8a8ec3db29", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_default", "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_gzip", "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression:compression_none"], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:compression", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver", "path": "docs/guides/resources--global_log_receiver--properties--gcp_bucket_receiver--compression.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gcp_bucket_receiver", "compression"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/gcp_bucket_receiver/compression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gcp_bucket_receiver.compression for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_bucket_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- [gcp_bucket_receiver](resources--global_log_receiver--properties--gcp_bucket_receiver.md)
- gcp_bucket_receiver.compression

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

- [compression_default](resources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_default.md): complete subsection reference.

- [compression_gzip](resources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_gzip.md): complete subsection reference.

- [compression_none](resources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_none.md): complete subsection reference.

## Next pages

- [gcp_bucket_receiver.compression.compression_default](resources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_default.md)
- [gcp_bucket_receiver.compression.compression_gzip](resources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_gzip.md)
- [gcp_bucket_receiver.compression.compression_none](resources--global_log_receiver--properties--gcp_bucket_receiver--compression--compression_none.md)
- [gcp_bucket_receiver](resources--global_log_receiver--properties--gcp_bucket_receiver.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
