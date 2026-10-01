---
page_title: "advanced_options"
subcategory: "Load Balancing"
description: "advanced_options for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 12725, "body_sha256": "sha256:2b11734d35e4b6874a60f783a90dbc1a1818b659dc88eb2d5dbe2cdad22280fc", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:advanced_options:auto_http_config", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:circuit_breaker", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:default_circuit_breaker", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_circuit_breaker", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_outlier_detection", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_proxy_protocol", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:disable_subsets", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_lb_source_ip_persistence", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:enable_subsets", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http1_config", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:http2_options", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:no_panic_threshold", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:no_request_limit_per_connection", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:outlier_detection", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:proxy_protocol_v1", "xcsh-docs:data-sources:origin_pool:properties:advanced_options:proxy_protocol_v2"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:advanced_options", "parent_id": "xcsh-docs:data-sources:origin_pool:reference", "path": "documentation/data-sources/origin_pool/properties/advanced_options/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["advanced_options"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/advanced_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "advanced_options for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
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

- [auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/auto_http_config/): complete subsection reference.

- [circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/circuit_breaker/): complete subsection reference.

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

- [default_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/default_circuit_breaker/): complete subsection reference.

- [disable_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_circuit_breaker/): complete subsection reference.

- [disable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_lb_source_ip_persistence/): complete subsection reference.

- [disable_outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_outlier_detection/): complete subsection reference.

- [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_proxy_protocol/): complete subsection reference.

- [disable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_subsets/): complete subsection reference.

- [enable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_lb_source_ip_persistence/): complete subsection reference.

- [enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/): complete subsection reference.

- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/): complete subsection reference.

- [http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http2_options/): complete subsection reference.

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

- [no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/no_panic_threshold/): complete subsection reference.

- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/no_request_limit_per_connection/): complete subsection reference.

- [outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/outlier_detection/): complete subsection reference.

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

- [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/proxy_protocol_v1/): complete subsection reference.

- [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/proxy_protocol_v2/): complete subsection reference.

## Next pages

- [advanced_options.auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/auto_http_config/)
- [advanced_options.circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/circuit_breaker/)
- [advanced_options.default_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/default_circuit_breaker/)
- [advanced_options.disable_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_circuit_breaker/)
- [advanced_options.disable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_lb_source_ip_persistence/)
- [advanced_options.disable_outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_outlier_detection/)
- [advanced_options.disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_proxy_protocol/)
- [advanced_options.disable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/disable_subsets/)
- [advanced_options.enable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_lb_source_ip_persistence/)
- [advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/enable_subsets/)
- [advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http1_config/)
- [advanced_options.http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/http2_options/)
- [advanced_options.no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/no_panic_threshold/)
- [advanced_options.no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/no_request_limit_per_connection/)
- [advanced_options.outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/outlier_detection/)
- [advanced_options.proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/proxy_protocol_v1/)
- [advanced_options.proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/advanced_options/proxy_protocol_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/)
