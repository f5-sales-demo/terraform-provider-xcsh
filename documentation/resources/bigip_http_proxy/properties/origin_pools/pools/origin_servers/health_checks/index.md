---
page_title: "origin_pools.pools.origin_servers.health_checks"
subcategory: ""
description: "Origin Server Health Checks."
xcsh_docs: {"aliases": ["backend servers", "origin pools pools origin servers health checks", "origin servers", "upstream servers"], "body_bytes": 8307, "body_sha256": "sha256:526570a2bdd9b2333ef34c797d9d5e46accca911c35417a3d19d952621f8c3b7", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers", "path": "documentation/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1231223200203212-2223102120201210-0132023001333121-0310321020212013-2303231100301333-3231130313200203-0020112010001202-2200002131110303", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-001.md", "relationships": [{"anchor": "schema-origin_pools--pools--origin_servers--health_checks--healthy_threshold", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "type": "requires"}, {"anchor": "schema-origin_pools--pools--origin_servers--health_checks--interval", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "type": "requires"}, {"anchor": "schema-origin_pools--pools--origin_servers--health_checks--timeout", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "type": "requires"}, {"anchor": "schema-origin_pools--pools--origin_servers--health_checks--unhealthy_threshold", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks:RequiredObjectAttributes:health_check,healthy_threshold,interval,timeout,unhealthy_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks"], "schema_version": 1, "sections": [{"aliases": ["health check"], "anchor": "section", "description": "List of Health Checks.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks.health_check:ConflictingListObjectAttributes:icmp_health_check,tcp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:icmp_health_check", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_pools.pools.origin_servers.health_checks.health_check:ConflictingListObjectAttributes:icmp_health_check,tcp_health_check", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks:health_check:tcp_health_check", "type": "conflicts"}], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "health_check"], "syntax": "block", "type": "object"}, {"aliases": ["healthy threshold", "login success", "succeeded", "success", "successful"], "anchor": "schema-origin_pools--pools--origin_servers--health_checks--healthy_threshold", "description": "Number of successful responses before declaring healthy. In other words, this is the number of healthy health checks required before a host is marked healthy. Note that during startup, only a single successful health check is required to mark a host healthy.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "healthy_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["interval"], "anchor": "schema-origin_pools--pools--origin_servers--health_checks--interval", "description": "Time interval in seconds between two health check requests.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["duration", "login success", "operation timeout", "succeeded", "success", "successful", "timeout"], "anchor": "schema-origin_pools--pools--origin_servers--health_checks--timeout", "description": "Timeout in seconds to wait for successful response. In other words, it is the time to wait for a health check response. If the timeout is reached the health check attempt will be considered a failure.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["unhealthy threshold"], "anchor": "schema-origin_pools--pools--origin_servers--health_checks--unhealthy_threshold", "description": "Number of failed responses before declaring unhealthy. In other words, this is the number of unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check if a host responds with 503 this threshold is ignored and the host is considered unhealthy immediately.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools:pools:origin_servers:health_checks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_pools", "pools", "origin_servers", "health_checks", "unhealthy_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Origin Server Health Checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_pools.pools.origin_servers.health_checks

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/)
- [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- origin_pools.pools.origin_servers.health_checks

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for health checks.

Upstream description:

Origin Server Health Checks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("health_check",
    "healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold")}
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
health_checks {
  # Configure direct properties listed below.
}
```

## Direct properties

- [health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/): complete subsection reference.

<a id="schema-origin_pools--pools--origin_servers--health_checks--healthy_threshold"></a>

### healthy_threshold property

Type: `"number"`. Optional.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Upstream description:

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

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

<a id="schema-origin_pools--pools--origin_servers--health_checks--interval"></a>

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

<a id="schema-origin_pools--pools--origin_servers--health_checks--timeout"></a>

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

<a id="schema-origin_pools--pools--origin_servers--health_checks--unhealthy_threshold"></a>

### unhealthy_threshold property

Type: `"number"`. Optional.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check
if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Upstream description:

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check
if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

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

- [origin_pools.pools.origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/)
- [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
