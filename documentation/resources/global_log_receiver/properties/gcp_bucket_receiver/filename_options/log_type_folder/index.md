---
page_title: "gcp_bucket_receiver.filename_options.log_type_folder"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["gcp bucket receiver filename options log type folder"], "body_bytes": 1280, "body_sha256": "sha256:c7363740d474863103171a0d83611b600bcf46b2347177151c5b8004d5d950bd", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options:log_type_folder", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:gcp_bucket_receiver:filename_options", "path": "documentation/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/log_type_folder/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3022233221120233-1000320331230333-2010030300302011-1232223000120223-1330323012121211-2311101311131002-1132323023020223-3322333322230223", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp_bucket_receiver", "filename_options", "log_type_folder"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/log_type_folder/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp_bucket_receiver.filename_options.log_type_folder

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/)
- [gcp_bucket_receiver.filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/filename_options/)
- gcp_bucket_receiver.filename_options.log_type_folder

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for log type folder.

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
log_type_folder = {}
```

This is an empty object or choice marker. It has no direct properties.
