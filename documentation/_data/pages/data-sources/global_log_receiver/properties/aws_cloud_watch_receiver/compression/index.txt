---
page_title: "aws_cloud_watch_receiver.compression"
subcategory: ""
description: "Compression Type."
xcsh_docs: {"aliases": ["aws cloud watch receiver compression"], "body_bytes": 2614, "body_sha256": "sha256:1f034c37a41d7de0d02dae6cb80d89d7b1692cfbc445bc8bcd78a67a40992a5a", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_default", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_gzip", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_none"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver", "path": "documentation/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2032202232313301-1121003121131023-2131120222223220-0120202302023332-3100110313000323-2231331232011100-2221300013333022-3312201021331113", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_cloud_watch_receiver", "compression"], "schema_version": 1, "sections": [{"aliases": ["aws cloud watch receiver compression compression default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "compression", "compression_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws cloud watch receiver compression compression gzip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_gzip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "compression", "compression_gzip"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws cloud watch receiver compression compression none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "compression", "compression_none"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Compression Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
