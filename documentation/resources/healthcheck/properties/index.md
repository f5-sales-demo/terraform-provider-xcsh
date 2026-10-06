---
page_title: "Property reference"
subcategory: "Monitoring"
description: "Property reference for xcsh_healthcheck."
xcsh_docs: {"aliases": ["healthcheck"], "body_bytes": 20041, "body_sha256": "sha256:a12f8e84f2ba5e0a193d66e72a513f9a81d11e8a1bc928acf9f589ef32f76e56", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:healthcheck:properties:default_jitter", "xcsh-docs:resources:healthcheck:properties:http_health_check", "xcsh-docs:resources:healthcheck:properties:tcp_health_check", "xcsh-docs:resources:healthcheck:properties:timeouts", "xcsh-docs:resources:healthcheck:properties:udp_icmp_health_check"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:reference", "parent_id": "xcsh-docs:resources:healthcheck:fundamentals", "path": "documentation/resources/healthcheck/properties/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1221303312230102-1010122132313112-0112020000032300-3233231221210023-0111020022032002-3313323101102012-2003313311113022-2311311312333213", "registry_path": "docs/guides/resources--healthcheck--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["default jitter"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:healthcheck:properties:default_jitter", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_jitter"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["healthy threshold", "succeeded", "success", "successful"], "anchor": "schema-healthy_threshold", "description": "Number of successful responses before declaring healthy. In other words, this is the number of healthy health checks required before a host is marked healthy. Note that during startup, only a single successful health check is required to mark a host healthy.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["healthy_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["http health check", "succeeded", "success", "successful"], "anchor": "section", "description": "Healthy if \"GET\" method on URL \"HTTP(s)://<host>/<path>\" with optional \"<header>\" returns success. \"host\" is not used for DNS resolution. It is used as HTTP Header in the request.", "document_id": "xcsh-docs:resources:healthcheck:properties:http_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-http_health_check--host_header", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:host_header,use_origin_server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:healthcheck:properties:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_health_check:ConflictingObjectAttributes:host_header,use_origin_server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:healthcheck:properties:http_health_check:use_origin_server_name", "type": "conflicts"}, {"anchor": "schema-http_health_check--path", "enforcement": "provider-schema", "group": "http_health_check:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:healthcheck:properties:http_health_check", "type": "requires"}], "schema_path": ["http_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["interval"], "anchor": "schema-interval", "description": "Time interval in seconds between two healthcheck requests.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["jitter percent"], "anchor": "schema-jitter_percent", "description": "Exclusive with Specify a custom jitter value as a percentage of the health check interval. Valid values are 0 (to disable jitter) and 10 to 50.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["jitter_percent"], "syntax": "attribute", "type": "number"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["succeeded", "success", "successful", "tcp health check"], "anchor": "section", "description": "Healthy if TCP connection is successful and response payload matches <expected_response>", "document_id": "xcsh-docs:resources:healthcheck:properties:tcp_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tcp_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "succeeded", "success", "successful", "timeout"], "anchor": "schema-timeout", "description": "Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:healthcheck:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["udp icmp health check"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:healthcheck:properties:udp_icmp_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["udp_icmp_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["unhealthy threshold"], "anchor": "schema-unhealthy_threshold", "description": "Number of failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy immediately.", "document_id": "xcsh-docs:resources:healthcheck:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["unhealthy_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [default_jitter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/default_jitter/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

<a id="schema-healthy_threshold"></a>

### healthy_threshold property

Type: `"number"`. Required.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy. Recommended: \`3\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "threshold",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "category": "threshold",
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-interval"></a>

### interval property

Type: `"number"`. Required.

Time interval in seconds between two healthcheck requests. Recommended: \`15\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="schema-jitter_percent"></a>

### jitter_percent property

Type: `"number"`. Optional, Computed.

Exclusive with \[default\_jitter\] Specify a custom jitter value as a percentage of the health check
interval. Valid values are 0 (to disable jitter) and 10 to 50. Server applies default when omitted.
Recommended: \`30\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 50},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Non-contiguous: {0} union [10, 50] — values 1-9 rejected by API",
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "multipleOf": 1,
    "ranges": [
      {
        "maximum": 0,
        "minimum": 0
      },
      {
        "maximum": 50,
        "minimum": 10
      }
    ]
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-50"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-50"
  }
}
```

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Healthcheck. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Healthcheck is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
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
  }
}
```

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/tcp_health_check/): complete subsection reference.

<a id="schema-timeout"></a>

### timeout property

Type: `"number"`. Required.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure. Recommended: \`3\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "API rejects timeout > 600 for healthchecks (global pattern says 3600)",
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/timeouts/): complete subsection reference.

- [udp_icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/udp_icmp_health_check/): complete subsection reference.

<a id="schema-unhealthy_threshold"></a>

### unhealthy_threshold property

Type: `"number"`. Required.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately. Recommended: \`1\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "threshold",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "category": "threshold",
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-annotations) |
| `default_jitter` | [default_jitter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/default_jitter/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-disable) |
| `healthy_threshold` | [healthy_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-healthy_threshold) |
| `http_health_check` | [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#section) |
| `http_health_check.expected_response` | [http_health_check.expected_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#schema-http_health_check--expected_response) |
| `http_health_check.expected_status_codes` | [http_health_check.expected_status_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#schema-http_health_check--expected_status_codes) |
| `http_health_check.headers` | [http_health_check.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#schema-http_health_check--headers) |
| `http_health_check.host_header` | [http_health_check.host_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#schema-http_health_check--host_header) |
| `http_health_check.path` | [http_health_check.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#schema-http_health_check--path) |
| `http_health_check.request_headers_to_remove` | [http_health_check.request_headers_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#schema-http_health_check--request_headers_to_remove) |
| `http_health_check.use_http2` | [http_health_check.use_http2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#schema-http_health_check--use_http2) |
| `http_health_check.use_origin_server_name` | [http_health_check.use_origin_server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/use_origin_server_name/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-id) |
| `interval` | [interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-interval) |
| `jitter_percent` | [jitter_percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-jitter_percent) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-namespace) |
| `tcp_health_check` | [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/tcp_health_check/#section) |
| `tcp_health_check.expected_response` | [tcp_health_check.expected_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/tcp_health_check/#schema-tcp_health_check--expected_response) |
| `tcp_health_check.send_payload` | [tcp_health_check.send_payload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/tcp_health_check/#schema-tcp_health_check--send_payload) |
| `timeout` | [timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-timeout) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/timeouts/#schema-timeouts--update) |
| `udp_icmp_health_check` | [udp_icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/udp_icmp_health_check/#section) |
| `unhealthy_threshold` | [unhealthy_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/#schema-unhealthy_threshold) |
