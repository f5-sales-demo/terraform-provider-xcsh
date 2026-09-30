---
page_title: "Property reference"
subcategory: "Load Balancing"
description: "Property reference for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 59828, "body_sha256": "sha256:3885531a066745708c563831524e8c57ce48f0b3665876f36ba3e3b6e7f38e6a", "canonical_id": "xcsh-docs:resources:origin_pool:reference", "child_ids": ["xcsh-docs:resources:origin_pool:properties:advanced_options", "xcsh-docs:resources:origin_pool:properties:automatic_port", "xcsh-docs:resources:origin_pool:properties:healthcheck", "xcsh-docs:resources:origin_pool:properties:lb_port", "xcsh-docs:resources:origin_pool:properties:no_tls", "xcsh-docs:resources:origin_pool:properties:origin_servers", "xcsh-docs:resources:origin_pool:properties:same_as_endpoint_port", "xcsh-docs:resources:origin_pool:properties:timeouts", "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type", "xcsh-docs:resources:origin_pool:properties:use_tls"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:reference", "parent_id": "xcsh-docs:resources:origin_pool:fundamentals", "path": "docs/guides/resources--origin_pool--reference.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- Property reference

## Direct properties

- [advanced_options](resources--origin_pool--properties--advanced_options.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [automatic_port](resources--origin_pool--properties--automatic_port.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

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

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-endpoint_selection"></a>

### endpoint_selection property

Type: `"string"`. Optional, Computed.

\[Enum: DISTRIBUTED|LOCAL\_ONLY|LOCAL\_PREFERRED\] Policy for selection of endpoints from local
site/remote site/both Consider both remote and local endpoints for load balancing LOCAL\_ONLY:
Consider only local endpoints for load balancing Enable this policy to load balance ONLY among
locally discovered endpoints Prefer the local endpoints for.. Possible values are \`DISTRIBUTED\`,
\`LOCAL\_ONLY\`, \`LOCAL\_PREFERRED\`. Defaults to \`DISTRIBUTED\`. Server applies default when
omitted.

Upstream description:

Policy for selection of endpoints from local site/remote site/both

Consider both remote and local endpoints for load balancing LOCAL\_ONLY: Consider only local
endpoints for load balancing Enable this policy to load balance ONLY among locally discovered
endpoints Prefer the local endpoints for load balancing. If local endpoints are not present remote
endpoints will be considered.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DISTRIBUTED",
    "LOCAL_ONLY",
    "LOCAL_PREFERRED"),
}
```

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

<a id="schema-health_check_port"></a>

### health_check_port property

Type: `"number"`. Optional, Computed.

\[OneOf: health\_check\_port, same\_as\_endpoint\_port\] Exclusive with \[same\_as\_endpoint\_port\]
Port used for performing health check.

Upstream description:

Exclusive with \[same\_as\_endpoint\_port\] Port used for performing health check.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "category": "networking",
      "confidence": 0.99,
      "note": "Asymmetry: port enforces [1,65535], health_check_port allows 0",
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [health_check_port](resources--origin_pool--reference.md#schema-health_check_port)
- [same_as_endpoint_port](resources--origin_pool--properties--same_as_endpoint_port.md#section)

Select alternatives according to the provider validators above.

- [healthcheck](resources--origin_pool--properties--healthcheck.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

- [lb_port](resources--origin_pool--properties--lb_port.md): complete subsection reference.

<a id="schema-loadbalancer_algorithm"></a>

### loadbalancer_algorithm property

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|LEAST\_REQUEST|RING\_HASH|RANDOM|LB\_OVERRIDE\] Different load balancing
algorithms supported When a connection to a endpoint in an upstream cluster is required, the load
balancer uses loadbalancer\_algorithm to determine which host is selected. - ROUND\_ROBIN:
ROUND\_ROBIN Policy in which each healthy/available upstream endpoint is selected in.. Possible
values are \`ROUND\_ROBIN\`, \`LEAST\_REQUEST\`, \`RING\_HASH\`, \`RANDOM\`, \`LB\_OVERRIDE\`.
Defaults to \`ROUND\_ROBIN\`. Server applies default when omitted.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ROUND_ROBIN",
    "LEAST_REQUEST",
    "RING_HASH",
    "RANDOM",
    "LB_OVERRIDE"),
}
```

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Origin Pool. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

Namespace where the Origin Pool is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [no_tls](resources--origin_pool--properties--no_tls.md): complete subsection reference.

- [origin_servers](resources--origin_pool--properties--origin_servers.md): complete subsection reference.

<a id="schema-port"></a>

### port property

Type: `"number"`. Optional, Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port. Recommended:
\`443\`.

Upstream description:

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [same_as_endpoint_port](resources--origin_pool--properties--same_as_endpoint_port.md): complete subsection reference.

- [timeouts](resources--origin_pool--properties--timeouts.md): complete subsection reference.

- [upstream_conn_pool_reuse_type](resources--origin_pool--properties--upstream_conn_pool_reuse_type.md): complete subsection reference.

- [use_tls](resources--origin_pool--properties--use_tls.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_options` | [advanced_options](resources--origin_pool--properties--advanced_options.md#section) |
| `advanced_options.auto_http_config` | [advanced_options.auto_http_config](resources--origin_pool--properties--advanced_options--auto_http_config.md#section) |
| `advanced_options.circuit_breaker` | [advanced_options.circuit_breaker](resources--origin_pool--properties--advanced_options--circuit_breaker.md#section) |
| `advanced_options.circuit_breaker.connection_limit` | [advanced_options.circuit_breaker.connection_limit](resources--origin_pool--properties--advanced_options--circuit_breaker.md#schema-advanced_options--circuit_breaker--connection_limit) |
| `advanced_options.circuit_breaker.max_requests` | [advanced_options.circuit_breaker.max_requests](resources--origin_pool--properties--advanced_options--circuit_breaker.md#schema-advanced_options--circuit_breaker--max_requests) |
| `advanced_options.circuit_breaker.pending_requests` | [advanced_options.circuit_breaker.pending_requests](resources--origin_pool--properties--advanced_options--circuit_breaker.md#schema-advanced_options--circuit_breaker--pending_requests) |
| `advanced_options.circuit_breaker.priority` | [advanced_options.circuit_breaker.priority](resources--origin_pool--properties--advanced_options--circuit_breaker.md#schema-advanced_options--circuit_breaker--priority) |
| `advanced_options.circuit_breaker.retries` | [advanced_options.circuit_breaker.retries](resources--origin_pool--properties--advanced_options--circuit_breaker.md#schema-advanced_options--circuit_breaker--retries) |
| `advanced_options.connection_timeout` | [advanced_options.connection_timeout](resources--origin_pool--properties--advanced_options.md#schema-advanced_options--connection_timeout) |
| `advanced_options.default_circuit_breaker` | [advanced_options.default_circuit_breaker](resources--origin_pool--properties--advanced_options--default_circuit_breaker.md#section) |
| `advanced_options.disable_circuit_breaker` | [advanced_options.disable_circuit_breaker](resources--origin_pool--properties--advanced_options--disable_circuit_breaker.md#section) |
| `advanced_options.disable_lb_source_ip_persistence` | [advanced_options.disable_lb_source_ip_persistence](resources--origin_pool--properties--advanced_options--disable_lb_source_ip_persistence.md#section) |
| `advanced_options.disable_outlier_detection` | [advanced_options.disable_outlier_detection](resources--origin_pool--properties--advanced_options--disable_outlier_detection.md#section) |
| `advanced_options.disable_proxy_protocol` | [advanced_options.disable_proxy_protocol](resources--origin_pool--properties--advanced_options--disable_proxy_protocol.md#section) |
| `advanced_options.disable_subsets` | [advanced_options.disable_subsets](resources--origin_pool--properties--advanced_options--disable_subsets.md#section) |
| `advanced_options.enable_lb_source_ip_persistence` | [advanced_options.enable_lb_source_ip_persistence](resources--origin_pool--properties--advanced_options--enable_lb_source_ip_persistence.md#section) |
| `advanced_options.enable_subsets` | [advanced_options.enable_subsets](resources--origin_pool--properties--advanced_options--enable_subsets.md#section) |
| `advanced_options.enable_subsets.any_endpoint` | [advanced_options.enable_subsets.any_endpoint](resources--origin_pool--properties--advanced_options--enable_subsets--any_endpoint.md#section) |
| `advanced_options.enable_subsets.default_subset` | [advanced_options.enable_subsets.default_subset](resources--origin_pool--properties--advanced_options--enable_subsets--default_subset.md#section) |
| `advanced_options.enable_subsets.default_subset.default_subset` | [advanced_options.enable_subsets.default_subset.default_subset](resources--origin_pool--properties--advanced_options--enable_subsets--default_subset--default_subset.md#section) |
| `advanced_options.enable_subsets.endpoint_subsets` | [advanced_options.enable_subsets.endpoint_subsets](resources--origin_pool--properties--advanced_options--enable_subsets--endpoint_subsets.md#section) |
| `advanced_options.enable_subsets.endpoint_subsets.keys` | [advanced_options.enable_subsets.endpoint_subsets.keys](resources--origin_pool--properties--advanced_options--enable_subsets--endpoint_subsets.md#schema-advanced_options--enable_subsets--endpoint_subsets--keys) |
| `advanced_options.enable_subsets.fail_request` | [advanced_options.enable_subsets.fail_request](resources--origin_pool--properties--advanced_options--enable_subsets--fail_request.md#section) |
| `advanced_options.http1_config` | [advanced_options.http1_config](resources--origin_pool--properties--advanced_options--http1_config.md#section) |
| `advanced_options.http1_config.header_transformation` | [advanced_options.http1_config.header_transformation](resources--origin_pool--properties--advanced_options--http1_config--header_transformation.md#section) |
| `advanced_options.http1_config.header_transformation.default_header_transformation` | [advanced_options.http1_config.header_transformation.default_header_transformation](resources--origin_pool--properties--advanced_options--http1_config--header_transformation--default_header_transformation.md#section) |
| `advanced_options.http1_config.header_transformation.preserve_case_header_transformation` | [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](resources--origin_pool--properties--advanced_options--http1_config--header_transformation--preserve_case_header_transformation.md#section) |
| `advanced_options.http1_config.header_transformation.proper_case_header_transformation` | [advanced_options.http1_config.header_transformation.proper_case_header_transformation](resources--origin_pool--properties--advanced_options--http1_config--header_transformation--proper_case_header_transformation.md#section) |
| `advanced_options.http2_options` | [advanced_options.http2_options](resources--origin_pool--properties--advanced_options--http2_options.md#section) |
| `advanced_options.http2_options.enabled` | [advanced_options.http2_options.enabled](resources--origin_pool--properties--advanced_options--http2_options.md#schema-advanced_options--http2_options--enabled) |
| `advanced_options.http_idle_timeout` | [advanced_options.http_idle_timeout](resources--origin_pool--properties--advanced_options.md#schema-advanced_options--http_idle_timeout) |
| `advanced_options.max_requests_per_connection` | [advanced_options.max_requests_per_connection](resources--origin_pool--properties--advanced_options.md#schema-advanced_options--max_requests_per_connection) |
| `advanced_options.no_panic_threshold` | [advanced_options.no_panic_threshold](resources--origin_pool--properties--advanced_options--no_panic_threshold.md#section) |
| `advanced_options.no_request_limit_per_connection` | [advanced_options.no_request_limit_per_connection](resources--origin_pool--properties--advanced_options--no_request_limit_per_connection.md#section) |
| `advanced_options.outlier_detection` | [advanced_options.outlier_detection](resources--origin_pool--properties--advanced_options--outlier_detection.md#section) |
| `advanced_options.outlier_detection.base_ejection_time` | [advanced_options.outlier_detection.base_ejection_time](resources--origin_pool--properties--advanced_options--outlier_detection.md#schema-advanced_options--outlier_detection--base_ejection_time) |
| `advanced_options.outlier_detection.consecutive_5xx` | [advanced_options.outlier_detection.consecutive_5xx](resources--origin_pool--properties--advanced_options--outlier_detection.md#schema-advanced_options--outlier_detection--consecutive_5xx) |
| `advanced_options.outlier_detection.consecutive_gateway_failure` | [advanced_options.outlier_detection.consecutive_gateway_failure](resources--origin_pool--properties--advanced_options--outlier_detection.md#schema-advanced_options--outlier_detection--consecutive_gateway_failure) |
| `advanced_options.outlier_detection.interval` | [advanced_options.outlier_detection.interval](resources--origin_pool--properties--advanced_options--outlier_detection.md#schema-advanced_options--outlier_detection--interval) |
| `advanced_options.outlier_detection.max_ejection_percent` | [advanced_options.outlier_detection.max_ejection_percent](resources--origin_pool--properties--advanced_options--outlier_detection.md#schema-advanced_options--outlier_detection--max_ejection_percent) |
| `advanced_options.panic_threshold` | [advanced_options.panic_threshold](resources--origin_pool--properties--advanced_options.md#schema-advanced_options--panic_threshold) |
| `advanced_options.proxy_protocol_v1` | [advanced_options.proxy_protocol_v1](resources--origin_pool--properties--advanced_options--proxy_protocol_v1.md#section) |
| `advanced_options.proxy_protocol_v2` | [advanced_options.proxy_protocol_v2](resources--origin_pool--properties--advanced_options--proxy_protocol_v2.md#section) |
| `annotations` | [annotations](resources--origin_pool--reference.md#schema-annotations) |
| `automatic_port` | [automatic_port](resources--origin_pool--properties--automatic_port.md#section) |
| `description` | [description](resources--origin_pool--reference.md#schema-description) |
| `disable` | [disable](resources--origin_pool--reference.md#schema-disable) |
| `endpoint_selection` | [endpoint_selection](resources--origin_pool--reference.md#schema-endpoint_selection) |
| `health_check_port` | [health_check_port](resources--origin_pool--reference.md#schema-health_check_port) |
| `healthcheck` | [healthcheck](resources--origin_pool--properties--healthcheck.md#section) |
| `healthcheck.name` | [healthcheck.name](resources--origin_pool--properties--healthcheck.md#schema-healthcheck--name) |
| `healthcheck.namespace` | [healthcheck.namespace](resources--origin_pool--properties--healthcheck.md#schema-healthcheck--namespace) |
| `healthcheck.tenant` | [healthcheck.tenant](resources--origin_pool--properties--healthcheck.md#schema-healthcheck--tenant) |
| `id` | [id](resources--origin_pool--reference.md#schema-id) |
| `labels` | [labels](resources--origin_pool--reference.md#schema-labels) |
| `lb_port` | [lb_port](resources--origin_pool--properties--lb_port.md#section) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](resources--origin_pool--reference.md#schema-loadbalancer_algorithm) |
| `name` | [name](resources--origin_pool--reference.md#schema-name) |
| `namespace` | [namespace](resources--origin_pool--reference.md#schema-namespace) |
| `no_tls` | [no_tls](resources--origin_pool--properties--no_tls.md#section) |
| `origin_servers` | [origin_servers](resources--origin_pool--properties--origin_servers.md#section) |
| `origin_servers.cbip_service` | [origin_servers.cbip_service](resources--origin_pool--properties--origin_servers--cbip_service.md#section) |
| `origin_servers.cbip_service.service_name` | [origin_servers.cbip_service.service_name](resources--origin_pool--properties--origin_servers--cbip_service.md#schema-origin_servers--cbip_service--service_name) |
| `origin_servers.consul_service` | [origin_servers.consul_service](resources--origin_pool--properties--origin_servers--consul_service.md#section) |
| `origin_servers.consul_service.inside_network` | [origin_servers.consul_service.inside_network](resources--origin_pool--properties--origin_servers--consul_service--inside_network.md#section) |
| `origin_servers.consul_service.outside_network` | [origin_servers.consul_service.outside_network](resources--origin_pool--properties--origin_servers--consul_service--outside_network.md#section) |
| `origin_servers.consul_service.service_name` | [origin_servers.consul_service.service_name](resources--origin_pool--properties--origin_servers--consul_service.md#schema-origin_servers--consul_service--service_name) |
| `origin_servers.consul_service.site_locator` | [origin_servers.consul_service.site_locator](resources--origin_pool--properties--origin_servers--consul_service--site_locator.md#section) |
| `origin_servers.consul_service.site_locator.site` | [origin_servers.consul_service.site_locator.site](resources--origin_pool--properties--origin_servers--consul_service--site_locator--site.md#section) |
| `origin_servers.consul_service.site_locator.site.name` | [origin_servers.consul_service.site_locator.site.name](resources--origin_pool--properties--origin_servers--consul_service--site_locator--site.md#schema-origin_servers--consul_service--site_locator--site--name) |
| `origin_servers.consul_service.site_locator.site.namespace` | [origin_servers.consul_service.site_locator.site.namespace](resources--origin_pool--properties--origin_servers--consul_service--site_locator--site.md#schema-origin_servers--consul_service--site_locator--site--namespace) |
| `origin_servers.consul_service.site_locator.site.tenant` | [origin_servers.consul_service.site_locator.site.tenant](resources--origin_pool--properties--origin_servers--consul_service--site_locator--site.md#schema-origin_servers--consul_service--site_locator--site--tenant) |
| `origin_servers.consul_service.site_locator.virtual_site` | [origin_servers.consul_service.site_locator.virtual_site](resources--origin_pool--properties--origin_servers--consul_service--site_locator--virtual_site.md#section) |
| `origin_servers.consul_service.site_locator.virtual_site.name` | [origin_servers.consul_service.site_locator.virtual_site.name](resources--origin_pool--properties--origin_servers--consul_service--site_locator--virtual_site.md#schema-origin_servers--consul_service--site_locator--virtual_site--name) |
| `origin_servers.consul_service.site_locator.virtual_site.namespace` | [origin_servers.consul_service.site_locator.virtual_site.namespace](resources--origin_pool--properties--origin_servers--consul_service--site_locator--virtual_site.md#schema-origin_servers--consul_service--site_locator--virtual_site--namespace) |
| `origin_servers.consul_service.site_locator.virtual_site.tenant` | [origin_servers.consul_service.site_locator.virtual_site.tenant](resources--origin_pool--properties--origin_servers--consul_service--site_locator--virtual_site.md#schema-origin_servers--consul_service--site_locator--virtual_site--tenant) |
| `origin_servers.consul_service.snat_pool` | [origin_servers.consul_service.snat_pool](resources--origin_pool--properties--origin_servers--consul_service--snat_pool.md#section) |
| `origin_servers.consul_service.snat_pool.no_snat_pool` | [origin_servers.consul_service.snat_pool.no_snat_pool](resources--origin_pool--properties--origin_servers--consul_service--snat_pool--no_snat_pool.md#section) |
| `origin_servers.consul_service.snat_pool.snat_pool` | [origin_servers.consul_service.snat_pool.snat_pool](resources--origin_pool--properties--origin_servers--consul_service--snat_pool--snat_pool.md#section) |
| `origin_servers.consul_service.snat_pool.snat_pool.prefixes` | [origin_servers.consul_service.snat_pool.snat_pool.prefixes](resources--origin_pool--properties--origin_servers--consul_service--snat_pool--snat_pool.md#schema-origin_servers--consul_service--snat_pool--snat_pool--prefixes) |
| `origin_servers.custom_endpoint_object` | [origin_servers.custom_endpoint_object](resources--origin_pool--properties--origin_servers--custom_endpoint_object.md#section) |
| `origin_servers.custom_endpoint_object.endpoint` | [origin_servers.custom_endpoint_object.endpoint](resources--origin_pool--properties--origin_servers--custom_endpoint_object--endpoint.md#section) |
| `origin_servers.custom_endpoint_object.endpoint.name` | [origin_servers.custom_endpoint_object.endpoint.name](resources--origin_pool--properties--origin_servers--custom_endpoint_object--endpoint.md#schema-origin_servers--custom_endpoint_object--endpoint--name) |
| `origin_servers.custom_endpoint_object.endpoint.namespace` | [origin_servers.custom_endpoint_object.endpoint.namespace](resources--origin_pool--properties--origin_servers--custom_endpoint_object--endpoint.md#schema-origin_servers--custom_endpoint_object--endpoint--namespace) |
| `origin_servers.custom_endpoint_object.endpoint.tenant` | [origin_servers.custom_endpoint_object.endpoint.tenant](resources--origin_pool--properties--origin_servers--custom_endpoint_object--endpoint.md#schema-origin_servers--custom_endpoint_object--endpoint--tenant) |
| `origin_servers.k8s_service` | [origin_servers.k8s_service](resources--origin_pool--properties--origin_servers--k8s_service.md#section) |
| `origin_servers.k8s_service.inside_network` | [origin_servers.k8s_service.inside_network](resources--origin_pool--properties--origin_servers--k8s_service--inside_network.md#section) |
| `origin_servers.k8s_service.outside_network` | [origin_servers.k8s_service.outside_network](resources--origin_pool--properties--origin_servers--k8s_service--outside_network.md#section) |
| `origin_servers.k8s_service.protocol` | [origin_servers.k8s_service.protocol](resources--origin_pool--properties--origin_servers--k8s_service.md#schema-origin_servers--k8s_service--protocol) |
| `origin_servers.k8s_service.service_name` | [origin_servers.k8s_service.service_name](resources--origin_pool--properties--origin_servers--k8s_service.md#schema-origin_servers--k8s_service--service_name) |
| `origin_servers.k8s_service.site_locator` | [origin_servers.k8s_service.site_locator](resources--origin_pool--properties--origin_servers--k8s_service--site_locator.md#section) |
| `origin_servers.k8s_service.site_locator.site` | [origin_servers.k8s_service.site_locator.site](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--site.md#section) |
| `origin_servers.k8s_service.site_locator.site.name` | [origin_servers.k8s_service.site_locator.site.name](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--k8s_service--site_locator--site--name) |
| `origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.k8s_service.site_locator.site.namespace](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--k8s_service--site_locator--site--namespace) |
| `origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.k8s_service.site_locator.site.tenant](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--k8s_service--site_locator--site--tenant) |
| `origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.k8s_service.site_locator.virtual_site](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--virtual_site.md#section) |
| `origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.k8s_service.site_locator.virtual_site.name](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--k8s_service--site_locator--virtual_site--name) |
| `origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--k8s_service--site_locator--virtual_site--namespace) |
| `origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--origin_pool--properties--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--k8s_service--site_locator--virtual_site--tenant) |
| `origin_servers.k8s_service.snat_pool` | [origin_servers.k8s_service.snat_pool](resources--origin_pool--properties--origin_servers--k8s_service--snat_pool.md#section) |
| `origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.k8s_service.snat_pool.no_snat_pool](resources--origin_pool--properties--origin_servers--k8s_service--snat_pool--no_snat_pool.md#section) |
| `origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.k8s_service.snat_pool.snat_pool](resources--origin_pool--properties--origin_servers--k8s_service--snat_pool--snat_pool.md#section) |
| `origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--origin_pool--properties--origin_servers--k8s_service--snat_pool--snat_pool.md#schema-origin_servers--k8s_service--snat_pool--snat_pool--prefixes) |
| `origin_servers.k8s_service.vk8s_networks` | [origin_servers.k8s_service.vk8s_networks](resources--origin_pool--properties--origin_servers--k8s_service--vk8s_networks.md#section) |
| `origin_servers.labels` | [origin_servers.labels](resources--origin_pool--properties--origin_servers.md#schema-origin_servers--labels) |
| `origin_servers.private_ip` | [origin_servers.private_ip](resources--origin_pool--properties--origin_servers--private_ip.md#section) |
| `origin_servers.private_ip.inside_network` | [origin_servers.private_ip.inside_network](resources--origin_pool--properties--origin_servers--private_ip--inside_network.md#section) |
| `origin_servers.private_ip.ip` | [origin_servers.private_ip.ip](resources--origin_pool--properties--origin_servers--private_ip.md#schema-origin_servers--private_ip--ip) |
| `origin_servers.private_ip.outside_network` | [origin_servers.private_ip.outside_network](resources--origin_pool--properties--origin_servers--private_ip--outside_network.md#section) |
| `origin_servers.private_ip.segment` | [origin_servers.private_ip.segment](resources--origin_pool--properties--origin_servers--private_ip--segment.md#section) |
| `origin_servers.private_ip.segment.name` | [origin_servers.private_ip.segment.name](resources--origin_pool--properties--origin_servers--private_ip--segment.md#schema-origin_servers--private_ip--segment--name) |
| `origin_servers.private_ip.segment.namespace` | [origin_servers.private_ip.segment.namespace](resources--origin_pool--properties--origin_servers--private_ip--segment.md#schema-origin_servers--private_ip--segment--namespace) |
| `origin_servers.private_ip.segment.tenant` | [origin_servers.private_ip.segment.tenant](resources--origin_pool--properties--origin_servers--private_ip--segment.md#schema-origin_servers--private_ip--segment--tenant) |
| `origin_servers.private_ip.site_locator` | [origin_servers.private_ip.site_locator](resources--origin_pool--properties--origin_servers--private_ip--site_locator.md#section) |
| `origin_servers.private_ip.site_locator.site` | [origin_servers.private_ip.site_locator.site](resources--origin_pool--properties--origin_servers--private_ip--site_locator--site.md#section) |
| `origin_servers.private_ip.site_locator.site.name` | [origin_servers.private_ip.site_locator.site.name](resources--origin_pool--properties--origin_servers--private_ip--site_locator--site.md#schema-origin_servers--private_ip--site_locator--site--name) |
| `origin_servers.private_ip.site_locator.site.namespace` | [origin_servers.private_ip.site_locator.site.namespace](resources--origin_pool--properties--origin_servers--private_ip--site_locator--site.md#schema-origin_servers--private_ip--site_locator--site--namespace) |
| `origin_servers.private_ip.site_locator.site.tenant` | [origin_servers.private_ip.site_locator.site.tenant](resources--origin_pool--properties--origin_servers--private_ip--site_locator--site.md#schema-origin_servers--private_ip--site_locator--site--tenant) |
| `origin_servers.private_ip.site_locator.virtual_site` | [origin_servers.private_ip.site_locator.virtual_site](resources--origin_pool--properties--origin_servers--private_ip--site_locator--virtual_site.md#section) |
| `origin_servers.private_ip.site_locator.virtual_site.name` | [origin_servers.private_ip.site_locator.virtual_site.name](resources--origin_pool--properties--origin_servers--private_ip--site_locator--virtual_site.md#schema-origin_servers--private_ip--site_locator--virtual_site--name) |
| `origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_servers.private_ip.site_locator.virtual_site.namespace](resources--origin_pool--properties--origin_servers--private_ip--site_locator--virtual_site.md#schema-origin_servers--private_ip--site_locator--virtual_site--namespace) |
| `origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_servers.private_ip.site_locator.virtual_site.tenant](resources--origin_pool--properties--origin_servers--private_ip--site_locator--virtual_site.md#schema-origin_servers--private_ip--site_locator--virtual_site--tenant) |
| `origin_servers.private_ip.snat_pool` | [origin_servers.private_ip.snat_pool](resources--origin_pool--properties--origin_servers--private_ip--snat_pool.md#section) |
| `origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_servers.private_ip.snat_pool.no_snat_pool](resources--origin_pool--properties--origin_servers--private_ip--snat_pool--no_snat_pool.md#section) |
| `origin_servers.private_ip.snat_pool.snat_pool` | [origin_servers.private_ip.snat_pool.snat_pool](resources--origin_pool--properties--origin_servers--private_ip--snat_pool--snat_pool.md#section) |
| `origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_servers.private_ip.snat_pool.snat_pool.prefixes](resources--origin_pool--properties--origin_servers--private_ip--snat_pool--snat_pool.md#schema-origin_servers--private_ip--snat_pool--snat_pool--prefixes) |
| `origin_servers.private_name` | [origin_servers.private_name](resources--origin_pool--properties--origin_servers--private_name.md#section) |
| `origin_servers.private_name.dns_name` | [origin_servers.private_name.dns_name](resources--origin_pool--properties--origin_servers--private_name.md#schema-origin_servers--private_name--dns_name) |
| `origin_servers.private_name.inside_network` | [origin_servers.private_name.inside_network](resources--origin_pool--properties--origin_servers--private_name--inside_network.md#section) |
| `origin_servers.private_name.outside_network` | [origin_servers.private_name.outside_network](resources--origin_pool--properties--origin_servers--private_name--outside_network.md#section) |
| `origin_servers.private_name.refresh_interval` | [origin_servers.private_name.refresh_interval](resources--origin_pool--properties--origin_servers--private_name.md#schema-origin_servers--private_name--refresh_interval) |
| `origin_servers.private_name.segment` | [origin_servers.private_name.segment](resources--origin_pool--properties--origin_servers--private_name--segment.md#section) |
| `origin_servers.private_name.segment.name` | [origin_servers.private_name.segment.name](resources--origin_pool--properties--origin_servers--private_name--segment.md#schema-origin_servers--private_name--segment--name) |
| `origin_servers.private_name.segment.namespace` | [origin_servers.private_name.segment.namespace](resources--origin_pool--properties--origin_servers--private_name--segment.md#schema-origin_servers--private_name--segment--namespace) |
| `origin_servers.private_name.segment.tenant` | [origin_servers.private_name.segment.tenant](resources--origin_pool--properties--origin_servers--private_name--segment.md#schema-origin_servers--private_name--segment--tenant) |
| `origin_servers.private_name.site_locator` | [origin_servers.private_name.site_locator](resources--origin_pool--properties--origin_servers--private_name--site_locator.md#section) |
| `origin_servers.private_name.site_locator.site` | [origin_servers.private_name.site_locator.site](resources--origin_pool--properties--origin_servers--private_name--site_locator--site.md#section) |
| `origin_servers.private_name.site_locator.site.name` | [origin_servers.private_name.site_locator.site.name](resources--origin_pool--properties--origin_servers--private_name--site_locator--site.md#schema-origin_servers--private_name--site_locator--site--name) |
| `origin_servers.private_name.site_locator.site.namespace` | [origin_servers.private_name.site_locator.site.namespace](resources--origin_pool--properties--origin_servers--private_name--site_locator--site.md#schema-origin_servers--private_name--site_locator--site--namespace) |
| `origin_servers.private_name.site_locator.site.tenant` | [origin_servers.private_name.site_locator.site.tenant](resources--origin_pool--properties--origin_servers--private_name--site_locator--site.md#schema-origin_servers--private_name--site_locator--site--tenant) |
| `origin_servers.private_name.site_locator.virtual_site` | [origin_servers.private_name.site_locator.virtual_site](resources--origin_pool--properties--origin_servers--private_name--site_locator--virtual_site.md#section) |
| `origin_servers.private_name.site_locator.virtual_site.name` | [origin_servers.private_name.site_locator.virtual_site.name](resources--origin_pool--properties--origin_servers--private_name--site_locator--virtual_site.md#schema-origin_servers--private_name--site_locator--virtual_site--name) |
| `origin_servers.private_name.site_locator.virtual_site.namespace` | [origin_servers.private_name.site_locator.virtual_site.namespace](resources--origin_pool--properties--origin_servers--private_name--site_locator--virtual_site.md#schema-origin_servers--private_name--site_locator--virtual_site--namespace) |
| `origin_servers.private_name.site_locator.virtual_site.tenant` | [origin_servers.private_name.site_locator.virtual_site.tenant](resources--origin_pool--properties--origin_servers--private_name--site_locator--virtual_site.md#schema-origin_servers--private_name--site_locator--virtual_site--tenant) |
| `origin_servers.private_name.snat_pool` | [origin_servers.private_name.snat_pool](resources--origin_pool--properties--origin_servers--private_name--snat_pool.md#section) |
| `origin_servers.private_name.snat_pool.no_snat_pool` | [origin_servers.private_name.snat_pool.no_snat_pool](resources--origin_pool--properties--origin_servers--private_name--snat_pool--no_snat_pool.md#section) |
| `origin_servers.private_name.snat_pool.snat_pool` | [origin_servers.private_name.snat_pool.snat_pool](resources--origin_pool--properties--origin_servers--private_name--snat_pool--snat_pool.md#section) |
| `origin_servers.private_name.snat_pool.snat_pool.prefixes` | [origin_servers.private_name.snat_pool.snat_pool.prefixes](resources--origin_pool--properties--origin_servers--private_name--snat_pool--snat_pool.md#schema-origin_servers--private_name--snat_pool--snat_pool--prefixes) |
| `origin_servers.public_ip` | [origin_servers.public_ip](resources--origin_pool--properties--origin_servers--public_ip.md#section) |
| `origin_servers.public_ip.ip` | [origin_servers.public_ip.ip](resources--origin_pool--properties--origin_servers--public_ip.md#schema-origin_servers--public_ip--ip) |
| `origin_servers.public_name` | [origin_servers.public_name](resources--origin_pool--properties--origin_servers--public_name.md#section) |
| `origin_servers.public_name.dns_name` | [origin_servers.public_name.dns_name](resources--origin_pool--properties--origin_servers--public_name.md#schema-origin_servers--public_name--dns_name) |
| `origin_servers.public_name.refresh_interval` | [origin_servers.public_name.refresh_interval](resources--origin_pool--properties--origin_servers--public_name.md#schema-origin_servers--public_name--refresh_interval) |
| `origin_servers.vn_private_ip` | [origin_servers.vn_private_ip](resources--origin_pool--properties--origin_servers--vn_private_ip.md#section) |
| `origin_servers.vn_private_ip.ip` | [origin_servers.vn_private_ip.ip](resources--origin_pool--properties--origin_servers--vn_private_ip.md#schema-origin_servers--vn_private_ip--ip) |
| `origin_servers.vn_private_ip.virtual_network` | [origin_servers.vn_private_ip.virtual_network](resources--origin_pool--properties--origin_servers--vn_private_ip--virtual_network.md#section) |
| `origin_servers.vn_private_ip.virtual_network.name` | [origin_servers.vn_private_ip.virtual_network.name](resources--origin_pool--properties--origin_servers--vn_private_ip--virtual_network.md#schema-origin_servers--vn_private_ip--virtual_network--name) |
| `origin_servers.vn_private_ip.virtual_network.namespace` | [origin_servers.vn_private_ip.virtual_network.namespace](resources--origin_pool--properties--origin_servers--vn_private_ip--virtual_network.md#schema-origin_servers--vn_private_ip--virtual_network--namespace) |
| `origin_servers.vn_private_ip.virtual_network.tenant` | [origin_servers.vn_private_ip.virtual_network.tenant](resources--origin_pool--properties--origin_servers--vn_private_ip--virtual_network.md#schema-origin_servers--vn_private_ip--virtual_network--tenant) |
| `origin_servers.vn_private_name` | [origin_servers.vn_private_name](resources--origin_pool--properties--origin_servers--vn_private_name.md#section) |
| `origin_servers.vn_private_name.dns_name` | [origin_servers.vn_private_name.dns_name](resources--origin_pool--properties--origin_servers--vn_private_name.md#schema-origin_servers--vn_private_name--dns_name) |
| `origin_servers.vn_private_name.private_network` | [origin_servers.vn_private_name.private_network](resources--origin_pool--properties--origin_servers--vn_private_name--private_network.md#section) |
| `origin_servers.vn_private_name.private_network.name` | [origin_servers.vn_private_name.private_network.name](resources--origin_pool--properties--origin_servers--vn_private_name--private_network.md#schema-origin_servers--vn_private_name--private_network--name) |
| `origin_servers.vn_private_name.private_network.namespace` | [origin_servers.vn_private_name.private_network.namespace](resources--origin_pool--properties--origin_servers--vn_private_name--private_network.md#schema-origin_servers--vn_private_name--private_network--namespace) |
| `origin_servers.vn_private_name.private_network.tenant` | [origin_servers.vn_private_name.private_network.tenant](resources--origin_pool--properties--origin_servers--vn_private_name--private_network.md#schema-origin_servers--vn_private_name--private_network--tenant) |
| `port` | [port](resources--origin_pool--reference.md#schema-port) |
| `same_as_endpoint_port` | [same_as_endpoint_port](resources--origin_pool--properties--same_as_endpoint_port.md#section) |
| `timeouts` | [timeouts](resources--origin_pool--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--origin_pool--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--origin_pool--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--origin_pool--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--origin_pool--properties--timeouts.md#schema-timeouts--update) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](resources--origin_pool--properties--upstream_conn_pool_reuse_type.md#section) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--origin_pool--properties--upstream_conn_pool_reuse_type--disable_conn_pool_reuse.md#section) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--origin_pool--properties--upstream_conn_pool_reuse_type--enable_conn_pool_reuse.md#section) |
| `use_tls` | [use_tls](resources--origin_pool--properties--use_tls.md#section) |
| `use_tls.default_session_key_caching` | [use_tls.default_session_key_caching](resources--origin_pool--properties--use_tls--default_session_key_caching.md#section) |
| `use_tls.disable_session_key_caching` | [use_tls.disable_session_key_caching](resources--origin_pool--properties--use_tls--disable_session_key_caching.md#section) |
| `use_tls.disable_sni` | [use_tls.disable_sni](resources--origin_pool--properties--use_tls--disable_sni.md#section) |
| `use_tls.max_session_keys` | [use_tls.max_session_keys](resources--origin_pool--properties--use_tls.md#schema-use_tls--max_session_keys) |
| `use_tls.no_mtls` | [use_tls.no_mtls](resources--origin_pool--properties--use_tls--no_mtls.md#section) |
| `use_tls.skip_server_verification` | [use_tls.skip_server_verification](resources--origin_pool--properties--use_tls--skip_server_verification.md#section) |
| `use_tls.sni` | [use_tls.sni](resources--origin_pool--properties--use_tls.md#schema-use_tls--sni) |
| `use_tls.tls_config` | [use_tls.tls_config](resources--origin_pool--properties--use_tls--tls_config.md#section) |
| `use_tls.tls_config.custom_security` | [use_tls.tls_config.custom_security](resources--origin_pool--properties--use_tls--tls_config--custom_security.md#section) |
| `use_tls.tls_config.custom_security.cipher_suites` | [use_tls.tls_config.custom_security.cipher_suites](resources--origin_pool--properties--use_tls--tls_config--custom_security.md#schema-use_tls--tls_config--custom_security--cipher_suites) |
| `use_tls.tls_config.custom_security.max_version` | [use_tls.tls_config.custom_security.max_version](resources--origin_pool--properties--use_tls--tls_config--custom_security.md#schema-use_tls--tls_config--custom_security--max_version) |
| `use_tls.tls_config.custom_security.min_version` | [use_tls.tls_config.custom_security.min_version](resources--origin_pool--properties--use_tls--tls_config--custom_security.md#schema-use_tls--tls_config--custom_security--min_version) |
| `use_tls.tls_config.default_security` | [use_tls.tls_config.default_security](resources--origin_pool--properties--use_tls--tls_config--default_security.md#section) |
| `use_tls.tls_config.low_security` | [use_tls.tls_config.low_security](resources--origin_pool--properties--use_tls--tls_config--low_security.md#section) |
| `use_tls.tls_config.medium_security` | [use_tls.tls_config.medium_security](resources--origin_pool--properties--use_tls--tls_config--medium_security.md#section) |
| `use_tls.use_host_header_as_sni` | [use_tls.use_host_header_as_sni](resources--origin_pool--properties--use_tls--use_host_header_as_sni.md#section) |
| `use_tls.use_mtls` | [use_tls.use_mtls](resources--origin_pool--properties--use_tls--use_mtls.md#section) |
| `use_tls.use_mtls.tls_certificates` | [use_tls.use_mtls.tls_certificates](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md#section) |
| `use_tls.use_mtls.tls_certificates.certificate_url` | [use_tls.use_mtls.tls_certificates.certificate_url](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md#schema-use_tls--use_mtls--tls_certificates--certificate_url) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--custom_hash_algorithms.md#section) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--custom_hash_algorithms.md#schema-use_tls--use_mtls--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `use_tls.use_mtls.tls_certificates.description_spec` | [use_tls.use_mtls.tls_certificates.description_spec](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md#schema-use_tls--use_mtls--tls_certificates--description_spec) |
| `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` | [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--disable_ocsp_stapling.md#section) |
| `use_tls.use_mtls.tls_certificates.private_key` | [use_tls.use_mtls.tls_certificates.private_key](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key.md#section) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md#schema-use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md#schema-use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--location) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info.md#schema-use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md#section) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md#schema-use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--private_key--clear_secret_info.md#schema-use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--url) |
| `use_tls.use_mtls.tls_certificates.use_system_defaults` | [use_tls.use_mtls.tls_certificates.use_system_defaults](resources--origin_pool--properties--use_tls--use_mtls--tls_certificates--use_system_defaults.md#section) |
| `use_tls.use_mtls_obj` | [use_tls.use_mtls_obj](resources--origin_pool--properties--use_tls--use_mtls_obj.md#section) |
| `use_tls.use_mtls_obj.name` | [use_tls.use_mtls_obj.name](resources--origin_pool--properties--use_tls--use_mtls_obj.md#schema-use_tls--use_mtls_obj--name) |
| `use_tls.use_mtls_obj.namespace` | [use_tls.use_mtls_obj.namespace](resources--origin_pool--properties--use_tls--use_mtls_obj.md#schema-use_tls--use_mtls_obj--namespace) |
| `use_tls.use_mtls_obj.tenant` | [use_tls.use_mtls_obj.tenant](resources--origin_pool--properties--use_tls--use_mtls_obj.md#schema-use_tls--use_mtls_obj--tenant) |
| `use_tls.use_server_verification` | [use_tls.use_server_verification](resources--origin_pool--properties--use_tls--use_server_verification.md#section) |
| `use_tls.use_server_verification.trusted_ca` | [use_tls.use_server_verification.trusted_ca](resources--origin_pool--properties--use_tls--use_server_verification--trusted_ca.md#section) |
| `use_tls.use_server_verification.trusted_ca.name` | [use_tls.use_server_verification.trusted_ca.name](resources--origin_pool--properties--use_tls--use_server_verification--trusted_ca.md#schema-use_tls--use_server_verification--trusted_ca--name) |
| `use_tls.use_server_verification.trusted_ca.namespace` | [use_tls.use_server_verification.trusted_ca.namespace](resources--origin_pool--properties--use_tls--use_server_verification--trusted_ca.md#schema-use_tls--use_server_verification--trusted_ca--namespace) |
| `use_tls.use_server_verification.trusted_ca.tenant` | [use_tls.use_server_verification.trusted_ca.tenant](resources--origin_pool--properties--use_tls--use_server_verification--trusted_ca.md#schema-use_tls--use_server_verification--trusted_ca--tenant) |
| `use_tls.use_server_verification.trusted_ca_url` | [use_tls.use_server_verification.trusted_ca_url](resources--origin_pool--properties--use_tls--use_server_verification.md#schema-use_tls--use_server_verification--trusted_ca_url) |
| `use_tls.volterra_trusted_ca` | [use_tls.volterra_trusted_ca](resources--origin_pool--properties--use_tls--volterra_trusted_ca.md#section) |

## Next pages

- [advanced_options](resources--origin_pool--properties--advanced_options.md)
- [automatic_port](resources--origin_pool--properties--automatic_port.md)
- [healthcheck](resources--origin_pool--properties--healthcheck.md)
- [lb_port](resources--origin_pool--properties--lb_port.md)
- [no_tls](resources--origin_pool--properties--no_tls.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- [same_as_endpoint_port](resources--origin_pool--properties--same_as_endpoint_port.md)
- [timeouts](resources--origin_pool--properties--timeouts.md)
- [upstream_conn_pool_reuse_type](resources--origin_pool--properties--upstream_conn_pool_reuse_type.md)
- [use_tls](resources--origin_pool--properties--use_tls.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
