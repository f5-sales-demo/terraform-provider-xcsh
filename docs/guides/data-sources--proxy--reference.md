---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 120097, "body_sha256": "sha256:8b9c726ce50bf16d5c9b001c3ce8849d1e15ea904443c2626360cfec4c622437", "canonical_id": "xcsh-docs:data-sources:proxy:reference", "child_ids": ["xcsh-docs:data-sources:proxy:properties:active_forward_proxy_policies", "xcsh-docs:data-sources:proxy:properties:do_not_advertise", "xcsh-docs:data-sources:proxy:properties:dynamic_proxy", "xcsh-docs:data-sources:proxy:properties:http_proxy", "xcsh-docs:data-sources:proxy:properties:no_forward_proxy_policy", "xcsh-docs:data-sources:proxy:properties:no_interception", "xcsh-docs:data-sources:proxy:properties:site_local_inside_network", "xcsh-docs:data-sources:proxy:properties:site_local_network", "xcsh-docs:data-sources:proxy:properties:site_virtual_sites", "xcsh-docs:data-sources:proxy:properties:tls_intercept"], "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:reference", "parent_id": "xcsh-docs:data-sources:proxy:fundamentals", "path": "docs/guides/data-sources--proxy--reference.md", "provider_name": "proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md)
- Property reference

## Direct properties

- [active_forward_proxy_policies](data-sources--proxy--properties--active_forward_proxy_policies.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

<a id="schema-connection_timeout"></a>

### connection_timeout property

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Upstream description:

The timeout for new network connections to upstream server. This is specified in milliseconds. The
default value is 2000 (2 seconds)

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

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the Proxy.

Upstream description:

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

- [do_not_advertise](data-sources--proxy--properties--do_not_advertise.md): complete subsection reference.

- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md): complete subsection reference.

- [http_proxy](data-sources--proxy--properties--http_proxy.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Proxy.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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

Namespace where the Proxy exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

- [no_forward_proxy_policy](data-sources--proxy--properties--no_forward_proxy_policy.md): complete subsection reference.

- [no_interception](data-sources--proxy--properties--no_interception.md): complete subsection reference.

- [site_local_inside_network](data-sources--proxy--properties--site_local_inside_network.md): complete subsection reference.

- [site_local_network](data-sources--proxy--properties--site_local_network.md): complete subsection reference.

- [site_virtual_sites](data-sources--proxy--properties--site_virtual_sites.md): complete subsection reference.

- [tls_intercept](data-sources--proxy--properties--tls_intercept.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_forward_proxy_policies` | [active_forward_proxy_policies](data-sources--proxy--properties--active_forward_proxy_policies.md#section) |
| `active_forward_proxy_policies.forward_proxy_policies` | [active_forward_proxy_policies.forward_proxy_policies](data-sources--proxy--properties--active_forward_proxy_policies--forward_proxy_policies.md#section) |
| `active_forward_proxy_policies.forward_proxy_policies.name` | [active_forward_proxy_policies.forward_proxy_policies.name](data-sources--proxy--properties--active_forward_proxy_policies--forward_proxy_policies.md#schema-active_forward_proxy_policies--forward_proxy_policies--name) |
| `active_forward_proxy_policies.forward_proxy_policies.namespace` | [active_forward_proxy_policies.forward_proxy_policies.namespace](data-sources--proxy--properties--active_forward_proxy_policies--forward_proxy_policies.md#schema-active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `active_forward_proxy_policies.forward_proxy_policies.tenant` | [active_forward_proxy_policies.forward_proxy_policies.tenant](data-sources--proxy--properties--active_forward_proxy_policies--forward_proxy_policies.md#schema-active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `annotations` | [annotations](data-sources--proxy--reference.md#schema-annotations) |
| `connection_timeout` | [connection_timeout](data-sources--proxy--reference.md#schema-connection_timeout) |
| `description` | [description](data-sources--proxy--reference.md#schema-description) |
| `do_not_advertise` | [do_not_advertise](data-sources--proxy--properties--do_not_advertise.md#section) |
| `dynamic_proxy` | [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md#section) |
| `dynamic_proxy.disable_dns_masquerade` | [dynamic_proxy.disable_dns_masquerade](data-sources--proxy--properties--dynamic_proxy--disable_dns_masquerade.md#section) |
| `dynamic_proxy.domains` | [dynamic_proxy.domains](data-sources--proxy--properties--dynamic_proxy.md#schema-dynamic_proxy--domains) |
| `dynamic_proxy.enable_dns_masquerade` | [dynamic_proxy.enable_dns_masquerade](data-sources--proxy--properties--dynamic_proxy--enable_dns_masquerade.md#section) |
| `dynamic_proxy.http_proxy` | [dynamic_proxy.http_proxy](data-sources--proxy--properties--dynamic_proxy--http_proxy.md#section) |
| `dynamic_proxy.http_proxy.more_option` | [dynamic_proxy.http_proxy.more_option](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#section) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy` | [dynamic_proxy.http_proxy.more_option.buffer_policy](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--buffer_policy.md#section) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.http_proxy.more_option.buffer_policy.disabled](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--buffer_policy.md#schema-dynamic_proxy--http_proxy--more_option--buffer_policy--disabled) |
| `dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.http_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--buffer_policy.md#schema-dynamic_proxy--http_proxy--more_option--buffer_policy--max_request_bytes) |
| `dynamic_proxy.http_proxy.more_option.compression_params` | [dynamic_proxy.http_proxy.more_option.compression_params](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--compression_params.md#section) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_length` | [dynamic_proxy.http_proxy.more_option.compression_params.content_length](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--compression_params.md#schema-dynamic_proxy--http_proxy--more_option--compression_params--content_length) |
| `dynamic_proxy.http_proxy.more_option.compression_params.content_type` | [dynamic_proxy.http_proxy.more_option.compression_params.content_type](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--compression_params.md#schema-dynamic_proxy--http_proxy--more_option--compression_params--content_type) |
| `dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.http_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--compression_params.md#schema-dynamic_proxy--http_proxy--more_option--compression_params--disable_on_etag_header) |
| `dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.http_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--compression_params.md#schema-dynamic_proxy--http_proxy--more_option--compression_params--remove_accept_encoding_header) |
| `dynamic_proxy.http_proxy.more_option.custom_errors` | [dynamic_proxy.http_proxy.more_option.custom_errors](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--custom_errors) |
| `dynamic_proxy.http_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.http_proxy.more_option.disable_default_error_pages](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--disable_default_error_pages) |
| `dynamic_proxy.http_proxy.more_option.disable_path_normalize` | [dynamic_proxy.http_proxy.more_option.disable_path_normalize](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--disable_path_normalize.md#section) |
| `dynamic_proxy.http_proxy.more_option.enable_path_normalize` | [dynamic_proxy.http_proxy.more_option.enable_path_normalize](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--enable_path_normalize.md#section) |
| `dynamic_proxy.http_proxy.more_option.idle_timeout` | [dynamic_proxy.http_proxy.more_option.idle_timeout](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--idle_timeout) |
| `dynamic_proxy.http_proxy.more_option.max_request_header_size` | [dynamic_proxy.http_proxy.more_option.max_request_header_size](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--max_request_header_size) |
| `dynamic_proxy.http_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.http_proxy.more_option.max_requests_per_connection](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--max_requests_per_connection) |
| `dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.http_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--no_request_limit_per_connection.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--name) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--overwrite) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info--url) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_add--value) |
| `dynamic_proxy.http_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.request_cookies_to_remove](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--request_cookies_to_remove) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.append](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--append) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.name](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--name) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--location) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#section) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info--url) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.request_headers_to_add.value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--request_headers_to_add.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_add--value) |
| `dynamic_proxy.http_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.request_headers_to_remove](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--request_headers_to_remove) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_domain) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_expiry) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_httponly.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_partitioned.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_path) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--add_secure.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_domain.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_expiry.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_httponly.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_max_age.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_partitioned.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_path.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_samesite.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_secure.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--ignore_value.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--max_age_value) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--name) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--overwrite) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--samesite_lax.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--samesite_none.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--samesite_strict.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info--url) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_add--value) |
| `dynamic_proxy.http_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.http_proxy.more_option.response_cookies_to_remove](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--response_cookies_to_remove) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.append](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--append) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.name](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--name) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--location) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#section) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info--url) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.http_proxy.more_option.response_headers_to_add.value](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option--response_headers_to_add.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_add--value) |
| `dynamic_proxy.http_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.http_proxy.more_option.response_headers_to_remove](data-sources--proxy--properties--dynamic_proxy--http_proxy--more_option.md#schema-dynamic_proxy--http_proxy--more_option--response_headers_to_remove) |
| `dynamic_proxy.https_proxy` | [dynamic_proxy.https_proxy](data-sources--proxy--properties--dynamic_proxy--https_proxy.md#section) |
| `dynamic_proxy.https_proxy.more_option` | [dynamic_proxy.https_proxy.more_option](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#section) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy` | [dynamic_proxy.https_proxy.more_option.buffer_policy](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--buffer_policy.md#section) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.disabled` | [dynamic_proxy.https_proxy.more_option.buffer_policy.disabled](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--buffer_policy.md#schema-dynamic_proxy--https_proxy--more_option--buffer_policy--disabled) |
| `dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes` | [dynamic_proxy.https_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--buffer_policy.md#schema-dynamic_proxy--https_proxy--more_option--buffer_policy--max_request_bytes) |
| `dynamic_proxy.https_proxy.more_option.compression_params` | [dynamic_proxy.https_proxy.more_option.compression_params](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--compression_params.md#section) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_length` | [dynamic_proxy.https_proxy.more_option.compression_params.content_length](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--compression_params.md#schema-dynamic_proxy--https_proxy--more_option--compression_params--content_length) |
| `dynamic_proxy.https_proxy.more_option.compression_params.content_type` | [dynamic_proxy.https_proxy.more_option.compression_params.content_type](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--compression_params.md#schema-dynamic_proxy--https_proxy--more_option--compression_params--content_type) |
| `dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header` | [dynamic_proxy.https_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--compression_params.md#schema-dynamic_proxy--https_proxy--more_option--compression_params--disable_on_etag_header) |
| `dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header` | [dynamic_proxy.https_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--compression_params.md#schema-dynamic_proxy--https_proxy--more_option--compression_params--remove_accept_encoding_header) |
| `dynamic_proxy.https_proxy.more_option.custom_errors` | [dynamic_proxy.https_proxy.more_option.custom_errors](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--custom_errors) |
| `dynamic_proxy.https_proxy.more_option.disable_default_error_pages` | [dynamic_proxy.https_proxy.more_option.disable_default_error_pages](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--disable_default_error_pages) |
| `dynamic_proxy.https_proxy.more_option.disable_path_normalize` | [dynamic_proxy.https_proxy.more_option.disable_path_normalize](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--disable_path_normalize.md#section) |
| `dynamic_proxy.https_proxy.more_option.enable_path_normalize` | [dynamic_proxy.https_proxy.more_option.enable_path_normalize](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--enable_path_normalize.md#section) |
| `dynamic_proxy.https_proxy.more_option.idle_timeout` | [dynamic_proxy.https_proxy.more_option.idle_timeout](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--idle_timeout) |
| `dynamic_proxy.https_proxy.more_option.max_request_header_size` | [dynamic_proxy.https_proxy.more_option.max_request_header_size](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--max_request_header_size) |
| `dynamic_proxy.https_proxy.more_option.max_requests_per_connection` | [dynamic_proxy.https_proxy.more_option.max_requests_per_connection](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--max_requests_per_connection) |
| `dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection` | [dynamic_proxy.https_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--no_request_limit_per_connection.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_add--name) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_add--overwrite) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info--url) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_add--value) |
| `dynamic_proxy.https_proxy.more_option.request_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.request_cookies_to_remove](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--request_cookies_to_remove) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.append](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_add--append) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.name](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_add--name) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--location) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#section) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info--url) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.request_headers_to_add.value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--request_headers_to_add.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_add--value) |
| `dynamic_proxy.https_proxy.more_option.request_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.request_headers_to_remove](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--request_headers_to_remove) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--add_domain) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--add_expiry) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--add_httponly.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--add_partitioned.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--add_path) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--add_secure.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_domain.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_expiry.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_httponly.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_max_age.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_partitioned.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_path.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_samesite.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_secure.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--ignore_value.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--max_age_value) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--name) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--overwrite) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--samesite_lax.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--samesite_none.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--samesite_strict.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info--url) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_cookies_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_add--value) |
| `dynamic_proxy.https_proxy.more_option.response_cookies_to_remove` | [dynamic_proxy.https_proxy.more_option.response_cookies_to_remove](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--response_cookies_to_remove) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.append` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.append](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_add--append) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.name` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.name](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_add--name) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--location) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#section) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info--url) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_add.value` | [dynamic_proxy.https_proxy.more_option.response_headers_to_add.value](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option--response_headers_to_add.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_add--value) |
| `dynamic_proxy.https_proxy.more_option.response_headers_to_remove` | [dynamic_proxy.https_proxy.more_option.response_headers_to_remove](data-sources--proxy--properties--dynamic_proxy--https_proxy--more_option.md#schema-dynamic_proxy--https_proxy--more_option--response_headers_to_remove) |
| `dynamic_proxy.https_proxy.tls_params` | [dynamic_proxy.https_proxy.tls_params](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params.md#section) |
| `dynamic_proxy.https_proxy.tls_params.no_mtls` | [dynamic_proxy.https_proxy.tls_params.no_mtls](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--no_mtls.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates` | [dynamic_proxy.https_proxy.tls_params.tls_certificates](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.certificate_url](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md#schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--certificate_url) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--custom_hash_algorithms.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--custom_hash_algorithms.md#schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.description_spec](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates.md#schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--description_spec) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--disable_ocsp_stapling.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info--location) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info.md#schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info.md#schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info.url](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info.md#schema-dynamic_proxy--https_proxy--tls_params--tls_certificates--private_key--clear_secret_info--url) |
| `dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults` | [dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_certificates--use_system_defaults.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_config` | [dynamic_proxy.https_proxy.tls_params.tls_config](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--custom_security.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.cipher_suites](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--custom_security.md#schema-dynamic_proxy--https_proxy--tls_params--tls_config--custom_security--cipher_suites) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.max_version](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--custom_security.md#schema-dynamic_proxy--https_proxy--tls_params--tls_config--custom_security--max_version) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version` | [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security.min_version](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--custom_security.md#schema-dynamic_proxy--https_proxy--tls_params--tls_config--custom_security--min_version) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.default_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--default_security.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.low_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--low_security.md#section) |
| `dynamic_proxy.https_proxy.tls_params.tls_config.medium_security` | [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--tls_config--medium_security.md#section) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls` | [dynamic_proxy.https_proxy.tls_params.use_mtls](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md#section) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional` | [dynamic_proxy.https_proxy.tls_params.use_mtls.client_certificate_optional](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--client_certificate_optional) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--crl.md#section) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.name](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--crl.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--crl--name) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.namespace](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--crl.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--crl--namespace) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.crl.tenant](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--crl.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--crl--tenant) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl` | [dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--no_crl.md#section) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca.md#section) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.name](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--name) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.namespace](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--namespace) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca.tenant](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca--tenant) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url` | [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca_url](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--trusted_ca_url) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--xfcc_disabled.md#section) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--xfcc_options.md#section) |
| `dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements` | [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options.xfcc_header_elements](data-sources--proxy--properties--dynamic_proxy--https_proxy--tls_params--use_mtls--xfcc_options.md#schema-dynamic_proxy--https_proxy--tls_params--use_mtls--xfcc_options--xfcc_header_elements) |
| `dynamic_proxy.sni_proxy` | [dynamic_proxy.sni_proxy](data-sources--proxy--properties--dynamic_proxy--sni_proxy.md#section) |
| `dynamic_proxy.sni_proxy.idle_timeout` | [dynamic_proxy.sni_proxy.idle_timeout](data-sources--proxy--properties--dynamic_proxy--sni_proxy.md#schema-dynamic_proxy--sni_proxy--idle_timeout) |
| `http_proxy` | [http_proxy](data-sources--proxy--properties--http_proxy.md#section) |
| `http_proxy.enable_http` | [http_proxy.enable_http](data-sources--proxy--properties--http_proxy--enable_http.md#section) |
| `http_proxy.more_option` | [http_proxy.more_option](data-sources--proxy--properties--http_proxy--more_option.md#section) |
| `http_proxy.more_option.buffer_policy` | [http_proxy.more_option.buffer_policy](data-sources--proxy--properties--http_proxy--more_option--buffer_policy.md#section) |
| `http_proxy.more_option.buffer_policy.disabled` | [http_proxy.more_option.buffer_policy.disabled](data-sources--proxy--properties--http_proxy--more_option--buffer_policy.md#schema-http_proxy--more_option--buffer_policy--disabled) |
| `http_proxy.more_option.buffer_policy.max_request_bytes` | [http_proxy.more_option.buffer_policy.max_request_bytes](data-sources--proxy--properties--http_proxy--more_option--buffer_policy.md#schema-http_proxy--more_option--buffer_policy--max_request_bytes) |
| `http_proxy.more_option.compression_params` | [http_proxy.more_option.compression_params](data-sources--proxy--properties--http_proxy--more_option--compression_params.md#section) |
| `http_proxy.more_option.compression_params.content_length` | [http_proxy.more_option.compression_params.content_length](data-sources--proxy--properties--http_proxy--more_option--compression_params.md#schema-http_proxy--more_option--compression_params--content_length) |
| `http_proxy.more_option.compression_params.content_type` | [http_proxy.more_option.compression_params.content_type](data-sources--proxy--properties--http_proxy--more_option--compression_params.md#schema-http_proxy--more_option--compression_params--content_type) |
| `http_proxy.more_option.compression_params.disable_on_etag_header` | [http_proxy.more_option.compression_params.disable_on_etag_header](data-sources--proxy--properties--http_proxy--more_option--compression_params.md#schema-http_proxy--more_option--compression_params--disable_on_etag_header) |
| `http_proxy.more_option.compression_params.remove_accept_encoding_header` | [http_proxy.more_option.compression_params.remove_accept_encoding_header](data-sources--proxy--properties--http_proxy--more_option--compression_params.md#schema-http_proxy--more_option--compression_params--remove_accept_encoding_header) |
| `http_proxy.more_option.custom_errors` | [http_proxy.more_option.custom_errors](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--custom_errors) |
| `http_proxy.more_option.disable_default_error_pages` | [http_proxy.more_option.disable_default_error_pages](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--disable_default_error_pages) |
| `http_proxy.more_option.disable_path_normalize` | [http_proxy.more_option.disable_path_normalize](data-sources--proxy--properties--http_proxy--more_option--disable_path_normalize.md#section) |
| `http_proxy.more_option.enable_path_normalize` | [http_proxy.more_option.enable_path_normalize](data-sources--proxy--properties--http_proxy--more_option--enable_path_normalize.md#section) |
| `http_proxy.more_option.idle_timeout` | [http_proxy.more_option.idle_timeout](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--idle_timeout) |
| `http_proxy.more_option.max_request_header_size` | [http_proxy.more_option.max_request_header_size](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--max_request_header_size) |
| `http_proxy.more_option.max_requests_per_connection` | [http_proxy.more_option.max_requests_per_connection](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--max_requests_per_connection) |
| `http_proxy.more_option.no_request_limit_per_connection` | [http_proxy.more_option.no_request_limit_per_connection](data-sources--proxy--properties--http_proxy--more_option--no_request_limit_per_connection.md#section) |
| `http_proxy.more_option.request_cookies_to_add` | [http_proxy.more_option.request_cookies_to_add](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add.md#section) |
| `http_proxy.more_option.request_cookies_to_add.name` | [http_proxy.more_option.request_cookies_to_add.name](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add.md#schema-http_proxy--more_option--request_cookies_to_add--name) |
| `http_proxy.more_option.request_cookies_to_add.overwrite` | [http_proxy.more_option.request_cookies_to_add.overwrite](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add.md#schema-http_proxy--more_option--request_cookies_to_add--overwrite) |
| `http_proxy.more_option.request_cookies_to_add.secret_value` | [http_proxy.more_option.request_cookies_to_add.secret_value](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value.md#section) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#section) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--request_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#section) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#schema-http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info.md#schema-http_proxy--more_option--request_cookies_to_add--secret_value--clear_secret_info--url) |
| `http_proxy.more_option.request_cookies_to_add.value` | [http_proxy.more_option.request_cookies_to_add.value](data-sources--proxy--properties--http_proxy--more_option--request_cookies_to_add.md#schema-http_proxy--more_option--request_cookies_to_add--value) |
| `http_proxy.more_option.request_cookies_to_remove` | [http_proxy.more_option.request_cookies_to_remove](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--request_cookies_to_remove) |
| `http_proxy.more_option.request_headers_to_add` | [http_proxy.more_option.request_headers_to_add](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add.md#section) |
| `http_proxy.more_option.request_headers_to_add.append` | [http_proxy.more_option.request_headers_to_add.append](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add.md#schema-http_proxy--more_option--request_headers_to_add--append) |
| `http_proxy.more_option.request_headers_to_add.name` | [http_proxy.more_option.request_headers_to_add.name](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add.md#schema-http_proxy--more_option--request_headers_to_add--name) |
| `http_proxy.more_option.request_headers_to_add.secret_value` | [http_proxy.more_option.request_headers_to_add.secret_value](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add--secret_value.md#section) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#section) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--location) |
| `http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--request_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#section) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#schema-http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info.md#schema-http_proxy--more_option--request_headers_to_add--secret_value--clear_secret_info--url) |
| `http_proxy.more_option.request_headers_to_add.value` | [http_proxy.more_option.request_headers_to_add.value](data-sources--proxy--properties--http_proxy--more_option--request_headers_to_add.md#schema-http_proxy--more_option--request_headers_to_add--value) |
| `http_proxy.more_option.request_headers_to_remove` | [http_proxy.more_option.request_headers_to_remove](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--request_headers_to_remove) |
| `http_proxy.more_option.response_cookies_to_add` | [http_proxy.more_option.response_cookies_to_add](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md#section) |
| `http_proxy.more_option.response_cookies_to_add.add_domain` | [http_proxy.more_option.response_cookies_to_add.add_domain](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md#schema-http_proxy--more_option--response_cookies_to_add--add_domain) |
| `http_proxy.more_option.response_cookies_to_add.add_expiry` | [http_proxy.more_option.response_cookies_to_add.add_expiry](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md#schema-http_proxy--more_option--response_cookies_to_add--add_expiry) |
| `http_proxy.more_option.response_cookies_to_add.add_httponly` | [http_proxy.more_option.response_cookies_to_add.add_httponly](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--add_httponly.md#section) |
| `http_proxy.more_option.response_cookies_to_add.add_partitioned` | [http_proxy.more_option.response_cookies_to_add.add_partitioned](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--add_partitioned.md#section) |
| `http_proxy.more_option.response_cookies_to_add.add_path` | [http_proxy.more_option.response_cookies_to_add.add_path](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md#schema-http_proxy--more_option--response_cookies_to_add--add_path) |
| `http_proxy.more_option.response_cookies_to_add.add_secure` | [http_proxy.more_option.response_cookies_to_add.add_secure](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--add_secure.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_domain` | [http_proxy.more_option.response_cookies_to_add.ignore_domain](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_domain.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_expiry` | [http_proxy.more_option.response_cookies_to_add.ignore_expiry](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_expiry.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_httponly` | [http_proxy.more_option.response_cookies_to_add.ignore_httponly](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_httponly.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_max_age` | [http_proxy.more_option.response_cookies_to_add.ignore_max_age](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_max_age.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_partitioned` | [http_proxy.more_option.response_cookies_to_add.ignore_partitioned](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_partitioned.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_path` | [http_proxy.more_option.response_cookies_to_add.ignore_path](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_path.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_samesite` | [http_proxy.more_option.response_cookies_to_add.ignore_samesite](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_samesite.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_secure` | [http_proxy.more_option.response_cookies_to_add.ignore_secure](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_secure.md#section) |
| `http_proxy.more_option.response_cookies_to_add.ignore_value` | [http_proxy.more_option.response_cookies_to_add.ignore_value](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--ignore_value.md#section) |
| `http_proxy.more_option.response_cookies_to_add.max_age_value` | [http_proxy.more_option.response_cookies_to_add.max_age_value](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md#schema-http_proxy--more_option--response_cookies_to_add--max_age_value) |
| `http_proxy.more_option.response_cookies_to_add.name` | [http_proxy.more_option.response_cookies_to_add.name](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md#schema-http_proxy--more_option--response_cookies_to_add--name) |
| `http_proxy.more_option.response_cookies_to_add.overwrite` | [http_proxy.more_option.response_cookies_to_add.overwrite](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md#schema-http_proxy--more_option--response_cookies_to_add--overwrite) |
| `http_proxy.more_option.response_cookies_to_add.samesite_lax` | [http_proxy.more_option.response_cookies_to_add.samesite_lax](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--samesite_lax.md#section) |
| `http_proxy.more_option.response_cookies_to_add.samesite_none` | [http_proxy.more_option.response_cookies_to_add.samesite_none](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--samesite_none.md#section) |
| `http_proxy.more_option.response_cookies_to_add.samesite_strict` | [http_proxy.more_option.response_cookies_to_add.samesite_strict](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--samesite_strict.md#section) |
| `http_proxy.more_option.response_cookies_to_add.secret_value` | [http_proxy.more_option.response_cookies_to_add.secret_value](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value.md#section) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#section) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--location) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--response_cookies_to_add--secret_value--blindfold_secret_info--store_provider) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#section) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#schema-http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info--provider_ref) |
| `http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info.md#schema-http_proxy--more_option--response_cookies_to_add--secret_value--clear_secret_info--url) |
| `http_proxy.more_option.response_cookies_to_add.value` | [http_proxy.more_option.response_cookies_to_add.value](data-sources--proxy--properties--http_proxy--more_option--response_cookies_to_add.md#schema-http_proxy--more_option--response_cookies_to_add--value) |
| `http_proxy.more_option.response_cookies_to_remove` | [http_proxy.more_option.response_cookies_to_remove](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--response_cookies_to_remove) |
| `http_proxy.more_option.response_headers_to_add` | [http_proxy.more_option.response_headers_to_add](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add.md#section) |
| `http_proxy.more_option.response_headers_to_add.append` | [http_proxy.more_option.response_headers_to_add.append](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add.md#schema-http_proxy--more_option--response_headers_to_add--append) |
| `http_proxy.more_option.response_headers_to_add.name` | [http_proxy.more_option.response_headers_to_add.name](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add.md#schema-http_proxy--more_option--response_headers_to_add--name) |
| `http_proxy.more_option.response_headers_to_add.secret_value` | [http_proxy.more_option.response_headers_to_add.secret_value](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add--secret_value.md#section) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#section) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--decryption_provider) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.location](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--location) |
| `http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [http_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info.md#schema-http_proxy--more_option--response_headers_to_add--secret_value--blindfold_secret_info--store_provider) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#section) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.provider_ref](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#schema-http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info--provider_ref) |
| `http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url` | [http_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info.url](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info.md#schema-http_proxy--more_option--response_headers_to_add--secret_value--clear_secret_info--url) |
| `http_proxy.more_option.response_headers_to_add.value` | [http_proxy.more_option.response_headers_to_add.value](data-sources--proxy--properties--http_proxy--more_option--response_headers_to_add.md#schema-http_proxy--more_option--response_headers_to_add--value) |
| `http_proxy.more_option.response_headers_to_remove` | [http_proxy.more_option.response_headers_to_remove](data-sources--proxy--properties--http_proxy--more_option.md#schema-http_proxy--more_option--response_headers_to_remove) |
| `id` | [id](data-sources--proxy--reference.md#schema-id) |
| `labels` | [labels](data-sources--proxy--reference.md#schema-labels) |
| `name` | [name](data-sources--proxy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--proxy--reference.md#schema-namespace) |
| `no_forward_proxy_policy` | [no_forward_proxy_policy](data-sources--proxy--properties--no_forward_proxy_policy.md#section) |
| `no_interception` | [no_interception](data-sources--proxy--properties--no_interception.md#section) |
| `site_local_inside_network` | [site_local_inside_network](data-sources--proxy--properties--site_local_inside_network.md#section) |
| `site_local_network` | [site_local_network](data-sources--proxy--properties--site_local_network.md#section) |
| `site_virtual_sites` | [site_virtual_sites](data-sources--proxy--properties--site_virtual_sites.md#section) |
| `site_virtual_sites.advertise_where` | [site_virtual_sites.advertise_where](data-sources--proxy--properties--site_virtual_sites--advertise_where.md#section) |
| `site_virtual_sites.advertise_where.port` | [site_virtual_sites.advertise_where.port](data-sources--proxy--properties--site_virtual_sites--advertise_where.md#schema-site_virtual_sites--advertise_where--port) |
| `site_virtual_sites.advertise_where.site` | [site_virtual_sites.advertise_where.site](data-sources--proxy--properties--site_virtual_sites--advertise_where--site.md#section) |
| `site_virtual_sites.advertise_where.site.ip` | [site_virtual_sites.advertise_where.site.ip](data-sources--proxy--properties--site_virtual_sites--advertise_where--site.md#schema-site_virtual_sites--advertise_where--site--ip) |
| `site_virtual_sites.advertise_where.site.network` | [site_virtual_sites.advertise_where.site.network](data-sources--proxy--properties--site_virtual_sites--advertise_where--site.md#schema-site_virtual_sites--advertise_where--site--network) |
| `site_virtual_sites.advertise_where.site.site` | [site_virtual_sites.advertise_where.site.site](data-sources--proxy--properties--site_virtual_sites--advertise_where--site--site.md#section) |
| `site_virtual_sites.advertise_where.site.site.name` | [site_virtual_sites.advertise_where.site.site.name](data-sources--proxy--properties--site_virtual_sites--advertise_where--site--site.md#schema-site_virtual_sites--advertise_where--site--site--name) |
| `site_virtual_sites.advertise_where.site.site.namespace` | [site_virtual_sites.advertise_where.site.site.namespace](data-sources--proxy--properties--site_virtual_sites--advertise_where--site--site.md#schema-site_virtual_sites--advertise_where--site--site--namespace) |
| `site_virtual_sites.advertise_where.site.site.tenant` | [site_virtual_sites.advertise_where.site.site.tenant](data-sources--proxy--properties--site_virtual_sites--advertise_where--site--site.md#schema-site_virtual_sites--advertise_where--site--site--tenant) |
| `site_virtual_sites.advertise_where.use_default_port` | [site_virtual_sites.advertise_where.use_default_port](data-sources--proxy--properties--site_virtual_sites--advertise_where--use_default_port.md#section) |
| `site_virtual_sites.advertise_where.virtual_site` | [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site.md#section) |
| `site_virtual_sites.advertise_where.virtual_site.network` | [site_virtual_sites.advertise_where.virtual_site.network](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site.md#schema-site_virtual_sites--advertise_where--virtual_site--network) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site` | [site_virtual_sites.advertise_where.virtual_site.virtual_site](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site--virtual_site.md#section) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.name` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.name](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site--virtual_site.md#schema-site_virtual_sites--advertise_where--virtual_site--virtual_site--name) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.namespace](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site--virtual_site.md#schema-site_virtual_sites--advertise_where--virtual_site--virtual_site--namespace) |
| `site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant` | [site_virtual_sites.advertise_where.virtual_site.virtual_site.tenant](data-sources--proxy--properties--site_virtual_sites--advertise_where--virtual_site--virtual_site.md#schema-site_virtual_sites--advertise_where--virtual_site--virtual_site--tenant) |
| `tls_intercept` | [tls_intercept](data-sources--proxy--properties--tls_intercept.md#section) |
| `tls_intercept.custom_certificate` | [tls_intercept.custom_certificate](data-sources--proxy--properties--tls_intercept--custom_certificate.md#section) |
| `tls_intercept.custom_certificate.certificate_url` | [tls_intercept.custom_certificate.certificate_url](data-sources--proxy--properties--tls_intercept--custom_certificate.md#schema-tls_intercept--custom_certificate--certificate_url) |
| `tls_intercept.custom_certificate.custom_hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--proxy--properties--tls_intercept--custom_certificate--custom_hash_algorithms.md#section) |
| `tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](data-sources--proxy--properties--tls_intercept--custom_certificate--custom_hash_algorithms.md#schema-tls_intercept--custom_certificate--custom_hash_algorithms--hash_algorithms) |
| `tls_intercept.custom_certificate.description_spec` | [tls_intercept.custom_certificate.description_spec](data-sources--proxy--properties--tls_intercept--custom_certificate.md#schema-tls_intercept--custom_certificate--description_spec) |
| `tls_intercept.custom_certificate.disable_ocsp_stapling` | [tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--proxy--properties--tls_intercept--custom_certificate--disable_ocsp_stapling.md#section) |
| `tls_intercept.custom_certificate.private_key` | [tls_intercept.custom_certificate.private_key](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key.md#section) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md#section) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md#schema-tls_intercept--custom_certificate--private_key--blindfold_secret_info--decryption_provider) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md#schema-tls_intercept--custom_certificate--private_key--blindfold_secret_info--location) |
| `tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--blindfold_secret_info.md#schema-tls_intercept--custom_certificate--private_key--blindfold_secret_info--store_provider) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info` | [tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--clear_secret_info.md#section) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--clear_secret_info.md#schema-tls_intercept--custom_certificate--private_key--clear_secret_info--provider_ref) |
| `tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [tls_intercept.custom_certificate.private_key.clear_secret_info.url](data-sources--proxy--properties--tls_intercept--custom_certificate--private_key--clear_secret_info.md#schema-tls_intercept--custom_certificate--private_key--clear_secret_info--url) |
| `tls_intercept.custom_certificate.use_system_defaults` | [tls_intercept.custom_certificate.use_system_defaults](data-sources--proxy--properties--tls_intercept--custom_certificate--use_system_defaults.md#section) |
| `tls_intercept.enable_for_all_domains` | [tls_intercept.enable_for_all_domains](data-sources--proxy--properties--tls_intercept--enable_for_all_domains.md#section) |
| `tls_intercept.policy` | [tls_intercept.policy](data-sources--proxy--properties--tls_intercept--policy.md#section) |
| `tls_intercept.policy.interception_rules` | [tls_intercept.policy.interception_rules](data-sources--proxy--properties--tls_intercept--policy--interception_rules.md#section) |
| `tls_intercept.policy.interception_rules.disable_interception` | [tls_intercept.policy.interception_rules.disable_interception](data-sources--proxy--properties--tls_intercept--policy--interception_rules--disable_interception.md#section) |
| `tls_intercept.policy.interception_rules.domain_match` | [tls_intercept.policy.interception_rules.domain_match](data-sources--proxy--properties--tls_intercept--policy--interception_rules--domain_match.md#section) |
| `tls_intercept.policy.interception_rules.domain_match.exact_value` | [tls_intercept.policy.interception_rules.domain_match.exact_value](data-sources--proxy--properties--tls_intercept--policy--interception_rules--domain_match.md#schema-tls_intercept--policy--interception_rules--domain_match--exact_value) |
| `tls_intercept.policy.interception_rules.domain_match.regex_value` | [tls_intercept.policy.interception_rules.domain_match.regex_value](data-sources--proxy--properties--tls_intercept--policy--interception_rules--domain_match.md#schema-tls_intercept--policy--interception_rules--domain_match--regex_value) |
| `tls_intercept.policy.interception_rules.domain_match.suffix_value` | [tls_intercept.policy.interception_rules.domain_match.suffix_value](data-sources--proxy--properties--tls_intercept--policy--interception_rules--domain_match.md#schema-tls_intercept--policy--interception_rules--domain_match--suffix_value) |
| `tls_intercept.policy.interception_rules.enable_interception` | [tls_intercept.policy.interception_rules.enable_interception](data-sources--proxy--properties--tls_intercept--policy--interception_rules--enable_interception.md#section) |
| `tls_intercept.trusted_ca_url` | [tls_intercept.trusted_ca_url](data-sources--proxy--properties--tls_intercept.md#schema-tls_intercept--trusted_ca_url) |
| `tls_intercept.volterra_certificate` | [tls_intercept.volterra_certificate](data-sources--proxy--properties--tls_intercept--volterra_certificate.md#section) |
| `tls_intercept.volterra_trusted_ca` | [tls_intercept.volterra_trusted_ca](data-sources--proxy--properties--tls_intercept--volterra_trusted_ca.md#section) |

## Next pages

- [active_forward_proxy_policies](data-sources--proxy--properties--active_forward_proxy_policies.md)
- [do_not_advertise](data-sources--proxy--properties--do_not_advertise.md)
- [dynamic_proxy](data-sources--proxy--properties--dynamic_proxy.md)
- [http_proxy](data-sources--proxy--properties--http_proxy.md)
- [no_forward_proxy_policy](data-sources--proxy--properties--no_forward_proxy_policy.md)
- [no_interception](data-sources--proxy--properties--no_interception.md)
- [site_local_inside_network](data-sources--proxy--properties--site_local_inside_network.md)
- [site_local_network](data-sources--proxy--properties--site_local_network.md)
- [site_virtual_sites](data-sources--proxy--properties--site_virtual_sites.md)
- [tls_intercept](data-sources--proxy--properties--tls_intercept.md)
- [xcsh_proxy](../data-sources/proxy.md)
