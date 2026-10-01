---
page_title: "device_list"
subcategory: ""
description: "device_list for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 947, "body_sha256": "sha256:673d7435b3120de87ea6a646888301acfcff096c9679d9d3c1a275adfb9b002c", "canonical_id": "xcsh-docs:resources:fleet:properties:device_list", "child_ids": ["xcsh-docs:resources:fleet:properties:device_list:devices"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:device_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--device_list.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["device_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/device_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "device_list for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# device_list

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- device_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add device for all interfaces belonging to this fleet.

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
device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [devices](resources--fleet--properties--device_list--devices.md): complete subsection reference.

## Next pages

- [device_list.devices](resources--fleet--properties--device_list--devices.md)
- [Property reference](resources--fleet--reference.md)
- [xcsh_fleet](../resources/fleet.md)
