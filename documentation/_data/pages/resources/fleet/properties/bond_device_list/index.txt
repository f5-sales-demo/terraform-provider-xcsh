---
page_title: "bond_device_list"
subcategory: ""
description: "bond_device_list for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1961, "body_sha256": "sha256:2b88fe5ced612d7a39112d50477740ee0c44ad23a01acb707e2cb1101e9311e5", "child_ids": ["xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:bond_device_list", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/bond_device_list/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["bond_device_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/bond_device_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bond_device_list for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- bond_device_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Upstream description:

List of bond devices for this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("bond_devices")}
```

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

- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/#section)
- [no_bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/no_bond_devices/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/): complete subsection reference.

## Next pages

- [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
