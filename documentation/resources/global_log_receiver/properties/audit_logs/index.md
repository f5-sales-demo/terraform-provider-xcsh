---
page_title: "audit_logs"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["audit logs"], "body_bytes": 1560, "body_sha256": "sha256:2ed7c77686649917a73487a0873c7413349f2bc6871560a6034cd8cb4e153df7", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:audit_logs", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/audit_logs/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3231131323212011-3030323112222132-1230222222202231-3212031220330031-2231302201012313-3103320303103003-3102113033100033-1132112103231012", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["audit_logs"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/audit_logs/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# audit_logs

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- audit_logs

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: audit\_logs, dns\_logs, request\_logs, security\_events\] Enable this option

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

OneOf alternatives in this subsection:

- [audit_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/audit_logs/#section)
- [dns_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/dns_logs/#section)
- [request_logs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/request_logs/#section)
- [security_events](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/security_events/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
audit_logs = {}
```

This is an empty object or choice marker. It has no direct properties.
