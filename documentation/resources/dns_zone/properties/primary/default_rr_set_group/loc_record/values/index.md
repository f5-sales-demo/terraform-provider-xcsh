---
page_title: "primary.default_rr_set_group.loc_record.values"
subcategory: "DNS"
description: "Configuration parameter for values"
xcsh_docs: {"aliases": ["primary default rr set group loc record values"], "body_bytes": 12541, "body_sha256": "sha256:07cf61c5131856254f5cf53af478b399ae1e371cf6911b6713d935c2d10a027b", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record", "path": "documentation/resources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2112121111131001-3303232101130313-1333110202023323-2101112231121012-1103313011133322-1033130001232233-0232023330011221-1223033200111230", "registry_path": "docs/guides/resources--dns_zone--reference--group-002.md", "relationships": [{"anchor": "schema-primary--default_rr_set_group--loc_record--values--altitude", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.loc_record.values:RequiredListObjectAttributes:altitude,latitude_degree,longitude_degree", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "type": "requires"}, {"anchor": "schema-primary--default_rr_set_group--loc_record--values--latitude_degree", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.loc_record.values:RequiredListObjectAttributes:altitude,latitude_degree,longitude_degree", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "type": "requires"}, {"anchor": "schema-primary--default_rr_set_group--loc_record--values--longitude_degree", "enforcement": "provider-schema", "group": "primary.default_rr_set_group.loc_record.values:RequiredListObjectAttributes:altitude,latitude_degree,longitude_degree", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "default_rr_set_group", "loc_record", "values"], "schema_version": 1, "sections": [{"aliases": ["altitude"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--altitude", "description": "Altitude in meters.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "altitude"], "syntax": "attribute", "type": "number"}, {"aliases": ["horizontal precision"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--horizontal_precision", "description": "Horizontal Precision in meters.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "horizontal_precision"], "syntax": "attribute", "type": "number"}, {"aliases": ["latitude degree"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--latitude_degree", "description": "Latitude degree, an integer between 0 and 90, including 0 and 90.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "latitude_degree"], "syntax": "attribute", "type": "number"}, {"aliases": ["latitude hemisphere"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--latitude_hemisphere", "description": "Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "latitude_hemisphere"], "syntax": "attribute", "type": "string"}, {"aliases": ["latitude minute"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--latitude_minute", "description": "Latitude minute, an integer between 0 and 59, including 0 and 59.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "latitude_minute"], "syntax": "attribute", "type": "number"}, {"aliases": ["latitude second"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--latitude_second", "description": "Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "latitude_second"], "syntax": "attribute", "type": "number"}, {"aliases": ["location diameter"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--location_diameter", "description": "Diameter of a sphere enclosing the described entity, in meters.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "location_diameter"], "syntax": "attribute", "type": "number"}, {"aliases": ["longitude degree"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--longitude_degree", "description": "Longitude degree, an integer between 0 and 180, including 0 and 180.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "longitude_degree"], "syntax": "attribute", "type": "number"}, {"aliases": ["longitude hemisphere"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--longitude_hemisphere", "description": "Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "longitude_hemisphere"], "syntax": "attribute", "type": "string"}, {"aliases": ["longitude minute"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--longitude_minute", "description": "Longitude minute, an integer between 0 and 59, including 0 and 59.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "longitude_minute"], "syntax": "attribute", "type": "number"}, {"aliases": ["longitude second"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--longitude_second", "description": "Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "longitude_second"], "syntax": "attribute", "type": "number"}, {"aliases": ["vertical precision"], "anchor": "schema-primary--default_rr_set_group--loc_record--values--vertical_precision", "description": "Vertical Precision in meters.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:default_rr_set_group:loc_record:values", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "default_rr_set_group", "loc_record", "values", "vertical_precision"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for values", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.default_rr_set_group.loc_record.values

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/)
- [primary.default_rr_set_group.loc_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/loc_record/)
- primary.default_rr_set_group.loc_record.values

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

LOC Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("altitude",
    "latitude_degree",
    "longitude_degree")}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-primary--default_rr_set_group--loc_record--values--altitude"></a>

### altitude property

Type: `"number"`. Optional.

Altitude. Altitude in meters.

Upstream description:

Altitude in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--horizontal_precision"></a>

### horizontal_precision property

Type: `"number"`. Optional.

Horizontal Precision. Horizontal Precision in meters.

Upstream description:

Horizontal Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--latitude_degree"></a>

### latitude_degree property

Type: `"number"`. Optional.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 90),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--latitude_hemisphere"></a>

### latitude_hemisphere property

Type: `"string"`. Optional.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Upstream description:

Latitude hemisphere can only be N or S

&#8203;- N: North Hemisphere

&#8203;- S: South Hemisphere.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("N",
    "S"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--latitude_minute"></a>

### latitude_minute property

Type: `"number"`. Optional.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--latitude_second"></a>

### latitude_second property

Type: `"number"`. Optional.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--location_diameter"></a>

### location_diameter property

Type: `"number"`. Optional.

Diameter of a sphere enclosing the described entity, in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--longitude_degree"></a>

### longitude_degree property

Type: `"number"`. Optional.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 180),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--longitude_hemisphere"></a>

### longitude_hemisphere property

Type: `"string"`. Optional.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Upstream description:

Longitude hemisphere can only be E or W

&#8203;- E: East Hemisphere

&#8203;- W: West Hemisphere.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("E",
    "W"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--longitude_minute"></a>

### longitude_minute property

Type: `"number"`. Optional.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--longitude_second"></a>

### longitude_second property

Type: `"number"`. Optional.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="schema-primary--default_rr_set_group--loc_record--values--vertical_precision"></a>

### vertical_precision property

Type: `"number"`. Optional.

Vertical Precision. Vertical Precision in meters.

Upstream description:

Vertical Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

## Next pages

- [primary.default_rr_set_group.loc_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/default_rr_set_group/loc_record/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
