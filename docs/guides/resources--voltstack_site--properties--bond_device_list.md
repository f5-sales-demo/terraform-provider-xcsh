---
page_title: "bond_device_list"
subcategory: ""
description: "bond_device_list for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 1641, "body_sha256": "sha256:57a8291dbbb8305ac65e7f55909fa112a1ce39bbee06858713b8841e678e9735", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "docs/guides/resources--voltstack_site--properties--bond_device_list.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bond_device_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/bond_device_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bond_device_list for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
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

- [bond_device_list](resources--voltstack_site--properties--bond_device_list.md#section)
- [no_bond_devices](resources--voltstack_site--properties--no_bond_devices.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bond_devices](resources--voltstack_site--properties--bond_device_list--bond_devices.md): complete subsection reference.

## Next pages

- [bond_device_list.bond_devices](resources--voltstack_site--properties--bond_device_list--bond_devices.md)
- [Property reference](resources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
