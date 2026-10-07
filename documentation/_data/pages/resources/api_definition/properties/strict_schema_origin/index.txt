---
page_title: "strict_schema_origin"
subcategory: "API Management"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "strict schema origin", "upstream servers"], "body_bytes": 963, "body_sha256": "sha256:76c6fdd01675fdfc38dd028bf5e66241579b42d2c19efed7b2391352e7328298", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:properties:strict_schema_origin", "parent_id": "xcsh-docs:resources:api_definition:reference", "path": "documentation/resources/api_definition/properties/strict_schema_origin/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2310302013203313-0231100001203113-1203121222003301-2311001121000010-0000132332111022-1101302100003131-0301013302332203-1033132331200113", "registry_path": "docs/guides/resources--api_definition--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["strict_schema_origin"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/properties/strict_schema_origin/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["api_definitionCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# strict_schema_origin

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/properties/)
- strict_schema_origin

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
strict_schema_origin = {}
```

This is an empty object or choice marker. It has no direct properties.
