---
page_title: "request_logs"
subcategory: ""
description: "Configuration for request logs with sampling choice. Allows selection between sampled (default) or unsampled (full) request logs."
xcsh_docs: {"aliases": ["request logs"], "body_bytes": 1893, "body_sha256": "sha256:74ac3404095976ae22927665f24c43e0b8585062d388c77dd79ad2af68ec7c57", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:request_logs:sampled", "xcsh-docs:data-sources:global_log_receiver:properties:request_logs:unsampled"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:request_logs", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/request_logs/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["request_logs"], "schema_version": 1, "sections": [{"aliases": ["sampled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:request_logs:sampled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_logs", "sampled"], "syntax": "attribute", "type": "object"}, {"aliases": ["unsampled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:request_logs:unsampled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_logs", "unsampled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/request_logs/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration for request logs with sampling choice. Allows selection between sampled (default) or unsampled (full) request logs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_logs

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- request_logs

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Upstream description:

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-sampling_choice": "[\"sampled\",\"unsampled\"]"
}
```

## Direct properties

- [sampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/request_logs/sampled/): complete subsection reference.

- [unsampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/request_logs/unsampled/): complete subsection reference.

## Next pages

- [request_logs.sampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/request_logs/sampled/)
- [request_logs.unsampled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/request_logs/unsampled/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
