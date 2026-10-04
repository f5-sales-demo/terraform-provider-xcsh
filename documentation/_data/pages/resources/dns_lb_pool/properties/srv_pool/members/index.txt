---
page_title: "srv_pool.members"
subcategory: ""
description: "Configuration parameter for members"
xcsh_docs: {"aliases": ["srv pool members"], "body_bytes": 9921, "body_sha256": "sha256:7c1ee99dca11255a43bf95c84a706c3666cb191af8575d648cba57872a133c83", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "parent_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool", "path": "documentation/resources/dns_lb_pool/properties/srv_pool/members/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2222100011301023-2013331101303131-0100021202102233-2232032010003231-1111232023232301-3013023232233113-2110301003003233-1333220210322232", "registry_path": "docs/guides/resources--dns_lb_pool--reference--group-001.md", "relationships": [{"anchor": "schema-srv_pool--members--port", "enforcement": "provider-schema", "group": "srv_pool.members:RequiredListObjectAttributes:port,priority,target,weight", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "type": "requires"}, {"anchor": "schema-srv_pool--members--priority", "enforcement": "provider-schema", "group": "srv_pool.members:RequiredListObjectAttributes:port,priority,target,weight", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "type": "requires"}, {"anchor": "schema-srv_pool--members--target", "enforcement": "provider-schema", "group": "srv_pool.members:RequiredListObjectAttributes:port,priority,target,weight", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "type": "requires"}, {"anchor": "schema-srv_pool--members--weight", "enforcement": "provider-schema", "group": "srv_pool.members:RequiredListObjectAttributes:port,priority,target,weight", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["srv_pool", "members"], "schema_version": 1, "sections": [{"aliases": ["srv pool members final translation"], "anchor": "schema-srv_pool--members--final_translation", "description": "If this flag is true, the SRV record will not be translated further.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["srv_pool", "members", "final_translation"], "syntax": "attribute", "type": "bool"}, {"aliases": ["srv pool members name"], "anchor": "schema-srv_pool--members--name", "description": "Pool member name.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["srv_pool", "members", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["srv pool members port"], "anchor": "schema-srv_pool--members--port", "description": "Port on which the service can be found.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["srv_pool", "members", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["srv pool members priority"], "anchor": "schema-srv_pool--members--priority", "description": "Priority of the target. A lower number indicates a higher preference.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["srv_pool", "members", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["srv pool members ratio"], "anchor": "schema-srv_pool--members--ratio", "description": "Configuration parameter for ratio", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["srv_pool", "members", "ratio"], "syntax": "attribute", "type": "number"}, {"aliases": ["srv pool members target"], "anchor": "schema-srv_pool--members--target", "description": "Domain name of the machine providing the service.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["srv_pool", "members", "target"], "syntax": "attribute", "type": "string"}, {"aliases": ["srv pool members weight"], "anchor": "schema-srv_pool--members--weight", "description": "Weight of the target. A higher number indicates a higher preference.", "document_id": "xcsh-docs:resources:dns_lb_pool:properties:srv_pool:members", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["srv_pool", "members", "weight"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/properties/srv_pool/members/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configuration parameter for members", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# srv_pool.members

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/)
- [srv_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/)
- srv_pool.members

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("port",
    "priority",
    "target",
    "weight")}
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

<a id="schema-srv_pool--members--final_translation"></a>

### final_translation property

Type: `"bool"`. Optional.

If this flag is true, the SRV record will not be translated further.

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

<a id="schema-srv_pool--members--name"></a>

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

<a id="schema-srv_pool--members--port"></a>

### port property

Type: `"number"`. Optional.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
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

<a id="schema-srv_pool--members--priority"></a>

### priority property

Type: `"number"`. Optional.

Priority of the target. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
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

<a id="schema-srv_pool--members--ratio"></a>

### ratio property

Type: `"number"`. Optional.

Load Balancing Ratio. Configuration parameter for ratio

Upstream description:

Configuration parameter for ratio

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="schema-srv_pool--members--target"></a>

### target property

Type: `"string"`. Optional.

Domain name of the machine providing the service.

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
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="schema-srv_pool--members--weight"></a>

### weight property

Type: `"number"`. Optional.

Weight of the target. A higher number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
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

## Next pages

- [srv_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/properties/srv_pool/)
- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
