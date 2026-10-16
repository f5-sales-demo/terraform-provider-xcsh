---
page_title: "gcp_bucket_receiver.batch.max_events_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["gcp bucket receiver batch max events disabled"], "body_bytes": 1222, "body_sha256": "sha256:1a6f70eacb2db48a4fe2574c69d4d10c98e66fa2b578e3311904504bbf061861", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch:max_events_disabled", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:batch", "path": "documentation/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/max_events_disabled/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0310031031220202-3201210220000022-0313033002111113-1220312002121310-3223303121231321-2330223002020310-2212023330031221-1311222112021220", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_bucket_receiver", "batch", "max_events_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/max_events_disabled/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_bucket_receiver.batch.max_events_disabled

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/)
- [gcp_bucket_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/batch/)
- gcp_bucket_receiver.batch.max_events_disabled

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
max_events_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.
