---
page_title: "bfd_disabled"
subcategory: "Infrastructure"
description: "bfd_disabled for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1196, "body_sha256": "sha256:f1134c1650b77cc71dd12491d0c7bbd82002c7e8587c89d5e098c46133dd0076", "canonical_id": "xcsh-docs:resources:site_mesh_group:properties:bfd_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:properties:bfd_disabled", "parent_id": "xcsh-docs:resources:site_mesh_group:reference", "path": "docs/guides/resources--site_mesh_group--properties--bfd_disabled.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bfd_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/properties/bfd_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bfd_disabled for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bfd_disabled

Breadcrumbs:

- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
- [Property reference](resources--site_mesh_group--reference.md)
- bfd_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: bfd\_disabled, bfd\_enabled\] Enable this option

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

- [bfd_disabled](resources--site_mesh_group--properties--bfd_disabled.md#section)
- [bfd_enabled](resources--site_mesh_group--properties--bfd_enabled.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bfd_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--site_mesh_group--reference.md)
- [xcsh_site_mesh_group](../resources/site_mesh_group.md)
