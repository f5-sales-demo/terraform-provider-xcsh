---
page_title: "custom_storage_config.storage_class_list.storage_classes"
subcategory: ""
description: "custom_storage_config.storage_class_list.storage_classes for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 9146, "body_sha256": "sha256:f81469b688345d583fb5619ecef9724d2f93828560415372c6fb1f169a224c69", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:custom_storage", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:hpe_storage", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:netapp_trident", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes:pure_service_orchestrator"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list:storage_classes", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_class_list", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["custom_storage_config", "storage_class_list", "storage_classes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_class_list.storage_classes for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_class_list.storage_classes

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/)
- custom_storage_config.storage_class_list.storage_classes

<a id="section"></a>

Type: `"list"`. Computed.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-custom_storage_config--storage_class_list--storage_classes--advanced_storage_parameters"></a>

### advanced_storage_parameters property

Type: `["map", "string"]`. Computed.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

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
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-custom_storage_config--storage_class_list--storage_classes--allow_volume_expansion"></a>

### allow_volume_expansion property

Type: `"bool"`. Computed.

Allow Volume Expansion. Allow volume expansion.

Upstream description:

Allow volume expansion.

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

- [custom_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/custom_storage/): complete subsection reference.

<a id="schema-custom_storage_config--storage_class_list--storage_classes--default_storage_class"></a>

### default_storage_class property

Type: `"bool"`. Computed.

Make this storage class default storage class for the K8s cluster.

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

<a id="schema-custom_storage_config--storage_class_list--storage_classes--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Storage Class Description. Description for this storage class.

- [hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/hpe_storage/): complete subsection reference.

- [netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/): complete subsection reference.

- [pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/pure_service_orchestrator/): complete subsection reference.

<a id="schema-custom_storage_config--storage_class_list--storage_classes--reclaim_policy"></a>

### reclaim_policy property

Type: `"string"`. Computed.

Policy configuration for this feature.

Upstream description:

Reclaim Policy.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="schema-custom_storage_config--storage_class_list--storage_classes--storage_class_name"></a>

### storage_class_name property

Type: `"string"`. Computed.

Name of the storage class as it will appear in K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="schema-custom_storage_config--storage_class_list--storage_classes--storage_device"></a>

### storage_device property

Type: `"string"`. Computed.

Storage device that this class will use. The Device name defined at previous step.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

## Next pages

- [custom_storage_config.storage_class_list.storage_classes.custom_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/custom_storage/)
- [custom_storage_config.storage_class_list.storage_classes.hpe_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/hpe_storage/)
- [custom_storage_config.storage_class_list.storage_classes.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/netapp_trident/)
- [custom_storage_config.storage_class_list.storage_classes.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/storage_classes/pure_service_orchestrator/)
- [custom_storage_config.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_class_list/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
