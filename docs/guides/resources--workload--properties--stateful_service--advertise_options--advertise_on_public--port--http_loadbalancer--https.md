---
page_title: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 15585, "body_sha256": "sha256:ea5960067096b3cf9b420571d61d9ede43694b938533d44c730cbfa188a6c133", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:coalescing_options", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:default_header", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:default_loadbalancer", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:disable_path_normalize", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:enable_path_normalize", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:http_protocol_options", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:non_default_loadbalancer", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:pass_through", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_cert_params", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https:tls_parameters"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer:https", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_on_public:port:http_loadbalancer", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "https"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_on_public/port/http_loadbalancer/https/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--properties--stateful_service--advertise_options--advertise_on_public.md)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https

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
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--add_hsts"></a>

### add_hsts property

Type: `"bool"`. Optional.

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

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--append_server_name"></a>

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

- [coalescing_options](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--coalescing_options.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--connection_idle_timeout"></a>

### connection_idle_timeout property

Type: `"number"`. Optional.

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

- [default_header](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--default_header.md): complete subsection reference.

- [default_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--default_loadbalancer.md): complete subsection reference.

- [disable_path_normalize](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--disable_path_normalize.md): complete subsection reference.

- [enable_path_normalize](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--enable_path_normalize.md): complete subsection reference.

- [http_protocol_options](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_redirect"></a>

### http_redirect property

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

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

- [non_default_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--non_default_loadbalancer.md): complete subsection reference.

- [pass_through](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--pass_through.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--port"></a>

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

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--port_ranges"></a>

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

<a id="schema-stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--server_name"></a>

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

- [tls_cert_params](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params.md): complete subsection reference.

- [tls_parameters](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--coalescing_options.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--default_header.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--default_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--disable_path_normalize.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--enable_path_normalize.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--http_protocol_options.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--non_default_loadbalancer.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--pass_through.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_cert_params.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer--https--tls_parameters.md)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [xcsh_workload](../resources/workload.md)
