---
page_title: "job.containers.readiness_check"
subcategory: "Container"
description: "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic."
xcsh_docs: {"aliases": ["job containers readiness check"], "body_bytes": 8770, "body_sha256": "sha256:2f9829d7cf7e30c06c0eaa6e49e9ac22d930a6b90337014a96d5d1a11102644d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:job:containers:readiness_check:exec_health_check", "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check:http_health_check", "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check:tcp_health_check"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check", "parent_id": "xcsh-docs:data-sources:workload:properties:job:containers", "path": "documentation/data-sources/workload/properties/job/containers/readiness_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2113313211202121-2132001213020303-1012200131230300-3221131110013000-3003212020020033-0003103011233202-3302111002202031-0302100200010010", "registry_path": "docs/guides/data-sources--workload--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers", "readiness_check"], "schema_version": 1, "sections": [{"aliases": ["exec health check"], "anchor": "section", "description": "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check:exec_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "readiness_check", "exec_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["healthy threshold", "succeeded", "success", "successful"], "anchor": "schema-job--containers--readiness_check--healthy_threshold", "description": "Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "healthy_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["http health check"], "anchor": "section", "description": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check:http_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "readiness_check", "http_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["initial delay"], "anchor": "schema-job--containers--readiness_check--initial_delay", "description": "Number of seconds after the container has started before health checks are initiated.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "initial_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["interval"], "anchor": "schema-job--containers--readiness_check--interval", "description": "Time interval in seconds between two health check requests.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["tcp health check"], "anchor": "section", "description": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "readiness_check", "tcp_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "operation timeout", "succeeded", "success", "successful", "timeout"], "anchor": "schema-job--containers--readiness_check--timeout", "description": "Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["unhealthy threshold"], "anchor": "schema-job--containers--readiness_check--unhealthy_threshold", "description": "Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy.", "document_id": "xcsh-docs:data-sources:workload:properties:job:containers:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "readiness_check", "unhealthy_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/job/containers/readiness_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.readiness_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/)
- job.containers.readiness_check

<a id="section"></a>

Type: `"single"`. Computed.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

## Direct properties

- [exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/exec_health_check/): complete subsection reference.

<a id="schema-job--containers--readiness_check--healthy_threshold"></a>

### healthy_threshold property

Type: `"number"`. Computed.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/http_health_check/): complete subsection reference.

<a id="schema-job--containers--readiness_check--initial_delay"></a>

### initial_delay property

Type: `"number"`. Computed.

Number of seconds after the container has started before health checks are initiated.

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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="schema-job--containers--readiness_check--interval"></a>

### interval property

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/tcp_health_check/): complete subsection reference.

<a id="schema-job--containers--readiness_check--timeout"></a>

### timeout property

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-job--containers--readiness_check--unhealthy_threshold"></a>

### unhealthy_threshold property

Type: `"number"`. Computed.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

## Next pages

- [job.containers.readiness_check.exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/exec_health_check/)
- [job.containers.readiness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/http_health_check/)
- [job.containers.readiness_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/readiness_check/tcp_health_check/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/job/containers/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
