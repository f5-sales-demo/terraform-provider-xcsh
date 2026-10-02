---
page_title: "simple_service.container.readiness_check"
subcategory: "Container"
description: "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic."
xcsh_docs: {"aliases": ["simple service container readiness check"], "body_bytes": 10220, "body_sha256": "sha256:15d85427c399f4182fbc65b9e4d9680892832ee3fb82f328b9503b1cdd86de4d", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:exec_health_check", "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:http_health_check", "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "parent_id": "xcsh-docs:resources:workload:properties:simple_service:container", "path": "documentation/resources/workload/properties/simple_service/container/readiness_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323", "registry_path": "docs/guides/resources--workload--reference--group-017.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "schema-simple_service--container--readiness_check--healthy_threshold", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "type": "requires"}, {"anchor": "schema-simple_service--container--readiness_check--interval", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "type": "requires"}, {"anchor": "schema-simple_service--container--readiness_check--timeout", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "type": "requires"}, {"anchor": "schema-simple_service--container--readiness_check--unhealthy_threshold", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "readiness_check"], "schema_version": 1, "sections": [{"aliases": ["exec health check"], "anchor": "section", "description": "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:exec_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--container--readiness_check--exec_health_check--command", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check.exec_health_check:RequiredObjectAttributes:command", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:exec_health_check", "type": "requires"}], "schema_path": ["simple_service", "container", "readiness_check", "exec_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["healthy threshold", "succeeded", "success", "successful"], "anchor": "schema-simple_service--container--readiness_check--healthy_threshold", "description": "Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "healthy_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["http health check"], "anchor": "section", "description": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:http_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--container--readiness_check--http_health_check--path", "enforcement": "provider-schema", "group": "simple_service.container.readiness_check.http_health_check:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:http_health_check", "type": "requires"}], "schema_path": ["simple_service", "container", "readiness_check", "http_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["initial delay"], "anchor": "schema-simple_service--container--readiness_check--initial_delay", "description": "Number of seconds after the container has started before health checks are initiated.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "initial_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["interval"], "anchor": "schema-simple_service--container--readiness_check--interval", "description": "Time interval in seconds between two health check requests.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["tcp health check"], "anchor": "section", "description": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check:tcp_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "tcp_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "succeeded", "success", "successful", "timeout"], "anchor": "schema-simple_service--container--readiness_check--timeout", "description": "Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["unhealthy threshold"], "anchor": "schema-simple_service--container--readiness_check--unhealthy_threshold", "description": "Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:readiness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "unhealthy_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/container/readiness_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.readiness_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/)
- simple_service.container.readiness_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
```

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

Terraform syntax:

```terraform
readiness_check {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/readiness_check/exec_health_check/): complete subsection reference.

<a id="schema-simple_service--container--readiness_check--healthy_threshold"></a>

### healthy_threshold property

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

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

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/readiness_check/http_health_check/): complete subsection reference.

<a id="schema-simple_service--container--readiness_check--initial_delay"></a>

### initial_delay property

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
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

<a id="schema-simple_service--container--readiness_check--interval"></a>

### interval property

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

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

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/readiness_check/tcp_health_check/): complete subsection reference.

<a id="schema-simple_service--container--readiness_check--timeout"></a>

### timeout property

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

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

<a id="schema-simple_service--container--readiness_check--unhealthy_threshold"></a>

### unhealthy_threshold property

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

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

- [simple_service.container.readiness_check.exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/readiness_check/exec_health_check/)
- [simple_service.container.readiness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/readiness_check/http_health_check/)
- [simple_service.container.readiness_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/readiness_check/tcp_health_check/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
