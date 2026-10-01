---
page_title: "disable_advanced_delivery"
subcategory: ""
description: "disable_advanced_delivery for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1413, "body_sha256": "sha256:71fa7ae0489e1b1aec8b6d9f1882ad592a040dc1c395420c6cffe2630425aa16", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:disable_advanced_delivery", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:disable_advanced_delivery", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "docs/guides/resources--securemesh_site_v2--properties--disable_advanced_delivery.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["disable_advanced_delivery"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/disable_advanced_delivery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_advanced_delivery for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_advanced_delivery

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- disable_advanced_delivery

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_advanced\_delivery, enable\_advanced\_delivery; Default:
disable\_advanced\_delivery\] Configuration parameter for disable advanced delivery.

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

- [disable_advanced_delivery](resources--securemesh_site_v2--properties--disable_advanced_delivery.md#section)
- [enable_advanced_delivery](resources--securemesh_site_v2--properties--enable_advanced_delivery.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_advanced_delivery = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
