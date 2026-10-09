---
page_title: "vmware.not_managed.node_list.interface_list.bond_interface"
subcategory: ""
description: "Bond devices configuration for fleet."
xcsh_docs: {"aliases": ["vmware not managed node list interface list bond interface"], "body_bytes": 7879, "body_sha256": "sha256:5d9c36123ed88cab36c7d36d5efc6c702348303e82106218b6aeb28838634f22", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface:active_backup", "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface:lacp"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/bond_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-018.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list.bond_interface:ConflictingObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface:active_backup", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list.bond_interface:ConflictingObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface:lacp", "type": "conflicts"}, {"anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--devices", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list.bond_interface:RequiredObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "type": "requires"}, {"anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--link_polling_interval", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list.bond_interface:RequiredObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "type": "requires"}, {"anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--link_up_delay", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list.bond_interface:RequiredObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "type": "requires"}, {"anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--name", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list.bond_interface:RequiredObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "bond_interface"], "schema_version": 1, "sections": [{"aliases": ["vmware not managed node list interface list bond interface active backup"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface:active_backup", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "bond_interface", "active_backup"], "syntax": "attribute", "type": "object"}, {"aliases": ["vmware not managed node list interface list bond interface devices"], "anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--devices", "description": "Ethernet devices that will make up this bond.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "bond_interface", "devices"], "syntax": "attribute", "type": "list"}, {"aliases": ["vmware not managed node list interface list bond interface lacp"], "anchor": "section", "description": "LACP parameters for the bond device.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface:lacp", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--lacp--rate", "enforcement": "provider-schema", "group": "vmware.not_managed.node_list.interface_list.bond_interface.lacp:RequiredObjectAttributes:rate", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface:lacp", "type": "requires"}], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "bond_interface", "lacp"], "syntax": "block", "type": "object"}, {"aliases": ["vmware not managed node list interface list bond interface link polling interval"], "anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--link_polling_interval", "description": "Link polling interval in milliseconds.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "bond_interface", "link_polling_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["vmware not managed node list interface list bond interface link up delay"], "anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--link_up_delay", "description": "Milliseconds wait before link is declared up.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "bond_interface", "link_up_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["vmware not managed node list interface list bond interface name"], "anchor": "schema-vmware--not_managed--node_list--interface_list--bond_interface--name", "description": "Name for the Bond. Ex 'bond0'", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "bond_interface", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/bond_interface/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Bond devices configuration for fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vmware.not_managed.node_list.interface_list.bond_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [vmware](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/)
- [vmware.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/)
- [vmware.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/)
- [vmware.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/)
- vmware.not_managed.node_list.interface_list.bond_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/bond_interface/active_backup/): complete subsection reference.

<a id="schema-vmware--not_managed--node_list--interface_list--bond_interface--devices"></a>

### devices property

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/bond_interface/lacp/): complete subsection reference.

<a id="schema-vmware--not_managed--node_list--interface_list--bond_interface--link_polling_interval"></a>

### link_polling_interval property

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="schema-vmware--not_managed--node_list--interface_list--bond_interface--link_up_delay"></a>

### link_up_delay property

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="schema-vmware--not_managed--node_list--interface_list--bond_interface--name"></a>

### name property

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```
