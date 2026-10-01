---
page_title: "hard_purge"
subcategory: ""
description: "hard_purge for xcsh_cdn_purge_command."
xcsh_docs: {"aliases": [], "body_bytes": 1197, "body_sha256": "sha256:618d00bc868186e978b6b4b790a35316b134d9e1b6467ca8a1de6d85e15c7508", "canonical_id": "xcsh-docs:resources:cdn_purge_command:properties:hard_purge", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_purge_command:properties:hard_purge", "parent_id": "xcsh-docs:resources:cdn_purge_command:reference", "path": "docs/guides/resources--cdn_purge_command--properties--hard_purge.md", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["hard_purge"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_purge_command/properties/hard_purge/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "hard_purge for xcsh_cdn_purge_command.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# hard_purge

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md)
- [Property reference](resources--cdn_purge_command--reference.md)
- hard_purge

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: hard\_purge, soft\_purge\] Enable this option

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

- [hard_purge](resources--cdn_purge_command--properties--hard_purge.md#section)
- [soft_purge](resources--cdn_purge_command--properties--soft_purge.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
hard_purge = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--cdn_purge_command--reference.md)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md)
