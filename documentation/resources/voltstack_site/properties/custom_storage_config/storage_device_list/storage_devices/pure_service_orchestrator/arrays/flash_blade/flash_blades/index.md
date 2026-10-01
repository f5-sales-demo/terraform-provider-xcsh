---
page_title: "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades"
subcategory: ""
description: "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 11177, "body_sha256": "sha256:022e9c72720eb69ea484cced6bcb509111da87f023606944bfd23fd4389b4e2e", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades:api_token"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade:flash_blades", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_blade", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/index.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_blade", "flash_blades"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/)
- [custom_storage_config.storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

For FlashBlades you must set the 'mgmt\_endpoint', 'api\_token' and nfs\_endpoint.

Upstream description:

For FlashBlades you must set the "mgmt\_endpoint", "api\_token" and nfs\_endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip"),
  validators.ConflictingListObjectAttributes("nfs_endpoint_dns_name",
    "nfs_endpoint_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
flash_blades {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/): complete subsection reference.

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--mgmt_dns_name"></a>

### mgmt_dns_name property

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--mgmt_ip"></a>

### mgmt_ip property

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--nfs_endpoint_dns_name"></a>

### nfs_endpoint_dns_name property

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_blade--flash_blades--nfs_endpoint_ip"></a>

### nfs_endpoint_ip property

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/flash_blades/api_token/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_blade/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
