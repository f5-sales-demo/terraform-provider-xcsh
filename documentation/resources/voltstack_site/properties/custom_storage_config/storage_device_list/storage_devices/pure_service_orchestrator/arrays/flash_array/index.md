---
page_title: "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array"
subcategory: ""
description: "Specify what storage flash arrays should be managed the plugin."
xcsh_docs: {"aliases": ["custom storage config storage device list storage devices pure service orchestrator arrays flash array"], "body_bytes": 9950, "body_sha256": "sha256:e3d4075bf1a74f2df4c66270ca52eccfa6c909661424462e81145f512e790b04", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321", "registry_path": "docs/guides/resources--voltstack_site--reference--group-007.md", "relationships": [{"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_fs_type", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array:RequiredObjectAttributes:default_fs_type,flash_arrays,iscsi_login_timeout,san_type", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "type": "requires"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--iscsi_login_timeout", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array:RequiredObjectAttributes:default_fs_type,flash_arrays,iscsi_login_timeout,san_type", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "type": "requires"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--san_type", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array:RequiredObjectAttributes:default_fs_type,flash_arrays,iscsi_login_timeout,san_type", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array:RequiredObjectAttributes:default_fs_type,flash_arrays,iscsi_login_timeout,san_type", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array"], "schema_version": 1, "sections": [{"aliases": ["default fs opt"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_fs_opt", "description": "Block volume default mkfs OPTIONS. Not recommended to change!", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "default_fs_opt"], "syntax": "attribute", "type": "string"}, {"aliases": ["default fs type"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_fs_type", "description": "Block volume default filesystem type. Not recommended to change!", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "default_fs_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["default mount opts"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_mount_opts", "description": "Block volume default filesystem mount OPTIONS. Not recommended to change!", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "default_mount_opts"], "syntax": "attribute", "type": "list"}, {"aliases": ["disable preempt attachments"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--disable_preempt_attachments", "description": "Enable/Disable attachment preemption!", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "disable_preempt_attachments"], "syntax": "attribute", "type": "bool"}, {"aliases": ["flash arrays"], "anchor": "section", "description": "For FlashArrays you must set the \"mgmt_endpoint\" and \"api_token\"", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--mgmt_dns_name", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays:ConflictingListObjectAttributes:mgmt_dns_name,mgmt_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--flash_arrays--mgmt_ip", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays:ConflictingListObjectAttributes:mgmt_dns_name,mgmt_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array:flash_arrays", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "flash_arrays"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "iscsi login timeout", "login", "login result", "operation timeout", "sign in"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--iscsi_login_timeout", "description": "ISCSI login timeout in seconds. Not recommended to change!", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "iscsi_login_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["san type"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--san_type", "description": "Block volume access protocol, either ISCSI or FC.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:pure_service_orchestrator:arrays:flash_array", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "pure_service_orchestrator", "arrays", "flash_array", "san_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specify what storage flash arrays should be managed the plugin.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/)
- [custom_storage_config.storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash arrays should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("default_fs_type",
    "flash_arrays",
    "iscsi_login_timeout",
    "san_type")}
```

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
flash_array {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_fs_opt"></a>

### default_fs_opt property

Type: `"string"`. Optional.

Block volume default mkfs OPTIONS. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_fs_type"></a>

### default_fs_type property

Type: `"string"`. Optional.

\[Enum: xfs|ext4\] Block volume default filesystem type. Not recommended to change!. Possible values
are \`xfs\`, \`ext4\`.

Upstream description:

Block volume default filesystem type. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("xfs",
    "ext4"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "xfs",
    "ext4"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--default_mount_opts"></a>

### default_mount_opts property

Type: `["list", "string"]`. Optional.

Block volume default filesystem mount OPTIONS. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--disable_preempt_attachments"></a>

### disable_preempt_attachments property

Type: `"bool"`. Optional.

Disable Preempt Attachments. Enable/Disable attachment preemption!

Upstream description:

Enable/Disable attachment preemption!

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

- [flash_arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/flash_arrays/): complete subsection reference.

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--iscsi_login_timeout"></a>

### iscsi_login_timeout property

Type: `"number"`. Optional.

ISCSI login timeout in seconds. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--pure_service_orchestrator--arrays--flash_array--san_type"></a>

### san_type property

Type: `"string"`. Optional.

\[Enum: ISCSI|FC\] Block volume access protocol, either ISCSI or FC. Possible values are \`ISCSI\`,
\`FC\`.

Upstream description:

Block volume access protocol, either ISCSI or FC.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ISCSI",
    "FC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ISCSI",
    "FC"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  }
}
```

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/flash_array/flash_arrays/)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/pure_service_orchestrator/arrays/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
