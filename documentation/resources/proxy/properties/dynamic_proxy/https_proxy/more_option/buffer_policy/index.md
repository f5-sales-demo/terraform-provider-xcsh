---
page_title: "dynamic_proxy.https_proxy.more_option.buffer_policy"
subcategory: ""
description: "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level"
xcsh_docs: {"aliases": ["dynamic proxy https proxy more option buffer policy"], "body_bytes": 2996, "body_sha256": "sha256:86bf97645db93602642d054bb1fe2314b29110c21f70d2896c3c24ecb5e92071", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:buffer_policy", "parent_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option", "path": "documentation/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/buffer_policy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0123133211122323-2023233032120021-3312312311300302-2000102323113120-2103010232021302-2001002012303122-0312102132203130-1031232032103123", "registry_path": "docs/guides/resources--proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dynamic_proxy", "https_proxy", "more_option", "buffer_policy"], "schema_version": 1, "sections": [{"aliases": ["dynamic proxy https proxy more option buffer policy disabled"], "anchor": "schema-dynamic_proxy--https_proxy--more_option--buffer_policy--disabled", "description": "Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:buffer_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "more_option", "buffer_policy", "disabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["dynamic proxy https proxy more option buffer policy max request bytes"], "anchor": "schema-dynamic_proxy--https_proxy--more_option--buffer_policy--max_request_bytes", "description": "The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response.", "document_id": "xcsh-docs:resources:proxy:properties:dynamic_proxy:https_proxy:more_option:buffer_policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dynamic_proxy", "https_proxy", "more_option", "buffer_policy", "max_request_bytes"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/buffer_policy/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["proxyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dynamic_proxy.https_proxy.more_option.buffer_policy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [dynamic_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/)
- [dynamic_proxy.https_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/)
- [dynamic_proxy.https_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/dynamic_proxy/https_proxy/more_option/)
- dynamic_proxy.https_proxy.more_option.buffer_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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

Terraform syntax:

```terraform
buffer_policy {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-dynamic_proxy--https_proxy--more_option--buffer_policy--disabled"></a>

### disabled property

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="schema-dynamic_proxy--https_proxy--more_option--buffer_policy--max_request_bytes"></a>

### max_request_bytes property

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```
