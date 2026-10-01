---
page_title: "storage_device_list"
subcategory: ""
description: "storage_device_list for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1040, "body_sha256": "sha256:36806a8e36f76607e2783cb9485cafccb49fbdef4467248a7b471eb1fe54705d", "canonical_id": "xcsh-docs:resources:fleet:properties:storage_device_list", "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--storage_device_list.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["storage_device_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "storage_device_list for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- storage_device_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this fleet.

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
storage_device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices](resources--fleet--properties--storage_device_list--storage_devices.md)
- [Property reference](resources--fleet--reference.md)
- [xcsh_fleet](../resources/fleet.md)
