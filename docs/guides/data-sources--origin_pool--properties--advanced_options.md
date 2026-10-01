---
page_title: "advanced_options"
subcategory: "Load Balancing"
description: "advanced_options for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 10817, "body_sha256": "sha256:c46ba296edb1b7ae688cc62068343f798b4834fbe92e7f8a1cb6e47f63b888a2", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:advanced_options:auto_http_config", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:circuit_breaker", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:default_circuit_breaker", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_circuit_breaker", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_outlier_detection", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_proxy_protocol", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_subsets", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_lb_source_ip_persistence", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http2_options", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:no_panic_threshold", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:no_request_limit_per_connection", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:outlier_detection", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:proxy_protocol_v1", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:proxy_protocol_v2"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "docs/guides/data-sources--origin_pool--properties--advanced_options.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["advanced_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- advanced_options

<a id="section"></a>

Type: `"single"`. Computed.

Configure Advanced OPTIONS for origin pool.

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

## Direct properties

- [auto_http_config](data-sources--origin_pool--properties--advanced_options--auto_http_config.md): complete subsection reference.

- [circuit_breaker](data-sources--origin_pool--properties--advanced_options--circuit_breaker.md): complete subsection reference.

<a id="schema-advanced_options--connection_timeout"></a>

### connection_timeout property

Type: `"number"`. Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds. Server applies default when omitted. Recommended:
\`2000\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

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

- [default_circuit_breaker](data-sources--origin_pool--properties--advanced_options--default_circuit_breaker.md): complete subsection reference.

- [disable_circuit_breaker](data-sources--origin_pool--properties--advanced_options--disable_circuit_breaker.md): complete subsection reference.

- [disable_lb_source_ip_persistence](data-sources--origin_pool--properties--advanced_options--disable_lb_source_ip_persistence.md): complete subsection reference.

- [disable_outlier_detection](data-sources--origin_pool--properties--advanced_options--disable_outlier_detection.md): complete subsection reference.

- [disable_proxy_protocol](data-sources--origin_pool--properties--advanced_options--disable_proxy_protocol.md): complete subsection reference.

- [disable_subsets](data-sources--origin_pool--properties--advanced_options--disable_subsets.md): complete subsection reference.

- [enable_lb_source_ip_persistence](data-sources--origin_pool--properties--advanced_options--enable_lb_source_ip_persistence.md): complete subsection reference.

- [enable_subsets](data-sources--origin_pool--properties--advanced_options--enable_subsets.md): complete subsection reference.

- [http1_config](data-sources--origin_pool--properties--advanced_options--http1_config.md): complete subsection reference.

- [http2_options](data-sources--origin_pool--properties--advanced_options--http2_options.md): complete subsection reference.

<a id="schema-advanced_options--http_idle_timeout"></a>

### http_idle_timeout property

Type: `"number"`. Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Server applies default when omitted. Recommended: \`300000\`.

Upstream description:

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
This is specified in milliseconds. The default value is 5 minutes.

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

<a id="schema-advanced_options--max_requests_per_connection"></a>

### max_requests_per_connection property

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests allowed
per connection to the origin server. Enter a value &gt;=1 to define the request limit per
connection.

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

- [no_panic_threshold](data-sources--origin_pool--properties--advanced_options--no_panic_threshold.md): complete subsection reference.

- [no_request_limit_per_connection](data-sources--origin_pool--properties--advanced_options--no_request_limit_per_connection.md): complete subsection reference.

- [outlier_detection](data-sources--origin_pool--properties--advanced_options--outlier_detection.md): complete subsection reference.

<a id="schema-advanced_options--panic_threshold"></a>

### panic_threshold property

Type: `"number"`. Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for load balancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for load balancing ignoring its health status.

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

- [proxy_protocol_v1](data-sources--origin_pool--properties--advanced_options--proxy_protocol_v1.md): complete subsection reference.

- [proxy_protocol_v2](data-sources--origin_pool--properties--advanced_options--proxy_protocol_v2.md): complete subsection reference.

## Next pages

- [advanced_options.auto_http_config](data-sources--origin_pool--properties--advanced_options--auto_http_config.md)
- [advanced_options.circuit_breaker](data-sources--origin_pool--properties--advanced_options--circuit_breaker.md)
- [advanced_options.default_circuit_breaker](data-sources--origin_pool--properties--advanced_options--default_circuit_breaker.md)
- [advanced_options.disable_circuit_breaker](data-sources--origin_pool--properties--advanced_options--disable_circuit_breaker.md)
- [advanced_options.disable_lb_source_ip_persistence](data-sources--origin_pool--properties--advanced_options--disable_lb_source_ip_persistence.md)
- [advanced_options.disable_outlier_detection](data-sources--origin_pool--properties--advanced_options--disable_outlier_detection.md)
- [advanced_options.disable_proxy_protocol](data-sources--origin_pool--properties--advanced_options--disable_proxy_protocol.md)
- [advanced_options.disable_subsets](data-sources--origin_pool--properties--advanced_options--disable_subsets.md)
- [advanced_options.enable_lb_source_ip_persistence](data-sources--origin_pool--properties--advanced_options--enable_lb_source_ip_persistence.md)
- [advanced_options.enable_subsets](data-sources--origin_pool--properties--advanced_options--enable_subsets.md)
- [advanced_options.http1_config](data-sources--origin_pool--properties--advanced_options--http1_config.md)
- [advanced_options.http2_options](data-sources--origin_pool--properties--advanced_options--http2_options.md)
- [advanced_options.no_panic_threshold](data-sources--origin_pool--properties--advanced_options--no_panic_threshold.md)
- [advanced_options.no_request_limit_per_connection](data-sources--origin_pool--properties--advanced_options--no_request_limit_per_connection.md)
- [advanced_options.outlier_detection](data-sources--origin_pool--properties--advanced_options--outlier_detection.md)
- [advanced_options.proxy_protocol_v1](data-sources--origin_pool--properties--advanced_options--proxy_protocol_v1.md)
- [advanced_options.proxy_protocol_v2](data-sources--origin_pool--properties--advanced_options--proxy_protocol_v2.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
