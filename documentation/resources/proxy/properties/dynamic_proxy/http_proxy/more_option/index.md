---
page_title: "dynamic_proxy.http_proxy.more_option"
subcategory: ""
description: "This defines various OPTIONS to define a route."
xcsh_docs: {"aliases": ["dynamic proxy http proxy more option"], "body_bytes": 17702, "body_sha256": "sha256:f0d080c1d13bd7beb46ea37f34ada0566b697694e756f84f0a7399bac546894f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:buffer_policy", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:compression_params", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:disable_path_normalize", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:enable_path_normalize", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:no_request_limit_per_connection", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy", "path": "documentation/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3010122303200211-1222020320103311-0221100322131212-1011330220333232-1132103312221022-3132110222033213-3121010101320213-3230132212112232", "registry_path": "docs/guides/resources--proxy--reference--group-001.md", "relationships": [{"anchor": "schema-dynamic_proxy--http_proxy--more_option--max_requests_per_connection", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:disable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:enable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option:ConflictingObjectAttributes:max_requests_per_connection,no_request_limit_per_connection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:no_request_limit_per_connection", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "http_proxy", "more_option"], "schema_version": 1, "sections": [{"aliases": ["buffer policy"], "anchor": "section", "description": "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:buffer_policy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "buffer_policy"], "syntax": "block", "type": "object"}, {"aliases": ["compression params"], "anchor": "section", "description": "Enables loadbalancer to compress dispatched data from an upstream service upon client request. The content is compressed and then sent to the client with the appropriate headers if either response and request allow. Only GZIP compression is supported. By default compression will be skipped when: A request does NOT", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:compression_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-dynamic_proxy--http_proxy--more_option--compression_params--content_length", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.compression_params:RequiredObjectAttributes:content_length", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:compression_params", "type": "requires"}], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "compression_params"], "syntax": "block", "type": "object"}, {"aliases": ["custom errors"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--custom_errors", "description": "Map of integer error codes as keys and string values that can be used to provide custom HTTP pages for each error code. Key of the map can be either response code class or HTTP Error code. Response code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response code class 5 -- for", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "custom_errors"], "syntax": "attribute", "type": "map"}, {"aliases": ["disable default error pages"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--disable_default_error_pages", "description": "Disable the use of default F5XC error pages.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "disable_default_error_pages"], "syntax": "attribute", "type": "bool"}, {"aliases": ["disable path normalize"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:disable_path_normalize", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "disable_path_normalize"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable path normalize"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:enable_path_normalize", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "enable_path_normalize"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "idle timeout", "operation timeout"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--idle_timeout", "description": "The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header has been received, otherwise the stream is reset.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "idle_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["max request header size"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--max_request_header_size", "description": "The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers share the same advertise_policy, the highest value configured across all such load balancers is used for all the load", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "max_request_header_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["max requests per connection"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--max_requests_per_connection", "description": "Exclusive with Sets the maximum number of requests a downstream client can send over a single connection to Envoy. Enter a value >=1 to define the request limit per connection.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "max_requests_per_connection"], "syntax": "attribute", "type": "number"}, {"aliases": ["no request limit per connection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:no_request_limit_per_connection", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "no_request_limit_per_connection"], "syntax": "attribute", "type": "object"}, {"aliases": ["request cookies to add"], "anchor": "section", "description": "Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--value", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--name", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_cookies_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_cookies_to_add", "type": "requires"}], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "request_cookies_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["request cookies to remove"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_remove", "description": "List of keys of Cookies to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "request_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["request headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--value", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--name", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.request_headers_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:request_headers_to_add", "type": "requires"}], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "request_headers_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["request headers to remove"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--request_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "request_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["response cookies to add"], "anchor": "section", "description": "Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_domain", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_domain,ignore_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_expiry", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_expiry,ignore_expiry", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_path", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_path,ignore_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--max_age_value", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_max_age,max_age_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--value", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--value", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:add_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_partitioned,ignore_partitioned", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:add_partitioned", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:add_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_domain,ignore_domain", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_domain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_expiry,ignore_expiry", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_expiry", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_httponly,ignore_httponly", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_httponly", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_max_age,max_age_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_max_age", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_partitioned,ignore_partitioned", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_partitioned", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_path,ignore_path", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_path", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_samesite", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:add_secure,ignore_secure", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_secure", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,secret_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:ignore_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_lax", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_lax", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_none", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_samesite,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:samesite_lax,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:samesite_none,samesite_strict", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:samesite_strict", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:ignore_value,secret_value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--name", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_cookies_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_cookies_to_add", "type": "requires"}], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_cookies_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["response cookies to remove"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_remove", "description": "List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["response headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--value", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_headers_to_add:ConflictingListObjectAttributes:secret_value,value", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add:secret_value", "type": "conflicts"}, {"anchor": "schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--name", "enforcement": "provider-schema", "group": "dynamic_proxy.http_proxy.more_option.response_headers_to_add:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option:response_headers_to_add", "type": "requires"}], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_headers_to_add"], "syntax": "block", "type": "object"}, {"aliases": ["response headers to remove"], "anchor": "schema-dynamic_proxy--http_proxy--more_option--response_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP response being sent towards downstream.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:http_proxy:more_option", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "http_proxy", "more_option", "response_headers_to_remove"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines various OPTIONS to define a route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.http_proxy.more_option

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/)
- dynamic_proxy.http_proxy.more_option

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

## Direct properties

- [buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/buffer_policy/): complete subsection reference.

- [compression_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/compression_params/): complete subsection reference.

<a id="schema-dynamic_proxy--http_proxy--more_option--custom_errors"></a>

### custom_errors property

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="schema-dynamic_proxy--http_proxy--more_option--disable_default_error_pages"></a>

### disable_default_error_pages property

Type: `"bool"`. Optional.

Disable the use of default F5XC error pages.

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

- [disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/disable_path_normalize/): complete subsection reference.

- [enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/enable_path_normalize/): complete subsection reference.

<a id="schema-dynamic_proxy--http_proxy--more_option--idle_timeout"></a>

### idle_timeout property

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="schema-dynamic_proxy--http_proxy--more_option--max_request_header_size"></a>

### max_request_header_size property

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

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

<a id="schema-dynamic_proxy--http_proxy--more_option--max_requests_per_connection"></a>

### max_requests_per_connection property

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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

- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/no_request_limit_per_connection/): complete subsection reference.

- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/): complete subsection reference.

<a id="schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_remove"></a>

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

- [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_headers_to_add/): complete subsection reference.

<a id="schema-dynamic_proxy--http_proxy--more_option--request_headers_to_remove"></a>

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

- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_cookies_to_add/): complete subsection reference.

<a id="schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_remove"></a>

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

- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_headers_to_add/): complete subsection reference.

<a id="schema-dynamic_proxy--http_proxy--more_option--response_headers_to_remove"></a>

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

## Next pages

- [dynamic_proxy.http_proxy.more_option.buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/buffer_policy/)
- [dynamic_proxy.http_proxy.more_option.compression_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/compression_params/)
- [dynamic_proxy.http_proxy.more_option.disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/disable_path_normalize/)
- [dynamic_proxy.http_proxy.more_option.enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/enable_path_normalize/)
- [dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/no_request_limit_per_connection/)
- [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_cookies_to_add/)
- [dynamic_proxy.http_proxy.more_option.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/request_headers_to_add/)
- [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_cookies_to_add/)
- [dynamic_proxy.http_proxy.more_option.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/more_option/response_headers_to_add/)
- [dynamic_proxy.http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/http_proxy/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
