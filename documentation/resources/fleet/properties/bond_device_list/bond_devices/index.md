---
page_title: "bond_device_list.bond_devices"
subcategory: ""
description: "List of bond devices."
xcsh_docs: {"aliases": ["bond device list bond devices"], "body_bytes": 8338, "body_sha256": "sha256:f2612b144c3b60c014d088bd7b6a6eb78380f4ec2c105c3dc965e66e46d48bd8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:active_backup", "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:lacp"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "parent_id": "xcsh-docs:resources:fleet:properties:bond_device_list", "path": "documentation/resources/fleet/properties/bond_device_list/bond_devices/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:active_backup", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:ConflictingListObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:lacp", "type": "conflicts"}, {"anchor": "schema-bond_device_list--bond_devices--devices", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_polling_interval", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--link_up_delay", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}, {"anchor": "schema-bond_device_list--bond_devices--name", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices:RequiredListObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bond_device_list", "bond_devices"], "schema_version": 1, "sections": [{"aliases": ["active backup"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:active_backup", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "active_backup"], "syntax": "attribute", "type": "object"}, {"aliases": ["devices"], "anchor": "schema-bond_device_list--bond_devices--devices", "description": "Ethernet devices that will make up this bond.", "document_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "devices"], "syntax": "attribute", "type": "list"}, {"aliases": ["lacp"], "anchor": "section", "description": "LACP parameters for the bond device.", "document_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:lacp", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bond_device_list--bond_devices--lacp--rate", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices.lacp:RequiredObjectAttributes:rate", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices:lacp", "type": "requires"}], "schema_path": ["bond_device_list", "bond_devices", "lacp"], "syntax": "block", "type": "object"}, {"aliases": ["link polling interval"], "anchor": "schema-bond_device_list--bond_devices--link_polling_interval", "description": "Link polling interval in milliseconds.", "document_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "link_polling_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["link up delay"], "anchor": "schema-bond_device_list--bond_devices--link_up_delay", "description": "Milliseconds wait before link is declared up.", "document_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "link_up_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["name"], "anchor": "schema-bond_device_list--bond_devices--name", "description": "Name for the Bond. Ex 'bond0'", "document_id": "xcsh-docs:resources:fleet:properties:bond_device_list:bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/bond_device_list/bond_devices/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of bond devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list.bond_devices

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/)
- bond_device_list.bond_devices

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Bond Devices. List of bond devices.

Upstream description:

List of bond devices.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingListObjectAttributes("active_backup",
    "lacp")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
bond_devices {
  # Configure direct properties listed below.
}
```

## Direct properties

- [active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/active_backup/): complete subsection reference.

<a id="schema-bond_device_list--bond_devices--devices"></a>

### devices property

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
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

- [lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/lacp/): complete subsection reference.

<a id="schema-bond_device_list--bond_devices--link_polling_interval"></a>

### link_polling_interval property

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-bond_device_list--bond_devices--link_up_delay"></a>

### link_up_delay property

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-bond_device_list--bond_devices--name"></a>

### name property

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [bond_device_list.bond_devices.active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/active_backup/)
- [bond_device_list.bond_devices.lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/bond_devices/lacp/)
- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/bond_device_list/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
