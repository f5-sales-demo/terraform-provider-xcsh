---
page_title: "mixed_schema_origin"
subcategory: "API Management"
description: "mixed_schema_origin for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 1289, "body_sha256": "sha256:f8e780981649d2b97aa81649c362c71c17a11e25722971dc7d91c23b805f5397", "canonical_id": "xcsh-docs:resources:api_definition:properties:mixed_schema_origin", "child_ids": [], "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:properties:mixed_schema_origin", "parent_id": "xcsh-docs:resources:api_definition:reference", "path": "docs/guides/resources--api_definition--properties--mixed_schema_origin.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mixed_schema_origin"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/properties/mixed_schema_origin/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mixed_schema_origin for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mixed_schema_origin

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md)
- [Property reference](resources--api_definition--reference.md)
- mixed_schema_origin

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: mixed\_schema\_origin, strict\_schema\_origin\] Configuration parameter for mixed schema
origin.

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

OneOf alternatives in this subsection:

- [mixed_schema_origin](resources--api_definition--properties--mixed_schema_origin.md#section)
- [strict_schema_origin](resources--api_definition--properties--strict_schema_origin.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
mixed_schema_origin = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--api_definition--reference.md)
- [xcsh_api_definition](../resources/api_definition.md)
