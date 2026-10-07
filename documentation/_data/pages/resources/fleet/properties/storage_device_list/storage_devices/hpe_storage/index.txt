---
page_title: "storage_device_list.storage_devices.hpe_storage"
subcategory: ""
description: "Device configuration for HPE Storage."
xcsh_docs: {"aliases": ["storage device list storage devices hpe storage"], "body_bytes": 7279, "body_sha256": "sha256:693836c7ae168c87ecb5ac109c966d7a29192c91ee9291fdbf918bfd5f9e44c0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:password"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0331020220211030-1311313113332322-3022303221300221-1112213221132313-3213332132202230-0111213113301310-3031100032111013-1121203331210313", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--hpe_storage--api_server_port", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage:RequiredObjectAttributes:api_server_port,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "requires"}, {"anchor": "schema-storage_device_list--storage_devices--hpe_storage--username", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage:RequiredObjectAttributes:api_server_port,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "hpe_storage"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices hpe storage api server port"], "anchor": "schema-storage_device_list--storage_devices--hpe_storage--api_server_port", "description": "Enter Storage Server Port.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "api_server_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["storage device list storage devices hpe storage iscsi chap password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage.iscsi_chap_password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:iscsi_chap_password:clear_secret_info", "type": "conflicts"}], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_password"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices hpe storage iscsi chap user"], "anchor": "schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_user", "description": "Chap Username to connect to the HPE storage.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "iscsi_chap_user"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices hpe storage password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:password", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.hpe_storage.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "password"], "syntax": "block", "type": "object"}, {"aliases": ["storage device list storage devices hpe storage storage server ip address"], "anchor": "schema-storage_device_list--storage_devices--hpe_storage--storage_server_ip_address", "description": "Enter storage server IP address.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "storage_server_ip_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices hpe storage storage server name"], "anchor": "schema-storage_device_list--storage_devices--hpe_storage--storage_server_name", "description": "Enter storage server Name.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "storage_server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices hpe storage username"], "anchor": "schema-storage_device_list--storage_devices--hpe_storage--username", "description": "Username to connect to the HPE storage management IP.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:hpe_storage", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "hpe_storage", "username"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Device configuration for HPE Storage.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["fleetCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.hpe_storage

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- storage_device_list.storage_devices.hpe_storage

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for hpe storage.

Additional upstream details:

Device configuration for HPE Storage.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("api_server_port",
    "username")}
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
hpe_storage {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-storage_device_list--storage_devices--hpe_storage--api_server_port"></a>

### api_server_port property

Type: `"number"`. Optional.

Storage server Port. Enter Storage Server Port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [iscsi_chap_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/iscsi_chap_password/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--hpe_storage--iscsi_chap_user"></a>

### iscsi_chap_user property

Type: `"string"`. Optional.

Chap Username to connect to the HPE storage.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/hpe_storage/password/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--hpe_storage--storage_server_ip_address"></a>

### storage_server_ip_address property

Type: `"string"`. Optional.

Storage Server IP address. Enter storage server IP address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-storage_device_list--storage_devices--hpe_storage--storage_server_name"></a>

### storage_server_name property

Type: `"string"`. Optional.

Storage Server Name. Enter storage server Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="schema-storage_device_list--storage_devices--hpe_storage--username"></a>

### username property

Type: `"string"`. Optional.

Username to connect to the HPE storage management IP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```
