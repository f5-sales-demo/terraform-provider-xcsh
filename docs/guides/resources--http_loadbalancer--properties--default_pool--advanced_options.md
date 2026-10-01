---
page_title: "default_pool.advanced_options"
subcategory: "Load Balancing"
description: "default_pool.advanced_options for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 13871, "body_sha256": "sha256:a73cfc85f2f16b5fbf2a00bfef8c1c7a1ff8df6038a118cdd030b7107f493882", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:auto_http_config", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:circuit_breaker", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:default_circuit_breaker", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_circuit_breaker", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_lb_source_ip_persistence", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_outlier_detection", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_proxy_protocol", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:disable_subsets", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_lb_source_ip_persistence", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:enable_subsets", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http1_config", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:http2_options", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:no_panic_threshold", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:no_request_limit_per_connection", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:outlier_detection", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:proxy_protocol_v1", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options:proxy_protocol_v2"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:advanced_options", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool", "path": "docs/guides/resources--http_loadbalancer--properties--default_pool--advanced_options.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["default_pool", "advanced_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/advanced_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "default_pool.advanced_options for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.advanced_options

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- default_pool.advanced_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configure Advanced OPTIONS for origin pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_http_config",
    "http1_config"),
  validators.ConflictingObjectAttributes("auto_http_config",
    "http2_options"),
  validators.ConflictingObjectAttributes("circuit_breaker",
    "default_circuit_breaker"),
  validators.ConflictingObjectAttributes("circuit_breaker",
    "disable_circuit_breaker"),
  validators.ConflictingObjectAttributes("default_circuit_breaker",
    "disable_circuit_breaker"),
  validators.ConflictingObjectAttributes("disable_lb_source_ip_persistence",
    "enable_lb_source_ip_persistence"),
  validators.ConflictingObjectAttributes("disable_outlier_detection",
    "outlier_detection"),
  validators.ConflictingObjectAttributes("disable_proxy_protocol",
    "proxy_protocol_v1"),
  validators.ConflictingObjectAttributes("disable_proxy_protocol",
    "proxy_protocol_v2"),
  validators.ConflictingObjectAttributes("disable_subsets",
    "enable_subsets"),
  validators.ConflictingObjectAttributes("http1_config",
    "http2_options"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection"),
  validators.ConflictingObjectAttributes("no_panic_threshold",
    "panic_threshold"),
  validators.ConflictingObjectAttributes("proxy_protocol_v1",
    "proxy_protocol_v2")}
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
  "x-ves-oneof-field-circuit_breaker_choice": "[\"circuit_breaker\",\"default_circuit_breaker\",\"disable_circuit_breaker\"]",
  "x-ves-oneof-field-http_protocol_type": "[\"auto_http_config\",\"http1_config\",\"http2_options\"]",
  "x-ves-oneof-field-lb_source_ip_persistence_choice": "[\"disable_lb_source_ip_persistence\",\"enable_lb_source_ip_persistence\"]",
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-outlier_detection_choice": "[\"disable_outlier_detection\",\"outlier_detection\"]",
  "x-ves-oneof-field-panic_threshold_type": "[\"no_panic_threshold\",\"panic_threshold\"]",
  "x-ves-oneof-field-proxy_protocol_choice": "[\"disable_proxy_protocol\",\"proxy_protocol_v1\",\"proxy_protocol_v2\"]",
  "x-ves-oneof-field-subset_choice": "[\"disable_subsets\",\"enable_subsets\"]"
}
```

Terraform syntax:

```terraform
advanced_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto_http_config](resources--http_loadbalancer--properties--default_pool--advanced_options--auto_http_config.md): complete subsection reference.

- [circuit_breaker](resources--http_loadbalancer--properties--default_pool--advanced_options--circuit_breaker.md): complete subsection reference.

<a id="schema-default_pool--advanced_options--connection_timeout"></a>

### connection_timeout property

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(1800000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

- [default_circuit_breaker](resources--http_loadbalancer--properties--default_pool--advanced_options--default_circuit_breaker.md): complete subsection reference.

- [disable_circuit_breaker](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_circuit_breaker.md): complete subsection reference.

- [disable_lb_source_ip_persistence](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_lb_source_ip_persistence.md): complete subsection reference.

- [disable_outlier_detection](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_outlier_detection.md): complete subsection reference.

- [disable_proxy_protocol](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_proxy_protocol.md): complete subsection reference.

- [disable_subsets](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_subsets.md): complete subsection reference.

- [enable_lb_source_ip_persistence](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_lb_source_ip_persistence.md): complete subsection reference.

- [enable_subsets](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets.md): complete subsection reference.

- [http1_config](resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config.md): complete subsection reference.

- [http2_options](resources--http_loadbalancer--properties--default_pool--advanced_options--http2_options.md): complete subsection reference.

<a id="schema-default_pool--advanced_options--http_idle_timeout"></a>

### http_idle_timeout property

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Upstream description:

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="schema-default_pool--advanced_options--max_requests_per_connection"></a>

### max_requests_per_connection property

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [no_panic_threshold](resources--http_loadbalancer--properties--default_pool--advanced_options--no_panic_threshold.md): complete subsection reference.

- [no_request_limit_per_connection](resources--http_loadbalancer--properties--default_pool--advanced_options--no_request_limit_per_connection.md): complete subsection reference.

- [outlier_detection](resources--http_loadbalancer--properties--default_pool--advanced_options--outlier_detection.md): complete subsection reference.

<a id="schema-default_pool--advanced_options--panic_threshold"></a>

### panic_threshold property

Type: `"number"`. Optional.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [proxy_protocol_v1](resources--http_loadbalancer--properties--default_pool--advanced_options--proxy_protocol_v1.md): complete subsection reference.

- [proxy_protocol_v2](resources--http_loadbalancer--properties--default_pool--advanced_options--proxy_protocol_v2.md): complete subsection reference.

## Next pages

- [default_pool.advanced_options.auto_http_config](resources--http_loadbalancer--properties--default_pool--advanced_options--auto_http_config.md)
- [default_pool.advanced_options.circuit_breaker](resources--http_loadbalancer--properties--default_pool--advanced_options--circuit_breaker.md)
- [default_pool.advanced_options.default_circuit_breaker](resources--http_loadbalancer--properties--default_pool--advanced_options--default_circuit_breaker.md)
- [default_pool.advanced_options.disable_circuit_breaker](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_circuit_breaker.md)
- [default_pool.advanced_options.disable_lb_source_ip_persistence](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_lb_source_ip_persistence.md)
- [default_pool.advanced_options.disable_outlier_detection](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_outlier_detection.md)
- [default_pool.advanced_options.disable_proxy_protocol](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_proxy_protocol.md)
- [default_pool.advanced_options.disable_subsets](resources--http_loadbalancer--properties--default_pool--advanced_options--disable_subsets.md)
- [default_pool.advanced_options.enable_lb_source_ip_persistence](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_lb_source_ip_persistence.md)
- [default_pool.advanced_options.enable_subsets](resources--http_loadbalancer--properties--default_pool--advanced_options--enable_subsets.md)
- [default_pool.advanced_options.http1_config](resources--http_loadbalancer--properties--default_pool--advanced_options--http1_config.md)
- [default_pool.advanced_options.http2_options](resources--http_loadbalancer--properties--default_pool--advanced_options--http2_options.md)
- [default_pool.advanced_options.no_panic_threshold](resources--http_loadbalancer--properties--default_pool--advanced_options--no_panic_threshold.md)
- [default_pool.advanced_options.no_request_limit_per_connection](resources--http_loadbalancer--properties--default_pool--advanced_options--no_request_limit_per_connection.md)
- [default_pool.advanced_options.outlier_detection](resources--http_loadbalancer--properties--default_pool--advanced_options--outlier_detection.md)
- [default_pool.advanced_options.proxy_protocol_v1](resources--http_loadbalancer--properties--default_pool--advanced_options--proxy_protocol_v1.md)
- [default_pool.advanced_options.proxy_protocol_v2](resources--http_loadbalancer--properties--default_pool--advanced_options--proxy_protocol_v2.md)
- [default_pool](resources--http_loadbalancer--properties--default_pool.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
