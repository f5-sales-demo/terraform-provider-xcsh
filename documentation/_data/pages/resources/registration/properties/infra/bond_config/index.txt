---
page_title: "infra.bond_config"
subcategory: ""
description: "Bond device configuration for VPM registration."
xcsh_docs: {"aliases": ["infra bond config"], "body_bytes": 5572, "body_sha256": "sha256:b18d9449466e0686839b388afa248afc6da854b203da00840ce23fa278593f08", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra:bond_config", "parent_id": "xcsh-docs:resources:registration:properties:infra", "path": "documentation/resources/registration/properties/infra/bond_config/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0121201130321321-3231311301122033-2230210013220010-2002333120112302-0120303122023013-0132320302232111-1221102101312231-0201133202131022", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [{"anchor": "schema-infra--bond_config--interfaces", "enforcement": "provider-schema", "group": "infra.bond_config:RequiredObjectAttributes:interfaces,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:infra:bond_config", "type": "requires"}, {"anchor": "schema-infra--bond_config--name", "enforcement": "provider-schema", "group": "infra.bond_config:RequiredObjectAttributes:interfaces,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:infra:bond_config", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["infra", "bond_config"], "schema_version": 1, "sections": [{"aliases": ["infra bond config interfaces"], "anchor": "schema-infra--bond_config--interfaces", "description": "Configuration parameter for interfaces", "document_id": "xcsh-docs:resources:registration:properties:infra:bond_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "bond_config", "interfaces"], "syntax": "attribute", "type": "list"}, {"aliases": ["infra bond config mode"], "anchor": "schema-infra--bond_config--mode", "description": "Bonding mode for bond device configuration Bond mode is not specified Active-backup bond mode (one interface active, others as backup) IEEE 802.3ad Dynamic link aggregation (LACP)", "document_id": "xcsh-docs:resources:registration:properties:infra:bond_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "bond_config", "mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["infra bond config name"], "anchor": "schema-infra--bond_config--name", "description": "Human-readable name for the resource", "document_id": "xcsh-docs:resources:registration:properties:infra:bond_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "bond_config", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/bond_config/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Bond device configuration for VPM registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra.bond_config

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- infra.bond_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bond device configuration for VPM registration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces",
    "name")}
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
bond_config {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--bond_config--interfaces"></a>

### interfaces property

Type: `["list", "string"]`. Optional.

Member Interfaces. Configuration parameter for interfaces

Upstream description:

Configuration parameter for interfaces

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-infra--bond_config--mode"></a>

### mode property

Type: `"string"`. Optional.

\[Enum: BOND\_MODE\_UNSPECIFIED|ACTIVE\_BACKUP|LACP\_802\_3AD\] Bonding mode for bond device
configuration Bond mode is not specified Active-backup bond mode (one interface active, others as
backup) IEEE 802.3ad Dynamic link aggregation (LACP). Possible values are
\`BOND\_MODE\_UNSPECIFIED\`, \`ACTIVE\_BACKUP\`, \`LACP\_802\_3AD\`. Defaults to
\`BOND\_MODE\_UNSPECIFIED\`.

Upstream description:

Bonding mode for bond device configuration

Bond mode is not specified Active-backup bond mode (one interface active, others as backup) IEEE
802.3ad Dynamic link aggregation (LACP)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BOND_MODE_UNSPECIFIED",
  "enum": [
    "BOND_MODE_UNSPECIFIED",
    "ACTIVE_BACKUP",
    "LACP_802_3AD"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--bond_config--name"></a>

### name property

Type: `"string"`. Optional.

Bond Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
