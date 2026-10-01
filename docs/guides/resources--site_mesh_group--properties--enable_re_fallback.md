---
page_title: "enable_re_fallback"
subcategory: "Infrastructure"
description: "enable_re_fallback for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 936, "body_sha256": "sha256:efeaa70399ac2e3e935a3f54051757d4b80164a4b309622a90e3e5c169718823", "canonical_id": "xcsh-docs:resources:site_mesh_group:properties:enable_re_fallback", "child_ids": [], "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:enable_re_fallback", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "docs/guides/resources--site_mesh_group--properties--enable_re_fallback.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["enable_re_fallback"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/enable_re_fallback/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "enable_re_fallback for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_re_fallback

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
- [Property reference](resources--site_mesh_group--reference.md)
- enable_re_fallback

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable re fallback.

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
enable_re_fallback = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--site_mesh_group--reference.md)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
