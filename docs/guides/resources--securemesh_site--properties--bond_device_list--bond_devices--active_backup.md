---
page_title: "bond_device_list.bond_devices.active_backup"
subcategory: ""
description: "bond_device_list.bond_devices.active_backup for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1211, "body_sha256": "sha256:d1ebd4b6679cb2f1019a93cc79b3e42de2835023503eff49eae0eb75a698c4db", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices:active_backup", "child_ids": [], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices:active_backup", "parent_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices", "path": "docs/guides/resources--securemesh_site--properties--bond_device_list--bond_devices--active_backup.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bond_device_list", "bond_devices", "active_backup"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/bond_device_list/bond_devices/active_backup/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bond_device_list.bond_devices.active_backup for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list.bond_devices.active_backup

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [bond_device_list](resources--securemesh_site--properties--bond_device_list.md)
- [bond_device_list.bond_devices](resources--securemesh_site--properties--bond_device_list--bond_devices.md)
- bond_device_list.bond_devices.active_backup

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [bond_device_list.bond_devices](resources--securemesh_site--properties--bond_device_list--bond_devices.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
