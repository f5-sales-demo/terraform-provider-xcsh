---
page_title: "primary.default_rr_set_group.caa_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary default rr set group caa record values"], "body_bytes": 5475, "body_sha256": "sha256:06499876f9d584b72414bd7184749491088071f10bba4297c78a161c84e000a0", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record:values", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record", "path": "documentation/resources/dns_zone/properties/primary/default_rr_set_group/caa_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0323333130111320-0113011320030021-3130002132133023-0021030230003002-2310030310133132-2223211333030301-1023130220021312-2000013103033322", "registry_path": "docs/guides/resources--dns_zone--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "caa_record", "values"], "schema_version": 1, "sections": [{"aliases": ["flags"], "anchor": "schema-primary--default_rr_set_group--caa_record--values--flags", "description": "This flag should be an integer between 0 and 255.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "caa_record", "values", "flags"], "syntax": "attribute", "type": "number"}, {"aliases": ["tag"], "anchor": "schema-primary--default_rr_set_group--caa_record--values--tag", "description": "Tag for categorization and filtering", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "caa_record", "values", "tag"], "syntax": "attribute", "type": "string"}, {"aliases": ["value"], "anchor": "schema-primary--default_rr_set_group--caa_record--values--value", "description": "Configuration parameter for value", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:caa_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "caa_record", "values", "value"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/caa_record/values/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
