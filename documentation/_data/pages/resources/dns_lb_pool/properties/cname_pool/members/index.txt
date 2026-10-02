---
page_title: "cname_pool.members"
subcategory: ""
description: "Configuration parameter for members"
xcsh_docs: {"aliases": ["cname pool members"], "body_bytes": 7247, "body_sha256": "sha256:c9ab565cbda1cd0f0b0ae03119db1e25cdd36f11ebbecbbed86ed5c5f58b2f05", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "parent_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool", "path": "documentation/resources/dns_lb_pool/properties/cname_pool/members/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1201013212131122-1333312112201001-1233033132210330-1122201133023222-1200233333310033-0023320210210233-2211213313311002-3302101302303120", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [{"anchor": "schema-cname_pool--members--domain", "enforcement": "provider-schema", "group": "cname_pool.members:RequiredListObjectAttributes:domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cname_pool", "members"], "schema_version": 1, "sections": [{"aliases": ["domain"], "anchor": "schema-cname_pool--members--domain", "description": "Specifies the fully qualified domain name.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["final translation"], "anchor": "schema-cname_pool--members--final_translation", "description": "If this flag is true, the CNAME record will not be translated further.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "final_translation"], "syntax": "attribute", "type": "bool"}, {"aliases": ["name"], "anchor": "schema-cname_pool--members--name", "description": "Pool member name.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["priority"], "anchor": "schema-cname_pool--members--priority", "description": "Used if the pool’s load balancing mode is set to Priority. Determines the order in which traffic is routed to pool members. The lower the number, the higher the priority, making those members active while higher-numbered members act as backups.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["ratio"], "anchor": "schema-cname_pool--members--ratio", "description": "Used if the pool’s load balancing mode is set to Ratio-Member.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:cname_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cname_pool", "members", "ratio"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/cname_pool/members/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for members", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cname_pool.members

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/)
- cname_pool.members

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
members {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cname_pool--members--domain"></a>

### domain property

Type: `"string"`. Optional.

Specifies the fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="schema-cname_pool--members--final_translation"></a>

### final_translation property

Type: `"bool"`. Optional.

If this flag is true, the CNAME record will not be translated further.

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

<a id="schema-cname_pool--members--name"></a>

### name property

Type: `"string"`. Optional.

Name. Pool member name.

Upstream description:

Pool member name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-cname_pool--members--priority"></a>

### priority property

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Priority. Determines the order in which traffic is
routed to pool members. The lower the number, the higher the priority, making those members active
while higher-numbered members act as backups.

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
    "minimum": 0,
    "multipleOf": 1
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

<a id="schema-cname_pool--members--ratio"></a>

### ratio property

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Ratio-Member.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

## Next pages

- [cname_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/cname_pool/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
