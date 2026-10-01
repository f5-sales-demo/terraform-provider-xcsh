---
page_title: "http_health_check"
subcategory: "Monitoring"
description: "http_health_check for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 11919, "body_sha256": "sha256:a3e271d2f3d3f0b90a716ffe4c46cac8dfa513b7489989d51192216f3c1fd773", "child_ids": ["xcsh-docs:resources:healthcheck:properties:http_health_check:use_origin_server_name"], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:properties:http_health_check", "parent_id": "xcsh-docs:resources:healthcheck:reference", "path": "documentation/resources/healthcheck/properties/http_health_check/index.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["http_health_check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/properties/http_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "http_health_check for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_health_check

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/)
- http_health_check

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: http\_health\_check, tcp\_health\_check, udp\_icmp\_health\_check\] Healthy if 'GET' method
on URL 'HTTP(s)://&lt;host&gt;/&lt;path&gt;' with optional '&lt;header&gt;' returns success. 'host'
is not used for DNS resolution. It is used as HTTP Header in the request.

Upstream description:

Healthy if "GET" method on URL "HTTP(s)://&lt;host&gt;/&lt;path&gt;" with optional "&lt;header&gt;"
returns success. "host" is not used for DNS resolution. It is used as HTTP Header in the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path"),
  validators.ConflictingObjectAttributes("host_header",
    "use_origin_server_name")}
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
  "x-ves-oneof-field-host_header_choice": "[\"host_header\",\"use_origin_server_name\"]"
}
```

OneOf alternatives in this subsection:

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/#section)
- [tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/tcp_health_check/#section)
- [udp_icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/udp_icmp_health_check/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
http_health_check {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-http_health_check--expected_response"></a>

### expected_response property

Type: `"string"`. Optional, Computed.

Raw bytes expected in the response of HTTP health check. Input is to be given in Hex encoded format.
If left empty, then response body is not considered for evaluating health check status. Server
applies default when omitted.

Upstream description:

Raw bytes expected in the response of HTTP health check. Input is to be given in Hex encoded format.
If left empty, then response body is not considered for evaluating health check status.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="schema-http_health_check--expected_status_codes"></a>

### expected_status_codes property

Type: `["list", "string"]`. Optional, Computed.

Specifies a list of HTTP response status codes considered healthy. To treat default HTTP expected
status code 200 as healthy, user has to configure it explicitly. This is a list of strings, each of
which is single HTTP status code or a range with start and end values separated by '-'. Defaults to
\`\[\]\`. Server applies default when omitted.

Upstream description:

Specifies a list of HTTP response status codes considered healthy. To treat default HTTP expected
status code 200 as healthy, user has to configure it explicitly. This is a list of strings, each of
which is single HTTP status code or a range with start and end values separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.http_status_range": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "10",
    "ves.io.schema.rules.repeated.items.string.min_len": "3",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_status_range": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "10",
    "ves.io.schema.rules.repeated.items.string.min_len": "3",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-http_health_check--headers"></a>

### headers property

Type: `["map", "string"]`. Optional, Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked cluster. This is a list of key-value pairs. Defaults to \`map\[\]\`. Server applies default
when omitted.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked cluster. This is a list of key-value pairs.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-http_health_check--host_header"></a>

### host_header property

Type: `"string"`. Optional.

Exclusive with \[use\_origin\_server\_name\] The value of the host header.

Upstream description:

Exclusive with \[use\_origin\_server\_name\] The value of the host header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="schema-http_health_check--path"></a>

### path property

Type: `"string"`. Optional.

Specifies the HTTP path that will be requested during health checking. Recommended: \`/\`.

Upstream description:

Specifies the HTTP path that will be requested during health checking.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="schema-http_health_check--request_headers_to_remove"></a>

### request_headers_to_remove property

Type: `["list", "string"]`. Optional, Computed.

Specifies a list of HTTP headers that should be removed from each request that is sent to the health
checked cluster. This is a list of keys of headers. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

Specifies a list of HTTP headers that should be removed from each request that is sent to the health
checked cluster. This is a list of keys of headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="schema-http_health_check--use_http2"></a>

### use_http2 property

Type: `"bool"`. Optional, Computed.

If set, health checks will be made using HTTP/2. Defaults to \`false\`. Server applies default when
omitted. Recommended: \`false\`.

Upstream description:

If set, health checks will be made using HTTP/2.

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

- [use_origin_server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/use_origin_server_name/): complete subsection reference.

## Next pages

- [http_health_check.use_origin_server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/use_origin_server_name/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
