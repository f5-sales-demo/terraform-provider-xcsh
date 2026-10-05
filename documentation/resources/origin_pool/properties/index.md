---
page_title: "Property reference"
subcategory: "Load Balancing"
description: "Property reference for xcsh_origin_pool."
xcsh_docs: {"aliases": ["backend servers", "origin pool", "origin servers", "upstream servers"], "body_bytes": 72542, "body_sha256": "sha256:71f56fa8192bd8fd29551c67862e3e1410fa99a8c040a6b57fc36d38cb786d19", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:advanced_options", "xcsh-docs:resources:origin_pool:properties:automatic_port", "xcsh-docs:resources:origin_pool:properties:healthcheck", "xcsh-docs:resources:origin_pool:properties:lb_port", "xcsh-docs:resources:origin_pool:properties:no_tls", "xcsh-docs:resources:origin_pool:properties:origin_servers", "xcsh-docs:resources:origin_pool:properties:same_as_endpoint_port", "xcsh-docs:resources:origin_pool:properties:timeouts", "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type", "xcsh-docs:resources:origin_pool:properties:use_tls"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:reference", "parent_id": "xcsh-docs:resources:origin_pool:fundamentals", "path": "documentation/resources/origin_pool/properties/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013", "registry_path": "docs/guides/resources--origin_pool--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["advanced options"], "anchor": "section", "description": "Configure Advanced OPTIONS for origin pool.", "document_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-advanced_options--max_requests_per_connection", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "type": "conflicts"}, {"anchor": "schema-advanced_options--panic_threshold", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:no_panic_threshold,panic_threshold", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:auto_http_config,http1_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:auto_http_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:auto_http_config,http2_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:auto_http_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:circuit_breaker,default_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:circuit_breaker,disable_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:circuit_breaker,default_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:default_circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:default_circuit_breaker,disable_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:default_circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:circuit_breaker,disable_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:default_circuit_breaker,disable_circuit_breaker", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_circuit_breaker", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_lb_source_ip_persistence,enable_lb_source_ip_persistence", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_lb_source_ip_persistence", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_outlier_detection,outlier_detection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_outlier_detection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_proxy_protocol,proxy_protocol_v1", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_proxy_protocol", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_proxy_protocol,proxy_protocol_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_proxy_protocol", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_subsets,enable_subsets", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:disable_subsets", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_lb_source_ip_persistence,enable_lb_source_ip_persistence", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_lb_source_ip_persistence", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_subsets,enable_subsets", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:enable_subsets", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:auto_http_config,http1_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:http1_config,http2_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http1_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:auto_http_config,http2_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http2_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:http1_config,http2_options", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:http2_options", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:no_panic_threshold,panic_threshold", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:no_panic_threshold", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:no_request_limit_per_connection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_outlier_detection,outlier_detection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:outlier_detection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_proxy_protocol,proxy_protocol_v1", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v1", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:proxy_protocol_v1,proxy_protocol_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v1", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:disable_proxy_protocol,proxy_protocol_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v2", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "advanced_options:ConflictingObjectAttributes:proxy_protocol_v1,proxy_protocol_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:advanced_options:proxy_protocol_v2", "type": "conflicts"}], "schema_path": ["advanced_options"], "syntax": "block", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["automatic port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:automatic_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["automatic_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["endpoint selection"], "anchor": "schema-endpoint_selection", "description": "Policy for selection of endpoints from local site/remote site/both Consider both remote and local endpoints for load balancing LOCAL_ONLY: Consider only local endpoints for load balancing Enable this policy to load balance ONLY among locally discovered endpoints Prefer the local endpoints for load balancing. If local", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint_selection"], "syntax": "attribute", "type": "string"}, {"aliases": ["health check port"], "anchor": "schema-health_check_port", "description": "Exclusive with Port used for performing health check.", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["health_check_port"], "syntax": "attribute", "type": "number"}, {"aliases": ["healthcheck"], "anchor": "section", "description": "Reference to healthcheck configuration objects.", "document_id": "xcsh-docs:resources:origin_pool:properties:healthcheck", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-healthcheck--name", "enforcement": "provider-schema", "group": "healthcheck:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:healthcheck", "type": "requires"}], "schema_path": ["healthcheck"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["lb port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:lb_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["lb_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["loadbalancer algorithm"], "anchor": "schema-loadbalancer_algorithm", "description": "Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster is required, the load balancer uses loadbalancer_algorithm to determine which host is selected. - ROUND_ROBIN: ROUND_ROBIN Policy in which each healthy/available upstream endpoint is selected in round robin order. -", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["loadbalancer_algorithm"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["no tls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:no_tls", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_tls"], "syntax": "attribute", "type": "object"}, {"aliases": ["backend servers", "origin servers", "upstream servers"], "anchor": "section", "description": "List of origin servers in this pool.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,consul_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,custom_endpoint_object", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,k8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:cbip_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,consul_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,custom_endpoint_object", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,k8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:consul_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,custom_endpoint_object", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,custom_endpoint_object", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,k8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:custom_endpoint_object", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,k8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,k8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,k8s_service", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:k8s_service", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_name,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_name,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_name,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_name,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_name,public_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_ip,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_ip,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_name,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_ip,public_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_name,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_name,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:public_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_name,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_ip,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_name,vn_private_ip", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:vn_private_ip,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:cbip_service,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:consul_service,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:custom_endpoint_object,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:k8s_service,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_ip,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:private_name,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_ip,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:public_name,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers:ConflictingListObjectAttributes:vn_private_ip,vn_private_name", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:vn_private_name", "type": "conflicts"}], "schema_path": ["origin_servers"], "syntax": "block", "type": "object"}, {"aliases": ["port"], "anchor": "schema-port", "description": "Exclusive with Endpoint service is available on this port.", "document_id": "xcsh-docs:resources:origin_pool:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["port"], "syntax": "attribute", "type": "number"}, {"aliases": ["same as endpoint port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:same_as_endpoint_port", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["same_as_endpoint_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "lifecycle timeout", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:origin_pool:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["upstream conn pool reuse type"], "anchor": "section", "description": "Select upstream connection pool reuse state for every downstream connection. This configuration choice is for HTTP(S) LB only.", "document_id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "upstream_conn_pool_reuse_type:ConflictingObjectAttributes:disable_conn_pool_reuse,enable_conn_pool_reuse", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type:disable_conn_pool_reuse", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "upstream_conn_pool_reuse_type:ConflictingObjectAttributes:disable_conn_pool_reuse,enable_conn_pool_reuse", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:upstream_conn_pool_reuse_type:enable_conn_pool_reuse", "type": "conflicts"}], "schema_path": ["upstream_conn_pool_reuse_type"], "syntax": "block", "type": "object"}, {"aliases": ["use tls"], "anchor": "section", "description": "Upstream TLS Parameters.", "document_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-use_tls--max_session_keys", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "type": "conflicts"}, {"anchor": "schema-use_tls--max_session_keys", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "type": "conflicts"}, {"anchor": "schema-use_tls--sni", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "type": "conflicts"}, {"anchor": "schema-use_tls--sni", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:default_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:default_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:default_session_key_caching,disable_session_key_caching", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:disable_session_key_caching,max_session_keys", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:disable_session_key_caching", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:disable_sni,sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:disable_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:no_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:skip_server_verification,use_server_verification", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:skip_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:disable_sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:sni,use_host_header_as_sni", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_host_header_as_sni", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:use_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:no_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:use_mtls,use_mtls_obj", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls_obj", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:skip_server_verification,use_server_verification", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_server_verification", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:skip_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:volterra_trusted_ca", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls:ConflictingObjectAttributes:use_server_verification,volterra_trusted_ca", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:volterra_trusted_ca", "type": "conflicts"}], "schema_path": ["use_tls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Property reference for xcsh_origin_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- Property reference

## Direct properties

- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/): complete subsection reference.

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

- [automatic_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/automatic_port/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-health_check_port)
- [same_as_endpoint_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/same_as_endpoint_port/#section)

Select alternatives according to the provider validators above.

- [healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/healthcheck/): complete subsection reference.

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

- [lb_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/lb_port/): complete subsection reference.

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

- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/no_tls/): complete subsection reference.

- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [same_as_endpoint_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/same_as_endpoint_port/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/timeouts/): complete subsection reference.

- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/): complete subsection reference.

- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_options` | [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/#section) |
| `advanced_options.auto_http_config` | [advanced_options.auto_http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/auto_http_config/#section) |
| `advanced_options.circuit_breaker` | [advanced_options.circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/circuit_breaker/#section) |
| `advanced_options.circuit_breaker.connection_limit` | [advanced_options.circuit_breaker.connection_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/circuit_breaker/#schema-advanced_options--circuit_breaker--connection_limit) |
| `advanced_options.circuit_breaker.max_requests` | [advanced_options.circuit_breaker.max_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/circuit_breaker/#schema-advanced_options--circuit_breaker--max_requests) |
| `advanced_options.circuit_breaker.pending_requests` | [advanced_options.circuit_breaker.pending_requests](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/circuit_breaker/#schema-advanced_options--circuit_breaker--pending_requests) |
| `advanced_options.circuit_breaker.priority` | [advanced_options.circuit_breaker.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/circuit_breaker/#schema-advanced_options--circuit_breaker--priority) |
| `advanced_options.circuit_breaker.retries` | [advanced_options.circuit_breaker.retries](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/circuit_breaker/#schema-advanced_options--circuit_breaker--retries) |
| `advanced_options.connection_timeout` | [advanced_options.connection_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/#schema-advanced_options--connection_timeout) |
| `advanced_options.default_circuit_breaker` | [advanced_options.default_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/default_circuit_breaker/#section) |
| `advanced_options.disable_circuit_breaker` | [advanced_options.disable_circuit_breaker](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_circuit_breaker/#section) |
| `advanced_options.disable_lb_source_ip_persistence` | [advanced_options.disable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_lb_source_ip_persistence/#section) |
| `advanced_options.disable_outlier_detection` | [advanced_options.disable_outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_outlier_detection/#section) |
| `advanced_options.disable_proxy_protocol` | [advanced_options.disable_proxy_protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_proxy_protocol/#section) |
| `advanced_options.disable_subsets` | [advanced_options.disable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/disable_subsets/#section) |
| `advanced_options.enable_lb_source_ip_persistence` | [advanced_options.enable_lb_source_ip_persistence](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_lb_source_ip_persistence/#section) |
| `advanced_options.enable_subsets` | [advanced_options.enable_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/#section) |
| `advanced_options.enable_subsets.any_endpoint` | [advanced_options.enable_subsets.any_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/any_endpoint/#section) |
| `advanced_options.enable_subsets.default_subset` | [advanced_options.enable_subsets.default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/default_subset/#section) |
| `advanced_options.enable_subsets.default_subset.default_subset` | [advanced_options.enable_subsets.default_subset.default_subset](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/default_subset/default_subset/#section) |
| `advanced_options.enable_subsets.endpoint_subsets` | [advanced_options.enable_subsets.endpoint_subsets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/endpoint_subsets/#section) |
| `advanced_options.enable_subsets.endpoint_subsets.keys` | [advanced_options.enable_subsets.endpoint_subsets.keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/endpoint_subsets/#schema-advanced_options--enable_subsets--endpoint_subsets--keys) |
| `advanced_options.enable_subsets.fail_request` | [advanced_options.enable_subsets.fail_request](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/enable_subsets/fail_request/#section) |
| `advanced_options.http1_config` | [advanced_options.http1_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/#section) |
| `advanced_options.http1_config.header_transformation` | [advanced_options.http1_config.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/#section) |
| `advanced_options.http1_config.header_transformation.default_header_transformation` | [advanced_options.http1_config.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/default_header_transformation/#section) |
| `advanced_options.http1_config.header_transformation.preserve_case_header_transformation` | [advanced_options.http1_config.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/preserve_case_header_transformation/#section) |
| `advanced_options.http1_config.header_transformation.proper_case_header_transformation` | [advanced_options.http1_config.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http1_config/header_transformation/proper_case_header_transformation/#section) |
| `advanced_options.http2_options` | [advanced_options.http2_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http2_options/#section) |
| `advanced_options.http2_options.enabled` | [advanced_options.http2_options.enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/http2_options/#schema-advanced_options--http2_options--enabled) |
| `advanced_options.http_idle_timeout` | [advanced_options.http_idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/#schema-advanced_options--http_idle_timeout) |
| `advanced_options.max_requests_per_connection` | [advanced_options.max_requests_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/#schema-advanced_options--max_requests_per_connection) |
| `advanced_options.no_panic_threshold` | [advanced_options.no_panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/no_panic_threshold/#section) |
| `advanced_options.no_request_limit_per_connection` | [advanced_options.no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/no_request_limit_per_connection/#section) |
| `advanced_options.outlier_detection` | [advanced_options.outlier_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/outlier_detection/#section) |
| `advanced_options.outlier_detection.base_ejection_time` | [advanced_options.outlier_detection.base_ejection_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/outlier_detection/#schema-advanced_options--outlier_detection--base_ejection_time) |
| `advanced_options.outlier_detection.consecutive_5xx` | [advanced_options.outlier_detection.consecutive_5xx](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/outlier_detection/#schema-advanced_options--outlier_detection--consecutive_5xx) |
| `advanced_options.outlier_detection.consecutive_gateway_failure` | [advanced_options.outlier_detection.consecutive_gateway_failure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/outlier_detection/#schema-advanced_options--outlier_detection--consecutive_gateway_failure) |
| `advanced_options.outlier_detection.interval` | [advanced_options.outlier_detection.interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/outlier_detection/#schema-advanced_options--outlier_detection--interval) |
| `advanced_options.outlier_detection.max_ejection_percent` | [advanced_options.outlier_detection.max_ejection_percent](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/outlier_detection/#schema-advanced_options--outlier_detection--max_ejection_percent) |
| `advanced_options.panic_threshold` | [advanced_options.panic_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/#schema-advanced_options--panic_threshold) |
| `advanced_options.proxy_protocol_v1` | [advanced_options.proxy_protocol_v1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/proxy_protocol_v1/#section) |
| `advanced_options.proxy_protocol_v2` | [advanced_options.proxy_protocol_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/proxy_protocol_v2/#section) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-annotations) |
| `automatic_port` | [automatic_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/automatic_port/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-disable) |
| `endpoint_selection` | [endpoint_selection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-endpoint_selection) |
| `health_check_port` | [health_check_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-health_check_port) |
| `healthcheck` | [healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/healthcheck/#section) |
| `healthcheck.name` | [healthcheck.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/healthcheck/#schema-healthcheck--name) |
| `healthcheck.namespace` | [healthcheck.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/healthcheck/#schema-healthcheck--namespace) |
| `healthcheck.tenant` | [healthcheck.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/healthcheck/#schema-healthcheck--tenant) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-labels) |
| `lb_port` | [lb_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/lb_port/#section) |
| `loadbalancer_algorithm` | [loadbalancer_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-loadbalancer_algorithm) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-namespace) |
| `no_tls` | [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/no_tls/#section) |
| `origin_servers` | [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/#section) |
| `origin_servers.cbip_service` | [origin_servers.cbip_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/cbip_service/#section) |
| `origin_servers.cbip_service.service_name` | [origin_servers.cbip_service.service_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/cbip_service/#schema-origin_servers--cbip_service--service_name) |
| `origin_servers.consul_service` | [origin_servers.consul_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/#section) |
| `origin_servers.consul_service.inside_network` | [origin_servers.consul_service.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/inside_network/#section) |
| `origin_servers.consul_service.outside_network` | [origin_servers.consul_service.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/outside_network/#section) |
| `origin_servers.consul_service.service_name` | [origin_servers.consul_service.service_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/#schema-origin_servers--consul_service--service_name) |
| `origin_servers.consul_service.site_locator` | [origin_servers.consul_service.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/#section) |
| `origin_servers.consul_service.site_locator.site` | [origin_servers.consul_service.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/site/#section) |
| `origin_servers.consul_service.site_locator.site.name` | [origin_servers.consul_service.site_locator.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/site/#schema-origin_servers--consul_service--site_locator--site--name) |
| `origin_servers.consul_service.site_locator.site.namespace` | [origin_servers.consul_service.site_locator.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/site/#schema-origin_servers--consul_service--site_locator--site--namespace) |
| `origin_servers.consul_service.site_locator.site.tenant` | [origin_servers.consul_service.site_locator.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/site/#schema-origin_servers--consul_service--site_locator--site--tenant) |
| `origin_servers.consul_service.site_locator.virtual_site` | [origin_servers.consul_service.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/virtual_site/#section) |
| `origin_servers.consul_service.site_locator.virtual_site.name` | [origin_servers.consul_service.site_locator.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/virtual_site/#schema-origin_servers--consul_service--site_locator--virtual_site--name) |
| `origin_servers.consul_service.site_locator.virtual_site.namespace` | [origin_servers.consul_service.site_locator.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/virtual_site/#schema-origin_servers--consul_service--site_locator--virtual_site--namespace) |
| `origin_servers.consul_service.site_locator.virtual_site.tenant` | [origin_servers.consul_service.site_locator.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/site_locator/virtual_site/#schema-origin_servers--consul_service--site_locator--virtual_site--tenant) |
| `origin_servers.consul_service.snat_pool` | [origin_servers.consul_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/#section) |
| `origin_servers.consul_service.snat_pool.no_snat_pool` | [origin_servers.consul_service.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/no_snat_pool/#section) |
| `origin_servers.consul_service.snat_pool.snat_pool` | [origin_servers.consul_service.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/snat_pool/#section) |
| `origin_servers.consul_service.snat_pool.snat_pool.prefixes` | [origin_servers.consul_service.snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/consul_service/snat_pool/snat_pool/#schema-origin_servers--consul_service--snat_pool--snat_pool--prefixes) |
| `origin_servers.custom_endpoint_object` | [origin_servers.custom_endpoint_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/custom_endpoint_object/#section) |
| `origin_servers.custom_endpoint_object.endpoint` | [origin_servers.custom_endpoint_object.endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/custom_endpoint_object/endpoint/#section) |
| `origin_servers.custom_endpoint_object.endpoint.name` | [origin_servers.custom_endpoint_object.endpoint.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/custom_endpoint_object/endpoint/#schema-origin_servers--custom_endpoint_object--endpoint--name) |
| `origin_servers.custom_endpoint_object.endpoint.namespace` | [origin_servers.custom_endpoint_object.endpoint.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/custom_endpoint_object/endpoint/#schema-origin_servers--custom_endpoint_object--endpoint--namespace) |
| `origin_servers.custom_endpoint_object.endpoint.tenant` | [origin_servers.custom_endpoint_object.endpoint.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/custom_endpoint_object/endpoint/#schema-origin_servers--custom_endpoint_object--endpoint--tenant) |
| `origin_servers.k8s_service` | [origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/#section) |
| `origin_servers.k8s_service.inside_network` | [origin_servers.k8s_service.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/inside_network/#section) |
| `origin_servers.k8s_service.outside_network` | [origin_servers.k8s_service.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/outside_network/#section) |
| `origin_servers.k8s_service.protocol` | [origin_servers.k8s_service.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/#schema-origin_servers--k8s_service--protocol) |
| `origin_servers.k8s_service.service_name` | [origin_servers.k8s_service.service_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/#schema-origin_servers--k8s_service--service_name) |
| `origin_servers.k8s_service.site_locator` | [origin_servers.k8s_service.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/#section) |
| `origin_servers.k8s_service.site_locator.site` | [origin_servers.k8s_service.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/site/#section) |
| `origin_servers.k8s_service.site_locator.site.name` | [origin_servers.k8s_service.site_locator.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/site/#schema-origin_servers--k8s_service--site_locator--site--name) |
| `origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.k8s_service.site_locator.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/site/#schema-origin_servers--k8s_service--site_locator--site--namespace) |
| `origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.k8s_service.site_locator.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/site/#schema-origin_servers--k8s_service--site_locator--site--tenant) |
| `origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.k8s_service.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/virtual_site/#section) |
| `origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.k8s_service.site_locator.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_servers--k8s_service--site_locator--virtual_site--name) |
| `origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.k8s_service.site_locator.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_servers--k8s_service--site_locator--virtual_site--namespace) |
| `origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.k8s_service.site_locator.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_servers--k8s_service--site_locator--virtual_site--tenant) |
| `origin_servers.k8s_service.snat_pool` | [origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/snat_pool/#section) |
| `origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.k8s_service.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/snat_pool/no_snat_pool/#section) |
| `origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.k8s_service.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/snat_pool/snat_pool/#section) |
| `origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.k8s_service.snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/snat_pool/snat_pool/#schema-origin_servers--k8s_service--snat_pool--snat_pool--prefixes) |
| `origin_servers.k8s_service.vk8s_networks` | [origin_servers.k8s_service.vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/k8s_service/vk8s_networks/#section) |
| `origin_servers.labels` | [origin_servers.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/#schema-origin_servers--labels) |
| `origin_servers.private_ip` | [origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/#section) |
| `origin_servers.private_ip.inside_network` | [origin_servers.private_ip.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/inside_network/#section) |
| `origin_servers.private_ip.ip` | [origin_servers.private_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/#schema-origin_servers--private_ip--ip) |
| `origin_servers.private_ip.outside_network` | [origin_servers.private_ip.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/outside_network/#section) |
| `origin_servers.private_ip.segment` | [origin_servers.private_ip.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/segment/#section) |
| `origin_servers.private_ip.segment.name` | [origin_servers.private_ip.segment.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/segment/#schema-origin_servers--private_ip--segment--name) |
| `origin_servers.private_ip.segment.namespace` | [origin_servers.private_ip.segment.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/segment/#schema-origin_servers--private_ip--segment--namespace) |
| `origin_servers.private_ip.segment.tenant` | [origin_servers.private_ip.segment.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/segment/#schema-origin_servers--private_ip--segment--tenant) |
| `origin_servers.private_ip.site_locator` | [origin_servers.private_ip.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/#section) |
| `origin_servers.private_ip.site_locator.site` | [origin_servers.private_ip.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/site/#section) |
| `origin_servers.private_ip.site_locator.site.name` | [origin_servers.private_ip.site_locator.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/site/#schema-origin_servers--private_ip--site_locator--site--name) |
| `origin_servers.private_ip.site_locator.site.namespace` | [origin_servers.private_ip.site_locator.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/site/#schema-origin_servers--private_ip--site_locator--site--namespace) |
| `origin_servers.private_ip.site_locator.site.tenant` | [origin_servers.private_ip.site_locator.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/site/#schema-origin_servers--private_ip--site_locator--site--tenant) |
| `origin_servers.private_ip.site_locator.virtual_site` | [origin_servers.private_ip.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/virtual_site/#section) |
| `origin_servers.private_ip.site_locator.virtual_site.name` | [origin_servers.private_ip.site_locator.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/virtual_site/#schema-origin_servers--private_ip--site_locator--virtual_site--name) |
| `origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_servers.private_ip.site_locator.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/virtual_site/#schema-origin_servers--private_ip--site_locator--virtual_site--namespace) |
| `origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_servers.private_ip.site_locator.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/site_locator/virtual_site/#schema-origin_servers--private_ip--site_locator--virtual_site--tenant) |
| `origin_servers.private_ip.snat_pool` | [origin_servers.private_ip.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/snat_pool/#section) |
| `origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_servers.private_ip.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/snat_pool/no_snat_pool/#section) |
| `origin_servers.private_ip.snat_pool.snat_pool` | [origin_servers.private_ip.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/snat_pool/snat_pool/#section) |
| `origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_servers.private_ip.snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_ip/snat_pool/snat_pool/#schema-origin_servers--private_ip--snat_pool--snat_pool--prefixes) |
| `origin_servers.private_name` | [origin_servers.private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/#section) |
| `origin_servers.private_name.dns_name` | [origin_servers.private_name.dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/#schema-origin_servers--private_name--dns_name) |
| `origin_servers.private_name.inside_network` | [origin_servers.private_name.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/inside_network/#section) |
| `origin_servers.private_name.outside_network` | [origin_servers.private_name.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/outside_network/#section) |
| `origin_servers.private_name.refresh_interval` | [origin_servers.private_name.refresh_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/#schema-origin_servers--private_name--refresh_interval) |
| `origin_servers.private_name.segment` | [origin_servers.private_name.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/segment/#section) |
| `origin_servers.private_name.segment.name` | [origin_servers.private_name.segment.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/segment/#schema-origin_servers--private_name--segment--name) |
| `origin_servers.private_name.segment.namespace` | [origin_servers.private_name.segment.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/segment/#schema-origin_servers--private_name--segment--namespace) |
| `origin_servers.private_name.segment.tenant` | [origin_servers.private_name.segment.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/segment/#schema-origin_servers--private_name--segment--tenant) |
| `origin_servers.private_name.site_locator` | [origin_servers.private_name.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/#section) |
| `origin_servers.private_name.site_locator.site` | [origin_servers.private_name.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/site/#section) |
| `origin_servers.private_name.site_locator.site.name` | [origin_servers.private_name.site_locator.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/site/#schema-origin_servers--private_name--site_locator--site--name) |
| `origin_servers.private_name.site_locator.site.namespace` | [origin_servers.private_name.site_locator.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/site/#schema-origin_servers--private_name--site_locator--site--namespace) |
| `origin_servers.private_name.site_locator.site.tenant` | [origin_servers.private_name.site_locator.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/site/#schema-origin_servers--private_name--site_locator--site--tenant) |
| `origin_servers.private_name.site_locator.virtual_site` | [origin_servers.private_name.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/virtual_site/#section) |
| `origin_servers.private_name.site_locator.virtual_site.name` | [origin_servers.private_name.site_locator.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/virtual_site/#schema-origin_servers--private_name--site_locator--virtual_site--name) |
| `origin_servers.private_name.site_locator.virtual_site.namespace` | [origin_servers.private_name.site_locator.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/virtual_site/#schema-origin_servers--private_name--site_locator--virtual_site--namespace) |
| `origin_servers.private_name.site_locator.virtual_site.tenant` | [origin_servers.private_name.site_locator.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/virtual_site/#schema-origin_servers--private_name--site_locator--virtual_site--tenant) |
| `origin_servers.private_name.snat_pool` | [origin_servers.private_name.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/snat_pool/#section) |
| `origin_servers.private_name.snat_pool.no_snat_pool` | [origin_servers.private_name.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/snat_pool/no_snat_pool/#section) |
| `origin_servers.private_name.snat_pool.snat_pool` | [origin_servers.private_name.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/snat_pool/snat_pool/#section) |
| `origin_servers.private_name.snat_pool.snat_pool.prefixes` | [origin_servers.private_name.snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/snat_pool/snat_pool/#schema-origin_servers--private_name--snat_pool--snat_pool--prefixes) |
| `origin_servers.public_ip` | [origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/public_ip/#section) |
| `origin_servers.public_ip.ip` | [origin_servers.public_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/public_ip/#schema-origin_servers--public_ip--ip) |
| `origin_servers.public_name` | [origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/public_name/#section) |
| `origin_servers.public_name.dns_name` | [origin_servers.public_name.dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/public_name/#schema-origin_servers--public_name--dns_name) |
| `origin_servers.public_name.refresh_interval` | [origin_servers.public_name.refresh_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/public_name/#schema-origin_servers--public_name--refresh_interval) |
| `origin_servers.vn_private_ip` | [origin_servers.vn_private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_ip/#section) |
| `origin_servers.vn_private_ip.ip` | [origin_servers.vn_private_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_ip/#schema-origin_servers--vn_private_ip--ip) |
| `origin_servers.vn_private_ip.virtual_network` | [origin_servers.vn_private_ip.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_ip/virtual_network/#section) |
| `origin_servers.vn_private_ip.virtual_network.name` | [origin_servers.vn_private_ip.virtual_network.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_ip/virtual_network/#schema-origin_servers--vn_private_ip--virtual_network--name) |
| `origin_servers.vn_private_ip.virtual_network.namespace` | [origin_servers.vn_private_ip.virtual_network.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_ip/virtual_network/#schema-origin_servers--vn_private_ip--virtual_network--namespace) |
| `origin_servers.vn_private_ip.virtual_network.tenant` | [origin_servers.vn_private_ip.virtual_network.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_ip/virtual_network/#schema-origin_servers--vn_private_ip--virtual_network--tenant) |
| `origin_servers.vn_private_name` | [origin_servers.vn_private_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/#section) |
| `origin_servers.vn_private_name.dns_name` | [origin_servers.vn_private_name.dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/#schema-origin_servers--vn_private_name--dns_name) |
| `origin_servers.vn_private_name.private_network` | [origin_servers.vn_private_name.private_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/private_network/#section) |
| `origin_servers.vn_private_name.private_network.name` | [origin_servers.vn_private_name.private_network.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/private_network/#schema-origin_servers--vn_private_name--private_network--name) |
| `origin_servers.vn_private_name.private_network.namespace` | [origin_servers.vn_private_name.private_network.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/private_network/#schema-origin_servers--vn_private_name--private_network--namespace) |
| `origin_servers.vn_private_name.private_network.tenant` | [origin_servers.vn_private_name.private_network.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/vn_private_name/private_network/#schema-origin_servers--vn_private_name--private_network--tenant) |
| `port` | [port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/#schema-port) |
| `same_as_endpoint_port` | [same_as_endpoint_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/same_as_endpoint_port/#section) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/timeouts/#schema-timeouts--update) |
| `upstream_conn_pool_reuse_type` | [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/#section) |
| `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/disable_conn_pool_reuse/#section) |
| `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` | [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/enable_conn_pool_reuse/#section) |
| `use_tls` | [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/#section) |
| `use_tls.default_session_key_caching` | [use_tls.default_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/default_session_key_caching/#section) |
| `use_tls.disable_session_key_caching` | [use_tls.disable_session_key_caching](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/disable_session_key_caching/#section) |
| `use_tls.disable_sni` | [use_tls.disable_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/disable_sni/#section) |
| `use_tls.max_session_keys` | [use_tls.max_session_keys](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/#schema-use_tls--max_session_keys) |
| `use_tls.no_mtls` | [use_tls.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/no_mtls/#section) |
| `use_tls.skip_server_verification` | [use_tls.skip_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/skip_server_verification/#section) |
| `use_tls.sni` | [use_tls.sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/#schema-use_tls--sni) |
| `use_tls.tls_config` | [use_tls.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/#section) |
| `use_tls.tls_config.custom_security` | [use_tls.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/custom_security/#section) |
| `use_tls.tls_config.custom_security.cipher_suites` | [use_tls.tls_config.custom_security.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/custom_security/#schema-use_tls--tls_config--custom_security--cipher_suites) |
| `use_tls.tls_config.custom_security.max_version` | [use_tls.tls_config.custom_security.max_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/custom_security/#schema-use_tls--tls_config--custom_security--max_version) |
| `use_tls.tls_config.custom_security.min_version` | [use_tls.tls_config.custom_security.min_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/custom_security/#schema-use_tls--tls_config--custom_security--min_version) |
| `use_tls.tls_config.default_security` | [use_tls.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/default_security/#section) |
| `use_tls.tls_config.low_security` | [use_tls.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/low_security/#section) |
| `use_tls.tls_config.medium_security` | [use_tls.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/tls_config/medium_security/#section) |
| `use_tls.use_host_header_as_sni` | [use_tls.use_host_header_as_sni](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_host_header_as_sni/#section) |
| `use_tls.use_mtls` | [use_tls.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/#section) |
| `use_tls.use_mtls.tls_certificates` | [use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/#section) |
| `use_tls.use_mtls.tls_certificates.certificate_url` | [use_tls.use_mtls.tls_certificates.certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/#schema-use_tls--use_mtls--tls_certificates--certificate_url) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/custom_hash_algorithms/#section) |
| `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` | [use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/custom_hash_algorithms/#schema-use_tls--use_mtls--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `use_tls.use_mtls.tls_certificates.description_spec` | [use_tls.use_mtls.tls_certificates.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/#schema-use_tls--use_mtls--tls_certificates--description_spec) |
| `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` | [use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/disable_ocsp_stapling/#section) |
| `use_tls.use_mtls.tls_certificates.private_key` | [use_tls.use_mtls.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/#section) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/#section) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/#schema-use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/#schema-use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--location) |
| `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` | [use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/blindfold_secret_info/#schema-use_tls--use_mtls--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/#section) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/#schema-use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` | [use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/private_key/clear_secret_info/#schema-use_tls--use_mtls--tls_certificates--private_key--clear_secret_info--url) |
| `use_tls.use_mtls.tls_certificates.use_system_defaults` | [use_tls.use_mtls.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/use_system_defaults/#section) |
| `use_tls.use_mtls_obj` | [use_tls.use_mtls_obj](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls_obj/#section) |
| `use_tls.use_mtls_obj.name` | [use_tls.use_mtls_obj.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls_obj/#schema-use_tls--use_mtls_obj--name) |
| `use_tls.use_mtls_obj.namespace` | [use_tls.use_mtls_obj.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls_obj/#schema-use_tls--use_mtls_obj--namespace) |
| `use_tls.use_mtls_obj.tenant` | [use_tls.use_mtls_obj.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls_obj/#schema-use_tls--use_mtls_obj--tenant) |
| `use_tls.use_server_verification` | [use_tls.use_server_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_server_verification/#section) |
| `use_tls.use_server_verification.trusted_ca` | [use_tls.use_server_verification.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_server_verification/trusted_ca/#section) |
| `use_tls.use_server_verification.trusted_ca.name` | [use_tls.use_server_verification.trusted_ca.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_server_verification/trusted_ca/#schema-use_tls--use_server_verification--trusted_ca--name) |
| `use_tls.use_server_verification.trusted_ca.namespace` | [use_tls.use_server_verification.trusted_ca.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_server_verification/trusted_ca/#schema-use_tls--use_server_verification--trusted_ca--namespace) |
| `use_tls.use_server_verification.trusted_ca.tenant` | [use_tls.use_server_verification.trusted_ca.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_server_verification/trusted_ca/#schema-use_tls--use_server_verification--trusted_ca--tenant) |
| `use_tls.use_server_verification.trusted_ca_url` | [use_tls.use_server_verification.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_server_verification/#schema-use_tls--use_server_verification--trusted_ca_url) |
| `use_tls.volterra_trusted_ca` | [use_tls.volterra_trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/volterra_trusted_ca/#section) |

## Next pages

- [advanced_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/advanced_options/)
- [automatic_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/automatic_port/)
- [healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/healthcheck/)
- [lb_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/lb_port/)
- [no_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/no_tls/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [same_as_endpoint_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/same_as_endpoint_port/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/timeouts/)
- [upstream_conn_pool_reuse_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/upstream_conn_pool_reuse_type/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
