---
page_title: "strict_schema_origin"
subcategory: "API Management"
description: "strict_schema_origin for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 911, "body_sha256": "sha256:42606ad906c58347c6ef057618e77967639dbc9181d8eb8817649ce97f80232d", "canonical_id": "xcsh-docs:resources:api_definition:properties:strict_schema_origin", "child_ids": [], "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:properties:strict_schema_origin", "parent_id": "xcsh-docs:resources:api_definition:reference", "path": "docs/guides/resources--api_definition--properties--strict_schema_origin.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["strict_schema_origin"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/properties/strict_schema_origin/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "strict_schema_origin for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# strict_schema_origin

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md)
- [Property reference](resources--api_definition--reference.md)
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

- [Property reference](resources--api_definition--reference.md)
- [xcsh_api_definition](../resources/api_definition.md)
