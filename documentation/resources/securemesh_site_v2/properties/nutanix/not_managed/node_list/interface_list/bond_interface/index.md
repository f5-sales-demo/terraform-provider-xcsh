---
page_title: "nutanix.not_managed.node_list.interface_list.bond_interface"
subcategory: ""
description: "Bond devices configuration for fleet."
xcsh_docs: {"aliases": ["nutanix not managed node list interface list bond interface"], "body_bytes": 8661, "body_sha256": "sha256:17473c7f07ef9412b2ce8ce7a3e1614f09d653133b8ad5f5780da9fe21832040", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface:active_backup", "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface:lacp"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/bond_interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "nutanix.not_managed.node_list.interface_list.bond_interface:ConflictingObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface:active_backup", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "nutanix.not_managed.node_list.interface_list.bond_interface:ConflictingObjectAttributes:active_backup,lacp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface:lacp", "type": "conflicts"}, {"anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--devices", "enforcement": "provider-schema", "group": "nutanix.not_managed.node_list.interface_list.bond_interface:RequiredObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "type": "requires"}, {"anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--link_polling_interval", "enforcement": "provider-schema", "group": "nutanix.not_managed.node_list.interface_list.bond_interface:RequiredObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "type": "requires"}, {"anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--link_up_delay", "enforcement": "provider-schema", "group": "nutanix.not_managed.node_list.interface_list.bond_interface:RequiredObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "type": "requires"}, {"anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--name", "enforcement": "provider-schema", "group": "nutanix.not_managed.node_list.interface_list.bond_interface:RequiredObjectAttributes:devices,link_polling_interval,link_up_delay,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "bond_interface"], "schema_version": 1, "sections": [{"aliases": ["nutanix not managed node list interface list bond interface active backup"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface:active_backup", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "bond_interface", "active_backup"], "syntax": "attribute", "type": "object"}, {"aliases": ["nutanix not managed node list interface list bond interface devices"], "anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--devices", "description": "Ethernet devices that will make up this bond.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "bond_interface", "devices"], "syntax": "attribute", "type": "list"}, {"aliases": ["nutanix not managed node list interface list bond interface lacp"], "anchor": "section", "description": "LACP parameters for the bond device.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface:lacp", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--lacp--rate", "enforcement": "provider-schema", "group": "nutanix.not_managed.node_list.interface_list.bond_interface.lacp:RequiredObjectAttributes:rate", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface:lacp", "type": "requires"}], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "bond_interface", "lacp"], "syntax": "block", "type": "object"}, {"aliases": ["nutanix not managed node list interface list bond interface link polling interval"], "anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--link_polling_interval", "description": "Link polling interval in milliseconds.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "bond_interface", "link_polling_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["nutanix not managed node list interface list bond interface link up delay"], "anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--link_up_delay", "description": "Milliseconds wait before link is declared up.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "bond_interface", "link_up_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["nutanix not managed node list interface list bond interface name"], "anchor": "schema-nutanix--not_managed--node_list--interface_list--bond_interface--name", "description": "Name for the Bond. Ex 'bond0'", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:bond_interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "bond_interface", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/bond_interface/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Bond devices configuration for fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# nutanix.not_managed.node_list.interface_list.bond_interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [nutanix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/)
- [nutanix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/)
- [nutanix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/)
- [nutanix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/)
- nutanix.not_managed.node_list.interface_list.bond_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
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

- [active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/bond_interface/active_backup/): complete subsection reference.

<a id="schema-nutanix--not_managed--node_list--interface_list--bond_interface--devices"></a>

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

- [lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/bond_interface/lacp/): complete subsection reference.

<a id="schema-nutanix--not_managed--node_list--interface_list--bond_interface--link_polling_interval"></a>

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

<a id="schema-nutanix--not_managed--node_list--interface_list--bond_interface--link_up_delay"></a>

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

<a id="schema-nutanix--not_managed--node_list--interface_list--bond_interface--name"></a>

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

- [nutanix.not_managed.node_list.interface_list.bond_interface.active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/bond_interface/active_backup/)
- [nutanix.not_managed.node_list.interface_list.bond_interface.lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/bond_interface/lacp/)
- [nutanix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
