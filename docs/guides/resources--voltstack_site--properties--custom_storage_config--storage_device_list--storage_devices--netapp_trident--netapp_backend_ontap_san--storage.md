---
page_title: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage"
subcategory: ""
description: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 4891, "body_sha256": "sha256:ab4d614723069c437e562deb2347253bcc7558767109652bc94172d153a1f68a", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage:volume_defaults"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/storage/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_device_list](resources--voltstack_site--properties--custom_storage_config--storage_device_list.md)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--zone"></a>

### zone property

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage--volume_defaults.md)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--properties--custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
