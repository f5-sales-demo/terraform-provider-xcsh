---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_cluster."
xcsh_docs: {"aliases": ["cluster"], "body_bytes": 54212, "body_sha256": "sha256:6995f21fe963a08366ce656962cb3c5ecffe05c8885e91bde0b94d593dd01a0c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:cluster:properties:auto_http_config", "xcsh-docs:resources:cluster:properties:circuit_breaker", "xcsh-docs:resources:cluster:properties:default_subset", "xcsh-docs:resources:cluster:properties:disable_proxy_protocol", "xcsh-docs:resources:cluster:properties:endpoint_subsets", "xcsh-docs:resources:cluster:properties:endpoints", "xcsh-docs:resources:cluster:properties:health_checks", "xcsh-docs:resources:cluster:properties:http1_config", "xcsh-docs:resources:cluster:properties:http2_options", "xcsh-docs:resources:cluster:properties:no_panic_threshold", "xcsh-docs:resources:cluster:properties:no_request_limit_per_connection", "xcsh-docs:resources:cluster:properties:outlier_detection", "xcsh-docs:resources:cluster:properties:proxy_protocol_v1", "xcsh-docs:resources:cluster:properties:proxy_protocol_v2", "xcsh-docs:resources:cluster:properties:timeouts", "xcsh-docs:resources:cluster:properties:tls_parameters", "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:reference", "parent_id": "xcsh-docs:resources:cluster:fundamentals", "path": "documentation/resources/cluster/properties/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321", "registry_path": "docs/guides/resources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["auto http config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:auto_http_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["auto_http_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["circuit breaker"], "anchor": "section", "description": "CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if the failures reach a certain threshold, automatically fail subsequent requests which allows to apply back pressure on downstream quickly.", "document_id": "xcsh-docs:resources:cluster:properties:circuit_breaker", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["circuit_breaker"], "syntax": "block", "type": "object"}, {"aliases": ["connection timeout", "duration"], "anchor": "schema-connection_timeout", "description": "The timeout for new network connections to endpoints in the cluster. This is specified in milliseconds. The default value is 2 seconds.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connection_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["default subset"], "anchor": "section", "description": "List of key-value pairs that define default subset. This subset can be referred in fallback_policy which gets used when route specifies no metadata or no subset matching the metadata exists.", "document_id": "xcsh-docs:resources:cluster:properties:default_subset", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["default_subset"], "syntax": "block", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable proxy protocol"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:disable_proxy_protocol", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_proxy_protocol"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint selection"], "anchor": "schema-endpoint_selection", "description": "Policy for selection of endpoints from local site/remote site/both Consider both remote and local endpoints for load balancing LOCAL_ONLY: Consider only local endpoints for load balancing Enable this policy to load balance ONLY among locally discovered endpoints Prefer the local endpoints for load balancing. If local", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_selection"], "syntax": "attribute", "type": "string"}, {"aliases": ["endpoint subsets"], "anchor": "section", "description": "Cluster may be configured to divide its endpoints into subsets based on metadata attached to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be selected by the load balancer. Endpoint_subsets is list of subsets for this cluster. Each entry in this list has definition for a", "document_id": "xcsh-docs:resources:cluster:properties:endpoint_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-endpoint_subsets--keys", "enforcement": "provider-schema", "group": "endpoint_subsets:RequiredListObjectAttributes:keys", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:endpoint_subsets", "type": "requires"}], "schema_path": ["endpoint_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["endpoints"], "anchor": "section", "description": "List of references to all endpoint objects that belong to this cluster.", "document_id": "xcsh-docs:resources:cluster:properties:endpoints", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["endpoints"], "syntax": "block", "type": "object"}, {"aliases": ["fallback policy"], "anchor": "schema-fallback_policy", "description": "Enumeration for SubsetFallbackPolicy if subset match is not found. The request fails as if the cluster had no endpoint matching the subset policy Any cluster endpoint may be selected if the cluster had no endpoint matching the subset policy Load balancing is done over endpoints matching default_subset if the cluster", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["fallback_policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["health checks"], "anchor": "section", "description": "List of references to healthcheck object for this cluster.", "document_id": "xcsh-docs:resources:cluster:properties:health_checks", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["health_checks"], "syntax": "block", "type": "object"}, {"aliases": ["http1 config"], "anchor": "section", "description": "HTTP/1.1 Protocol OPTIONS for upstream connections.", "document_id": "xcsh-docs:resources:cluster:properties:http1_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http1_config"], "syntax": "block", "type": "object"}, {"aliases": ["http2 options"], "anchor": "section", "description": "Http2 Protocol OPTIONS for upstream connections.", "document_id": "xcsh-docs:resources:cluster:properties:http2_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["http2_options"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "http idle timeout"], "anchor": "schema-http_idle_timeout", "description": "The idle timeout for upstream connection pool connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is specified in", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_idle_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["loadbalancer algorithm"], "anchor": "schema-loadbalancer_algorithm", "description": "Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster is required, the load balancer uses loadbalancer_algorithm to determine which host is selected. - ROUND_ROBIN: ROUND_ROBIN Policy in which each healthy/available upstream endpoint is selected in round robin order. -", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["loadbalancer_algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["max requests per connection"], "anchor": "schema-max_requests_per_connection", "description": "Exclusive with Sets the maximum number of requests allowed per connection to the origin server. Enter a value >=1 to define the request limit per connection.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["max_requests_per_connection"], "syntax": "attribute", "type": "number"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["no panic threshold"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:no_panic_threshold", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_panic_threshold"], "syntax": "attribute", "type": "object"}, {"aliases": ["no request limit per connection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:no_request_limit_per_connection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_request_limit_per_connection"], "syntax": "attribute", "type": "object"}, {"aliases": ["outlier detection"], "anchor": "section", "description": "Outlier detection and ejection is the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Outlier detection is a form of passive health checking. Algorithm 1. A endpoint is determined to be an", "document_id": "xcsh-docs:resources:cluster:properties:outlier_detection", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["outlier_detection"], "syntax": "block", "type": "object"}, {"aliases": ["panic threshold"], "anchor": "schema-panic_threshold", "description": "Exclusive with Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be considered for loadbalancing ignoring its health status.", "document_id": "xcsh-docs:resources:cluster:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["panic_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["proxy protocol v1"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:proxy_protocol_v1", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_protocol_v1"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy protocol v2"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cluster:properties:proxy_protocol_v2", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_protocol_v2"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:cluster:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["tls parameters"], "anchor": "section", "description": "TLS configuration for upstream connections.", "document_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tls_parameters--max_session_keys", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--max_session_keys", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--sni", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "schema-tls_parameters--sni", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:cert_params,common_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:cert_params,common_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:common_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:tls_parameters:use_host_header_as_sni", "type": "conflicts"}], "schema_path": ["tls_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["upstream conn pool reuse type"], "anchor": "section", "description": "Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only.", "document_id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "upstream_conn_pool_reuse_type:ConflictingObjectAttributes:disable_conn_pool_reuse,enable_conn_pool_reuse", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "upstream_conn_pool_reuse_type:ConflictingObjectAttributes:disable_conn_pool_reuse,enable_conn_pool_reuse", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cluster:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse", "type": "conflicts"}], "schema_path": ["upstream_conn_pool_reuse_type"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Property reference for xcsh_cluster.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- Property reference

## Direct properties

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
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

- [auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/auto_http_config/): complete subsection reference.

- [circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/circuit_breaker/): complete subsection reference.

<a id="schema-connection_timeout"></a>

### connection_timeout property

Type: `"number"`. Optional, Computed.

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The seconds. Defaults to \`2\`.

Upstream description:

The timeout for new network connections to endpoints in the cluster. This is specified in
milliseconds. The default value is 2 seconds.

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/default_subset/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/disable_proxy_protocol/): complete subsection reference.

<a id="schema-endpoint_selection"></a>

### endpoint_selection property

Type: `"string"`. Optional, Computed.

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

- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoint_subsets/): complete subsection reference.

- [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoints/): complete subsection reference.

<a id="schema-fallback_policy"></a>

### fallback_policy property

Type: `"string"`. Optional, Computed.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NO_FALLBACK",
    "ANY_ENDPOINT",
    "DEFAULT_SUBSET"),
}
```

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

- [health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/health_checks/): complete subsection reference.

- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/): complete subsection reference.

- [http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http2_options/): complete subsection reference.

<a id="schema-http_idle_timeout"></a>

### http_idle_timeout property

Type: `"number"`. Optional, Computed.

The idle timeout for upstream connection pool connections. The idle timeout is defined as the period
in which there are no active requests. When the idle timeout is reached the connection will be
closed.

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

<a id="schema-loadbalancer_algorithm"></a>

### loadbalancer_algorithm property

Type: `"string"`. Optional, Computed.

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

<a id="schema-max_requests_per_connection"></a>

### max_requests_per_connection property

Type: `"number"`. Optional, Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests allowed per connection to the origin server. Enter a value &gt;=1 to
define the request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\]

Sets the maximum number of requests allowed per connection to the origin server. Enter a value
&gt;=1 to define the request limit per connection.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [max_requests_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-max_requests_per_connection)
- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/no_request_limit_per_connection/#section)

Select alternatives according to the provider validators above.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Cluster. Must be unique within the namespace.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Namespace where the Cluster is created.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/no_panic_threshold/): complete subsection reference.

- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/no_request_limit_per_connection/): complete subsection reference.

- [outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/outlier_detection/): complete subsection reference.

<a id="schema-panic_threshold"></a>

### panic_threshold property

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_panic\_threshold\] Configure a threshold (percentage of unhealthy endpoints)
below which all endpoints will be considered for loadbalancing ignoring its health status.

Upstream description:

Exclusive with \[no\_panic\_threshold\]

Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be
considered for loadbalancing ignoring its health status.

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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v1/): complete subsection reference.

- [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v2/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/timeouts/): complete subsection reference.

- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/): complete subsection reference.

- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/upstream_conn_pool_reuse_type/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-annotations) |
| `auto_http_config` | [auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/auto_http_config/#section) |
| `circuit_breaker` | [circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/circuit_breaker/#section) |
| `circuit_breaker.connection_limit` | [circuit_breaker.connection_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/circuit_breaker/#schema-circuit_breaker--connection_limit) |
| `circuit_breaker.max_requests` | [circuit_breaker.max_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/circuit_breaker/#schema-circuit_breaker--max_requests) |
| `circuit_breaker.pending_requests` | [circuit_breaker.pending_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/circuit_breaker/#schema-circuit_breaker--pending_requests) |
| `circuit_breaker.priority` | [circuit_breaker.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/circuit_breaker/#schema-circuit_breaker--priority) |
| `circuit_breaker.retries` | [circuit_breaker.retries](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/circuit_breaker/#schema-circuit_breaker--retries) |
| `connection_timeout` | [connection_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-connection_timeout) |
| `default_subset` | [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/default_subset/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-disable) |
| `disable_proxy_protocol` | [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/disable_proxy_protocol/#section) |
| `endpoint_selection` | [endpoint_selection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-endpoint_selection) |
| `endpoint_subsets` | [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoint_subsets/#section) |
| `endpoint_subsets.keys` | [endpoint_subsets.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoint_subsets/#schema-endpoint_subsets--keys) |
| `endpoints` | [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoints/#section) |
| `endpoints.kind` | [endpoints.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoints/#schema-endpoints--kind) |
| `endpoints.name` | [endpoints.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoints/#schema-endpoints--name) |
| `endpoints.namespace` | [endpoints.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoints/#schema-endpoints--namespace) |
| `endpoints.tenant` | [endpoints.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoints/#schema-endpoints--tenant) |
| `endpoints.uid` | [endpoints.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoints/#schema-endpoints--uid) |
| `fallback_policy` | [fallback_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-fallback_policy) |
| `health_checks` | [health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/health_checks/#section) |
| `health_checks.kind` | [health_checks.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/health_checks/#schema-health_checks--kind) |
| `health_checks.name` | [health_checks.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/health_checks/#schema-health_checks--name) |
| `health_checks.namespace` | [health_checks.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/health_checks/#schema-health_checks--namespace) |
| `health_checks.tenant` | [health_checks.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/health_checks/#schema-health_checks--tenant) |
| `health_checks.uid` | [health_checks.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/health_checks/#schema-health_checks--uid) |
| `http1_config` | [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/#section) |
| `http1_config.header_transformation` | [http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/#section) |
| `http1_config.header_transformation.default_header_transformation` | [http1_config.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/default_header_transformation/#section) |
| `http1_config.header_transformation.preserve_case_header_transformation` | [http1_config.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/preserve_case_header_transformation/#section) |
| `http1_config.header_transformation.proper_case_header_transformation` | [http1_config.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/header_transformation/proper_case_header_transformation/#section) |
| `http2_options` | [http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http2_options/#section) |
| `http2_options.enabled` | [http2_options.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http2_options/#schema-http2_options--enabled) |
| `http_idle_timeout` | [http_idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-http_idle_timeout) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-labels) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-loadbalancer_algorithm) |
| `max_requests_per_connection` | [max_requests_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-max_requests_per_connection) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-namespace) |
| `no_panic_threshold` | [no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/no_panic_threshold/#section) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/no_request_limit_per_connection/#section) |
| `outlier_detection` | [outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/outlier_detection/#section) |
| `outlier_detection.base_ejection_time` | [outlier_detection.base_ejection_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/outlier_detection/#schema-outlier_detection--base_ejection_time) |
| `outlier_detection.consecutive_5xx` | [outlier_detection.consecutive_5xx](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/outlier_detection/#schema-outlier_detection--consecutive_5xx) |
| `outlier_detection.consecutive_gateway_failure` | [outlier_detection.consecutive_gateway_failure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/outlier_detection/#schema-outlier_detection--consecutive_gateway_failure) |
| `outlier_detection.interval` | [outlier_detection.interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/outlier_detection/#schema-outlier_detection--interval) |
| `outlier_detection.max_ejection_percent` | [outlier_detection.max_ejection_percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/outlier_detection/#schema-outlier_detection--max_ejection_percent) |
| `panic_threshold` | [panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/#schema-panic_threshold) |
| `proxy_protocol_v1` | [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v1/#section) |
| `proxy_protocol_v2` | [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v2/#section) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/timeouts/#schema-timeouts--update) |
| `tls_parameters` | [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/#section) |
| `tls_parameters.cert_params` | [tls_parameters.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/#section) |
| `tls_parameters.cert_params.certificates` | [tls_parameters.cert_params.certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/certificates/#section) |
| `tls_parameters.cert_params.certificates.kind` | [tls_parameters.cert_params.certificates.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/certificates/#schema-tls_parameters--cert_params--certificates--kind) |
| `tls_parameters.cert_params.certificates.name` | [tls_parameters.cert_params.certificates.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/certificates/#schema-tls_parameters--cert_params--certificates--name) |
| `tls_parameters.cert_params.certificates.namespace` | [tls_parameters.cert_params.certificates.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/certificates/#schema-tls_parameters--cert_params--certificates--namespace) |
| `tls_parameters.cert_params.certificates.tenant` | [tls_parameters.cert_params.certificates.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/certificates/#schema-tls_parameters--cert_params--certificates--tenant) |
| `tls_parameters.cert_params.certificates.uid` | [tls_parameters.cert_params.certificates.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/certificates/#schema-tls_parameters--cert_params--certificates--uid) |
| `tls_parameters.cert_params.cipher_suites` | [tls_parameters.cert_params.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/#schema-tls_parameters--cert_params--cipher_suites) |
| `tls_parameters.cert_params.maximum_protocol_version` | [tls_parameters.cert_params.maximum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/#schema-tls_parameters--cert_params--maximum_protocol_version) |
| `tls_parameters.cert_params.minimum_protocol_version` | [tls_parameters.cert_params.minimum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/#schema-tls_parameters--cert_params--minimum_protocol_version) |
| `tls_parameters.cert_params.skip_server_verification` | [tls_parameters.cert_params.skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/skip_server_verification/#section) |
| `tls_parameters.cert_params.tls_validation_params` | [tls_parameters.cert_params.tls_validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/#section) |
| `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` | [tls_parameters.cert_params.tls_validation_params.skip_hostname_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/#schema-tls_parameters--cert_params--tls_validation_params--skip_hostname_verification) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca` | [tls_parameters.cert_params.tls_validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/#section) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#section) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` | [tls_parameters.cert_params.tls_validation_params.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/#schema-tls_parameters--cert_params--tls_validation_params--trusted_ca_url) |
| `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` | [tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/tls_validation_params/#schema-tls_parameters--cert_params--tls_validation_params--verify_subject_alt_names) |
| `tls_parameters.cert_params.volterra_trusted_ca` | [tls_parameters.cert_params.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/volterra_trusted_ca/#section) |
| `tls_parameters.common_params` | [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/#section) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--cipher_suites) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--maximum_protocol_version) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--minimum_protocol_version) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/#section) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/#schema-tls_parameters--common_params--tls_certificates--certificate_url) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/#section) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/#schema-tls_parameters--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/#schema-tls_parameters--common_params--tls_certificates--description_spec) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/#section) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--url) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/#section) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/#section) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--skip_hostname_verification) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--trusted_ca_url) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--verify_subject_alt_names) |
| `tls_parameters.default_session_key_caching` | [tls_parameters.default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/default_session_key_caching/#section) |
| `tls_parameters.disable_session_key_caching` | [tls_parameters.disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/disable_session_key_caching/#section) |
| `tls_parameters.disable_sni` | [tls_parameters.disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/disable_sni/#section) |
| `tls_parameters.max_session_keys` | [tls_parameters.max_session_keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/#schema-tls_parameters--max_session_keys) |
| `tls_parameters.sni` | [tls_parameters.sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/#schema-tls_parameters--sni) |
| `tls_parameters.use_host_header_as_sni` | [tls_parameters.use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/use_host_header_as_sni/#section) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/upstream_conn_pool_reuse_type/#section) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/#section) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/#section) |

## Next pages

- [auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/auto_http_config/)
- [circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/circuit_breaker/)
- [default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/default_subset/)
- [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/disable_proxy_protocol/)
- [endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoint_subsets/)
- [endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/endpoints/)
- [health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/health_checks/)
- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http1_config/)
- [http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/http2_options/)
- [no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/no_panic_threshold/)
- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/no_request_limit_per_connection/)
- [outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/outlier_detection/)
- [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v1/)
- [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/proxy_protocol_v2/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/timeouts/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/)
- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/upstream_conn_pool_reuse_type/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
