---
page_title: "custom_storage_config"
subcategory: ""
description: "custom_storage_config for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 4250, "body_sha256": "sha256:d9ef485bcfb1d6b0f27809e85bf1fffe61e05a114d9664d0221dcb1bfbadf7a0", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:default_storage_class", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:no_static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:no_storage_device", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:no_storage_interfaces", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_class_list", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config", "parent_id": "xcsh-docs:resources:voltstack_site:reference", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- custom_storage_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_storage\_config, default\_storage\_config; Default: default\_storage\_config\]
VssStorageConfiguration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_storage_class",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_storage_device",
    "storage_device_list"),
  validators.ConflictingObjectAttributes("no_storage_interfaces",
    "storage_interface_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage_class\",\"storage_class_list\"]",
  "x-ves-oneof-field-storage_device_choice": "[\"no_storage_device\",\"storage_device_list\"]",
  "x-ves-oneof-field-storage_interface_choice": "[\"no_storage_interfaces\",\"storage_interface_list\"]"
}
```

OneOf alternatives in this subsection:

- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md#section)
- [default_storage_config](resources--voltstack_site--properties--default_storage_config.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_storage_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_storage_class](resources--voltstack_site--properties--custom_storage_config--default_storage_class.md): complete subsection reference.

- [no_static_routes](resources--voltstack_site--properties--custom_storage_config--no_static_routes.md): complete subsection reference.

- [no_storage_device](resources--voltstack_site--properties--custom_storage_config--no_storage_device.md): complete subsection reference.

- [no_storage_interfaces](resources--voltstack_site--properties--custom_storage_config--no_storage_interfaces.md): complete subsection reference.

- [static_routes](resources--voltstack_site--properties--custom_storage_config--static_routes.md): complete subsection reference.

- [storage_class_list](resources--voltstack_site--properties--custom_storage_config--storage_class_list.md): complete subsection reference.

- [storage_device_list](resources--voltstack_site--properties--custom_storage_config--storage_device_list.md): complete subsection reference.

- [storage_interface_list](resources--voltstack_site--properties--custom_storage_config--storage_interface_list.md): complete subsection reference.

## Next pages

- [custom_storage_config.default_storage_class](resources--voltstack_site--properties--custom_storage_config--default_storage_class.md)
- [custom_storage_config.no_static_routes](resources--voltstack_site--properties--custom_storage_config--no_static_routes.md)
- [custom_storage_config.no_storage_device](resources--voltstack_site--properties--custom_storage_config--no_storage_device.md)
- [custom_storage_config.no_storage_interfaces](resources--voltstack_site--properties--custom_storage_config--no_storage_interfaces.md)
- [custom_storage_config.static_routes](resources--voltstack_site--properties--custom_storage_config--static_routes.md)
- [custom_storage_config.storage_class_list](resources--voltstack_site--properties--custom_storage_config--storage_class_list.md)
- [custom_storage_config.storage_device_list](resources--voltstack_site--properties--custom_storage_config--storage_device_list.md)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--properties--custom_storage_config--storage_interface_list.md)
- [Property reference](resources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
