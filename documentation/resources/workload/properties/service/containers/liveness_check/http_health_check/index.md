---
page_title: "service.containers.liveness_check.http_health_check"
subcategory: "Container"
description: "HTTPHealthCheckType describes a health check based on HTTP GET requests."
xcsh_docs: {"aliases": ["service containers liveness check http health check"], "body_bytes": 6446, "body_sha256": "sha256:fb53551ed901186874511db3d06575f955d09f3d00473c12b128ce958e78da57", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check:port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check", "parent_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check", "path": "documentation/resources/workload/properties/service/containers/liveness_check/http_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1113113011220122-3021002211200132-1030321211022322-2321112333102103-1311312130011022-1322220102130220-1221131203113220-0310132223103300", "registry_path": "docs/guides/resources--workload--reference--group-015.md", "relationships": [{"anchor": "schema-service--containers--liveness_check--http_health_check--path", "enforcement": "provider-schema", "group": "service.containers.liveness_check.http_health_check:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "containers", "liveness_check", "http_health_check"], "schema_version": 1, "sections": [{"aliases": ["service containers liveness check http health check headers"], "anchor": "schema-service--containers--liveness_check--http_health_check--headers", "description": "Specifies a list of HTTP headers that should be added to each request that is sent to the health checked container. This is a list of key-value pairs.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "liveness_check", "http_health_check", "headers"], "syntax": "attribute", "type": "map"}, {"aliases": ["service containers liveness check http health check host header"], "anchor": "schema-service--containers--liveness_check--http_health_check--host_header", "description": "The value of the host header in the HTTP health check request.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "liveness_check", "http_health_check", "host_header"], "syntax": "attribute", "type": "string"}, {"aliases": ["service containers liveness check http health check path"], "anchor": "schema-service--containers--liveness_check--http_health_check--path", "description": "Path to access on the HTTP server.", "document_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "containers", "liveness_check", "http_health_check", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["service containers liveness check http health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check:port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--containers--liveness_check--http_health_check--port--name", "enforcement": "provider-schema", "group": "service.containers.liveness_check.http_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check:port", "type": "conflicts"}, {"anchor": "schema-service--containers--liveness_check--http_health_check--port--num", "enforcement": "provider-schema", "group": "service.containers.liveness_check.http_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:containers:liveness_check:http_health_check:port", "type": "conflicts"}], "schema_path": ["service", "containers", "liveness_check", "http_health_check", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/containers/liveness_check/http_health_check/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.containers.liveness_check.http_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/)
- [service.containers.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/liveness_check/)
- service.containers.liveness_check.http_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-service--containers--liveness_check--http_health_check--headers"></a>

### headers property

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":256,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"256\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"2048\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":2048,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "2048",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 2048,
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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-service--containers--liveness_check--http_health_check--host_header"></a>

### host_header property

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="schema-service--containers--liveness_check--http_health_check--path"></a>

### path property

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/liveness_check/http_health_check/port/): complete subsection reference.
