---
page_title: "s3_receiver.compression.compression_none"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["s3 receiver compression compression none"], "body_bytes": 1169, "body_sha256": "sha256:3c62b2d599f8b61f8e072b23d2b44e43c913e970442fe6367fe5945506f255cf", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:compression:compression_none", "parent_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:compression", "path": "documentation/data-sources/global_log_receiver/properties/s3_receiver/compression/compression_none/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3000302301201303-3102312231000030-2230122033111202-1100210120321131-3311002022330203-3223331330220131-2211300321013303-1221310330001320", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["s3_receiver", "compression", "compression_none"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/s3_receiver/compression/compression_none/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# s3_receiver.compression.compression_none

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [s3_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/)
- [s3_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/compression/)
- s3_receiver.compression.compression_none

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.
