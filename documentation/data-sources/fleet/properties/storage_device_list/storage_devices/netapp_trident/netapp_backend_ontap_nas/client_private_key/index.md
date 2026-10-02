---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "storage device list storage devices netapp trident netapp backend ontap nas client private key", "upstream servers"], "body_bytes": 3180, "body_sha256": "sha256:e8fad66f34d80b5660461c24fac11530a03b2084dd3497ed267aee65bf165505", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key:blindfold_secret_info", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "path": "documentation/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0100130312032100-0131313301301121-2230100323023301-2031022012110133-0102122311030123-3121233311203313-3100002020303213-1202010013011213", "registry_path": "docs/guides/data-sources--fleet--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "client_private_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "client_private_key", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "client_private_key", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/clear_secret_info/): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/blindfold_secret_info/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/clear_secret_info/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
