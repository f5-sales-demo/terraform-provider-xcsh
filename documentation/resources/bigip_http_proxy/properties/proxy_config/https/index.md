---
page_title: "proxy_config.https"
subcategory: ""
description: "Choice for selecting HTTP proxy with bring your own certificates."
xcsh_docs: {"aliases": ["proxy config https"], "body_bytes": 13623, "body_sha256": "sha256:23a8d3529f51021d9c13daca3e8eee239d3c41eb2204ead854b344e2578a0ef6", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_header", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_loadbalancer", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:disable_path_normalize", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:enable_path_normalize", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:non_default_loadbalancer", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "parent_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config", "path": "documentation/resources/bigip_http_proxy/properties/proxy_config/https/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3203121210323113-2101032200321202-0010303030311321-3220231332303132-1310103011003233-3221200330233120-1030332101023220-1003313201121111", "registry_path": "docs/guides/resources--bigip_http_proxy--reference--group-003.md", "relationships": [{"anchor": "schema-proxy_config--https--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,default_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--append_server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--port", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--port_ranges", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_header,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "schema-proxy_config--https--server_name", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:pass_through,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,default_header", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_header,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_header,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_loadbalancer,non_default_loadbalancer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_loadbalancer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:disable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:disable_path_normalize,enable_path_normalize", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:enable_path_normalize", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_loadbalancer,non_default_loadbalancer", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:non_default_loadbalancer", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:append_server_name,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:default_header,pass_through", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:pass_through,server_name", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:tls_cert_params,tls_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https:ConflictingObjectAttributes:tls_cert_params,tls_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_config", "https"], "schema_version": 1, "sections": [{"aliases": ["add hsts"], "anchor": "schema-proxy_config--https--add_hsts", "description": "Add HTTP Strict-Transport-Security response header.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "add_hsts"], "syntax": "attribute", "type": "bool"}, {"aliases": ["append server name"], "anchor": "schema-proxy_config--https--append_server_name", "description": "Exclusive with Define the header value for the header name “server”. If header value is already present, it is not overwritten and passed as-is.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "append_server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["coalescing options"], "anchor": "section", "description": "TLS connection coalescing configuration (not compatible with mTLS)", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.coalescing_options:ConflictingObjectAttributes:default_coalescing,strict_coalescing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:default_coalescing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.coalescing_options:ConflictingObjectAttributes:default_coalescing,strict_coalescing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:coalescing_options:strict_coalescing", "type": "conflicts"}], "schema_path": ["proxy_config", "https", "coalescing_options"], "syntax": "block", "type": "object"}, {"aliases": ["connection idle timeout", "duration", "operation timeout"], "anchor": "schema-proxy_config--https--connection_idle_timeout", "description": "The idle timeout for downstream connections. The idle timeout is defined as the period in which there are no active requests. When the idle timeout is reached the connection will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is specified in milliseconds.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "connection_idle_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["default header"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_header", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "default_header"], "syntax": "attribute", "type": "object"}, {"aliases": ["default loadbalancer"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:default_loadbalancer", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "default_loadbalancer"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable path normalize"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:disable_path_normalize", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "disable_path_normalize"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable path normalize"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:enable_path_normalize", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "enable_path_normalize"], "syntax": "attribute", "type": "object"}, {"aliases": ["http protocol options"], "anchor": "section", "description": "HTTP protocol configuration OPTIONS for downstream connections.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_only,http_protocol_enable_v1_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_only,http_protocol_enable_v2_only", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_only", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_only,http_protocol_enable_v1_v2", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_v2", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_v2,http_protocol_enable_v2_only", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v1_v2", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_only,http_protocol_enable_v2_only", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v2_only", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.http_protocol_options:ConflictingObjectAttributes:http_protocol_enable_v1_v2,http_protocol_enable_v2_only", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:http_protocol_options:http_protocol_enable_v2_only", "type": "conflicts"}], "schema_path": ["proxy_config", "https", "http_protocol_options"], "syntax": "block", "type": "object"}, {"aliases": ["http redirect"], "anchor": "schema-proxy_config--https--http_redirect", "description": "Redirect HTTP traffic to HTTPS.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "http_redirect"], "syntax": "attribute", "type": "bool"}, {"aliases": ["non default loadbalancer"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:non_default_loadbalancer", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "non_default_loadbalancer"], "syntax": "attribute", "type": "object"}, {"aliases": ["pass through"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:pass_through", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "pass_through"], "syntax": "attribute", "type": "object"}, {"aliases": ["port"], "anchor": "schema-proxy_config--https--port", "description": "Exclusive with HTTPS port to Listen.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["port ranges"], "anchor": "schema-proxy_config--https--port_ranges", "description": "Exclusive with A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by \"-\".", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "port_ranges"], "syntax": "attribute", "type": "string"}, {"aliases": ["server name"], "anchor": "schema-proxy_config--https--server_name", "description": "Exclusive with Define the header value for the header name “server”. This will overwrite existing values, if any, for the server header.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_config", "https", "server_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["tls cert params"], "anchor": "section", "description": "Select TLS Parameters and Certificates.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_cert_params:RequiredObjectAttributes:certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_cert_params:certificates", "type": "requires"}], "schema_path": ["proxy_config", "https", "tls_cert_params"], "syntax": "block", "type": "object"}, {"aliases": ["tls parameters"], "anchor": "section", "description": "Inline TLS parameters.", "document_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_parameters:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:no_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_parameters:ConflictingObjectAttributes:no_mtls,use_mtls", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:use_mtls", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "proxy_config.https.tls_parameters:RequiredObjectAttributes:tls_certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config:https:tls_parameters:tls_certificates", "type": "requires"}], "schema_path": ["proxy_config", "https", "tls_parameters"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/proxy_config/https/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Choice for selecting HTTP proxy with bring your own certificates.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_config.https

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- proxy_config.https

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

<a id="schema-proxy_config--https--add_hsts"></a>

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

<a id="schema-proxy_config--https--append_server_name"></a>

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

- [coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/): complete subsection reference.

<a id="schema-proxy_config--https--connection_idle_timeout"></a>

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

- [default_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/default_header/): complete subsection reference.

- [default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/default_loadbalancer/): complete subsection reference.

- [disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/disable_path_normalize/): complete subsection reference.

- [enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/enable_path_normalize/): complete subsection reference.

- [http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/): complete subsection reference.

<a id="schema-proxy_config--https--http_redirect"></a>

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

- [non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/non_default_loadbalancer/): complete subsection reference.

- [pass_through](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/pass_through/): complete subsection reference.

<a id="schema-proxy_config--https--port"></a>

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

<a id="schema-proxy_config--https--port_ranges"></a>

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

<a id="schema-proxy_config--https--server_name"></a>

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

- [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/): complete subsection reference.

- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/): complete subsection reference.

## Next pages

- [proxy_config.https.coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/)
- [proxy_config.https.default_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/default_header/)
- [proxy_config.https.default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/default_loadbalancer/)
- [proxy_config.https.disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/disable_path_normalize/)
- [proxy_config.https.enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/enable_path_normalize/)
- [proxy_config.https.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/)
- [proxy_config.https.non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/non_default_loadbalancer/)
- [proxy_config.https.pass_through](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/pass_through/)
- [proxy_config.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/)
- [proxy_config.https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/properties/proxy_config/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bigip_http_proxy/)
