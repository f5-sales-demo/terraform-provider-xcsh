---
page_title: "job.containers.liveness_check"
subcategory: "Container"
description: "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic."
xcsh_docs: {"aliases": ["job containers liveness check"], "body_bytes": 9981, "body_sha256": "sha256:35cd3fd259e6b40a7b006012a0910883cef167dad90f9cdfc1b2d905f8aaa2aa", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:job:containers:liveness_check:exec_health_check", "xcsh-docs:resources:workload:properties:job:containers:liveness_check:http_health_check", "xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "parent_id": "xcsh-docs:resources:workload:properties:job:containers", "path": "documentation/resources/workload/properties/job/containers/liveness_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1212102312001320-2330002220103201-0331323102112032-0231200202213223-0111102010022023-3221030302200030-0103233001300201-2012131231021122", "registry_path": "docs/guides/resources--workload--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "job.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers.liveness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers.liveness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "job.containers.liveness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "schema-job--containers--liveness_check--healthy_threshold", "enforcement": "provider-schema", "group": "job.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "type": "requires"}, {"anchor": "schema-job--containers--liveness_check--interval", "enforcement": "provider-schema", "group": "job.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "type": "requires"}, {"anchor": "schema-job--containers--liveness_check--timeout", "enforcement": "provider-schema", "group": "job.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "type": "requires"}, {"anchor": "schema-job--containers--liveness_check--unhealthy_threshold", "enforcement": "provider-schema", "group": "job.containers.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["job", "containers", "liveness_check"], "schema_version": 1, "sections": [{"aliases": ["exec health check"], "anchor": "section", "description": "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:exec_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--containers--liveness_check--exec_health_check--command", "enforcement": "provider-schema", "group": "job.containers.liveness_check.exec_health_check:RequiredObjectAttributes:command", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:exec_health_check", "type": "requires"}], "schema_path": ["job", "containers", "liveness_check", "exec_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["healthy threshold", "login success", "succeeded", "success", "successful"], "anchor": "schema-job--containers--liveness_check--healthy_threshold", "description": "Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "liveness_check", "healthy_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["http health check"], "anchor": "section", "description": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:http_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-job--containers--liveness_check--http_health_check--path", "enforcement": "provider-schema", "group": "job.containers.liveness_check.http_health_check:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:http_health_check", "type": "requires"}], "schema_path": ["job", "containers", "liveness_check", "http_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["initial delay"], "anchor": "schema-job--containers--liveness_check--initial_delay", "description": "Number of seconds after the container has started before health checks are initiated.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "liveness_check", "initial_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["interval"], "anchor": "schema-job--containers--liveness_check--interval", "description": "Time interval in seconds between two health check requests.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "liveness_check", "interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["tcp health check"], "anchor": "section", "description": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check:tcp_health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["job", "containers", "liveness_check", "tcp_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "login success", "operation timeout", "succeeded", "success", "successful", "timeout"], "anchor": "schema-job--containers--liveness_check--timeout", "description": "Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "liveness_check", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["unhealthy threshold"], "anchor": "schema-job--containers--liveness_check--unhealthy_threshold", "description": "Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy.", "document_id": "xcsh-docs:resources:workload:properties:job:containers:liveness_check", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["job", "containers", "liveness_check", "unhealthy_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/job/containers/liveness_check/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# job.containers.liveness_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [job](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/)
- job.containers.liveness_check

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
liveness_check {
  # Configure direct properties listed below.
}
```

## Direct properties

- [exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/liveness_check/exec_health_check/): complete subsection reference.

<a id="schema-job--containers--liveness_check--healthy_threshold"></a>

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

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/liveness_check/http_health_check/): complete subsection reference.

<a id="schema-job--containers--liveness_check--initial_delay"></a>

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

<a id="schema-job--containers--liveness_check--interval"></a>

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

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/liveness_check/tcp_health_check/): complete subsection reference.

<a id="schema-job--containers--liveness_check--timeout"></a>

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

<a id="schema-job--containers--liveness_check--unhealthy_threshold"></a>

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

- [job.containers.liveness_check.exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/liveness_check/exec_health_check/)
- [job.containers.liveness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/liveness_check/http_health_check/)
- [job.containers.liveness_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/liveness_check/tcp_health_check/)
- [job.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/job/containers/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
