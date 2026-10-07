---
page_title: "cdn_loadbalancer"
subcategory: ""
description: "Set the scope of the API Group to a specific CDN Loadbalancer."
xcsh_docs: {"aliases": ["cdn loadbalancer"], "body_bytes": 875, "body_sha256": "sha256:5e6343ae70fa9ed2c8dea26f3a01b510b2c1719d83f17b98a404dd819a52f2b9", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_api_group:properties:cdn_loadbalancer", "parent_id": "xcsh-docs:data-sources:app_api_group:reference", "path": "documentation/data-sources/app_api_group/properties/cdn_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1122213102230112-1112022003110130-3220331113222231-1102000331113300-0311030101111100-2010310330333033-3103121011001321-3320100122100111", "registry_path": "docs/guides/data-sources--app_api_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cdn_loadbalancer"], "schema_version": 1, "sections": [{"aliases": ["cdn loadbalancer cdn loadbalancer"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:app_api_group:properties:cdn_loadbalancer:cdn_loadbalancer", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cdn_loadbalancer", "cdn_loadbalancer"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/properties/cdn_loadbalancer/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Set the scope of the API Group to a specific CDN Loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cdn_loadbalancer

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/)
- cdn_loadbalancer

<a id="section"></a>

Type: `"single"`. Computed.

Set the scope of the API Group to a specific CDN Loadbalancer.

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

- [cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/properties/cdn_loadbalancer/cdn_loadbalancer/): complete subsection reference.
