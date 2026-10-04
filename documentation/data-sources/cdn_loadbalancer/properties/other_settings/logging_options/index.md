---
page_title: "other_settings.logging_options"
subcategory: "Load Balancing"
description: "This defines various OPTIONS related to logging."
xcsh_docs: {"aliases": ["other settings logging options"], "body_bytes": 1994, "body_sha256": "sha256:e0e6d8dfb3257a5ff88aaac2c7cef453661a3cca9a24e7ffce94f6202f6dbb95", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "path": "documentation/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2223103101102133-1321201122230001-0323100123120022-2010212213323231-0212002021131321-2330330122301331-0201233222123130-0011222322220210", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["other_settings", "logging_options"], "schema_version": 1, "sections": [{"aliases": ["other settings logging options client log options"], "anchor": "section", "description": "List of headers to Log.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "logging_options", "client_log_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "other settings logging options origin log options", "upstream servers"], "anchor": "section", "description": "List of headers to Log.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "logging_options", "origin_log_options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "This defines various OPTIONS related to logging.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings.logging_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [other_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/)
- other_settings.logging_options

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS related to logging.

Upstream description:

This defines various OPTIONS related to logging.

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

## Direct properties

- [client_log_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/client_log_options/): complete subsection reference.

- [origin_log_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/origin_log_options/): complete subsection reference.

## Next pages

- [other_settings.logging_options.client_log_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/client_log_options/)
- [other_settings.logging_options.origin_log_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/origin_log_options/)
- [other_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
