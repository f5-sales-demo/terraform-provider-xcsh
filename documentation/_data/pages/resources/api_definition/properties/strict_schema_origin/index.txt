---
page_title: "strict_schema_origin"
subcategory: "API Management"
description: "strict_schema_origin for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 1218, "body_sha256": "sha256:b2e30a4d7e84bc31e054b2fdba120fedfb0e99f6de26d84c4c3968d6017e55b3", "child_ids": [], "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:properties:strict_schema_origin", "parent_id": "xcsh-docs:resources:api_definition:reference", "path": "documentation/resources/api_definition/properties/strict_schema_origin/index.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["strict_schema_origin"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/properties/strict_schema_origin/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "strict_schema_origin for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
strict_schema_origin = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/properties/)
- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/)
