---
page_title: "custom_storage_config"
subcategory: ""
description: "custom_storage_config for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 4837, "body_sha256": "sha256:97db65dea1166afa4b0c2ac1b3e4b4335739202caa3f27ccd4f5993cae2404f5", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:default_storage_class", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:no_static_routes", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:no_storage_device", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:no_storage_interfaces", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:static_routes", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["custom_storage_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
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

- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/#section)
- [default_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/default_storage_config/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [default_storage_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/default_storage_class/): complete subsection reference.

- [no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/no_static_routes/): complete subsection reference.

- [no_storage_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/no_storage_device/): complete subsection reference.

- [no_storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/no_storage_interfaces/): complete subsection reference.

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/static_routes/): complete subsection reference.

- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/): complete subsection reference.

- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_device_list/): complete subsection reference.

- [storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/): complete subsection reference.

## Next pages

- [custom_storage_config.default_storage_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/default_storage_class/)
- [custom_storage_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/no_static_routes/)
- [custom_storage_config.no_storage_device](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/no_storage_device/)
- [custom_storage_config.no_storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/no_storage_interfaces/)
- [custom_storage_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/static_routes/)
- [custom_storage_config.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/)
- [custom_storage_config.storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_device_list/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
