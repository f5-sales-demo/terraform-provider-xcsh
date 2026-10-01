---
page_title: "disabled"
subcategory: ""
description: "disabled for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 1118, "body_sha256": "sha256:cdc36c597ddaeb94155a4c4831ac1826938e0984c2fed8c01c91dc55f6c90837", "canonical_id": "xcsh-docs:resources:cloud_link:properties:disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:disabled", "parent_id": "xcsh-docs:resources:cloud_link:reference", "path": "docs/guides/resources--cloud_link--properties--disabled.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disabled for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disabled

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md)
- [Property reference](resources--cloud_link--reference.md)
- disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, enabled\] Enable this option

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

- [disabled](resources--cloud_link--properties--disabled.md#section)
- [enabled](resources--cloud_link--properties--enabled.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cloud_link--reference.md)
- [xcsh_cloud_link](../resources/cloud_link.md)
