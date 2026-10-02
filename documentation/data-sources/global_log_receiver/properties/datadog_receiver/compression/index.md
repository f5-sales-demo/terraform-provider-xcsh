---
page_title: "datadog_receiver.compression"
subcategory: ""
description: "Compression Type."
xcsh_docs: {"aliases": ["datadog receiver compression"], "body_bytes": 2494, "body_sha256": "sha256:b5b0e297f2b9603967e26b1e4e3f9fec75dc08061b69731a8be8a1bb41f2fc82", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:compression:compression_default", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:compression:compression_gzip", "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:compression:compression_none"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:compression", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver", "path": "documentation/data-sources/global_log_receiver/properties/datadog_receiver/compression/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0111332233101201-3133220000000011-2101103331333201-0011011330330330-2320103123303111-1230120233123202-2133312211313103-1002223122102203", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["datadog_receiver", "compression"], "schema_version": 1, "sections": [{"aliases": ["compression default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:compression:compression_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["datadog_receiver", "compression", "compression_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["compression gzip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:compression:compression_gzip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["datadog_receiver", "compression", "compression_gzip"], "syntax": "attribute", "type": "object"}, {"aliases": ["compression none"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:datadog_receiver:compression:compression_none", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["datadog_receiver", "compression", "compression_none"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/datadog_receiver/compression/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Compression Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# datadog_receiver.compression

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/)
- datadog_receiver.compression

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

- [compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/compression/compression_default/): complete subsection reference.

- [compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/compression/compression_gzip/): complete subsection reference.

- [compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/compression/compression_none/): complete subsection reference.

## Next pages

- [datadog_receiver.compression.compression_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/compression/compression_default/)
- [datadog_receiver.compression.compression_gzip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/compression/compression_gzip/)
- [datadog_receiver.compression.compression_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/compression/compression_none/)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
