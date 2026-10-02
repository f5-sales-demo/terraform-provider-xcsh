---
page_title: "kafka_receiver.batch.timeout_seconds_default"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["duration", "kafka receiver batch timeout seconds default", "operation timeout"], "body_bytes": 1497, "body_sha256": "sha256:f78ac1d0481072fbc5716b3ab6ef90d99856f8b9b442e98299a7069ce649302c", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch:timeout_seconds_default", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:kafka_receiver:batch", "path": "documentation/resources/global_log_receiver/properties/kafka_receiver/batch/timeout_seconds_default/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0230322320212102-1231230032123311-2012002232222313-3231003111003330-1012202300222321-2210102220210002-3023113333111130-0000302233010212", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kafka_receiver", "batch", "timeout_seconds_default"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/kafka_receiver/batch/timeout_seconds_default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kafka_receiver.batch.timeout_seconds_default

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/)
- [kafka_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/)
- kafka_receiver.batch.timeout_seconds_default

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

- [kafka_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/batch/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
