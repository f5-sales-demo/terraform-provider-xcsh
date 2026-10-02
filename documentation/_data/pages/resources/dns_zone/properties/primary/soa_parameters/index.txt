---
page_title: "primary.soa_parameters"
subcategory: "DNS"
description: "Configuration parameter for soa parameters."
xcsh_docs: {"aliases": ["primary soa parameters"], "body_bytes": 6669, "body_sha256": "sha256:3bb14641b06b80bd67c6ab4a224f42123a63604e83b09cd75d13a0aa0bc99081", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "parent_id": "xcsh-docs:resources:dns_zone:properties:primary", "path": "documentation/resources/dns_zone/properties/primary/soa_parameters/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1020202022030312-2130233001011230-0123121101212023-0220103223030203-2021100102202011-2332311013003301-3000013203013302-0121111103000130", "registry_path": "docs/guides/resources--dns_zone--reference--group-003.md", "relationships": [{"anchor": "schema-primary--soa_parameters--refresh", "enforcement": "provider-schema", "group": "primary.soa_parameters:RequiredObjectAttributes:refresh,retry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "type": "requires"}, {"anchor": "schema-primary--soa_parameters--retry", "enforcement": "provider-schema", "group": "primary.soa_parameters:RequiredObjectAttributes:refresh,retry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["primary", "soa_parameters"], "schema_version": 1, "sections": [{"aliases": ["expire"], "anchor": "schema-primary--soa_parameters--expire", "description": "Expire value indicates when secondary nameservers should stop answering request for this zone if primary does not respond.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "expire"], "syntax": "attribute", "type": "number"}, {"aliases": ["negative ttl"], "anchor": "schema-primary--soa_parameters--negative_ttl", "description": "Negative TTL value indicates how long to cache non-existent resource record for this zone.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "negative_ttl"], "syntax": "attribute", "type": "number"}, {"aliases": ["refresh"], "anchor": "schema-primary--soa_parameters--refresh", "description": "Refresh value indicates when secondary nameservers should query for the SOA record to detect zone changes.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "refresh"], "syntax": "attribute", "type": "number"}, {"aliases": ["retry"], "anchor": "schema-primary--soa_parameters--retry", "description": "Retry value indicates when secondary nameservers should retry to request the serial number if primary does not respond.", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "retry"], "syntax": "attribute", "type": "number"}, {"aliases": ["ttl"], "anchor": "schema-primary--soa_parameters--ttl", "description": "SOA record time to live (in seconds)", "document_id": "xcsh-docs:resources:dns_zone:properties:primary:soa_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["primary", "soa_parameters", "ttl"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/properties/primary/soa_parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for soa parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# primary.soa_parameters

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- primary.soa_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for soa parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("refresh",
    "retry")}
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
soa_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-primary--soa_parameters--expire"></a>

### expire property

Type: `"number"`. Optional.

Expire value indicates when secondary nameservers should stop answering request for this zone if
primary does not respond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-primary--soa_parameters--negative_ttl"></a>

### negative_ttl property

Type: `"number"`. Optional.

Negative TTL value indicates how long to cache non-existent resource record for this zone.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-primary--soa_parameters--refresh"></a>

### refresh property

Type: `"number"`. Optional.

Refresh value indicates when secondary nameservers should query for the SOA record to detect zone
changes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(3600, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-primary--soa_parameters--retry"></a>

### retry property

Type: `"number"`. Optional.

Retry value indicates when secondary nameservers should retry to request the serial number if
primary does not respond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="schema-primary--soa_parameters--ttl"></a>

### ttl property

Type: `"number"`. Optional.

TTL. SOA record time to live (in seconds)

Upstream description:

SOA record time to live (in seconds)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

## Next pages

- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/primary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/)
