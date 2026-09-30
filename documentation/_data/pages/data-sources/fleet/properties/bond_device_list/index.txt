---
page_title: "bond_device_list"
subcategory: ""
description: "bond_device_list for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1610, "body_sha256": "sha256:007bb2dc703619c9ac4b8c448cf83a9ccca44a6ad87061355adf2e0ffc1a34bf", "child_ids": ["xcsh-docs:data-sources:fleet:properties:bond_device_list:bond_devices"], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:bond_device_list", "parent_id": "xcsh-docs:data-sources:fleet:reference", "path": "documentation/data-sources/fleet/properties/bond_device_list/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["bond_device_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/bond_device_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bond_device_list for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bond_device_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- bond_device_list

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Upstream description:

List of bond devices for this fleet.

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

- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/bond_device_list/#section)
- [no_bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/no_bond_devices/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/bond_device_list/bond_devices/): complete subsection reference.

## Next pages

- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/bond_device_list/bond_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
