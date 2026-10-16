---
page_title: "splunk_receiver.batch.max_bytes_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["splunk receiver batch max bytes disabled"], "body_bytes": 1195, "body_sha256": "sha256:87ad20ad8e888e88ad98e7a4cf6d8fc2e51e846567f32e80d2101689bd5b8843", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:batch:max_bytes_disabled", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:splunk_receiver:batch", "path": "documentation/resources/global_log_receiver/properties/splunk_receiver/batch/max_bytes_disabled/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2231210211002131-2100003222320311-2132101201122032-0211120312220210-2200303132122202-3111120030210130-0213221102332330-0003323302322232", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["splunk_receiver", "batch", "max_bytes_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/splunk_receiver/batch/max_bytes_disabled/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# splunk_receiver.batch.max_bytes_disabled

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [splunk_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/)
- [splunk_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/batch/)
- splunk_receiver.batch.max_bytes_disabled

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
