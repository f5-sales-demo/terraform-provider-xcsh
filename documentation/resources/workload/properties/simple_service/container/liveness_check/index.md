---
page_title: "simple_service.container.liveness_check"
subcategory: "Container"
description: "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic."
xcsh_docs: {"aliases": ["simple service container liveness check"], "body_bytes": 8581, "body_sha256": "sha256:e7afaac172d733de43144ac9a2299fdda47c280a166af0eacc1121ec2b7de7b4", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:exec_health_check", "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:http_health_check", "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "parent_id": "xcsh-docs:resources:workload:properties:simple_service:container", "path": "documentation/resources/workload/properties/simple_service/container/liveness_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:exec_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:ConflictingObjectAttributes:exec_health_check,http_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:http_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:ConflictingObjectAttributes:exec_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:ConflictingObjectAttributes:http_health_check,tcp_health_check", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check", "type": "conflicts"}, {"anchor": "schema-simple_service--container--liveness_check--healthy_threshold", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "type": "requires"}, {"anchor": "schema-simple_service--container--liveness_check--interval", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "type": "requires"}, {"anchor": "schema-simple_service--container--liveness_check--timeout", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "type": "requires"}, {"anchor": "schema-simple_service--container--liveness_check--unhealthy_threshold", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check:RequiredObjectAttributes:healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "liveness_check"], "schema_version": 1, "sections": [{"aliases": ["simple service container liveness check exec health check"], "anchor": "section", "description": "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:exec_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--container--liveness_check--exec_health_check--command", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check.exec_health_check:RequiredObjectAttributes:command", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:exec_health_check", "type": "requires"}], "schema_path": ["simple_service", "container", "liveness_check", "exec_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["simple service container liveness check healthy threshold", "succeeded", "success", "successful"], "anchor": "schema-simple_service--container--liveness_check--healthy_threshold", "description": "Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "healthy_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["simple service container liveness check http health check"], "anchor": "section", "description": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:http_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-simple_service--container--liveness_check--http_health_check--path", "enforcement": "provider-schema", "group": "simple_service.container.liveness_check.http_health_check:RequiredObjectAttributes:path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:http_health_check", "type": "requires"}], "schema_path": ["simple_service", "container", "liveness_check", "http_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["simple service container liveness check initial delay"], "anchor": "schema-simple_service--container--liveness_check--initial_delay", "description": "Number of seconds after the container has started before health checks are initiated.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "initial_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["simple service container liveness check interval"], "anchor": "schema-simple_service--container--liveness_check--interval", "description": "Time interval in seconds between two health check requests.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["simple service container liveness check tcp health check"], "anchor": "section", "description": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check:tcp_health_check", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "tcp_health_check"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "simple service container liveness check timeout", "succeeded", "success", "successful"], "anchor": "schema-simple_service--container--liveness_check--timeout", "description": "Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["simple service container liveness check unhealthy threshold"], "anchor": "schema-simple_service--container--liveness_check--unhealthy_threshold", "description": "Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy.", "document_id": "xcsh-docs:resources:workload:properties:simple_service:container:liveness_check", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "liveness_check", "unhealthy_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/simple_service/container/liveness_check/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.liveness_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/)
- simple_service.container.liveness_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/liveness_check/exec_health_check/): complete subsection reference.

<a id="schema-simple_service--container--liveness_check--healthy_threshold"></a>

### healthy_threshold property

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/liveness_check/http_health_check/): complete subsection reference.

<a id="schema-simple_service--container--liveness_check--initial_delay"></a>

### initial_delay property

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="schema-simple_service--container--liveness_check--interval"></a>

### interval property

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/simple_service/container/liveness_check/tcp_health_check/): complete subsection reference.

<a id="schema-simple_service--container--liveness_check--timeout"></a>

### timeout property

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="schema-simple_service--container--liveness_check--unhealthy_threshold"></a>

### unhealthy_threshold property

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```
