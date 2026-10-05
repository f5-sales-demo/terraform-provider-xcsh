---
page_title: "more_option"
subcategory: "Load Balancing"
description: "This defines various OPTIONS to define a route."
xcsh_docs: {"aliases": ["more option"], "body_bytes": 16273, "body_sha256": "sha256:10e0e35957c76717cc12f52b9116bc231786d96b22bd8d67dc45f83cfab0d3d3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:more_option:buffer_policy", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:compression_params", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:disable_path_normalize", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:enable_path_normalize", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:no_request_limit_per_connection", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:request_cookies_to_add", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:request_headers_to_add", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_headers_to_add"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/more_option/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["more_option"], "schema_version": 1, "sections": [{"aliases": ["more option buffer policy"], "anchor": "section", "description": "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:buffer_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["more_option", "buffer_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option compression params"], "anchor": "section", "description": "Enables loadbalancer to compress dispatched data from an upstream service upon client request. The content is compressed and then sent to the client with the appropriate headers if either response and request allow. Only GZIP compression is supported. By default compression will be skipped when: A request does NOT", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:compression_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["more_option", "compression_params"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option custom errors"], "anchor": "schema-more_option--custom_errors", "description": "Map of integer error codes as keys and string values that can be used to provide custom HTTP pages for each error code. Key of the map can be either response code class or HTTP Error code. Response code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response code class 5 -- for", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "custom_errors"], "syntax": "attribute", "type": "map"}, {"aliases": ["more option disable default error pages"], "anchor": "schema-more_option--disable_default_error_pages", "description": "Disable the use of default F5XC error pages.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "disable_default_error_pages"], "syntax": "attribute", "type": "bool"}, {"aliases": ["more option disable path normalize"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:disable_path_normalize", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "disable_path_normalize"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option enable path normalize"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:enable_path_normalize", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "enable_path_normalize"], "syntax": "attribute", "type": "object"}, {"aliases": ["duration", "more option idle timeout"], "anchor": "schema-more_option--idle_timeout", "description": "The amount of time that a stream can exist without upstream or downstream activity, in milliseconds. The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header has been received, otherwise the stream is reset.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "idle_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["more option max request header size"], "anchor": "schema-more_option--max_request_header_size", "description": "The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers share the same advertise_policy, the highest value configured across all such load balancers is used for all the load", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "max_request_header_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["more option max requests per connection"], "anchor": "schema-more_option--max_requests_per_connection", "description": "Exclusive with Sets the maximum number of requests a downstream client can send over a single connection to Envoy. Enter a value >=1 to define the request limit per connection.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "max_requests_per_connection"], "syntax": "attribute", "type": "number"}, {"aliases": ["more option no request limit per connection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:no_request_limit_per_connection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "no_request_limit_per_connection"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option request cookies to add"], "anchor": "section", "description": "Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:request_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["more_option", "request_cookies_to_add"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option request cookies to remove"], "anchor": "schema-more_option--request_cookies_to_remove", "description": "List of keys of Cookies to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "request_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["more option request headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers specified at this level are applied after headers from matched Route are applied.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:request_headers_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["more_option", "request_headers_to_add"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option request headers to remove"], "anchor": "schema-more_option--request_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP request being sent towards upstream.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "request_headers_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["more option response cookies to add"], "anchor": "section", "description": "Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response being sent towards downstream. Cookies specified at this level are applied after cookies from matched Route are applied.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_cookies_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["more_option", "response_cookies_to_add"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option response cookies to remove"], "anchor": "schema-more_option--response_cookies_to_remove", "description": "List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire set-cookie header will be removed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_cookies_to_remove"], "syntax": "attribute", "type": "list"}, {"aliases": ["more option response headers to add"], "anchor": "section", "description": "Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers specified at this level are applied after headers from matched Route are applied.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option:response_headers_to_add", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["more_option", "response_headers_to_add"], "syntax": "attribute", "type": "object"}, {"aliases": ["more option response headers to remove"], "anchor": "schema-more_option--response_headers_to_remove", "description": "List of keys of Headers to be removed from the HTTP response being sent towards downstream.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:more_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["more_option", "response_headers_to_remove"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/more_option/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This defines various OPTIONS to define a route.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# more_option

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- more_option

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

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

## Direct properties

- [buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/buffer_policy/): complete subsection reference.

- [compression_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/compression_params/): complete subsection reference.

<a id="schema-more_option--custom_errors"></a>

### custom_errors property

Type: `["map", "string"]`. Computed.

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
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "ranges": [
        [
          3,
          3
        ],
        [
          4,
          4
        ],
        [
          5,
          5
        ],
        [
          300,
          599
        ]
      ],
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "65536",
      "ves.io.schema.rules.map.values.string.uri_ref": "true"
    },
    "values": {
      "format": "uri-reference",
      "maxLength": 65536,
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

<a id="schema-more_option--disable_default_error_pages"></a>

### disable_default_error_pages property

Type: `"bool"`. Computed.

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

- [disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/disable_path_normalize/): complete subsection reference.

- [enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/enable_path_normalize/): complete subsection reference.

<a id="schema-more_option--idle_timeout"></a>

### idle_timeout property

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="schema-more_option--max_request_header_size"></a>

### max_request_header_size property

Type: `"number"`. Computed.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="schema-more_option--max_requests_per_connection"></a>

### max_requests_per_connection property

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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

- [no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/no_request_limit_per_connection/): complete subsection reference.

- [request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/request_cookies_to_add/): complete subsection reference.

<a id="schema-more_option--request_cookies_to_remove"></a>

### request_cookies_to_remove property

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/request_headers_to_add/): complete subsection reference.

<a id="schema-more_option--request_headers_to_remove"></a>

### request_headers_to_remove property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/): complete subsection reference.

<a id="schema-more_option--response_cookies_to_remove"></a>

### response_cookies_to_remove property

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_headers_to_add/): complete subsection reference.

<a id="schema-more_option--response_headers_to_remove"></a>

### response_headers_to_remove property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [more_option.buffer_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/buffer_policy/)
- [more_option.compression_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/compression_params/)
- [more_option.disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/disable_path_normalize/)
- [more_option.enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/enable_path_normalize/)
- [more_option.no_request_limit_per_connection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/no_request_limit_per_connection/)
- [more_option.request_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/request_cookies_to_add/)
- [more_option.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/request_headers_to_add/)
- [more_option.response_cookies_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_cookies_to_add/)
- [more_option.response_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/more_option/response_headers_to_add/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
