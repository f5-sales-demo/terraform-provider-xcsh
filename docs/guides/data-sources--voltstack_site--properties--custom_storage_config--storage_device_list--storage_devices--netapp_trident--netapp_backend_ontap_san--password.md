---
page_title: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password"
subcategory: ""
description: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3140, "body_sha256": "sha256:84fdc4affde32040d4472613ec274b63771a44d10403f42df89213e28b7d77cd", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:password", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:password:blindfold_secret_info", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:password:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:password", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "path": "docs/guides/data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "password"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/password/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_storage_config](data-sources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list.md)
- [custom_storage_config.storage_device_list.storage_devices](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--clear_secret_info.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--blindfold_secret_info.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--password--clear_secret_info.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](data-sources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
