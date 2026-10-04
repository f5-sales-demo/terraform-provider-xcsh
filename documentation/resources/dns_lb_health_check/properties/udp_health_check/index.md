---
page_title: "udp_health_check"
subcategory: ""
description: "Configuration parameter for udp health check."
xcsh_docs: {"aliases": ["udp health check"], "body_bytes": 5947, "body_sha256": "sha256:99d673194352cd295f5a0e359847531cd968d52a7e51fbc8e8c6fa16e0b8a17f", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "parent_id": "xcsh-docs:resources:dns_lb_health_check:reference", "path": "documentation/resources/dns_lb_health_check/properties/udp_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1033111102330311-0210031231001012-0311000130202221-2102102013102323-2102322000022300-3023201333202310-3212111232132212-1013112131001321", "registry_path": "docs/guides/resources--dns_lb_health_check--reference--group-001.md", "relationships": [{"anchor": "schema-udp_health_check--health_check_port", "enforcement": "provider-schema", "group": "udp_health_check:RequiredObjectAttributes:health_check_port,receive,send", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "type": "requires"}, {"anchor": "schema-udp_health_check--receive", "enforcement": "provider-schema", "group": "udp_health_check:RequiredObjectAttributes:health_check_port,receive,send", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "type": "requires"}, {"anchor": "schema-udp_health_check--send", "enforcement": "provider-schema", "group": "udp_health_check:RequiredObjectAttributes:health_check_port,receive,send", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["udp_health_check"], "schema_version": 1, "sections": [{"aliases": ["udp health check health check port"], "anchor": "schema-udp_health_check--health_check_port", "description": "Port used for performing health check.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["udp_health_check", "health_check_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["udp health check health check secondary port"], "anchor": "schema-udp_health_check--health_check_secondary_port", "description": "Secondary port used for performing health check. If included, both ports must be healthy for the health check to pass.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["udp_health_check", "health_check_secondary_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["udp health check receive"], "anchor": "schema-udp_health_check--receive", "description": "UDP response to be matched. It can be a regex.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["udp_health_check", "receive"], "syntax": "attribute", "type": "string"}, {"aliases": ["udp health check send"], "anchor": "schema-udp_health_check--send", "description": "UDP payload.", "document_id": "xcsh-docs:resources:dns_lb_health_check:properties:udp_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["udp_health_check", "send"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/properties/udp_health_check/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configuration parameter for udp health check.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# udp_health_check

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/)
- udp_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for udp health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check_port",
    "receive",
    "send")}
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
udp_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-udp_health_check--health_check_port"></a>

### health_check_port property

Type: `"number"`. Optional.

Health Check Port. Port used for performing health check.

Upstream description:

Port used for performing health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-udp_health_check--health_check_secondary_port"></a>

### health_check_secondary_port property

Type: `"number"`. Optional.

Secondary port used for performing health check. If included, both ports must be healthy for the
health check to pass.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-udp_health_check--receive"></a>

### receive property

Type: `"string"`. Optional.

UDP response to be matched. It can be a regex.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-udp_health_check--send"></a>

### send property

Type: `"string"`. Optional.

Send String. UDP payload.

Upstream description:

UDP payload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/)
- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
