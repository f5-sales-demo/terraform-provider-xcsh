---
page_title: "kafka_receiver.compression.compression_default"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["kafka receiver compression compression default"], "body_bytes": 1246, "body_sha256": "sha256:9bb23a3c853c99c8defa85c359b243f08efcff1948644da3c1bb0952344bd1c0", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression:compression_default", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:compression", "path": "documentation/resources/global_log_receiver/properties/kafka_receiver/compression/compression_default/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1000012120303202-3220313100012230-1030231031303133-0332001321203210-1201302323002223-1210333321031321-0020133031110102-3200303213121330", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kafka_receiver", "compression", "compression_default"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/kafka_receiver/compression/compression_default/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver.compression.compression_default

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/)
- [kafka_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/compression/)
- kafka_receiver.compression.compression_default

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for compression default.

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

Terraform syntax:

```terraform
compression_default = {}
```

This is an empty object or choice marker. It has no direct properties.
