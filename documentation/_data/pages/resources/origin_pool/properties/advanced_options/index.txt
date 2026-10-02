---
page_title: "advanced_options"
subcategory: "Load Balancing"
description: "Configure Advanced OPTIONS for origin pool."
xcsh_docs: {"aliases": ["advanced options", "backend servers", "origin servers", "upstream servers"], "body_bytes": 14675, "body_sha256": "sha256:fa9a5b2b3b294df7fabff119be1b94716372808b04e6d8ae7fa97a67e91b29e2", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:advanced_options:auto_http_config", "xcsh-docs:resources:origin_pool:properties:advanced_options:circuit_breaker", "xcsh-docs:resources:origin_pool:properties:advanced_options:default_circuit_breaker", "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_circuit_breaker", "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_outlier_detection", "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_proxy_protocol", "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_subsets", "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_lb_source_ip_persistence", "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets", "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "xcsh-docs:resources:origin_pool:properties:advanced_options:http2_options", "xcsh-docs:resources:origin_pool:properties:advanced_options:no_panic_threshold", "xcsh-docs:resources:origin_pool:properties:advanced_options:no_request_limit_per_connection", "xcsh-docs:resources:origin_pool:properties:advanced_options:outlier_detection", "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v1", "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v2"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "parent_id": "xcsh-docs:resources:origin_pool:reference", "path": "documentation/resources/origin_pool/properties/advanced_options/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1023311220110031-0231320022131021-0033212101303010-3313122003030112-0232231312233323-0112101220121031-0120211031032111-3200312011212231", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [{"anchor": "schema-advanced_options--max_requests_per_connection", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "type": "conflicts"}, {"anchor": "schema-advanced_options--panic_threshold", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:no_panic_threshold,panic_threshold", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:auto_http_config,http1_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:auto_http_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:auto_http_config,http2_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:auto_http_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:circuit_breaker,default_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:circuit_breaker,disable_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:circuit_breaker,default_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:default_circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:default_circuit_breaker,disable_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:default_circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:circuit_breaker,disable_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:default_circuit_breaker,disable_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_lb_source_ip_persistence,enable_lb_source_ip_persistence", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_outlier_detection,outlier_detection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_outlier_detection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_proxy_protocol,proxy_protocol_v1", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_proxy_protocol", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_proxy_protocol,proxy_protocol_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_proxy_protocol", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_subsets,enable_subsets", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_subsets", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_lb_source_ip_persistence,enable_lb_source_ip_persistence", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_lb_source_ip_persistence", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_subsets,enable_subsets", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:auto_http_config,http1_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:http1_config,http2_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:auto_http_config,http2_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http2_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:http1_config,http2_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http2_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:no_panic_threshold,panic_threshold", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:no_panic_threshold", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:no_request_limit_per_connection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_outlier_detection,outlier_detection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:outlier_detection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_proxy_protocol,proxy_protocol_v1", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v1", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:proxy_protocol_v1,proxy_protocol_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v1", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_proxy_protocol,proxy_protocol_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v2", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:proxy_protocol_v1,proxy_protocol_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v2", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["advanced_options"], "schema_version": 1, "sections": [{"aliases": ["auto http config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:auto_http_config", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "auto_http_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["circuit breaker"], "anchor": "section", "description": "CircuitBreaker provides a mechanism for watching failures in upstream connections or requests and if the failures reach a certain threshold, automatically fail subsequent requests which allows to apply back pressure on downstream quickly.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:circuit_breaker", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "circuit_breaker"], "syntax": "block", "type": "object"}, {"aliases": ["connection timeout", "duration", "operation timeout"], "anchor": "schema-advanced_options--connection_timeout", "description": "The timeout for new network connections to endpoints in the cluster. This is specified in milliseconds. The default value is 2 seconds.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "connection_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["default circuit breaker"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:default_circuit_breaker", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "default_circuit_breaker"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable circuit breaker"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_circuit_breaker", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "disable_circuit_breaker"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable lb source ip persistence"], "anchor": "section", "description": "IP address configuration", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "disable_lb_source_ip_persistence"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable outlier detection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_outlier_detection", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "disable_outlier_detection"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable proxy protocol"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_proxy_protocol", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "disable_proxy_protocol"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable subsets"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_subsets", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "disable_subsets"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable lb source ip persistence"], "anchor": "section", "description": "IP address configuration", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_lb_source_ip_persistence", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "enable_lb_source_ip_persistence"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "enable subsets", "origin servers", "upstream servers"], "anchor": "section", "description": "Configure subset OPTIONS for origin pool.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,default_subset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:any_endpoint", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,default_subset", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:default_subset,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:default_subset", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:any_endpoint,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:fail_request", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:ConflictingObjectAttributes:default_subset,fail_request", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:fail_request", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options.enable_subsets:RequiredObjectAttributes:endpoint_subsets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets:endpoint_subsets", "type": "requires"}], "schema_path": ["advanced_options", "enable_subsets"], "syntax": "block", "type": "object"}, {"aliases": ["http1 config"], "anchor": "section", "description": "HTTP/1.1 Protocol OPTIONS for upstream connections.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "http1_config"], "syntax": "block", "type": "object"}, {"aliases": ["http2 options"], "anchor": "section", "description": "Http2 Protocol OPTIONS for upstream connections.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http2_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "http2_options"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "http idle timeout", "operation timeout"], "anchor": "schema-advanced_options--http_idle_timeout", "description": "The idle timeout for upstream connection pool connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is specified in", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "http_idle_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["backend servers", "max requests per connection", "origin servers", "upstream servers"], "anchor": "schema-advanced_options--max_requests_per_connection", "description": "Exclusive with Sets the maximum number of requests allowed per connection to the origin server. Enter a value >=1 to define the request limit per connection.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "max_requests_per_connection"], "syntax": "attribute", "type": "number"}, {"aliases": ["no panic threshold"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:no_panic_threshold", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "no_panic_threshold"], "syntax": "attribute", "type": "object"}, {"aliases": ["no request limit per connection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:no_request_limit_per_connection", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "no_request_limit_per_connection"], "syntax": "attribute", "type": "object"}, {"aliases": ["outlier detection"], "anchor": "section", "description": "Outlier detection and ejection is the process of dynamically determining whether some number of hosts in an upstream cluster are performing unlike the others and removing them from the healthy load balancing set. Outlier detection is a form of passive health checking. Algorithm 1. A endpoint is determined to be an", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:outlier_detection", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_options", "outlier_detection"], "syntax": "block", "type": "object"}, {"aliases": ["panic threshold"], "anchor": "schema-advanced_options--panic_threshold", "description": "Exclusive with Configure a threshold (percentage of unhealthy endpoints) below which all endpoints will be considered for load balancing ignoring its health status.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "panic_threshold"], "syntax": "attribute", "type": "number"}, {"aliases": ["proxy protocol v1"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v1", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "proxy_protocol_v1"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy protocol v2"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v2", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["advanced_options", "proxy_protocol_v2"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/advanced_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configure Advanced OPTIONS for origin pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# advanced_options

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- advanced_options

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

- [auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/auto_http_config/): complete subsection reference.

- [circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/circuit_breaker/): complete subsection reference.

<a id="schema-advanced_options--connection_timeout"></a>

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

- [default_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/default_circuit_breaker/): complete subsection reference.

- [disable_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_circuit_breaker/): complete subsection reference.

- [disable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_lb_source_ip_persistence/): complete subsection reference.

- [disable_outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_outlier_detection/): complete subsection reference.

- [disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_proxy_protocol/): complete subsection reference.

- [disable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_subsets/): complete subsection reference.

- [enable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_lb_source_ip_persistence/): complete subsection reference.

- [enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/): complete subsection reference.

- [http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/): complete subsection reference.

- [http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http2_options/): complete subsection reference.

<a id="schema-advanced_options--http_idle_timeout"></a>

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

<a id="schema-advanced_options--max_requests_per_connection"></a>

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

- [no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/no_panic_threshold/): complete subsection reference.

- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/no_request_limit_per_connection/): complete subsection reference.

- [outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/outlier_detection/): complete subsection reference.

<a id="schema-advanced_options--panic_threshold"></a>

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

- [proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/proxy_protocol_v1/): complete subsection reference.

- [proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/proxy_protocol_v2/): complete subsection reference.

## Next pages

- [advanced_options.auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/auto_http_config/)
- [advanced_options.circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/circuit_breaker/)
- [advanced_options.default_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/default_circuit_breaker/)
- [advanced_options.disable_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_circuit_breaker/)
- [advanced_options.disable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_lb_source_ip_persistence/)
- [advanced_options.disable_outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_outlier_detection/)
- [advanced_options.disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_proxy_protocol/)
- [advanced_options.disable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_subsets/)
- [advanced_options.enable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_lb_source_ip_persistence/)
- [advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/)
- [advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/)
- [advanced_options.http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http2_options/)
- [advanced_options.no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/no_panic_threshold/)
- [advanced_options.no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/no_request_limit_per_connection/)
- [advanced_options.outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/outlier_detection/)
- [advanced_options.proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/proxy_protocol_v1/)
- [advanced_options.proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/proxy_protocol_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
