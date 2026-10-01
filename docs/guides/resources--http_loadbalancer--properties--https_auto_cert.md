---
page_title: "https_auto_cert"
subcategory: "Load Balancing"
description: "https_auto_cert for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 12592, "body_sha256": "sha256:613c4883ecf0c52e86e5de5c2b575911e5df87031b19cff5083562c5725f5381", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:coalescing_options", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:default_header", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:default_loadbalancer", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:disable_path_normalize", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:enable_path_normalize", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:http_protocol_options", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:no_mtls", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:non_default_loadbalancer", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:pass_through", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:tls_config", "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:use_mtls"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--https_auto_cert.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_auto_cert"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https_auto_cert/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_auto_cert for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- https_auto_cert

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-https_auto_cert--add_hsts"></a>

### add_hsts property

Type: `"bool"`. Optional, Computed.

Add HTTP Strict-Transport-Security response header. Defaults to \`false\`. Server applies default
when omitted.

Upstream description:

Add HTTP Strict-Transport-Security response header.

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

<a id="schema-https_auto_cert--append_server_name"></a>

### append_server_name property

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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

- [coalescing_options](resources--http_loadbalancer--properties--https_auto_cert--coalescing_options.md): complete subsection reference.

<a id="schema-https_auto_cert--connection_idle_timeout"></a>

### connection_idle_timeout property

Type: `"number"`. Optional, Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Server
applies default when omitted.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
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

- [default_header](resources--http_loadbalancer--properties--https_auto_cert--default_header.md): complete subsection reference.

- [default_loadbalancer](resources--http_loadbalancer--properties--https_auto_cert--default_loadbalancer.md): complete subsection reference.

- [disable_path_normalize](resources--http_loadbalancer--properties--https_auto_cert--disable_path_normalize.md): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--properties--https_auto_cert--enable_path_normalize.md): complete subsection reference.

- [http_protocol_options](resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options.md): complete subsection reference.

<a id="schema-https_auto_cert--http_redirect"></a>

### http_redirect property

Type: `"bool"`. Optional, Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS. Defaults to \`false\`. Server applies
default when omitted.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [no_mtls](resources--http_loadbalancer--properties--https_auto_cert--no_mtls.md): complete subsection reference.

- [non_default_loadbalancer](resources--http_loadbalancer--properties--https_auto_cert--non_default_loadbalancer.md): complete subsection reference.

- [pass_through](resources--http_loadbalancer--properties--https_auto_cert--pass_through.md): complete subsection reference.

<a id="schema-https_auto_cert--port"></a>

### port property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-https_auto_cert--port_ranges"></a>

### port_ranges property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="schema-https_auto_cert--server_name"></a>

### server_name property

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_config](resources--http_loadbalancer--properties--https_auto_cert--tls_config.md): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--properties--https_auto_cert--use_mtls.md): complete subsection reference.

## Next pages

- [https_auto_cert.coalescing_options](resources--http_loadbalancer--properties--https_auto_cert--coalescing_options.md)
- [https_auto_cert.default_header](resources--http_loadbalancer--properties--https_auto_cert--default_header.md)
- [https_auto_cert.default_loadbalancer](resources--http_loadbalancer--properties--https_auto_cert--default_loadbalancer.md)
- [https_auto_cert.disable_path_normalize](resources--http_loadbalancer--properties--https_auto_cert--disable_path_normalize.md)
- [https_auto_cert.enable_path_normalize](resources--http_loadbalancer--properties--https_auto_cert--enable_path_normalize.md)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--properties--https_auto_cert--http_protocol_options.md)
- [https_auto_cert.no_mtls](resources--http_loadbalancer--properties--https_auto_cert--no_mtls.md)
- [https_auto_cert.non_default_loadbalancer](resources--http_loadbalancer--properties--https_auto_cert--non_default_loadbalancer.md)
- [https_auto_cert.pass_through](resources--http_loadbalancer--properties--https_auto_cert--pass_through.md)
- [https_auto_cert.tls_config](resources--http_loadbalancer--properties--https_auto_cert--tls_config.md)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--properties--https_auto_cert--use_mtls.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
