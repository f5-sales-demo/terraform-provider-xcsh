---
page_title: "mixed_schema_origin"
subcategory: "API Management"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "mixed schema origin", "origin servers", "upstream servers"], "body_bytes": 1294, "body_sha256": "sha256:0ae0a302296b0beba99feaf0dc47751793c040a206161c31971ee91463d2e0c6", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_definition:properties:mixed_schema_origin", "parent_id": "xcsh-docs:data-sources:api_definition:reference", "path": "documentation/data-sources/api_definition/properties/mixed_schema_origin/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-1113312301230031-0230200300100013-1130332120030100-0010201203220110-1001312011120333-1130232132203320-3220122122131201-2330322330002012", "registry_path": "docs/guides/data-sources--api_definition--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["mixed_schema_origin"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_definition/properties/mixed_schema_origin/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["api_definitionCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mixed_schema_origin

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/)
- mixed_schema_origin

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: mixed\_schema\_origin, strict\_schema\_origin\] Configuration parameter for mixed schema
origin.

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

- [mixed_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/mixed_schema_origin/#section)
- [strict_schema_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/properties/strict_schema_origin/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
