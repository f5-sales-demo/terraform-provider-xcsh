---
page_title: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap"
subcategory: ""
description: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 5383, "body_sha256": "sha256:cae7b7691930ac039f3372eb33b0ee2f195c304aefb2b5a27aea2e5aa8f88b63", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_initiator_secret", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap:chap_target_initiator_secret"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_device_list](resources--voltstack_site--properties--custom_storage_config--storage_device_list.md)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.

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

Terraform syntax:

```terraform
use_chap {
  # Configure direct properties listed below.
}
```

## Direct properties

- [chap_initiator_secret](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret.md): complete subsection reference.

- [chap_target_initiator_secret](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_username"></a>

### chap_target_username property

Type: `"string"`. Optional.

Target username. Required if useCHAP=true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_username"></a>

### chap_username property

Type: `"string"`. Optional.

Inbound username. Required if useCHAP=true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_initiator_secret.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--use_chap--chap_target_initiator_secret.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
