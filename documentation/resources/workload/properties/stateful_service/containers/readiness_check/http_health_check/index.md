---
page_title: "stateful_service.containers.readiness_check.http_health_check"
subcategory: "Container"
description: "HTTPHealthCheckType describes a health check based on HTTP GET requests."
xcsh_docs: {"aliases": ["stateful service containers readiness check http health check"], "body_bytes": 5849, "body_sha256": "sha256:8bf182aee1688fc94832bc41138e7d2810a3b11c4b569560f057654115cadf00", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check:port"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check", "path": "documentation/resources/workload/properties/stateful_service/containers/readiness_check/http_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3103313332112223-2130033020321112-0322301232002211-2223322202123210-2331330121022122-1311321200133031-0301021310022213-3322311200230001", "registry_path": "docs/guides/resources--workload--reference--group-028.md", "relationships": [{"anchor": "schema-stateful_service--containers--readiness_check--http_health_check--path", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check.http_health_check:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "containers", "readiness_check", "http_health_check"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "schema-stateful_service--containers--readiness_check--http_health_check--headers", "description": "Specifies a list of HTTP headers that should be added to each request that is sent to the health checked container. This is a list of key-value pairs.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "readiness_check", "http_health_check", "headers"], "syntax": "attribute", "type": "map"}, {"aliases": ["host header"], "anchor": "schema-stateful_service--containers--readiness_check--http_health_check--host_header", "description": "The value of the host header in the HTTP health check request.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "readiness_check", "http_health_check", "host_header"], "syntax": "attribute", "type": "string"}, {"aliases": ["path"], "anchor": "schema-stateful_service--containers--readiness_check--http_health_check--path", "description": "Path to access on the HTTP server.", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "containers", "readiness_check", "http_health_check", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check:port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-stateful_service--containers--readiness_check--http_health_check--port--name", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check.http_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check:port", "type": "conflicts"}, {"anchor": "schema-stateful_service--containers--readiness_check--http_health_check--port--num", "enforcement": "provider-schema", "group": "stateful_service.containers.readiness_check.http_health_check.port:ConflictingObjectAttributes:name,num", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:stateful_service:containers:readiness_check:http_health_check:port", "type": "conflicts"}], "schema_path": ["stateful_service", "containers", "readiness_check", "http_health_check", "port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/containers/readiness_check/http_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.containers.readiness_check.http_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/)
- [stateful_service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/)
- [stateful_service.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/)
- stateful_service.containers.readiness_check.http_health_check

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

<a id="schema-stateful_service--containers--readiness_check--http_health_check--headers"></a>

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

<a id="schema-stateful_service--containers--readiness_check--http_health_check--host_header"></a>

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

<a id="schema-stateful_service--containers--readiness_check--http_health_check--path"></a>

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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/http_health_check/port/): complete subsection reference.

## Next pages

- [stateful_service.containers.readiness_check.http_health_check.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/http_health_check/port/)
- [stateful_service.containers.readiness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/stateful_service/containers/readiness_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
