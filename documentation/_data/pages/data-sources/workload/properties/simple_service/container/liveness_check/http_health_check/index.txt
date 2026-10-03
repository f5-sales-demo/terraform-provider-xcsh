---
page_title: "simple_service.container.liveness_check.http_health_check"
subcategory: "Container"
description: "HTTPHealthCheckType describes a health check based on HTTP GET requests."
xcsh_docs: {"aliases": ["simple service container liveness check http health check"], "body_bytes": 5268, "body_sha256": "sha256:86ffd76b6615a7cdea9b1447e656e06871a307254356c7b92cdf86fa0831b57f", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:http_health_check:port"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:http_health_check", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check", "path": "documentation/data-sources/workload/properties/simple_service/container/liveness_check/http_health_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3222103322132103-0310130002332210-0100322111130121-3100012213200222-2013230001020130-0013213201013233-2113031123302120-2021110321101112", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "liveness_check", "http_health_check"], "schema_version": 1, "sections": [{"aliases": ["simple service container liveness check http health check headers"], "anchor": "schema-simple_service--container--liveness_check--http_health_check--headers", "description": "Specifies a list of HTTP headers that should be added to each request that is sent to the health checked container. This is a list of key-value pairs.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:http_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "http_health_check", "headers"], "syntax": "attribute", "type": "map"}, {"aliases": ["simple service container liveness check http health check host header"], "anchor": "schema-simple_service--container--liveness_check--http_health_check--host_header", "description": "The value of the host header in the HTTP health check request.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:http_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "http_health_check", "host_header"], "syntax": "attribute", "type": "string"}, {"aliases": ["simple service container liveness check http health check path"], "anchor": "schema-simple_service--container--liveness_check--http_health_check--path", "description": "Path to access on the HTTP server.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:http_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "http_health_check", "path"], "syntax": "attribute", "type": "string"}, {"aliases": ["simple service container liveness check http health check port"], "anchor": "section", "description": "Port", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:liveness_check:http_health_check:port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "http_health_check", "port"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/container/liveness_check/http_health_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.liveness_check.http_health_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/)
- [simple_service.container.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/liveness_check/)
- simple_service.container.liveness_check.http_health_check

<a id="section"></a>

Type: `"single"`. Computed.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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

## Direct properties

<a id="schema-simple_service--container--liveness_check--http_health_check--headers"></a>

### headers property

Type: `["map", "string"]`. Computed.

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

<a id="schema-simple_service--container--liveness_check--http_health_check--host_header"></a>

### host_header property

Type: `"string"`. Computed.

The value of the host header in the HTTP health check request.

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

<a id="schema-simple_service--container--liveness_check--http_health_check--path"></a>

### path property

Type: `"string"`. Computed.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

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

- [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/liveness_check/http_health_check/port/): complete subsection reference.

## Next pages

- [simple_service.container.liveness_check.http_health_check.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/liveness_check/http_health_check/port/)
- [simple_service.container.liveness_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/liveness_check/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
