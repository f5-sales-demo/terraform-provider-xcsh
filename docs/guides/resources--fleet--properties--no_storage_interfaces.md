---
page_title: "no_storage_interfaces"
subcategory: ""
description: "no_storage_interfaces for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1271, "body_sha256": "sha256:aa0acca60620d891b596ca1454a47b9901349347f04753b7002540a8fd6ffb7e", "canonical_id": "xcsh-docs:resources:fleet:properties:no_storage_interfaces", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:no_storage_interfaces", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--no_storage_interfaces.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["no_storage_interfaces"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/no_storage_interfaces/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "no_storage_interfaces for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# no_storage_interfaces

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- no_storage_interfaces

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_interfaces, storage\_interface\_list; Default: no\_storage\_interfaces\]
Configuration parameter for no storage interfaces.

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

- [no_storage_interfaces](resources--fleet--properties--no_storage_interfaces.md#section)
- [storage_interface_list](resources--fleet--properties--storage_interface_list.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_interfaces = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--fleet--reference.md)
- [xcsh_fleet](../resources/fleet.md)
