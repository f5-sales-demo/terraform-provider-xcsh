---
page_title: "proxy_config.https"
subcategory: ""
description: "proxy_config.https for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 10627, "body_sha256": "sha256:79076571a8a08a325a338c9fd33b23a8a91bf7daf034e38111acd8df8438504b", "canonical_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:coalescing_options", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:default_header", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:default_loadbalancer", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:disable_path_normalize", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:enable_path_normalize", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:non_default_loadbalancer", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:pass_through", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https:tls_parameters"], "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config:https", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config", "path": "docs/guides/data-sources--bigip_http_proxy--properties--proxy_config--https.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["proxy_config", "https"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_config/https/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "proxy_config.https for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# proxy_config.https

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
- [Property reference](data-sources--bigip_http_proxy--reference.md)
- [proxy_config](data-sources--bigip_http_proxy--properties--proxy_config.md)
- proxy_config.https

<a id="section"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

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

## Direct properties

<a id="schema-proxy_config--https--add_hsts"></a>

### add_hsts property

Type: `"bool"`. Computed.

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

<a id="schema-proxy_config--https--append_server_name"></a>

### append_server_name property

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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

- [coalescing_options](data-sources--bigip_http_proxy--properties--proxy_config--https--coalescing_options.md): complete subsection reference.

<a id="schema-proxy_config--https--connection_idle_timeout"></a>

### connection_idle_timeout property

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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

- [default_header](data-sources--bigip_http_proxy--properties--proxy_config--https--default_header.md): complete subsection reference.

- [default_loadbalancer](data-sources--bigip_http_proxy--properties--proxy_config--https--default_loadbalancer.md): complete subsection reference.

- [disable_path_normalize](data-sources--bigip_http_proxy--properties--proxy_config--https--disable_path_normalize.md): complete subsection reference.

- [enable_path_normalize](data-sources--bigip_http_proxy--properties--proxy_config--https--enable_path_normalize.md): complete subsection reference.

- [http_protocol_options](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options.md): complete subsection reference.

<a id="schema-proxy_config--https--http_redirect"></a>

### http_redirect property

Type: `"bool"`. Computed.

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

- [non_default_loadbalancer](data-sources--bigip_http_proxy--properties--proxy_config--https--non_default_loadbalancer.md): complete subsection reference.

- [pass_through](data-sources--bigip_http_proxy--properties--proxy_config--https--pass_through.md): complete subsection reference.

<a id="schema-proxy_config--https--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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

<a id="schema-proxy_config--https--port_ranges"></a>

### port_ranges property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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

<a id="schema-proxy_config--https--server_name"></a>

### server_name property

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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

- [tls_cert_params](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md): complete subsection reference.

- [tls_parameters](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters.md): complete subsection reference.

## Next pages

- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--properties--proxy_config--https--coalescing_options.md)
- [proxy_config.https.default_header](data-sources--bigip_http_proxy--properties--proxy_config--https--default_header.md)
- [proxy_config.https.default_loadbalancer](data-sources--bigip_http_proxy--properties--proxy_config--https--default_loadbalancer.md)
- [proxy_config.https.disable_path_normalize](data-sources--bigip_http_proxy--properties--proxy_config--https--disable_path_normalize.md)
- [proxy_config.https.enable_path_normalize](data-sources--bigip_http_proxy--properties--proxy_config--https--enable_path_normalize.md)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options.md)
- [proxy_config.https.non_default_loadbalancer](data-sources--bigip_http_proxy--properties--proxy_config--https--non_default_loadbalancer.md)
- [proxy_config.https.pass_through](data-sources--bigip_http_proxy--properties--proxy_config--https--pass_through.md)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--properties--proxy_config--https--tls_parameters.md)
- [proxy_config](data-sources--bigip_http_proxy--properties--proxy_config.md)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md)
