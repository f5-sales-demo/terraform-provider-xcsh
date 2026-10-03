---
page_title: "primary.default_rr_set_group.caa_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary default rr set group caa record values"], "body_bytes": 5475, "body_sha256": "sha256:06499876f9d584b72414bd7184749491088071f10bba4297c78a161c84e000a0", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record:values", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record", "path": "documentation/resources/dns_zone/properties/primary/default_rr_set_group/caa_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0323333130111320-0113011320030021-3130002132133023-0021030230003002-2310030310133132-2223211333030301-1023130220021312-2000013103033322", "registry_path": "docs/guides/resources--dns_zone--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "caa_record", "values"], "schema_version": 1, "sections": [{"aliases": ["primary default rr set group caa record values flags"], "anchor": "schema-primary--default_rr_set_group--caa_record--values--flags", "description": "This flag should be an integer between 0 and 255.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "caa_record", "values", "flags"], "syntax": "attribute", "type": "number"}, {"aliases": ["primary default rr set group caa record values tag"], "anchor": "schema-primary--default_rr_set_group--caa_record--values--tag", "description": "Tag for categorization and filtering", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "caa_record", "values", "tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary default rr set group caa record values value"], "anchor": "schema-primary--default_rr_set_group--caa_record--values--value", "description": "Configuration parameter for value", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "caa_record", "values", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/caa_record/values/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.caa_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_rr_set_group.caa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/caa_record/)
- primary.default_rr_set_group.caa_record.values

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

CAA Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--default_rr_set_group--caa_record--values--flags"></a>

### flags property

Type: `"number"`. Optional.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="schema-primary--default_rr_set_group--caa_record--values--tag"></a>

### tag property

Type: `"string"`. Optional.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("issue",
    "issuewild",
    "iodef"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="schema-primary--default_rr_set_group--caa_record--values--value"></a>

### value property

Type: `"string"`. Optional.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [primary.default_rr_set_group.caa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/caa_record/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
