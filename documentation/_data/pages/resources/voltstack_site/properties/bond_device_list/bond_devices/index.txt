---
page_title: "bond_device_list.bond_devices"
subcategory: ""
description: "List of bond devices."
xcsh_docs: {"aliases": ["bond device list bond devices"], "body_bytes": 8437, "body_sha256": "sha256:a4fe04c3422bd3c9e0d00841198fb0ae085cc1b95e73bb88158e4c68a36e50b2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices:active_backup", "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices:lacp"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "parent_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list", "path": "documentation/resources/voltstack_site/properties/bond_device_list/bond_devices/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3131101003030033-1101101112120022-2101210232301130-3233011101203111-0133323313111201-2121330200333033-1312112233000101-0113022100202020", "registry_path": "docs/guides/resources--voltstack_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bond_device_list", "bond_devices"], "schema_version": 1, "sections": [{"aliases": ["active backup"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices:active_backup", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "active_backup"], "syntax": "attribute", "type": "object"}, {"aliases": ["devices"], "anchor": "schema-bond_device_list--bond_devices--devices", "description": "Ethernet devices that will make up this bond.", "document_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "devices"], "syntax": "attribute", "type": "list"}, {"aliases": ["lacp"], "anchor": "section", "description": "LACP parameters for the bond device.", "document_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices:lacp", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bond_device_list--bond_devices--lacp--rate", "enforcement": "provider-schema", "group": "bond_device_list.bond_devices.lacp:RequiredObjectAttributes:rate", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices:lacp", "type": "requires"}], "schema_path": ["bond_device_list", "bond_devices", "lacp"], "syntax": "block", "type": "object"}, {"aliases": ["link polling interval"], "anchor": "schema-bond_device_list--bond_devices--link_polling_interval", "description": "Link polling interval in milliseconds.", "document_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "link_polling_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["link up delay"], "anchor": "schema-bond_device_list--bond_devices--link_up_delay", "description": "Milliseconds wait before link is declared up.", "document_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "link_up_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["name"], "anchor": "schema-bond_device_list--bond_devices--name", "description": "Name for the Bond. Ex 'bond0'", "document_id": "xcsh-docs:resources:voltstack_site:properties:bond_device_list:bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bond_device_list", "bond_devices", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/bond_device_list/bond_devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of bond devices.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bond_device_list.bond_devices

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/)
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

- [active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/bond_devices/active_backup/): complete subsection reference.

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

- [lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/bond_devices/lacp/): complete subsection reference.

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

- [bond_device_list.bond_devices.active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/bond_devices/active_backup/)
- [bond_device_list.bond_devices.lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/bond_devices/lacp/)
- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/bond_device_list/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
