---
page_title: "other_settings"
subcategory: "Load Balancing"
description: "Other Settings."
xcsh_docs: {"aliases": ["other settings"], "body_bytes": 1435, "body_sha256": "sha256:125e74809c2ec70352c6d6629e4c95fd31824caf50214e0051508e768ba9311f", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options", "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/other_settings/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["other_settings"], "schema_version": 1, "sections": [{"aliases": ["other settings add location"], "anchor": "schema-other_settings--add_location", "description": "X-example: true Appends header x-F5 Distributed Cloud-location = <RE-site-name> in responses.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["other_settings", "add_location"], "syntax": "attribute", "type": "bool"}, {"aliases": ["other settings header options"], "anchor": "section", "description": "This defines various OPTIONS related to request/response headers.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:header_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "header_options"], "syntax": "attribute", "type": "object"}, {"aliases": ["other settings logging options"], "anchor": "section", "description": "This defines various OPTIONS related to logging.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:other_settings:logging_options", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "logging_options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/other_settings/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Other Settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- other_settings

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for other settings.

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

<a id="schema-other_settings--add_location"></a>

### add_location property

Type: `"bool"`. Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses.

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

- [header_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/header_options/): complete subsection reference.

- [logging_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/other_settings/logging_options/): complete subsection reference.
