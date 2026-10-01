---
page_title: "custom_storage_config"
subcategory: ""
description: "custom_storage_config for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3727, "body_sha256": "sha256:b4fe9546ed2950749a82440c62ff87545390aa5c6a355177f3fff351c3bb2437", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:default_storage_class", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:no_static_routes", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:no_storage_device", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:no_storage_interfaces", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:static_routes", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "docs/guides/data-sources--voltstack_site--properties--custom_storage_config.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- custom_storage_config

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_storage\_config, default\_storage\_config; Default: default\_storage\_config\]
VssStorageConfiguration.

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

- [custom_storage_config](data-sources--voltstack_site--properties--custom_storage_config.md#section)
- [default_storage_config](data-sources--voltstack_site--properties--default_storage_config.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [default_storage_class](data-sources--voltstack_site--properties--custom_storage_config--default_storage_class.md): complete subsection reference.

- [no_static_routes](data-sources--voltstack_site--properties--custom_storage_config--no_static_routes.md): complete subsection reference.

- [no_storage_device](data-sources--voltstack_site--properties--custom_storage_config--no_storage_device.md): complete subsection reference.

- [no_storage_interfaces](data-sources--voltstack_site--properties--custom_storage_config--no_storage_interfaces.md): complete subsection reference.

- [static_routes](data-sources--voltstack_site--properties--custom_storage_config--static_routes.md): complete subsection reference.

- [storage_class_list](data-sources--voltstack_site--properties--custom_storage_config--storage_class_list.md): complete subsection reference.

- [storage_device_list](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list.md): complete subsection reference.

- [storage_interface_list](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list.md): complete subsection reference.

## Next pages

- [custom_storage_config.default_storage_class](data-sources--voltstack_site--properties--custom_storage_config--default_storage_class.md)
- [custom_storage_config.no_static_routes](data-sources--voltstack_site--properties--custom_storage_config--no_static_routes.md)
- [custom_storage_config.no_storage_device](data-sources--voltstack_site--properties--custom_storage_config--no_storage_device.md)
- [custom_storage_config.no_storage_interfaces](data-sources--voltstack_site--properties--custom_storage_config--no_storage_interfaces.md)
- [custom_storage_config.static_routes](data-sources--voltstack_site--properties--custom_storage_config--static_routes.md)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--properties--custom_storage_config--storage_class_list.md)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list.md)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
