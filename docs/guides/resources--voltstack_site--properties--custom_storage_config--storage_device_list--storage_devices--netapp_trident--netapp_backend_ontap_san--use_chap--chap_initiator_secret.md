---
page_title: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret"
subcategory: ""
description: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 3884, "body_sha256": "sha256:d31b64d2b65d3ad1a69ee095afbe96d4c1b55401dae98c695507d506284aea70", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:blindfold_secret_info", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret:clear_secret_info"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap", "chap_initiator_secret"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/chap_initiator_secret/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_device_list](resources--voltstack_site--properties--custom_storage_config--storage_device_list.md)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap.md)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
chap_initiator_secret {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--clear_secret_info.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--blindfold_secret_info.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret--clear_secret_info.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
