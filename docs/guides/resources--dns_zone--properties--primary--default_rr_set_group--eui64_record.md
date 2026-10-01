---
page_title: "primary.default_rr_set_group.eui64_record"
subcategory: "DNS"
description: "primary.default_rr_set_group.eui64_record for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 4401, "body_sha256": "sha256:fdd439ea8d0ec0ebe92e8bbb6c1e9bd50c6cf8f1c8758dfbe7cbe1892ba489d6", "canonical_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:eui64_record", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:eui64_record", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group", "path": "docs/guides/resources--dns_zone--properties--primary--default_rr_set_group--eui64_record.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["primary", "default_rr_set_group", "eui64_record"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/eui64_record/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "primary.default_rr_set_group.eui64_record for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.eui64_record

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md)
- [Property reference](resources--dns_zone--reference.md)
- [primary](resources--dns_zone--properties--primary.md)
- [primary.default_rr_set_group](resources--dns_zone--properties--primary--default_rr_set_group.md)
- primary.default_rr_set_group.eui64_record

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui64 record.

Upstream description:

DNS EUI64 Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
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
eui64_record {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--default_rr_set_group--eui64_record--name"></a>

### name property

Type: `"string"`. Optional.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="schema-primary--default_rr_set_group--eui64_record--value"></a>

### value property

Type: `"string"`. Optional.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Upstream description:

A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(23, 23),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

## Next pages

- [primary.default_rr_set_group](resources--dns_zone--properties--primary--default_rr_set_group.md)
- [xcsh_dns_zone](../resources/dns_zone.md)
