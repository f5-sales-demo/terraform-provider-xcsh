---
page_title: "disable_re_fallback"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable re fallback"], "body_bytes": 1375, "body_sha256": "sha256:e12880d45a11cd9cc08efdcfb34bd0b8bc4e425c3ff60a98f538927412072d35", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:disable_re_fallback", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "documentation/resources/site_mesh_group/properties/disable_re_fallback/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-2300102311310103-1132033112330111-0232101032131033-2010201232322220-2230101212022313-0030303231303133-2111320001222103-2322220330311130", "registry_path": "docs/guides/resources--site_mesh_group--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_re_fallback"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/disable_re_fallback/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_re_fallback

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
- disable_re_fallback

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_re\_fallback, enable\_re\_fallback; Default: disable\_re\_fallback\] Configuration
parameter for disable re fallback.

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

- [disable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/disable_re_fallback/#section)
- [enable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/enable_re_fallback/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_re_fallback = {}
```

This is an empty object or choice marker. It has no direct properties.
