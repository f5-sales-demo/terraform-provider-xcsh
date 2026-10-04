---
page_title: "aws_cloud_watch_receiver.batch.timeout_seconds_default"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["aws cloud watch receiver batch timeout seconds default", "duration"], "body_bytes": 1577, "body_sha256": "sha256:1c345eee906cd4f8a0fdebf212a554f2b7a93e2d73b76ce4a9eb17acda851a4d", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch:timeout_seconds_default", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch", "path": "documentation/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/timeout_seconds_default/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1330020102111032-2101030322010323-2021110331100011-1112300133111322-0222202132023103-1103223301120303-3113033321322132-2332022230303200", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_cloud_watch_receiver", "batch", "timeout_seconds_default"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/timeout_seconds_default/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_cloud_watch_receiver.batch.timeout_seconds_default

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/)
- [aws_cloud_watch_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/)
- aws_cloud_watch_receiver.batch.timeout_seconds_default

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

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
timeout_seconds_default = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [aws_cloud_watch_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
