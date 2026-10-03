---
page_title: "http_proxy.more_option.buffer_policy"
subcategory: ""
description: "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level"
xcsh_docs: {"aliases": ["http proxy more option buffer policy"], "body_bytes": 3784, "body_sha256": "sha256:602d6de218d9dac946c3636f944a43d424a9acc3ef7ba7e3df9e1876d4125244", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:buffer_policy", "parent_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option", "path": "documentation/resources/proxy/properties/http_proxy/more_option/buffer_policy/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0300111113000322-0301122303123033-2331031131312202-2232013311303303-2130110102131033-3332231021130301-2213011301023020-3333033311323303", "registry_path": "docs/guides/resources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_proxy", "more_option", "buffer_policy"], "schema_version": 1, "sections": [{"aliases": ["http proxy more option buffer policy disabled"], "anchor": "schema-http_proxy--more_option--buffer_policy--disabled", "description": "Disable buffering for a particular route. This is useful when virtual-host has buffering, but we need to disable it on a specific route. The value of this field is ignored for virtual-host.", "document_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:buffer_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_proxy", "more_option", "buffer_policy", "disabled"], "syntax": "attribute", "type": "bool"}, {"aliases": ["http proxy more option buffer policy max request bytes"], "anchor": "schema-http_proxy--more_option--buffer_policy--max_request_bytes", "description": "The maximum request size that the filter will buffer before the connection manager will stop buffering and return a RequestEntityTooLarge (413) response.", "document_id": "xcsh-docs:resources:proxy:properties:http_proxy:more_option:buffer_policy", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["http_proxy", "more_option", "buffer_policy", "max_request_bytes"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/http_proxy/more_option/buffer_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Some upstream applications are not capable of handling streamed data. This config enables buffering the entire request before sending to upstream application. We can specify the maximum buffer size and buffer interval with this config. Buffering can be enabled and disabled at VirtualHost and Route levels Route level", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_proxy.more_option.buffer_policy

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/)
- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- http_proxy.more_option.buffer_policy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

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

<a id="schema-http_proxy--more_option--buffer_policy--disabled"></a>

### disabled property

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

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

<a id="schema-http_proxy--more_option--buffer_policy--max_request_bytes"></a>

### max_request_bytes property

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

## Next pages

- [http_proxy.more_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/http_proxy/more_option/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
