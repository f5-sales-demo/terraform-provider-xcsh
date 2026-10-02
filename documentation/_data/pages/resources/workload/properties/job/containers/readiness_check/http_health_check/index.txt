---
page_title: "job.containers.readiness_check.http_health_check"
subcategory: "Container"
description: "HTTPHealthCheckType describes a health check based on HTTP GET requests."
xcsh_docs: {"aliases": ["job containers readiness check http health check"], "body_bytes": 5641, "body_sha256": "sha256:48ee3e52770ee4eb3a27f248ec82eae939a576ccd6eb944381fafa421f9f3559", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check:port"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check", "parent_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check", "path": "documentation/resources/workload/properties/job/containers/readiness_check/http_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0233132103213321-3313220020312110-2101103210301101-0310230020133310-0222131330330222-3011213321020020-3122230111300132-0313022002201010", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [{"anchor": "schema-job--containers--readiness_check--http_health_check--path", "enforcement": "provider-schema", "group": "job.containers.readiness_check.http_health_check:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers", "readiness_check", "http_health_check"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "schema-job--containers--readiness_check--http_health_check--headers", "description": "Specifies a list of HTTP headers that should be added to each request that is sent to the health checked container. This is a list of key-value pairs.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "http_health_check", "headers"], "syntax": "attribute", "type": "map"}, {"aliases": ["host header"], "anchor": "schema-job--containers--readiness_check--http_health_check--host_header", "description": "The value of the host header in the HTTP health check request.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "http_health_check", "host_header"], "syntax": "attribute", "type": "string"}, {"aliases": ["path"], "anchor": "schema-job--containers--readiness_check--http_health_check--path", "description": "Path to access on the HTTP server.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "http_health_check", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check:port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--containers--readiness_check--http_health_check--port--name", "enforcement": "provider-schema", "group": "job.containers.readiness_check.http_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check:port", "type": "conflicts"}, {"anchor": "schema-job--containers--readiness_check--http_health_check--port--num", "enforcement": "provider-schema", "group": "job.containers.readiness_check.http_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:readiness_check:http_health_check:port", "type": "conflicts"}], "schema_path": ["job", "containers", "readiness_check", "http_health_check", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/containers/readiness_check/http_health_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.readiness_check.http_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/)
- [job.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/readiness_check/)
- job.containers.readiness_check.http_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
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

<a id="schema-job--containers--readiness_check--http_health_check--headers"></a>

### headers property

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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

<a id="schema-job--containers--readiness_check--http_health_check--host_header"></a>

### host_header property

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="schema-job--containers--readiness_check--http_health_check--path"></a>

### path property

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/readiness_check/http_health_check/port/): complete subsection reference.

## Next pages

- [job.containers.readiness_check.http_health_check.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/readiness_check/http_health_check/port/)
- [job.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/readiness_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
