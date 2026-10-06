---
page_title: "aws_cloud_watch_receiver.compression"
subcategory: ""
description: "Compression Type."
xcsh_docs: {"aliases": ["aws cloud watch receiver compression"], "body_bytes": 1681, "body_sha256": "sha256:abf1953aa42f7b476ab07f37c09ade329d2d8246503ede20fe22c7863b8cc9b3", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_default", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_gzip", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_none"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver", "path": "documentation/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_cloud_watch_receiver", "compression"], "schema_version": 1, "sections": [{"aliases": ["aws cloud watch receiver compression compression default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "compression", "compression_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws cloud watch receiver compression compression gzip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_gzip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "compression", "compression_gzip"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws cloud watch receiver compression compression none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "compression", "compression_none"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Compression Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_cloud_watch_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/)
- aws_cloud_watch_receiver.compression

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Additional upstream details:

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
