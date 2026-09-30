---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 43230, "body_sha256": "sha256:f846c26cf2cce053dae74ae7025dcfd9923cc8d5e66cda6dfbed6640f1d129c8", "canonical_id": "xcsh-docs:data-sources:cluster:reference", "child_ids": ["xcsh-docs:data-sources:cluster:properties:auto_http_config", "xcsh-docs:data-sources:cluster:properties:circuit_breaker", "xcsh-docs:data-sources:cluster:properties:default_subset", "xcsh-docs:data-sources:cluster:properties:disable_proxy_protocol", "xcsh-docs:data-sources:cluster:properties:endpoint_subsets", "xcsh-docs:data-sources:cluster:properties:endpoints", "xcsh-docs:data-sources:cluster:properties:health_checks", "xcsh-docs:data-sources:cluster:properties:http1_config", "xcsh-docs:data-sources:cluster:properties:http2_options", "xcsh-docs:data-sources:cluster:properties:no_panic_threshold", "xcsh-docs:data-sources:cluster:properties:no_request_limit_per_connection", "xcsh-docs:data-sources:cluster:properties:outlier_detection", "xcsh-docs:data-sources:cluster:properties:proxy_protocol_v1", "xcsh-docs:data-sources:cluster:properties:proxy_protocol_v2", "xcsh-docs:data-sources:cluster:properties:tls_parameters", "xcsh-docs:data-sources:cluster:properties:upstream_conn_pool_reuse_type"], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:reference", "parent_id": "xcsh-docs:data-sources:cluster:fundamentals", "path": "docs/guides/data-sources--cluster--reference.md", "provider_name": "cluster", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [auto_http_config](data-sources--cluster--properties--auto_http_config.md): complete subsection reference.

- [circuit_breaker](data-sources--cluster--properties--circuit_breaker.md): complete subsection reference.

<a id="schema-connection_timeout"></a>

### connection_timeout property

Type: `"number"`. Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The seconds. Defaults to \`2\`.

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

- [default_subset](data-sources--cluster--properties--default_subset.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Cluster.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [disable_proxy_protocol](data-sources--cluster--properties--disable_proxy_protocol.md): complete subsection reference.

<a id="schema-endpoint_selection"></a>

### endpoint_selection property

Type: `"string"`. Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`.

Upstream description:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Receipt-pinned upstream constraints:

```json
{
  "default": "DISTRIBUTED",
  "enum": [
    "DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [endpoint_subsets](data-sources--cluster--properties--endpoint_subsets.md): complete subsection reference.

- [endpoints](data-sources--cluster--properties--endpoints.md): complete subsection reference.

<a id="schema-fallback_policy"></a>

### fallback_policy property

Type: `"string"`. Computed.

\[Enum: NO\_FALLBACK|ANY\_ENDPOINT|DEFAULT\_SUBSET\] Enumeration for SubsetFallbackPolicy if subset
match is not found. The request fails as if the cluster had no endpoint matching the subset policy
Any cluster endpoint may be selected if the cluster had no endpoint matching the subset policy Load
balancing is done over endpoints matching.. Possible values are \`NO\_FALLBACK\`, \`ANY\_ENDPOINT\`,
\`DEFAULT\_SUBSET\`. Defaults to \`NO\_FALLBACK\`.

Upstream description:

Enumeration for SubsetFallbackPolicy if subset match is not found.

The request fails as if the cluster had no endpoint matching the subset policy Any cluster endpoint
may be selected if the cluster had no endpoint matching the subset policy Load balancing is done
over endpoints matching default\_subset if the cluster had no endpoint matching the subset policy.

Receipt-pinned upstream constraints:

```json
{
  "default": "NO_FALLBACK",
  "enum": [
    "NO_FALLBACK",
    "ANY_ENDPOINT",
    "DEFAULT_SUBSET"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [health_checks](data-sources--cluster--properties--health_checks.md): complete subsection reference.

- [http1_config](data-sources--cluster--properties--http1_config.md): complete subsection reference.

- [http2_options](data-sources--cluster--properties--http2_options.md): complete subsection reference.

<a id="schema-http_idle_timeout"></a>

### http_idle_timeout property

Type: `"number"`. Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-loadbalancer_algorithm"></a>

### loadbalancer_algorithm property

Type: `"string"`. Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`.

Upstream description:

Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster
is required, the load balancer uses loadbalancer\_algorithm to determine which host is selected.

&#8203;- ROUND\_ROBIN: ROUND\_ROBIN

Policy in which each healthy/available upstream endpoint is selected in round robin order. &#8203;-
LEAST\_REQUEST: LEAST\_REQUEST

Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests
&#8203;- RING\_HASH: RING\_HASH

Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of
the incoming request is calculated using request hash policy. The ring/modulo hash load balancer
implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a
circle such that the addition or removal of a host from the host set changes only affect 1/N
requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load
balancer is only effective when protocol routing is used that specifies a value to hash on. The
minimum ring size governs the replication factor for each host in the ring. For example, if the
minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times. &#8203;-
RANDOM: RANDOM

Policy in which each available upstream endpoint is selected in random order. The random load
balancer selects a random healthy host. The random load balancer generally performs better than
round robin if no health checking policy is configured. Random selection avoids bias towards the
host in the set that comes after a failed host. &#8203;- LB\_OVERRIDE: Load Balancer Override

Hash policy is taken from from the load balancer which is using this origin pool.

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-max_requests_per_connection"></a>

### max_requests_per_connection property

Type: `"number"`. Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests allowed per connection to the origin server. Enter a value &gt;=1 to
define the request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\]

Sets the maximum number of requests allowed per connection to the origin server. Enter a value
&gt;=1 to define the request limit per connection.

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

OneOf alternatives in this subsection:

- [max_requests_per_connection](data-sources--cluster--reference.md#schema-max_requests_per_connection)
- [no_request_limit_per_connection](data-sources--cluster--properties--no_request_limit_per_connection.md#section)

Select alternatives according to the provider validators above.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Cluster.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace where the Cluster exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_panic_threshold](data-sources--cluster--properties--no_panic_threshold.md): complete subsection reference.

- [no_request_limit_per_connection](data-sources--cluster--properties--no_request_limit_per_connection.md): complete subsection reference.

- [outlier_detection](data-sources--cluster--properties--outlier_detection.md): complete subsection reference.

<a id="schema-panic_threshold"></a>

### panic_threshold property

Type: `"number"`. Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for loadbalancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for loadbalancing ignoring its health status.

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

- [proxy_protocol_v1](data-sources--cluster--properties--proxy_protocol_v1.md): complete subsection reference.

- [proxy_protocol_v2](data-sources--cluster--properties--proxy_protocol_v2.md): complete subsection reference.

- [tls_parameters](data-sources--cluster--properties--tls_parameters.md): complete subsection reference.

- [upstream_conn_pool_reuse_type](data-sources--cluster--properties--upstream_conn_pool_reuse_type.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cluster--reference.md#schema-annotations) |
| `auto_http_config` | [auto_http_config](data-sources--cluster--properties--auto_http_config.md#section) |
| `circuit_breaker` | [circuit_breaker](data-sources--cluster--properties--circuit_breaker.md#section) |
| `circuit_breaker.connection_limit` | [circuit_breaker.connection_limit](data-sources--cluster--properties--circuit_breaker.md#schema-circuit_breaker--connection_limit) |
| `circuit_breaker.max_requests` | [circuit_breaker.max_requests](data-sources--cluster--properties--circuit_breaker.md#schema-circuit_breaker--max_requests) |
| `circuit_breaker.pending_requests` | [circuit_breaker.pending_requests](data-sources--cluster--properties--circuit_breaker.md#schema-circuit_breaker--pending_requests) |
| `circuit_breaker.priority` | [circuit_breaker.priority](data-sources--cluster--properties--circuit_breaker.md#schema-circuit_breaker--priority) |
| `circuit_breaker.retries` | [circuit_breaker.retries](data-sources--cluster--properties--circuit_breaker.md#schema-circuit_breaker--retries) |
| `connection_timeout` | [connection_timeout](data-sources--cluster--reference.md#schema-connection_timeout) |
| `default_subset` | [default_subset](data-sources--cluster--properties--default_subset.md#section) |
| `description` | [description](data-sources--cluster--reference.md#schema-description) |
| `disable_proxy_protocol` | [disable_proxy_protocol](data-sources--cluster--properties--disable_proxy_protocol.md#section) |
| `endpoint_selection` | [endpoint_selection](data-sources--cluster--reference.md#schema-endpoint_selection) |
| `endpoint_subsets` | [endpoint_subsets](data-sources--cluster--properties--endpoint_subsets.md#section) |
| `endpoint_subsets.keys` | [endpoint_subsets.keys](data-sources--cluster--properties--endpoint_subsets.md#schema-endpoint_subsets--keys) |
| `endpoints` | [endpoints](data-sources--cluster--properties--endpoints.md#section) |
| `endpoints.kind` | [endpoints.kind](data-sources--cluster--properties--endpoints.md#schema-endpoints--kind) |
| `endpoints.name` | [endpoints.name](data-sources--cluster--properties--endpoints.md#schema-endpoints--name) |
| `endpoints.namespace` | [endpoints.namespace](data-sources--cluster--properties--endpoints.md#schema-endpoints--namespace) |
| `endpoints.tenant` | [endpoints.tenant](data-sources--cluster--properties--endpoints.md#schema-endpoints--tenant) |
| `endpoints.uid` | [endpoints.uid](data-sources--cluster--properties--endpoints.md#schema-endpoints--uid) |
| `fallback_policy` | [fallback_policy](data-sources--cluster--reference.md#schema-fallback_policy) |
| `health_checks` | [health_checks](data-sources--cluster--properties--health_checks.md#section) |
| `health_checks.kind` | [health_checks.kind](data-sources--cluster--properties--health_checks.md#schema-health_checks--kind) |
| `health_checks.name` | [health_checks.name](data-sources--cluster--properties--health_checks.md#schema-health_checks--name) |
| `health_checks.namespace` | [health_checks.namespace](data-sources--cluster--properties--health_checks.md#schema-health_checks--namespace) |
| `health_checks.tenant` | [health_checks.tenant](data-sources--cluster--properties--health_checks.md#schema-health_checks--tenant) |
| `health_checks.uid` | [health_checks.uid](data-sources--cluster--properties--health_checks.md#schema-health_checks--uid) |
| `http1_config` | [http1_config](data-sources--cluster--properties--http1_config.md#section) |
| `http1_config.header_transformation` | [http1_config.header_transformation](data-sources--cluster--properties--http1_config--header_transformation.md#section) |
| `http1_config.header_transformation.default_header_transformation` | [http1_config.header_transformation.default_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--default_header_transformation.md#section) |
| `http1_config.header_transformation.preserve_case_header_transformation` | [http1_config.header_transformation.preserve_case_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--preserve_case_header_transformation.md#section) |
| `http1_config.header_transformation.proper_case_header_transformation` | [http1_config.header_transformation.proper_case_header_transformation](data-sources--cluster--properties--http1_config--header_transformation--proper_case_header_transformation.md#section) |
| `http2_options` | [http2_options](data-sources--cluster--properties--http2_options.md#section) |
| `http2_options.enabled` | [http2_options.enabled](data-sources--cluster--properties--http2_options.md#schema-http2_options--enabled) |
| `http_idle_timeout` | [http_idle_timeout](data-sources--cluster--reference.md#schema-http_idle_timeout) |
| `id` | [id](data-sources--cluster--reference.md#schema-id) |
| `labels` | [labels](data-sources--cluster--reference.md#schema-labels) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](data-sources--cluster--reference.md#schema-loadbalancer_algorithm) |
| `max_requests_per_connection` | [max_requests_per_connection](data-sources--cluster--reference.md#schema-max_requests_per_connection) |
| `name` | [name](data-sources--cluster--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--cluster--reference.md#schema-namespace) |
| `no_panic_threshold` | [no_panic_threshold](data-sources--cluster--properties--no_panic_threshold.md#section) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](data-sources--cluster--properties--no_request_limit_per_connection.md#section) |
| `outlier_detection` | [outlier_detection](data-sources--cluster--properties--outlier_detection.md#section) |
| `outlier_detection.base_ejection_time` | [outlier_detection.base_ejection_time](data-sources--cluster--properties--outlier_detection.md#schema-outlier_detection--base_ejection_time) |
| `outlier_detection.consecutive_5xx` | [outlier_detection.consecutive_5xx](data-sources--cluster--properties--outlier_detection.md#schema-outlier_detection--consecutive_5xx) |
| `outlier_detection.consecutive_gateway_failure` | [outlier_detection.consecutive_gateway_failure](data-sources--cluster--properties--outlier_detection.md#schema-outlier_detection--consecutive_gateway_failure) |
| `outlier_detection.interval` | [outlier_detection.interval](data-sources--cluster--properties--outlier_detection.md#schema-outlier_detection--interval) |
| `outlier_detection.max_ejection_percent` | [outlier_detection.max_ejection_percent](data-sources--cluster--properties--outlier_detection.md#schema-outlier_detection--max_ejection_percent) |
| `panic_threshold` | [panic_threshold](data-sources--cluster--reference.md#schema-panic_threshold) |
| `proxy_protocol_v1` | [proxy_protocol_v1](data-sources--cluster--properties--proxy_protocol_v1.md#section) |
| `proxy_protocol_v2` | [proxy_protocol_v2](data-sources--cluster--properties--proxy_protocol_v2.md#section) |
| `tls_parameters` | [tls_parameters](data-sources--cluster--properties--tls_parameters.md#section) |
| `tls_parameters.cert_params` | [tls_parameters.cert_params](data-sources--cluster--properties--tls_parameters--cert_params.md#section) |
| `tls_parameters.cert_params.certificates` | [tls_parameters.cert_params.certificates](data-sources--cluster--properties--tls_parameters--cert_params--certificates.md#section) |
| `tls_parameters.cert_params.certificates.kind` | [tls_parameters.cert_params.certificates.kind](data-sources--cluster--properties--tls_parameters--cert_params--certificates.md#schema-tls_parameters--cert_params--certificates--kind) |
| `tls_parameters.cert_params.certificates.name` | [tls_parameters.cert_params.certificates.name](data-sources--cluster--properties--tls_parameters--cert_params--certificates.md#schema-tls_parameters--cert_params--certificates--name) |
| `tls_parameters.cert_params.certificates.namespace` | [tls_parameters.cert_params.certificates.namespace](data-sources--cluster--properties--tls_parameters--cert_params--certificates.md#schema-tls_parameters--cert_params--certificates--namespace) |
| `tls_parameters.cert_params.certificates.tenant` | [tls_parameters.cert_params.certificates.tenant](data-sources--cluster--properties--tls_parameters--cert_params--certificates.md#schema-tls_parameters--cert_params--certificates--tenant) |
| `tls_parameters.cert_params.certificates.uid` | [tls_parameters.cert_params.certificates.uid](data-sources--cluster--properties--tls_parameters--cert_params--certificates.md#schema-tls_parameters--cert_params--certificates--uid) |
| `tls_parameters.cert_params.cipher_suites` | [tls_parameters.cert_params.cipher_suites](data-sources--cluster--properties--tls_parameters--cert_params.md#schema-tls_parameters--cert_params--cipher_suites) |
| `tls_parameters.cert_params.maximum_protocol_version` | [tls_parameters.cert_params.maximum_protocol_version](data-sources--cluster--properties--tls_parameters--cert_params.md#schema-tls_parameters--cert_params--maximum_protocol_version) |
| `tls_parameters.cert_params.minimum_protocol_version` | [tls_parameters.cert_params.minimum_protocol_version](data-sources--cluster--properties--tls_parameters--cert_params.md#schema-tls_parameters--cert_params--minimum_protocol_version) |
| `tls_parameters.cert_params.skip_server_verification` | [tls_parameters.cert_params.skip_server_verification](data-sources--cluster--properties--tls_parameters--cert_params--skip_server_verification.md#section) |
| `tls_parameters.cert_params.tls_validation_params` | [tls_parameters.cert_params.tls_validation_params](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params.md#section) |
| `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` | [tls_parameters.cert_params.tls_validation_params.skip_hostname_verification](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params.md#schema-tls_parameters--cert_params--tls_validation_params--skip_hostname_verification) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca` | [tls_parameters.cert_params.tls_validation_params.trusted_ca](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params--trusted_ca.md#section) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#section) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` | [tls_parameters.cert_params.tls_validation_params.trusted_ca_url](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params.md#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca_url) |
| `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` | [tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names](data-sources--cluster--properties--tls_parameters--cert_params--tls_validation_params.md#schema-tls_parameters--cert_params--tls_validation_params--verify_subject_alt_names) |
| `tls_parameters.cert_params.volterra_trusted_ca` | [tls_parameters.cert_params.volterra_trusted_ca](data-sources--cluster--properties--tls_parameters--cert_params--volterra_trusted_ca.md#section) |
| `tls_parameters.common_params` | [tls_parameters.common_params](data-sources--cluster--properties--tls_parameters--common_params.md#section) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](data-sources--cluster--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--cipher_suites) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](data-sources--cluster--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--maximum_protocol_version) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](data-sources--cluster--properties--tls_parameters--common_params.md#schema-tls_parameters--common_params--minimum_protocol_version) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates.md#section) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates.md#schema-tls_parameters--common_params--tls_certificates--certificate_url) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--custom_hash_algorithms.md#section) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--custom_hash_algorithms.md#schema-tls_parameters--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates.md#schema-tls_parameters--common_params--tls_certificates--description_spec) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--disable_ocsp_stapling.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#section) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--url) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--use_system_defaults.md#section) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](data-sources--cluster--properties--tls_parameters--common_params--validation_params.md#section) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](data-sources--cluster--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--skip_hostname_verification) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](data-sources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca.md#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--cluster--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](data-sources--cluster--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--trusted_ca_url) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](data-sources--cluster--properties--tls_parameters--common_params--validation_params.md#schema-tls_parameters--common_params--validation_params--verify_subject_alt_names) |
| `tls_parameters.default_session_key_caching` | [tls_parameters.default_session_key_caching](data-sources--cluster--properties--tls_parameters--default_session_key_caching.md#section) |
| `tls_parameters.disable_session_key_caching` | [tls_parameters.disable_session_key_caching](data-sources--cluster--properties--tls_parameters--disable_session_key_caching.md#section) |
| `tls_parameters.disable_sni` | [tls_parameters.disable_sni](data-sources--cluster--properties--tls_parameters--disable_sni.md#section) |
| `tls_parameters.max_session_keys` | [tls_parameters.max_session_keys](data-sources--cluster--properties--tls_parameters.md#schema-tls_parameters--max_session_keys) |
| `tls_parameters.sni` | [tls_parameters.sni](data-sources--cluster--properties--tls_parameters.md#schema-tls_parameters--sni) |
| `tls_parameters.use_host_header_as_sni` | [tls_parameters.use_host_header_as_sni](data-sources--cluster--properties--tls_parameters--use_host_header_as_sni.md#section) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](data-sources--cluster--properties--upstream_conn_pool_reuse_type.md#section) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--cluster--properties--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md#section) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--cluster--properties--upstream_conn_pool_reuse_type--enable_conn_pool_reuse.md#section) |

## Next pages

- [auto_http_config](data-sources--cluster--properties--auto_http_config.md)
- [circuit_breaker](data-sources--cluster--properties--circuit_breaker.md)
- [default_subset](data-sources--cluster--properties--default_subset.md)
- [disable_proxy_protocol](data-sources--cluster--properties--disable_proxy_protocol.md)
- [endpoint_subsets](data-sources--cluster--properties--endpoint_subsets.md)
- [endpoints](data-sources--cluster--properties--endpoints.md)
- [health_checks](data-sources--cluster--properties--health_checks.md)
- [http1_config](data-sources--cluster--properties--http1_config.md)
- [http2_options](data-sources--cluster--properties--http2_options.md)
- [no_panic_threshold](data-sources--cluster--properties--no_panic_threshold.md)
- [no_request_limit_per_connection](data-sources--cluster--properties--no_request_limit_per_connection.md)
- [outlier_detection](data-sources--cluster--properties--outlier_detection.md)
- [proxy_protocol_v1](data-sources--cluster--properties--proxy_protocol_v1.md)
- [proxy_protocol_v2](data-sources--cluster--properties--proxy_protocol_v2.md)
- [tls_parameters](data-sources--cluster--properties--tls_parameters.md)
- [upstream_conn_pool_reuse_type](data-sources--cluster--properties--upstream_conn_pool_reuse_type.md)
- [xcsh_cluster](../data-sources/cluster.md)
