---
page_title: "rules.client_region"
subcategory: ""
description: "rules.client_region for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 987, "body_sha256": "sha256:a6a8220bfbf3daa993e36a910dc135f87fda33bbe0c90e3d8fb777b635127833", "canonical_id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "child_ids": [], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:properties:rules:client_region", "parent_id": "xcsh-docs:resources:user_identification:properties:rules", "path": "docs/guides/resources--user_identification--properties--rules--client_region.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "client_region"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/properties/rules/client_region/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.client_region for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.client_region

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md)
- [Property reference](resources--user_identification--reference.md)
- [rules](resources--user_identification--properties--rules.md)
- rules.client_region

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
client_region = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules](resources--user_identification--properties--rules.md)
- [xcsh_user_identification](../resources/user_identification.md)
