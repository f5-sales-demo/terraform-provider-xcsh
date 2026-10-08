---
page_title: "strict_schema_origin"
subcategory: "API Management"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "strict schema origin", "upstream servers"], "body_bytes": 896, "body_sha256": "sha256:2dfa3da24616d3a25ed99b42fd362c60af173600aa9c3c66b21763a1899d89e0", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_definition:properties:strict_schema_origin", "parent_id": "xcsh-docs:data-sources:api_definition:reference", "path": "documentation/data-sources/api_definition/properties/strict_schema_origin/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1303030311022303-2032313222313311-1221131323011003-1031102310113210-2313022033101221-0001300200113222-2020213030301212-2210322332103133", "registry_path": "docs/guides/data-sources--api_definition--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["strict_schema_origin"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_definition/properties/strict_schema_origin/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["api_definitionCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# strict_schema_origin

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/)
- strict_schema_origin

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict schema origin. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

This is an empty object or choice marker. It has no direct properties.
