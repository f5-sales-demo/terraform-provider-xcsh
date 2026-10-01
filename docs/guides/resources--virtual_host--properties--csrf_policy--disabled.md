---
page_title: "csrf_policy.disabled"
subcategory: ""
description: "csrf_policy.disabled for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 959, "body_sha256": "sha256:8c8022495f2f8f156453a4d5edda1f5a6e071c32305b400ebd4c008c2cfa36f1", "canonical_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:disabled", "parent_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy", "path": "docs/guides/resources--virtual_host--properties--csrf_policy--disabled.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["csrf_policy", "disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/csrf_policy/disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "csrf_policy.disabled for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# csrf_policy.disabled

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [csrf_policy](resources--virtual_host--properties--csrf_policy.md)
- csrf_policy.disabled

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
disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [csrf_policy](resources--virtual_host--properties--csrf_policy.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
