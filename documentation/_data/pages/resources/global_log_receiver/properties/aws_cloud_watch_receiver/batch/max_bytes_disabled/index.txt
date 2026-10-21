---
page_title: "aws_cloud_watch_receiver.batch.max_bytes_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["aws cloud watch receiver batch max bytes disabled"], "body_bytes": 1249, "body_sha256": "sha256:929753ac6da15405ef9c88ab44a070ce8bb4c5f75e3d34f4a63b486a081e512a", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch:max_bytes_disabled", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch", "path": "documentation/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/max_bytes_disabled/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0112333202213111-3012310213332112-0303121102121312-2301013310132233-0220033011011113-0121022320133203-1202111022023010-0130330220113110", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_cloud_watch_receiver", "batch", "max_bytes_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/max_bytes_disabled/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_cloud_watch_receiver.batch.max_bytes_disabled

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/)
- [aws_cloud_watch_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/)
- aws_cloud_watch_receiver.batch.max_bytes_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
max_bytes_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
