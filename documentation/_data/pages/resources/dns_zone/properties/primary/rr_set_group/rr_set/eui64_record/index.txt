---
page_title: "primary.rr_set_group.rr_set.eui64_record"
subcategory: "DNS"
description: "DNS EUI64 Record."
xcsh_docs: {"aliases": ["primary rr set group rr set eui64 record"], "body_bytes": 4583, "body_sha256": "sha256:5c996ff6a2d0669602e89d49a757f6aedc47caf5cc236b7e1d10a4e30a6fa0f9", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:eui64_record", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set", "path": "documentation/resources/dns_zone/properties/primary/rr_set_group/rr_set/eui64_record/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1123232112333312-0301023300332122-0100233233030001-1023002032221030-0013031220223001-2301310111033331-0120133033311203-0112320013203033", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--eui64_record--value", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.eui64_record:RequiredObjectAttributes:value", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:eui64_record", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "eui64_record"], "schema_version": 1, "sections": [{"aliases": ["primary rr set group rr set eui64 record name"], "anchor": "schema-primary--rr_set_group--rr_set--eui64_record--name", "description": "EUI64 Record name, please provide only the specific subdomain or record name without the base domain.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:eui64_record", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "eui64_record", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary rr set group rr set eui64 record value"], "anchor": "schema-primary--rr_set_group--rr_set--eui64_record--value", "description": "A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:eui64_record", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "eui64_record", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/rr_set/eui64_record/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "DNS EUI64 Record.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.eui64_record

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/)
- primary.rr_set_group.rr_set.eui64_record

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui64 record.

Additional upstream details:

DNS EUI64 Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="schema-primary--rr_set_group--rr_set--eui64_record--name"></a>

### name property

Type: `"string"`. Optional.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="schema-primary--rr_set_group--rr_set--eui64_record--value"></a>

### value property

Type: `"string"`. Optional.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
