---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_virtual_host."
xcsh_docs: {"aliases": ["virtual host"], "body_bytes": 117957, "body_sha256": "sha256:045663b5966cc325e1a30955a636543d05ac9954895a4059ba6b9ffff3c21cfe", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:advertise_policies", "xcsh-docs:resources:virtual_host:properties:authentication", "xcsh-docs:resources:virtual_host:properties:buffer_policy", "xcsh-docs:resources:virtual_host:properties:captcha_challenge", "xcsh-docs:resources:virtual_host:properties:coalescing_options", "xcsh-docs:resources:virtual_host:properties:compression_params", "xcsh-docs:resources:virtual_host:properties:cors_policy", "xcsh-docs:resources:virtual_host:properties:csrf_policy", "xcsh-docs:resources:virtual_host:properties:default_header", "xcsh-docs:resources:virtual_host:properties:default_loadbalancer", "xcsh-docs:resources:virtual_host:properties:disable_path_normalize", "xcsh-docs:resources:virtual_host:properties:dynamic_reverse_proxy", "xcsh-docs:resources:virtual_host:properties:enable_path_normalize", "xcsh-docs:resources:virtual_host:properties:http_protocol_options", "xcsh-docs:resources:virtual_host:properties:js_challenge", "xcsh-docs:resources:virtual_host:properties:no_authentication", "xcsh-docs:resources:virtual_host:properties:no_challenge", "xcsh-docs:resources:virtual_host:properties:no_request_limit_per_connection", "xcsh-docs:resources:virtual_host:properties:non_default_loadbalancer", "xcsh-docs:resources:virtual_host:properties:pass_through", "xcsh-docs:resources:virtual_host:properties:rate_limiter_allowed_prefixes", "xcsh-docs:resources:virtual_host:properties:request_cookies_to_add", "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "xcsh-docs:resources:virtual_host:properties:response_headers_to_add", "xcsh-docs:resources:virtual_host:properties:retry_policy", "xcsh-docs:resources:virtual_host:properties:routes", "xcsh-docs:resources:virtual_host:properties:sensitive_data_policy", "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation", "xcsh-docs:resources:virtual_host:properties:timeouts", "xcsh-docs:resources:virtual_host:properties:tls_cert_params", "xcsh-docs:resources:virtual_host:properties:tls_parameters", "xcsh-docs:resources:virtual_host:properties:user_identification", "xcsh-docs:resources:virtual_host:properties:waf_type"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:reference", "parent_id": "xcsh-docs:resources:virtual_host:fundamentals", "path": "documentation/resources/virtual_host/properties/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0213223201331332-3131232310010210-0330121003120232-2120200222200312-0323331101112032-1020103022113222-1023223202200132-0323130101032220", "registry_path": "docs/guides/resources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["add location"], "anchor": "schema-add_location", "description": "X-example: true Appends header x-F5 Distributed Cloud-location = <RE-site-name> in responses. This configuration is ignored on CE sites.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["add_location"], "syntax": "attribute", "type": "bool"}, {"aliases": ["advertise policies"], "anchor": "section", "description": "Advertise Policy allows you to define networks or sites where you want a VIP for this virtual host to be advertised. Each Policy rule can have different parameters, like TLS configuration, ports, optionally IP address to be used for VIP. If advertise policy is not specified then no VIP is assigned for this virtual", "document_id": "xcsh-docs:resources:virtual_host:properties:advertise_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["advertise_policies"], "syntax": "block", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["append server name"], "anchor": "schema-append_server_name", "description": "Exclusive with Specifies the value to be used for Server header if it is not already present. If Server Header is already present it is not overwritten. It is just passed.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["append_server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication", "credential setup", "credentials"], "anchor": "section", "description": "Authentication related information. This allows to configure the URL to redirect after the authentication Authentication Object Reference, configuration of cookie params etc.", "document_id": "xcsh-docs:resources:virtual_host:properties:authentication", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-authentication--redirect_url", "enforcement": "provider-schema", "group": "authentication:ConflictingObjectAttributes:redirect_dynamic,redirect_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication:ConflictingObjectAttributes:cookie_params,use_auth_object_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication:ConflictingObjectAttributes:redirect_dynamic,redirect_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:redirect_dynamic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication:ConflictingObjectAttributes:cookie_params,use_auth_object_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:use_auth_object_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "authentication:RequiredObjectAttributes:auth_config", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:authentication:auth_config", "type": "requires"}], "schema_path": ["authentication"], "syntax": "block", "type": "object"}, {"aliases": ["buffer policy"], "anchor": "section", "description": "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level", "document_id": "xcsh-docs:resources:virtual_host:properties:buffer_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["buffer_policy"], "syntax": "block", "type": "object"}, {"aliases": ["captcha challenge", "login success", "succeeded", "success", "successful"], "anchor": "section", "description": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "document_id": "xcsh-docs:resources:virtual_host:properties:captcha_challenge", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-captcha_challenge--cookie_expiry", "enforcement": "provider-schema", "group": "captcha_challenge:RequiredObjectAttributes:cookie_expiry", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:captcha_challenge", "type": "requires"}], "schema_path": ["captcha_challenge"], "syntax": "block", "type": "object"}, {"aliases": ["coalescing options"], "anchor": "section", "description": "TLS connection coalescing configuration (not compatible with mTLS)", "document_id": "xcsh-docs:resources:virtual_host:properties:coalescing_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "coalescing_options:ConflictingObjectAttributes:default_coalescing,strict_coalescing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:coalescing_options:default_coalescing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "coalescing_options:ConflictingObjectAttributes:default_coalescing,strict_coalescing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:coalescing_options:strict_coalescing", "type": "conflicts"}], "schema_path": ["coalescing_options"], "syntax": "block", "type": "object"}, {"aliases": ["compression params"], "anchor": "section", "description": "Enables loadbalancer to compress dispatched data from an upstream service upon client request. The content is compressed and then sent to the client with the appropriate headers if either response and request allow. Only GZIP compression is supported. By default compression will be skipped when: A request does NOT", "document_id": "xcsh-docs:resources:virtual_host:properties:compression_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-compression_params--content_length", "enforcement": "provider-schema", "group": "compression_params:RequiredObjectAttributes:content_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:compression_params", "type": "requires"}], "schema_path": ["compression_params"], "syntax": "block", "type": "object"}, {"aliases": ["connection idle timeout", "duration", "operation timeout"], "anchor": "schema-connection_idle_timeout", "description": "The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is specified in milliseconds.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["connection_idle_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["backend servers", "cors policy", "duration", "operation timeout", "origin servers", "upstream servers"], "anchor": "section", "description": "Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route level configuration takes precedence. An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre)", "document_id": "xcsh-docs:resources:virtual_host:properties:cors_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cors_policy"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "csrf policy", "origin servers", "upstream servers"], "anchor": "section", "description": "To mitigate CSRF attack , the policy checks where a request is coming from to determine if the request's origin is the same as its destination.the policy relies on two pieces of information used in determining if a request originated from the same host. 1. The origin that caused the user agent to issue the request", "document_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:all_load_balancer_domains", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,custom_domain_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:custom_domain_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:all_load_balancer_domains,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "csrf_policy:ConflictingObjectAttributes:custom_domain_list,disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:csrf_policy:disabled", "type": "conflicts"}], "schema_path": ["csrf_policy"], "syntax": "block", "type": "object"}, {"aliases": ["custom errors"], "anchor": "schema-custom_errors", "description": "Map of integer error codes as keys and string values that can be used to provide custom HTTP pages for each error code. Key of the map can be either response code class or HTTP Error code. Response code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response code class 5 -- for", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_errors"], "syntax": "attribute", "type": "map"}, {"aliases": ["default header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:default_header", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["default loadbalancer"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:default_loadbalancer", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_loadbalancer"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable default error pages"], "anchor": "schema-disable_default_error_pages", "description": "An option to specify whether to disable using default F5XC error pages.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_default_error_pages"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable dns resolve"], "anchor": "schema-disable_dns_resolve", "description": "Disable DNS resolution for domains specified in the virtual host When the virtual host is configured as Dynamive Resolve Proxy (DRP), disable DNS resolution for domains configured. This configuration is suitable for HTTP CONNECT proxy.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_dns_resolve"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable path normalize"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:disable_path_normalize", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_path_normalize"], "syntax": "attribute", "type": "object"}, {"aliases": ["domains"], "anchor": "schema-domains", "description": "A list of Domains (host/authority header) that will be matched to this Virtual Host. Wildcard hosts are supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names: www.example.com. 2. Domains starting with a Wildcard: *.example.com. Not supported Domains: - Just a Wildcard: * - A", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["domains"], "syntax": "attribute", "type": "list"}, {"aliases": ["dynamic reverse proxy"], "anchor": "section", "description": "In this mode of proxy, virtual host will resolve the destination endpoint dynamically. The dynamic resolution is done using a predefined field in the request. This predefined field depends on the ProxyType configured on the Virtual Host. For HTTP traffic, i.e. With ProxyType as HTTP_PROXY or HTTPS_PROXY, virtual host", "document_id": "xcsh-docs:resources:virtual_host:properties:dynamic_reverse_proxy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_reverse_proxy"], "syntax": "block", "type": "object"}, {"aliases": ["enable path normalize"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:enable_path_normalize", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_path_normalize"], "syntax": "attribute", "type": "object"}, {"aliases": ["http protocol options"], "anchor": "section", "description": "HTTP protocol configuration OPTIONS for downstream connections.", "document_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_only,http_protocol_enable_v1_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_only,http_protocol_enable_v2_only", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_only", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_only,http_protocol_enable_v1_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_v2", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_v2,http_protocol_enable_v2_only", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v1_v2", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_only,http_protocol_enable_v2_only", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v2_only", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_v2,http_protocol_enable_v2_only", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:http_protocol_options:http_protocol_enable_v2_only", "type": "conflicts"}], "schema_path": ["http_protocol_options"], "syntax": "block", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["duration", "idle timeout", "operation timeout"], "anchor": "schema-idle_timeout", "description": "Idle timeout is the amount of time that the loadbalancer will allow a stream to exist with no upstream or downstream activity. Idle timeout and Proxy Type: HTTP_PROXY, HTTPS_PROXY: Idle timer is started when the first byte is received on the connection. Each time an encode/decode event for headers or data is processed", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["idle_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["js challenge"], "anchor": "section", "description": "Enables loadbalancer to perform client browser compatibility test by redirecting to a page with Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do Javascript Challenge, it will", "document_id": "xcsh-docs:resources:virtual_host:properties:js_challenge", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-js_challenge--cookie_expiry", "enforcement": "provider-schema", "group": "js_challenge:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:js_challenge", "type": "requires"}, {"anchor": "schema-js_challenge--js_script_delay", "enforcement": "provider-schema", "group": "js_challenge:RequiredObjectAttributes:cookie_expiry,js_script_delay", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:js_challenge", "type": "requires"}], "schema_path": ["js_challenge"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["max request header size"], "anchor": "schema-max_request_header_size", "description": "The maximum request header size in KiB for incoming connections. If un-configured, the default max request headers allowed is 60 KiB. Requests that exceed this limit will receive a 431 response. The max configurable limit is 96 KiB, based on current implementation constraints. Note: a. This configuration parameter is", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["max_request_header_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["max requests per connection"], "anchor": "schema-max_requests_per_connection", "description": "Exclusive with Sets the maximum number of requests a downstream client can send over a single connection to Envoy. Enter a value >=1 to define the request limit per connection.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["max_requests_per_connection"], "syntax": "attribute", "type": "number"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication", "credential setup", "credentials", "no authentication"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:no_authentication", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_authentication"], "syntax": "attribute", "type": "object"}, {"aliases": ["no challenge"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:no_challenge", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_challenge"], "syntax": "attribute", "type": "object"}, {"aliases": ["no request limit per connection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:no_request_limit_per_connection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_request_limit_per_connection"], "syntax": "attribute", "type": "object"}, {"aliases": ["non default loadbalancer"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:non_default_loadbalancer", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["non_default_loadbalancer"], "syntax": "attribute", "type": "object"}, {"aliases": ["pass through"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:virtual_host:properties:pass_through", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["pass_through"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy"], "anchor": "schema-proxy", "description": "ProxyType tells the type of proxy to install for the virtual host. Only the following combination of VirtualHosts within same AdvertisePolicy is permitted (None of them should have \"*\" in domains when used with other VirtualHosts in same AdvertisePolicy) 1. Multiple TCP_PROXY_WITH_SNI and multiple HTTPS_PROXY 2.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["rate limiter allowed prefixes"], "anchor": "section", "description": "References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.", "document_id": "xcsh-docs:resources:virtual_host:properties:rate_limiter_allowed_prefixes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rate_limiter_allowed_prefixes"], "syntax": "block", "type": "object"}, {"aliases": ["request cookies to add"], "anchor": "section", "description": "Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "document_id": "xcsh-docs:resources:virtual_host:properties:request_cookies_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-request_cookies_to_add--value", "enforcement": "provider-schema", "group": "request_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_cookies_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-request_cookies_to_add--name", "enforcement": "provider-schema", "group": "request_cookies_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_cookies_to_add", "type": "requires"}], "schema_path": ["request_cookies_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["request cookies to remove"], "anchor": "schema-request_cookies_to_remove", "description": "List of keys of Cookies to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["request headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied.", "document_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-request_headers_to_add--value", "enforcement": "provider-schema", "group": "request_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "request_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-request_headers_to_add--name", "enforcement": "provider-schema", "group": "request_headers_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:request_headers_to_add", "type": "requires"}], "schema_path": ["request_headers_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["request headers to remove"], "anchor": "schema-request_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["request_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["response cookies to add"], "anchor": "section", "description": "Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "document_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-response_cookies_to_add--add_domain", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_domain,ignore_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-response_cookies_to_add--add_expiry", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_expiry,ignore_expiry", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-response_cookies_to_add--add_path", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_path,ignore_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-response_cookies_to_add--max_age_value", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_max_age,max_age_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-response_cookies_to_add--value", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-response_cookies_to_add--value", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:add_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_partitioned,ignore_partitioned", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:add_partitioned", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:add_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_domain,ignore_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_expiry,ignore_expiry", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_expiry", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_max_age,max_age_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_max_age", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_partitioned,ignore_partitioned", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_partitioned", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_path,ignore_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,secret_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:ignore_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,secret_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-response_cookies_to_add--name", "enforcement": "provider-schema", "group": "response_cookies_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_cookies_to_add", "type": "requires"}], "schema_path": ["response_cookies_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["response cookies to remove"], "anchor": "schema-response_cookies_to_remove", "description": "List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["response headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied.", "document_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-response_headers_to_add--value", "enforcement": "provider-schema", "group": "response_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "response_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-response_headers_to_add--name", "enforcement": "provider-schema", "group": "response_headers_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:response_headers_to_add", "type": "requires"}], "schema_path": ["response_headers_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["response headers to remove"], "anchor": "schema-response_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP response being sent towards downstream.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["response_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["retry policy"], "anchor": "section", "description": "Retry policy configuration for route destination.", "document_id": "xcsh-docs:resources:virtual_host:properties:retry_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-retry_policy--retry_condition", "enforcement": "provider-schema", "group": "retry_policy:RequiredObjectAttributes:retry_condition", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:retry_policy", "type": "requires"}], "schema_path": ["retry_policy"], "syntax": "block", "type": "object"}, {"aliases": ["routes"], "anchor": "section", "description": "The list of routes that will be matched, in order, for incoming requests. The first route that matches will be used. Currently route object is redundant in case of TCP proxy but required. For TCP_PROXY/TCP_PROXY_WITH_SNI/SMA_PROXY VirtualHosts, the route object only specifies the cluster/weighted-cluster as route", "document_id": "xcsh-docs:resources:virtual_host:properties:routes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes"], "syntax": "block", "type": "object"}, {"aliases": ["sensitive data policy"], "anchor": "section", "description": "References to sensitive_data_policy objects.", "document_id": "xcsh-docs:resources:virtual_host:properties:sensitive_data_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["sensitive_data_policy"], "syntax": "block", "type": "object"}, {"aliases": ["server name"], "anchor": "schema-server_name", "description": "Exclusive with Specifies the value to be used for Server header inserted in responses. This will overwrite existing values if any for Server Header.", "document_id": "xcsh-docs:resources:virtual_host:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["slow ddos mitigation"], "anchor": "section", "description": "\"Slow and low\" attacks tie up server resources, leaving none available for servicing requests from actual users.", "document_id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-slow_ddos_mitigation--request_timeout", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:ConflictingObjectAttributes:disable_request_timeout,request_timeout", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:ConflictingObjectAttributes:disable_request_timeout,request_timeout", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation:disable_request_timeout", "type": "conflicts"}, {"anchor": "schema-slow_ddos_mitigation--request_headers_timeout", "enforcement": "provider-schema", "group": "slow_ddos_mitigation:RequiredObjectAttributes:request_headers_timeout", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:slow_ddos_mitigation", "type": "requires"}], "schema_path": ["slow_ddos_mitigation"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:virtual_host:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["authentication", "cert", "certificate", "credential setup", "credentials", "existing certificates", "tls cert params", "tls certificates"], "anchor": "section", "description": "Certificate Parameters for authentication, TLS ciphers, and trust store.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_cert_params:ConflictingObjectAttributes:client_certificate_optional,client_certificate_required", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_cert_params:ConflictingObjectAttributes:client_certificate_optional,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_cert_params:ConflictingObjectAttributes:client_certificate_optional,client_certificate_required", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_required", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_cert_params:ConflictingObjectAttributes:client_certificate_required,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_required", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_cert_params:ConflictingObjectAttributes:client_certificate_optional,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:no_client_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_cert_params:ConflictingObjectAttributes:client_certificate_required,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:no_client_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_cert_params:RequiredObjectAttributes:certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:certificates", "type": "requires"}], "schema_path": ["tls_cert_params"], "syntax": "block", "type": "object"}, {"aliases": ["tls parameters"], "anchor": "section", "description": "TLS configuration for downstream connections.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,client_certificate_required", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:client_certificate_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:client_certificate_optional", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,client_certificate_required", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:client_certificate_required", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_required,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:client_certificate_required", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_optional,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:no_client_certificate", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_parameters:ConflictingObjectAttributes:client_certificate_required,no_client_certificate", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:no_client_certificate", "type": "conflicts"}], "schema_path": ["tls_parameters"], "syntax": "block", "type": "object"}, {"aliases": ["user identification"], "anchor": "section", "description": "A reference to user_identification object. The rules in the user_identification object are evaluated to determine the user identifier to be rate limited.", "document_id": "xcsh-docs:resources:virtual_host:properties:user_identification", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["user_identification"], "syntax": "block", "type": "object"}, {"aliases": ["waf type"], "anchor": "section", "description": "WAF instance will be pointing to an app_firewall object.", "document_id": "xcsh-docs:resources:virtual_host:properties:waf_type", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_type:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_type:ConflictingObjectAttributes:app_firewall,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:waf_type:app_firewall", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_type:ConflictingObjectAttributes:app_firewall,disable_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:waf_type:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_type:ConflictingObjectAttributes:disable_waf,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:waf_type:disable_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_type:ConflictingObjectAttributes:app_firewall,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:waf_type:inherit_waf", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_type:ConflictingObjectAttributes:disable_waf,inherit_waf", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:virtual_host:properties:waf_type:inherit_waf", "type": "conflicts"}], "schema_path": ["waf_type"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_virtual_host.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- Property reference

## Direct properties

<a id="schema-add_location"></a>

### add_location property

Type: `"bool"`. Optional, Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses. This configuration is ignored on CE sites.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.
This configuration is ignored on CE sites.

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

- [advertise_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/advertise_policies/): complete subsection reference.

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

<a id="schema-append_server_name"></a>

### append_server_name property

Type: `"string"`. Optional, Computed.

\[OneOf: append\_server\_name, default\_header, pass\_through, server\_name; Default:
default\_header\] Exclusive with \[default\_header pass\_through server\_name\] Specifies the value
to be used for Server header if it is not already present. If Server Header is already present it is
not overwritten. It is just passed.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Specifies the value to be used for
Server header if it is not already present. If Server Header is already present it is not
overwritten. It is just passed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

OneOf alternatives in this subsection:

- [append_server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-append_server_name)
- [default_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/default_header/#section)
- [pass_through](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/pass_through/#section)
- [server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-server_name)

Select alternatives according to the provider validators above.

- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/): complete subsection reference.

- [buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/buffer_policy/): complete subsection reference.

- [captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/captcha_challenge/): complete subsection reference.

- [coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/coalescing_options/): complete subsection reference.

- [compression_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/compression_params/): complete subsection reference.

<a id="schema-connection_idle_timeout"></a>

### connection_idle_timeout property

Type: `"number"`. Optional, Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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

- [cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/): complete subsection reference.

- [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/): complete subsection reference.

<a id="schema-custom_errors"></a>

### custom_errors property

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value is the uri\_ref. Currently supported URL schemes
is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify
this message as base64 encoded plain text message e.g. "Access Denied" or it can be HTML paragraph
or a body string encoded as base64 string E.g. "&lt;p&gt; Access Denied &lt;/p&gt;". Base64 encoded
string for this HTML is "PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==" Specific response code takes preference
when both response code and response code class matches for a request.

The configured custom errors are only applicable for loadbalancer generated errors. Errors returned
from upstream server is propagated as is.

F5XC provides default error pages for the errors generated by the loadbalancer. Content of these
pages are not editable. User has an option to disable the use of default F5XC error pages.

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
    "ves.io.schema.rules.map.keys.uint32.gte": "3",
    "ves.io.schema.rules.map.keys.uint32.lte": "599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.gte": "3",
    "ves.io.schema.rules.map.keys.uint32.lte": "599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

- [default_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/default_header/): complete subsection reference.

- [default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/default_loadbalancer/): complete subsection reference.

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

<a id="schema-disable_default_error_pages"></a>

### disable_default_error_pages property

Type: `"bool"`. Optional, Computed.

Option to specify whether to disable using default F5XC error pages.

Upstream description:

An option to specify whether to disable using default F5XC error pages.

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

<a id="schema-disable_dns_resolve"></a>

### disable_dns_resolve property

Type: `"bool"`. Optional, Computed.

Disable DNS resolution for domains specified in the virtual host When the virtual host is configured
as Dynamive Resolve Proxy (DRP), disable DNS resolution for domains configured. This configuration
is suitable for HTTP CONNECT proxy.

Upstream description:

Disable DNS resolution for domains specified in the virtual host

When the virtual host is configured as Dynamive Resolve Proxy (DRP), disable DNS resolution for
domains configured. This configuration is suitable for HTTP CONNECT proxy.

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

- [disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/disable_path_normalize/): complete subsection reference.

<a id="schema-domains"></a>

### domains property

Type: `["list", "string"]`. Optional.

List of domain names matched to this virtual host for routing incoming requests. Supports wildcard
patterns like \*.example.com for subdomain matching.

Upstream description:

A list of Domains (host/authority header) that will be matched to this Virtual Host. Wildcard hosts
are supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the virtual host proxy type is
TCP\_PROXY\_WITH\_SNI/HTTPS\_PROXY Domains also indicate the list of names for which DNS resolution
will be automatically resolved to IP addresses by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 33),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 33,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 33,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "33",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "33",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [dynamic_reverse_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/): complete subsection reference.

- [enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/enable_path_normalize/): complete subsection reference.

- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-idle_timeout"></a>

### idle_timeout property

Type: `"number"`. Optional, Computed.

Idle timeout is the amount of time that the loadbalancer will allow a stream to exist with no
upstream or downstream activity. Idle timeout and Proxy Type: HTTP\_PROXY, HTTPS\_PROXY: Idle timer
is started when the first byte is received on the connection. Each time an encode/decode event for..

Upstream description:

Idle timeout is the amount of time that the loadbalancer will allow a stream to exist with no
upstream or downstream activity.

Idle timeout and Proxy Type:

HTTP\_PROXY, HTTPS\_PROXY: Idle timer is started when the first byte is received on the connection.
Each time an encode/decode event for headers or data is processed for the stream, the timer will be
reset. If the timeout fires, the stream is terminated with a 504 (Gateway Timeout) error code if no
upstream response header has been received, otherwise a stream reset occurs. The default idle
timeout is 30 seconds

TCP PROXY, TCP\_PROXY\_WITH\_SNI, SMA\_PROXY: The idle timeout is defined as the period in which
there are no bytes sent or received on either the upstream or downstream connection. The default
idle timeout is 1 hour.

UDP PROXY: The idle timeout for sessions. Idle timeout is defined as the period in which there are
no datagrams sent or received on the session. The default if not specified is 1 minute.

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

- [js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/js_challenge/): complete subsection reference.

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

<a id="schema-max_request_header_size"></a>

### max_request_header_size property

Type: `"number"`. Optional, Computed.

The maximum request header size in KiB for incoming connections. If un-configured, the default max
request headers allowed is 60 KiB. Requests that exceed this limit will receive a 431 response.

Upstream description:

The maximum request header size in KiB for incoming connections.

If un-configured, the default max request headers allowed is 60 KiB.

Requests that exceed this limit will receive a 431 response.

The max configurable limit is 96 KiB, based on current implementation constraints.

Note: a. This configuration parameter is applicable only for HTTP\_PROXY and HTTPS\_PROXY b. When
multiple HTTP\_PROXY virtual hosts share the same advertise policy, the effective "maximum request
header size" for such virtual hosts is the highest value configured on any of the virtual hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="schema-max_requests_per_connection"></a>

### max_requests_per_connection property

Type: `"number"`. Optional, Computed.

\[OneOf: max\_requests\_per\_connection, no\_request\_limit\_per\_connection; Default:
no\_request\_limit\_per\_connection\] Exclusive with \[no\_request\_limit\_per\_connection\] Sets
the maximum number of requests a downstream client can send over a single connection to Envoy. Enter
a value &gt;=1 to define the request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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

OneOf alternatives in this subsection:

- [max_requests_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-max_requests_per_connection)
- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_request_limit_per_connection/#section)

Select alternatives according to the provider validators above.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Virtual Host. Must be unique within the namespace.

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

Namespace where the Virtual Host is created.

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

- [no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_authentication/): complete subsection reference.

- [no_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_challenge/): complete subsection reference.

- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_request_limit_per_connection/): complete subsection reference.

- [non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/non_default_loadbalancer/): complete subsection reference.

- [pass_through](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/pass_through/): complete subsection reference.

<a id="schema-proxy"></a>

### proxy property

Type: `"string"`. Optional, Computed.

\[Enum:
UDP\_PROXY|SMA\_PROXY|DNS\_PROXY|ZTNA\_PROXY|UZTNA\_PROXY|TMM\_HTTP\_PROXY|TMM\_HTTPS\_PROXY|TMM\_TCP\_PROXY|TMM\_UDP\_PROXY|TMM\_QUIC\_PROXY\]
ProxyType tells the type of proxy to install for the virtual host. Only the following combination of
VirtualHosts within same AdvertisePolicy is permitted (None of them should have '\*' in domains when
used with other VirtualHosts in same AdvertisePolicy) 1. Multiple TCP\_PROXY\_WITH\_SNI and..
Possible values are \`UDP\_PROXY\`, \`SMA\_PROXY\`, \`DNS\_PROXY\`, \`ZTNA\_PROXY\`,
\`UZTNA\_PROXY\`, \`TMM\_HTTP\_PROXY\`, \`TMM\_HTTPS\_PROXY\`, \`TMM\_TCP\_PROXY\`,
\`TMM\_UDP\_PROXY\`, \`TMM\_QUIC\_PROXY\`.

Upstream description:

ProxyType tells the type of proxy to install for the virtual host.

Only the following combination of VirtualHosts within same AdvertisePolicy is permitted (None of
them should have "\*" in domains when used with other VirtualHosts in same AdvertisePolicy)
&#8203;1. Multiple TCP\_PROXY\_WITH\_SNI and multiple HTTPS\_PROXY &#8203;2. Multiple HTTP\_PROXY
&#8203;3. Multiple HTTPS\_PROXY &#8203;4. Multiple TCP\_PROXY\_WITH\_SNI

HTTPS\_PROXY without TLS parameters is not permitted
HTTP\_PROXY/HTTPS\_PROXY/TCP\_PROXY\_WITH\_SNI/SMA\_PROXY with empty domains is not permitted
TCP\_PROXY\_WITH\_SNI/SMA\_PROXY should not have "\*" in domains

&#8203;- HTTP\_PROXY: HTTP\_PROXY

Install HTTP proxy. HTTP Proxy is the default proxy installed. &#8203;- TCP\_PROXY: TCP\_PROXY

Install TCP proxy &#8203;- TCP\_PROXY\_WITH\_SNI: TCP\_PROXY\_WITH\_SNI

Install TCP proxy with SNI Routing &#8203;- TLS\_TCP\_PROXY: TCP\_PROXY

Install TCP proxy &#8203;- TLS\_TCP\_PROXY\_WITH\_SNI: TCP\_PROXY\_WITH\_SNI

Install TCP proxy with SNI Routing &#8203;- HTTPS\_PROXY: HTTPS\_PROXY

Install HTTPS proxy &#8203;- UDP\_PROXY: UDP\_PROXY

Install UDP proxy &#8203;- SMA\_PROXY: SMA\_PROXY

Install Secret Management Access proxy &#8203;- DNS\_PROXY: DNS\_PROXY

Install DNS proxy &#8203;- ZTNA\_PROXY: ZTNA\_PROXY

Install ZTNA proxy.this is going to be deprecated with UZTNA\_PROXY. &#8203;- UZTNA\_PROXY:
UZTNA\_PROXY

Install UZTNA proxy &#8203;- TMM\_HTTP\_PROXY: TMM\_HTTP\_PROXY

Install TMM HTTP proxy for HTTP/1.1 and HTTP/2 traffic. Used by TMM proxy type. &#8203;-
TMM\_HTTPS\_PROXY: TMM\_HTTPS\_PROXY

Install TMM HTTPS proxy for HTTP/1.1 and HTTP/2 traffic. Used by TMM proxy type. &#8203;-
TMM\_TCP\_PROXY: TMM\_TCP\_PROXY

Install TMM TCP proxy for TCP traffic. Used by TMM proxy type. &#8203;- TMM\_UDP\_PROXY:
TMM\_UDP\_PROXY

Install TMM UDP proxy for UDP traffic. Used by TMM proxy type. &#8203;- TMM\_QUIC\_PROXY:
TMM\_QUIC\_PROXY

Install TMM QUIC proxy for HTTP/3 traffic. Used by TMM proxy type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UDP_PROXY",
    "SMA_PROXY",
    "DNS_PROXY",
    "ZTNA_PROXY",
    "UZTNA_PROXY",
    "TMM_HTTP_PROXY",
    "TMM_HTTPS_PROXY",
    "TMM_TCP_PROXY",
    "TMM_UDP_PROXY",
    "TMM_QUIC_PROXY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "UDP_PROXY",
    "SMA_PROXY",
    "DNS_PROXY",
    "ZTNA_PROXY",
    "UZTNA_PROXY",
    "TMM_HTTP_PROXY",
    "TMM_HTTPS_PROXY",
    "TMM_TCP_PROXY",
    "TMM_UDP_PROXY",
    "TMM_QUIC_PROXY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/rate_limiter_allowed_prefixes/): complete subsection reference.

- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/): complete subsection reference.

<a id="schema-request_cookies_to_remove"></a>

### request_cookies_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/): complete subsection reference.

<a id="schema-request_headers_to_remove"></a>

### request_headers_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/): complete subsection reference.

<a id="schema-response_cookies_to_remove"></a>

### response_cookies_to_remove property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/): complete subsection reference.

<a id="schema-response_headers_to_remove"></a>

### response_headers_to_remove property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/): complete subsection reference.

- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/routes/): complete subsection reference.

- [sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/sensitive_data_policy/): complete subsection reference.

<a id="schema-server_name"></a>

### server_name property

Type: `"string"`. Optional, Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Specifies the value to be used
for Server header inserted in responses. This will overwrite existing values if any for Server
Header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/slow_ddos_mitigation/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/timeouts/): complete subsection reference.

- [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/): complete subsection reference.

- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/): complete subsection reference.

- [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/user_identification/): complete subsection reference.

- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `add_location` | [add_location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-add_location) |
| `advertise_policies` | [advertise_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/advertise_policies/#section) |
| `advertise_policies.kind` | [advertise_policies.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/advertise_policies/#schema-advertise_policies--kind) |
| `advertise_policies.name` | [advertise_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/advertise_policies/#schema-advertise_policies--name) |
| `advertise_policies.namespace` | [advertise_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/advertise_policies/#schema-advertise_policies--namespace) |
| `advertise_policies.tenant` | [advertise_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/advertise_policies/#schema-advertise_policies--tenant) |
| `advertise_policies.uid` | [advertise_policies.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/advertise_policies/#schema-advertise_policies--uid) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-annotations) |
| `append_server_name` | [append_server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-append_server_name) |
| `authentication` | [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/#section) |
| `authentication.auth_config` | [authentication.auth_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/auth_config/#section) |
| `authentication.auth_config.kind` | [authentication.auth_config.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/auth_config/#schema-authentication--auth_config--kind) |
| `authentication.auth_config.name` | [authentication.auth_config.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/auth_config/#schema-authentication--auth_config--name) |
| `authentication.auth_config.namespace` | [authentication.auth_config.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/auth_config/#schema-authentication--auth_config--namespace) |
| `authentication.auth_config.tenant` | [authentication.auth_config.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/auth_config/#schema-authentication--auth_config--tenant) |
| `authentication.auth_config.uid` | [authentication.auth_config.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/auth_config/#schema-authentication--auth_config--uid) |
| `authentication.cookie_params` | [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/#section) |
| `authentication.cookie_params.auth_hmac` | [authentication.cookie_params.auth_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/#section) |
| `authentication.cookie_params.auth_hmac.prim_key` | [authentication.cookie_params.auth_hmac.prim_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/#section) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/blindfold_secret_info/#section) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/blindfold_secret_info/#schema-authentication--cookie_params--auth_hmac--prim_key--blindfold_secret_info--decryption_provider) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/blindfold_secret_info/#schema-authentication--cookie_params--auth_hmac--prim_key--blindfold_secret_info--location) |
| `authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.prim_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/blindfold_secret_info/#schema-authentication--cookie_params--auth_hmac--prim_key--blindfold_secret_info--store_provider) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/clear_secret_info/#section) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/clear_secret_info/#schema-authentication--cookie_params--auth_hmac--prim_key--clear_secret_info--provider_ref) |
| `authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.prim_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/prim_key/clear_secret_info/#schema-authentication--cookie_params--auth_hmac--prim_key--clear_secret_info--url) |
| `authentication.cookie_params.auth_hmac.prim_key_expiry` | [authentication.cookie_params.auth_hmac.prim_key_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/#schema-authentication--cookie_params--auth_hmac--prim_key_expiry) |
| `authentication.cookie_params.auth_hmac.sec_key` | [authentication.cookie_params.auth_hmac.sec_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/#section) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/#section) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/#schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--decryption_provider) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/#schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--location) |
| `authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider` | [authentication.cookie_params.auth_hmac.sec_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/blindfold_secret_info/#schema-authentication--cookie_params--auth_hmac--sec_key--blindfold_secret_info--store_provider) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/clear_secret_info/#section) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/clear_secret_info/#schema-authentication--cookie_params--auth_hmac--sec_key--clear_secret_info--provider_ref) |
| `authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url` | [authentication.cookie_params.auth_hmac.sec_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/sec_key/clear_secret_info/#schema-authentication--cookie_params--auth_hmac--sec_key--clear_secret_info--url) |
| `authentication.cookie_params.auth_hmac.sec_key_expiry` | [authentication.cookie_params.auth_hmac.sec_key_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/auth_hmac/#schema-authentication--cookie_params--auth_hmac--sec_key_expiry) |
| `authentication.cookie_params.cookie_expiry` | [authentication.cookie_params.cookie_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/#schema-authentication--cookie_params--cookie_expiry) |
| `authentication.cookie_params.cookie_refresh_interval` | [authentication.cookie_params.cookie_refresh_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/#schema-authentication--cookie_params--cookie_refresh_interval) |
| `authentication.cookie_params.kms_key_hmac` | [authentication.cookie_params.kms_key_hmac](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/kms_key_hmac/#section) |
| `authentication.cookie_params.session_expiry` | [authentication.cookie_params.session_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/#schema-authentication--cookie_params--session_expiry) |
| `authentication.redirect_dynamic` | [authentication.redirect_dynamic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/redirect_dynamic/#section) |
| `authentication.redirect_url` | [authentication.redirect_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/#schema-authentication--redirect_url) |
| `authentication.use_auth_object_config` | [authentication.use_auth_object_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/use_auth_object_config/#section) |
| `buffer_policy` | [buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/buffer_policy/#section) |
| `buffer_policy.disabled` | [buffer_policy.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/buffer_policy/#schema-buffer_policy--disabled) |
| `buffer_policy.max_request_bytes` | [buffer_policy.max_request_bytes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/buffer_policy/#schema-buffer_policy--max_request_bytes) |
| `captcha_challenge` | [captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/captcha_challenge/#section) |
| `captcha_challenge.cookie_expiry` | [captcha_challenge.cookie_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/captcha_challenge/#schema-captcha_challenge--cookie_expiry) |
| `captcha_challenge.custom_page` | [captcha_challenge.custom_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/captcha_challenge/#schema-captcha_challenge--custom_page) |
| `coalescing_options` | [coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/coalescing_options/#section) |
| `coalescing_options.default_coalescing` | [coalescing_options.default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/coalescing_options/default_coalescing/#section) |
| `coalescing_options.strict_coalescing` | [coalescing_options.strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/coalescing_options/strict_coalescing/#section) |
| `compression_params` | [compression_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/compression_params/#section) |
| `compression_params.content_length` | [compression_params.content_length](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/compression_params/#schema-compression_params--content_length) |
| `compression_params.content_type` | [compression_params.content_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/compression_params/#schema-compression_params--content_type) |
| `compression_params.disable_on_etag_header` | [compression_params.disable_on_etag_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/compression_params/#schema-compression_params--disable_on_etag_header) |
| `compression_params.remove_accept_encoding_header` | [compression_params.remove_accept_encoding_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/compression_params/#schema-compression_params--remove_accept_encoding_header) |
| `connection_idle_timeout` | [connection_idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-connection_idle_timeout) |
| `cors_policy` | [cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#section) |
| `cors_policy.allow_credentials` | [cors_policy.allow_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#schema-cors_policy--allow_credentials) |
| `cors_policy.allow_headers` | [cors_policy.allow_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#schema-cors_policy--allow_headers) |
| `cors_policy.allow_methods` | [cors_policy.allow_methods](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#schema-cors_policy--allow_methods) |
| `cors_policy.allow_origin` | [cors_policy.allow_origin](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#schema-cors_policy--allow_origin) |
| `cors_policy.allow_origin_regex` | [cors_policy.allow_origin_regex](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#schema-cors_policy--allow_origin_regex) |
| `cors_policy.disabled` | [cors_policy.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#schema-cors_policy--disabled) |
| `cors_policy.expose_headers` | [cors_policy.expose_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#schema-cors_policy--expose_headers) |
| `cors_policy.maximum_age` | [cors_policy.maximum_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/#schema-cors_policy--maximum_age) |
| `csrf_policy` | [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/#section) |
| `csrf_policy.all_load_balancer_domains` | [csrf_policy.all_load_balancer_domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/all_load_balancer_domains/#section) |
| `csrf_policy.custom_domain_list` | [csrf_policy.custom_domain_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/custom_domain_list/#section) |
| `csrf_policy.custom_domain_list.domains` | [csrf_policy.custom_domain_list.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/custom_domain_list/#schema-csrf_policy--custom_domain_list--domains) |
| `csrf_policy.disabled` | [csrf_policy.disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/disabled/#section) |
| `custom_errors` | [custom_errors](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-custom_errors) |
| `default_header` | [default_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/default_header/#section) |
| `default_loadbalancer` | [default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/default_loadbalancer/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-disable) |
| `disable_default_error_pages` | [disable_default_error_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-disable_default_error_pages) |
| `disable_dns_resolve` | [disable_dns_resolve](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-disable_dns_resolve) |
| `disable_path_normalize` | [disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/disable_path_normalize/#section) |
| `domains` | [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-domains) |
| `dynamic_reverse_proxy` | [dynamic_reverse_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/#section) |
| `dynamic_reverse_proxy.connection_timeout` | [dynamic_reverse_proxy.connection_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/#schema-dynamic_reverse_proxy--connection_timeout) |
| `dynamic_reverse_proxy.resolution_network` | [dynamic_reverse_proxy.resolution_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/resolution_network/#section) |
| `dynamic_reverse_proxy.resolution_network.kind` | [dynamic_reverse_proxy.resolution_network.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/resolution_network/#schema-dynamic_reverse_proxy--resolution_network--kind) |
| `dynamic_reverse_proxy.resolution_network.name` | [dynamic_reverse_proxy.resolution_network.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/resolution_network/#schema-dynamic_reverse_proxy--resolution_network--name) |
| `dynamic_reverse_proxy.resolution_network.namespace` | [dynamic_reverse_proxy.resolution_network.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/resolution_network/#schema-dynamic_reverse_proxy--resolution_network--namespace) |
| `dynamic_reverse_proxy.resolution_network.tenant` | [dynamic_reverse_proxy.resolution_network.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/resolution_network/#schema-dynamic_reverse_proxy--resolution_network--tenant) |
| `dynamic_reverse_proxy.resolution_network.uid` | [dynamic_reverse_proxy.resolution_network.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/resolution_network/#schema-dynamic_reverse_proxy--resolution_network--uid) |
| `dynamic_reverse_proxy.resolution_network_type` | [dynamic_reverse_proxy.resolution_network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/#schema-dynamic_reverse_proxy--resolution_network_type) |
| `dynamic_reverse_proxy.resolve_endpoint_dynamically` | [dynamic_reverse_proxy.resolve_endpoint_dynamically](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/#schema-dynamic_reverse_proxy--resolve_endpoint_dynamically) |
| `enable_path_normalize` | [enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/enable_path_normalize/#section) |
| `http_protocol_options` | [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/#section) |
| `http_protocol_options.http_protocol_enable_v1_only` | [http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/#section) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/#section) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/#section) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/#section) |
| `http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/#section) |
| `http_protocol_options.http_protocol_enable_v1_v2` | [http_protocol_options.http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v1_v2/#section) |
| `http_protocol_options.http_protocol_enable_v2_only` | [http_protocol_options.http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/http_protocol_enable_v2_only/#section) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-id) |
| `idle_timeout` | [idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-idle_timeout) |
| `js_challenge` | [js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/js_challenge/#section) |
| `js_challenge.cookie_expiry` | [js_challenge.cookie_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/js_challenge/#schema-js_challenge--cookie_expiry) |
| `js_challenge.custom_page` | [js_challenge.custom_page](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/js_challenge/#schema-js_challenge--custom_page) |
| `js_challenge.js_script_delay` | [js_challenge.js_script_delay](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/js_challenge/#schema-js_challenge--js_script_delay) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-labels) |
| `max_request_header_size` | [max_request_header_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-max_request_header_size) |
| `max_requests_per_connection` | [max_requests_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-max_requests_per_connection) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-namespace) |
| `no_authentication` | [no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_authentication/#section) |
| `no_challenge` | [no_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_challenge/#section) |
| `no_request_limit_per_connection` | [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_request_limit_per_connection/#section) |
| `non_default_loadbalancer` | [non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/non_default_loadbalancer/#section) |
| `pass_through` | [pass_through](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/pass_through/#section) |
| `proxy` | [proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-proxy) |
| `rate_limiter_allowed_prefixes` | [rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/rate_limiter_allowed_prefixes/#section) |
| `rate_limiter_allowed_prefixes.kind` | [rate_limiter_allowed_prefixes.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/rate_limiter_allowed_prefixes/#schema-rate_limiter_allowed_prefixes--kind) |
| `rate_limiter_allowed_prefixes.name` | [rate_limiter_allowed_prefixes.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/rate_limiter_allowed_prefixes/#schema-rate_limiter_allowed_prefixes--name) |
| `rate_limiter_allowed_prefixes.namespace` | [rate_limiter_allowed_prefixes.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/rate_limiter_allowed_prefixes/#schema-rate_limiter_allowed_prefixes--namespace) |
| `rate_limiter_allowed_prefixes.tenant` | [rate_limiter_allowed_prefixes.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/rate_limiter_allowed_prefixes/#schema-rate_limiter_allowed_prefixes--tenant) |
| `rate_limiter_allowed_prefixes.uid` | [rate_limiter_allowed_prefixes.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/rate_limiter_allowed_prefixes/#schema-rate_limiter_allowed_prefixes--uid) |
| `request_cookies_to_add` | [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/#section) |
| `request_cookies_to_add.name` | [request_cookies_to_add.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/#schema-request_cookies_to_add--name) |
| `request_cookies_to_add.overwrite` | [request_cookies_to_add.overwrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/#schema-request_cookies_to_add--overwrite) |
| `request_cookies_to_add.secret_value` | [request_cookies_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/#section) |
| `request_cookies_to_add.secret_value.blindfold_secret_info` | [request_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/blindfold_secret_info/#section) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/blindfold_secret_info/#schema-request_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.location` | [request_cookies_to_add.secret_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/blindfold_secret_info/#schema-request_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/blindfold_secret_info/#schema-request_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `request_cookies_to_add.secret_value.clear_secret_info` | [request_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/clear_secret_info/#section) |
| `request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [request_cookies_to_add.secret_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/clear_secret_info/#schema-request_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `request_cookies_to_add.secret_value.clear_secret_info.url` | [request_cookies_to_add.secret_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/secret_value/clear_secret_info/#schema-request_cookies_to_add--secret_value--clear_secret_info--url) |
| `request_cookies_to_add.value` | [request_cookies_to_add.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/#schema-request_cookies_to_add--value) |
| `request_cookies_to_remove` | [request_cookies_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-request_cookies_to_remove) |
| `request_headers_to_add` | [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/#section) |
| `request_headers_to_add.append` | [request_headers_to_add.append](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/#schema-request_headers_to_add--append) |
| `request_headers_to_add.name` | [request_headers_to_add.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/#schema-request_headers_to_add--name) |
| `request_headers_to_add.secret_value` | [request_headers_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/#section) |
| `request_headers_to_add.secret_value.blindfold_secret_info` | [request_headers_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/blindfold_secret_info/#section) |
| `request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/blindfold_secret_info/#schema-request_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `request_headers_to_add.secret_value.blindfold_secret_info.location` | [request_headers_to_add.secret_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/blindfold_secret_info/#schema-request_headers_to_add--secret_value--blindfold_secret_info--location) |
| `request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [request_headers_to_add.secret_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/blindfold_secret_info/#schema-request_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `request_headers_to_add.secret_value.clear_secret_info` | [request_headers_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/clear_secret_info/#section) |
| `request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [request_headers_to_add.secret_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/clear_secret_info/#schema-request_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `request_headers_to_add.secret_value.clear_secret_info.url` | [request_headers_to_add.secret_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/secret_value/clear_secret_info/#schema-request_headers_to_add--secret_value--clear_secret_info--url) |
| `request_headers_to_add.value` | [request_headers_to_add.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/#schema-request_headers_to_add--value) |
| `request_headers_to_remove` | [request_headers_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-request_headers_to_remove) |
| `response_cookies_to_add` | [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/#section) |
| `response_cookies_to_add.add_domain` | [response_cookies_to_add.add_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/#schema-response_cookies_to_add--add_domain) |
| `response_cookies_to_add.add_expiry` | [response_cookies_to_add.add_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/#schema-response_cookies_to_add--add_expiry) |
| `response_cookies_to_add.add_httponly` | [response_cookies_to_add.add_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/add_httponly/#section) |
| `response_cookies_to_add.add_partitioned` | [response_cookies_to_add.add_partitioned](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/add_partitioned/#section) |
| `response_cookies_to_add.add_path` | [response_cookies_to_add.add_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/#schema-response_cookies_to_add--add_path) |
| `response_cookies_to_add.add_secure` | [response_cookies_to_add.add_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/add_secure/#section) |
| `response_cookies_to_add.ignore_domain` | [response_cookies_to_add.ignore_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_domain/#section) |
| `response_cookies_to_add.ignore_expiry` | [response_cookies_to_add.ignore_expiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_expiry/#section) |
| `response_cookies_to_add.ignore_httponly` | [response_cookies_to_add.ignore_httponly](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_httponly/#section) |
| `response_cookies_to_add.ignore_max_age` | [response_cookies_to_add.ignore_max_age](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_max_age/#section) |
| `response_cookies_to_add.ignore_partitioned` | [response_cookies_to_add.ignore_partitioned](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_partitioned/#section) |
| `response_cookies_to_add.ignore_path` | [response_cookies_to_add.ignore_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_path/#section) |
| `response_cookies_to_add.ignore_samesite` | [response_cookies_to_add.ignore_samesite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_samesite/#section) |
| `response_cookies_to_add.ignore_secure` | [response_cookies_to_add.ignore_secure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_secure/#section) |
| `response_cookies_to_add.ignore_value` | [response_cookies_to_add.ignore_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/ignore_value/#section) |
| `response_cookies_to_add.max_age_value` | [response_cookies_to_add.max_age_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/#schema-response_cookies_to_add--max_age_value) |
| `response_cookies_to_add.name` | [response_cookies_to_add.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/#schema-response_cookies_to_add--name) |
| `response_cookies_to_add.overwrite` | [response_cookies_to_add.overwrite](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/#schema-response_cookies_to_add--overwrite) |
| `response_cookies_to_add.samesite_lax` | [response_cookies_to_add.samesite_lax](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/samesite_lax/#section) |
| `response_cookies_to_add.samesite_none` | [response_cookies_to_add.samesite_none](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/samesite_none/#section) |
| `response_cookies_to_add.samesite_strict` | [response_cookies_to_add.samesite_strict](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/samesite_strict/#section) |
| `response_cookies_to_add.secret_value` | [response_cookies_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/secret_value/#section) |
| `response_cookies_to_add.secret_value.blindfold_secret_info` | [response_cookies_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/secret_value/blindfold_secret_info/#section) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/secret_value/blindfold_secret_info/#schema-response_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.location` | [response_cookies_to_add.secret_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/secret_value/blindfold_secret_info/#schema-response_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/secret_value/blindfold_secret_info/#schema-response_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `response_cookies_to_add.secret_value.clear_secret_info` | [response_cookies_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/secret_value/clear_secret_info/#section) |
| `response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [response_cookies_to_add.secret_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/secret_value/clear_secret_info/#schema-response_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `response_cookies_to_add.secret_value.clear_secret_info.url` | [response_cookies_to_add.secret_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/secret_value/clear_secret_info/#schema-response_cookies_to_add--secret_value--clear_secret_info--url) |
| `response_cookies_to_add.value` | [response_cookies_to_add.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/#schema-response_cookies_to_add--value) |
| `response_cookies_to_remove` | [response_cookies_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-response_cookies_to_remove) |
| `response_headers_to_add` | [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/#section) |
| `response_headers_to_add.append` | [response_headers_to_add.append](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/#schema-response_headers_to_add--append) |
| `response_headers_to_add.name` | [response_headers_to_add.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/#schema-response_headers_to_add--name) |
| `response_headers_to_add.secret_value` | [response_headers_to_add.secret_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/#section) |
| `response_headers_to_add.secret_value.blindfold_secret_info` | [response_headers_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/blindfold_secret_info/#section) |
| `response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/blindfold_secret_info/#schema-response_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `response_headers_to_add.secret_value.blindfold_secret_info.location` | [response_headers_to_add.secret_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/blindfold_secret_info/#schema-response_headers_to_add--secret_value--blindfold_secret_info--location) |
| `response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [response_headers_to_add.secret_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/blindfold_secret_info/#schema-response_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `response_headers_to_add.secret_value.clear_secret_info` | [response_headers_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/clear_secret_info/#section) |
| `response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [response_headers_to_add.secret_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/clear_secret_info/#schema-response_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `response_headers_to_add.secret_value.clear_secret_info.url` | [response_headers_to_add.secret_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/secret_value/clear_secret_info/#schema-response_headers_to_add--secret_value--clear_secret_info--url) |
| `response_headers_to_add.value` | [response_headers_to_add.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/#schema-response_headers_to_add--value) |
| `response_headers_to_remove` | [response_headers_to_remove](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-response_headers_to_remove) |
| `retry_policy` | [retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/#section) |
| `retry_policy.back_off` | [retry_policy.back_off](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/back_off/#section) |
| `retry_policy.back_off.base_interval` | [retry_policy.back_off.base_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/back_off/#schema-retry_policy--back_off--base_interval) |
| `retry_policy.back_off.max_interval` | [retry_policy.back_off.max_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/back_off/#schema-retry_policy--back_off--max_interval) |
| `retry_policy.num_retries` | [retry_policy.num_retries](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/#schema-retry_policy--num_retries) |
| `retry_policy.per_try_timeout` | [retry_policy.per_try_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/#schema-retry_policy--per_try_timeout) |
| `retry_policy.retriable_status_codes` | [retry_policy.retriable_status_codes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/#schema-retry_policy--retriable_status_codes) |
| `retry_policy.retry_condition` | [retry_policy.retry_condition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/#schema-retry_policy--retry_condition) |
| `routes` | [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/routes/#section) |
| `routes.kind` | [routes.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/routes/#schema-routes--kind) |
| `routes.name` | [routes.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/routes/#schema-routes--name) |
| `routes.namespace` | [routes.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/routes/#schema-routes--namespace) |
| `routes.tenant` | [routes.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/routes/#schema-routes--tenant) |
| `routes.uid` | [routes.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/routes/#schema-routes--uid) |
| `sensitive_data_policy` | [sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/sensitive_data_policy/#section) |
| `sensitive_data_policy.kind` | [sensitive_data_policy.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/sensitive_data_policy/#schema-sensitive_data_policy--kind) |
| `sensitive_data_policy.name` | [sensitive_data_policy.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/sensitive_data_policy/#schema-sensitive_data_policy--name) |
| `sensitive_data_policy.namespace` | [sensitive_data_policy.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/sensitive_data_policy/#schema-sensitive_data_policy--namespace) |
| `sensitive_data_policy.tenant` | [sensitive_data_policy.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/sensitive_data_policy/#schema-sensitive_data_policy--tenant) |
| `sensitive_data_policy.uid` | [sensitive_data_policy.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/sensitive_data_policy/#schema-sensitive_data_policy--uid) |
| `server_name` | [server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/#schema-server_name) |
| `slow_ddos_mitigation` | [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/slow_ddos_mitigation/#section) |
| `slow_ddos_mitigation.disable_request_timeout` | [slow_ddos_mitigation.disable_request_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/slow_ddos_mitigation/disable_request_timeout/#section) |
| `slow_ddos_mitigation.request_headers_timeout` | [slow_ddos_mitigation.request_headers_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/slow_ddos_mitigation/#schema-slow_ddos_mitigation--request_headers_timeout) |
| `slow_ddos_mitigation.request_timeout` | [slow_ddos_mitigation.request_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/slow_ddos_mitigation/#schema-slow_ddos_mitigation--request_timeout) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/timeouts/#schema-timeouts--update) |
| `tls_cert_params` | [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/#section) |
| `tls_cert_params.certificates` | [tls_cert_params.certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/certificates/#section) |
| `tls_cert_params.certificates.kind` | [tls_cert_params.certificates.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/certificates/#schema-tls_cert_params--certificates--kind) |
| `tls_cert_params.certificates.name` | [tls_cert_params.certificates.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/certificates/#schema-tls_cert_params--certificates--name) |
| `tls_cert_params.certificates.namespace` | [tls_cert_params.certificates.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/certificates/#schema-tls_cert_params--certificates--namespace) |
| `tls_cert_params.certificates.tenant` | [tls_cert_params.certificates.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/certificates/#schema-tls_cert_params--certificates--tenant) |
| `tls_cert_params.certificates.uid` | [tls_cert_params.certificates.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/certificates/#schema-tls_cert_params--certificates--uid) |
| `tls_cert_params.cipher_suites` | [tls_cert_params.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/#schema-tls_cert_params--cipher_suites) |
| `tls_cert_params.client_certificate_optional` | [tls_cert_params.client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/client_certificate_optional/#section) |
| `tls_cert_params.client_certificate_required` | [tls_cert_params.client_certificate_required](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/client_certificate_required/#section) |
| `tls_cert_params.maximum_protocol_version` | [tls_cert_params.maximum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/#schema-tls_cert_params--maximum_protocol_version) |
| `tls_cert_params.minimum_protocol_version` | [tls_cert_params.minimum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/#schema-tls_cert_params--minimum_protocol_version) |
| `tls_cert_params.no_client_certificate` | [tls_cert_params.no_client_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/no_client_certificate/#section) |
| `tls_cert_params.validation_params` | [tls_cert_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/#section) |
| `tls_cert_params.validation_params.skip_hostname_verification` | [tls_cert_params.validation_params.skip_hostname_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/#schema-tls_cert_params--validation_params--skip_hostname_verification) |
| `tls_cert_params.validation_params.trusted_ca` | [tls_cert_params.validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/#section) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/trusted_ca_list/#section) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_cert_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_cert_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_cert_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_cert_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_cert_params.validation_params.trusted_ca.trusted_ca_list.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_cert_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_cert_params.validation_params.trusted_ca_url` | [tls_cert_params.validation_params.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/#schema-tls_cert_params--validation_params--trusted_ca_url) |
| `tls_cert_params.validation_params.verify_subject_alt_names` | [tls_cert_params.validation_params.verify_subject_alt_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/#schema-tls_cert_params--validation_params--verify_subject_alt_names) |
| `tls_cert_params.xfcc_header_elements` | [tls_cert_params.xfcc_header_elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/#schema-tls_cert_params--xfcc_header_elements) |
| `tls_parameters` | [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/#section) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/client_certificate_optional/#section) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/client_certificate_required/#section) |
| `tls_parameters.common_params` | [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/#section) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--cipher_suites) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--maximum_protocol_version) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/#schema-tls_parameters--common_params--minimum_protocol_version) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/#section) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/#schema-tls_parameters--common_params--tls_certificates--certificate_url) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/#section) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/custom_hash_algorithms/#schema-tls_parameters--common_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/#schema-tls_parameters--common_params--tls_certificates--description_spec) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/#section) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/blindfold_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#section) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/private_key/clear_secret_info/#schema-tls_parameters--common_params--tls_certificates--private_key--clear_secret_info--url) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/#section) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/#section) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--skip_hostname_verification) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#section) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--kind) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--name) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--namespace) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--tenant) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/#schema-tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list--uid) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--trusted_ca_url) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/validation_params/#schema-tls_parameters--common_params--validation_params--verify_subject_alt_names) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/no_client_certificate/#section) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/#schema-tls_parameters--xfcc_header_elements) |
| `user_identification` | [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/user_identification/#section) |
| `user_identification.kind` | [user_identification.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/user_identification/#schema-user_identification--kind) |
| `user_identification.name` | [user_identification.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/user_identification/#schema-user_identification--name) |
| `user_identification.namespace` | [user_identification.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/user_identification/#schema-user_identification--namespace) |
| `user_identification.tenant` | [user_identification.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/user_identification/#schema-user_identification--tenant) |
| `user_identification.uid` | [user_identification.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/user_identification/#schema-user_identification--uid) |
| `waf_type` | [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/#section) |
| `waf_type.app_firewall` | [waf_type.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/#section) |
| `waf_type.app_firewall.app_firewall` | [waf_type.app_firewall.app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/#section) |
| `waf_type.app_firewall.app_firewall.kind` | [waf_type.app_firewall.app_firewall.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/#schema-waf_type--app_firewall--app_firewall--kind) |
| `waf_type.app_firewall.app_firewall.name` | [waf_type.app_firewall.app_firewall.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/#schema-waf_type--app_firewall--app_firewall--name) |
| `waf_type.app_firewall.app_firewall.namespace` | [waf_type.app_firewall.app_firewall.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/#schema-waf_type--app_firewall--app_firewall--namespace) |
| `waf_type.app_firewall.app_firewall.tenant` | [waf_type.app_firewall.app_firewall.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/#schema-waf_type--app_firewall--app_firewall--tenant) |
| `waf_type.app_firewall.app_firewall.uid` | [waf_type.app_firewall.app_firewall.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/app_firewall/app_firewall/#schema-waf_type--app_firewall--app_firewall--uid) |
| `waf_type.disable_waf` | [waf_type.disable_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/disable_waf/#section) |
| `waf_type.inherit_waf` | [waf_type.inherit_waf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/inherit_waf/#section) |

## Next pages

- [advertise_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/advertise_policies/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/)
- [buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/buffer_policy/)
- [captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/captcha_challenge/)
- [coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/coalescing_options/)
- [compression_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/compression_params/)
- [cors_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/cors_policy/)
- [csrf_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/csrf_policy/)
- [default_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/default_header/)
- [default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/default_loadbalancer/)
- [disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/disable_path_normalize/)
- [dynamic_reverse_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/dynamic_reverse_proxy/)
- [enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/enable_path_normalize/)
- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/http_protocol_options/)
- [js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/js_challenge/)
- [no_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_authentication/)
- [no_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_challenge/)
- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/no_request_limit_per_connection/)
- [non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/non_default_loadbalancer/)
- [pass_through](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/pass_through/)
- [rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/rate_limiter_allowed_prefixes/)
- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_cookies_to_add/)
- [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/request_headers_to_add/)
- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_cookies_to_add/)
- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/response_headers_to_add/)
- [retry_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/retry_policy/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/routes/)
- [sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/sensitive_data_policy/)
- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/slow_ddos_mitigation/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/timeouts/)
- [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/)
- [user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/user_identification/)
- [waf_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/waf_type/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
