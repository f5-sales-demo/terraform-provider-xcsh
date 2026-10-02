---
page_title: "datadog_receiver.batch.max_bytes_disabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["datadog receiver batch max bytes disabled"], "body_bytes": 1498, "body_sha256": "sha256:f344f9ec33ac491c7f080d678ed420500bee5565fd2a6311578b2e93bf563981", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:batch:max_bytes_disabled", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:datadog_receiver:batch", "path": "documentation/resources/global_log_receiver/properties/datadog_receiver/batch/max_bytes_disabled/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3021020103021100-2032012202002331-1221212303102301-0001121023002033-2301310330113130-1200303111330100-3201203110102322-3103201222322102", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["datadog_receiver", "batch", "max_bytes_disabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/datadog_receiver/batch/max_bytes_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# datadog_receiver.batch.max_bytes_disabled

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/)
- [datadog_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/)
- datadog_receiver.batch.max_bytes_disabled

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
max_bytes_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [datadog_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/batch/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
