---
page_title: "aws_cloud_watch_receiver.compression"
subcategory: ""
description: "aws_cloud_watch_receiver.compression for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 2515, "body_sha256": "sha256:15cae1680b9c118e68415a2a1e1c265f7522acd2bb8388a34e599cda1334819b", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_default", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_gzip", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_none"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver", "path": "documentation/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["aws_cloud_watch_receiver", "compression"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_cloud_watch_receiver.compression for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# aws_cloud_watch_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/)
- aws_cloud_watch_receiver.compression

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

- [compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_default/): complete subsection reference.

- [compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_gzip/): complete subsection reference.

- [compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_none/): complete subsection reference.

## Next pages

- [aws_cloud_watch_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_default/)
- [aws_cloud_watch_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_gzip/)
- [aws_cloud_watch_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/compression_none/)
- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
