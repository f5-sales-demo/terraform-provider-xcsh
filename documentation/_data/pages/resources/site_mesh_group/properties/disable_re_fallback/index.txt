---
page_title: "disable_re_fallback"
subcategory: "Infrastructure"
description: "disable_re_fallback for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1633, "body_sha256": "sha256:359db0fa120c7f45c54c83fbb8e54707b3ae82a040d60f41b6d585f67ab71fce", "child_ids": [], "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:disable_re_fallback", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "documentation/resources/site_mesh_group/properties/disable_re_fallback/index.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["disable_re_fallback"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/disable_re_fallback/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "disable_re_fallback for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [disable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/disable_re_fallback/#section)
- [enable_re_fallback](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/enable_re_fallback/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_re_fallback = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
