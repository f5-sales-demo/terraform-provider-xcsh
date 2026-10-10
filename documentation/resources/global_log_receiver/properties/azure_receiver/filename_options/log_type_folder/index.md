---
page_title: "azure_receiver.filename_options.log_type_folder"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["azure receiver filename options log type folder"], "body_bytes": 1250, "body_sha256": "sha256:f109093fe1942b9b054d1ab88ea09dbcc396b9c0a147d479902ca44c4b7351e0", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options:log_type_folder", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:azure_receiver:filename_options", "path": "documentation/resources/global_log_receiver/properties/azure_receiver/filename_options/log_type_folder/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2313133202312310-3003023111113031-1311332220021230-3012030101123323-1312022021301302-3010032220132011-2333133101221310-3331211233213210", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure_receiver", "filename_options", "log_type_folder"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/azure_receiver/filename_options/log_type_folder/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure_receiver.filename_options.log_type_folder

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/)
- [azure_receiver.filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/filename_options/)
- azure_receiver.filename_options.log_type_folder

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
