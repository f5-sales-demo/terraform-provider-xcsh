---
page_title: "simple_service.container.readiness_check"
subcategory: "Container"
description: "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic."
xcsh_docs: {"aliases": ["simple service container readiness check"], "body_bytes": 8992, "body_sha256": "sha256:709b185dc0c6884030cc6924d937ebdd3cb1c444573588b83a5fa6435ee92e18", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:exec_health_check", "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:http_health_check", "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:tcp_health_check"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check", "parent_id": "xcsh-docs:data-sources:workload:properties:simple_service:container", "path": "documentation/data-sources/workload/properties/simple_service/container/readiness_check/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["simple_service", "container", "readiness_check"], "schema_version": 1, "sections": [{"aliases": ["simple service container readiness check exec health check"], "anchor": "section", "description": "ExecHealthCheckType describes a health check based on \"run in container\" action. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:exec_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "exec_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["simple service container readiness check healthy threshold", "succeeded", "success", "successful"], "anchor": "schema-simple_service--container--readiness_check--healthy_threshold", "description": "Number of consecutive successful responses after having failed before declaring healthy. In other words, this is the number of healthy health checks required before marking healthy. Note that during startup and liveliness, only a single successful health check is required to mark a container healthy.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "healthy_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["simple service container readiness check http health check"], "anchor": "section", "description": "HTTPHealthCheckType describes a health check based on HTTP GET requests.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:http_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "http_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["simple service container readiness check initial delay"], "anchor": "schema-simple_service--container--readiness_check--initial_delay", "description": "Number of seconds after the container has started before health checks are initiated.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "initial_delay"], "syntax": "attribute", "type": "number"}, {"aliases": ["simple service container readiness check interval"], "anchor": "schema-simple_service--container--readiness_check--interval", "description": "Time interval in seconds between two health check requests.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["simple service container readiness check tcp health check"], "anchor": "section", "description": "TCPHealthCheckType describes a health check based on opening a TCP connection.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check:tcp_health_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "tcp_health_check"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "simple service container readiness check timeout", "succeeded", "success", "successful"], "anchor": "schema-simple_service--container--readiness_check--timeout", "description": "Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["simple service container readiness check unhealthy threshold"], "anchor": "schema-simple_service--container--readiness_check--unhealthy_threshold", "description": "Number of consecutive failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a container is marked unhealthy.", "document_id": "xcsh-docs:data-sources:workload:properties:simple_service:container:readiness_check", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["simple_service", "container", "readiness_check", "unhealthy_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/simple_service/container/readiness_check/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "HealthCheckType describes a health check to be performed against a container to determine whether it has started up or is alive or ready to receive traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# simple_service.container.readiness_check

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [simple_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/)
- simple_service.container.readiness_check

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

- [exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/exec_health_check/): complete subsection reference.

<a id="schema-simple_service--container--readiness_check--healthy_threshold"></a>

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

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/http_health_check/): complete subsection reference.

<a id="schema-simple_service--container--readiness_check--initial_delay"></a>

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

<a id="schema-simple_service--container--readiness_check--interval"></a>

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

- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/tcp_health_check/): complete subsection reference.

<a id="schema-simple_service--container--readiness_check--timeout"></a>

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

<a id="schema-simple_service--container--readiness_check--unhealthy_threshold"></a>

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

## Next pages

- [simple_service.container.readiness_check.exec_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/exec_health_check/)
- [simple_service.container.readiness_check.http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/http_health_check/)
- [simple_service.container.readiness_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/readiness_check/tcp_health_check/)
- [simple_service.container](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/simple_service/container/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
