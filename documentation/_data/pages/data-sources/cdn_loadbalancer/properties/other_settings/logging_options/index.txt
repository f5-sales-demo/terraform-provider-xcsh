---
page_title: "other_settings.logging_options"
subcategory: "Load Balancing"
description: "This defines various OPTIONS related to logging."
xcsh_docs: {"aliases": ["other settings logging options"], "body_bytes": 1257, "body_sha256": "sha256:f4ce38875b8480e339de53d5cb1e6907a461d662191b6080a7535a85c3530fe6", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "path": "documentation/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2223103101102133-1321201122230001-0323100123120022-2010212213323231-0212002021131321-2330330122301331-0201233222123130-0011222322220210", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["other_settings", "logging_options"], "schema_version": 1, "sections": [{"aliases": ["other settings logging options client log options"], "anchor": "section", "description": "List of headers to Log.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "logging_options", "client_log_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "other settings logging options origin log options", "upstream servers"], "anchor": "section", "description": "List of headers to Log.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "logging_options", "origin_log_options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines various OPTIONS related to logging.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
