---
page_title: "primary.rr_set_group.rr_set.naptr_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary rr set group rr set naptr record values"], "body_bytes": 8979, "body_sha256": "sha256:2afbfb78dd39c124b1cfa9c9d4eba00bf85681d295f96195049729936baebf95", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record", "path": "documentation/resources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0100220222133303-1230232023133312-3132333010213303-0000021310121331-0121023031303003-2003023202002101-3320001002112000-1133232110312031", "registry_path": "docs/guides/resources--dns_zone--reference--group-003.md", "relationships": [{"anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--flags", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.naptr_record.values:RequiredListObjectAttributes:flags,order,preference", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "type": "requires"}, {"anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--order", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.naptr_record.values:RequiredListObjectAttributes:flags,order,preference", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "type": "requires"}, {"anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--preference", "enforcement": "provider-schema", "group": "primary.rr_set_group.rr_set.naptr_record.values:RequiredListObjectAttributes:flags,order,preference", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "values"], "schema_version": 1, "sections": [{"aliases": ["primary rr set group rr set naptr record values flags"], "anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--flags", "description": "Flag to control aspects of the rewriting and interpretation of the fields in the record. At this time only four flags, S/A/U/P, are defined.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "values", "flags"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary rr set group rr set naptr record values order"], "anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--order", "description": "Order in which the NAPTR records must be processed. A lower number indicates a higher preference.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "values", "order"], "syntax": "attribute", "type": "number"}, {"aliases": ["primary rr set group rr set naptr record values preference"], "anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--preference", "description": "Preference when records have the same order. A lower number indicates a higher preference.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "values", "preference"], "syntax": "attribute", "type": "number"}, {"aliases": ["primary rr set group rr set naptr record values regexp"], "anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--regexp", "description": "Regular expression to construct the next domain name to lookup.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "values", "regexp"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary rr set group rr set naptr record values replacement"], "anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--replacement", "description": "The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "values", "replacement"], "syntax": "attribute", "type": "string"}, {"aliases": ["primary rr set group rr set naptr record values service"], "anchor": "schema-primary--rr_set_group--rr_set--naptr_record--values--service", "description": "Specifies the service(s) available down this rewrite path.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:rr_set_group:rr_set:naptr_record:values", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "rr_set_group", "rr_set", "naptr_record", "values", "service"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.rr_set_group.rr_set.naptr_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/)
- [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/)
- [primary.rr_set_group.rr_set.naptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/)
- primary.rr_set_group.rr_set.naptr_record.values

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

NAPTR Value. Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("flags",
    "order",
    "preference")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
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

<a id="schema-primary--rr_set_group--rr_set--naptr_record--values--flags"></a>

### flags property

Type: `"string"`. Optional.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--naptr_record--values--order"></a>

### order property

Type: `"number"`. Optional.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--naptr_record--values--preference"></a>

### preference property

Type: `"number"`. Optional.

Preference when records have the same order. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--naptr_record--values--regexp"></a>

### regexp property

Type: `"string"`. Optional.

Regular expression to construct the next domain name to lookup.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--naptr_record--values--replacement"></a>

### replacement property

Type: `"string"`. Optional.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-primary--rr_set_group--rr_set--naptr_record--values--service"></a>

### service property

Type: `"string"`. Optional.

Specifies the service(s) available down this rewrite path.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```
