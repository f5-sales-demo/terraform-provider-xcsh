---
page_title: "proxy_advertisement.advertise_custom.advertise_where"
subcategory: ""
description: "Where should this load balancer be available."
xcsh_docs: {"aliases": ["proxy advertisement advertise custom advertise where"], "body_bytes": 8941, "body_sha256": "sha256:cf52debd38381b3506dae56c1a44473ec31ea417a23807b524230d5e40a8a243", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_dualstack_on_public", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_on_public", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:site", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:use_default_port", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise custom advertise where advertise dualstack on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_dualstack_on_public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_dualstack_on_public"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise custom advertise where advertise on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_on_public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_on_public"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise custom advertise where advertise v6 on public"], "anchor": "section", "description": "This defines a way to advertise a load balancer on public. If optional public_ip is provided, it will only be advertised on RE sites where that public_ip is available.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:advertise_v6_on_public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "advertise_v6_on_public"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise custom advertise where port"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--port", "description": "Exclusive with Port to Listen.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["proxy advertisement advertise custom advertise where port ranges"], "anchor": "schema-proxy_advertisement--advertise_custom--advertise_where--port_ranges", "description": "Exclusive with A string containing a comma separated list of port ranges. Each port range consists of a single port or two ports separated by \"-\".", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "port_ranges"], "syntax": "attribute", "type": "string"}, {"aliases": ["proxy advertisement advertise custom advertise where site"], "anchor": "section", "description": "This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise custom advertise where use default port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:use_default_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "use_default_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise custom advertise where virtual network"], "anchor": "section", "description": "Parameters to advertise on a given virtual network.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise custom advertise where virtual site"], "anchor": "section", "description": "This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise custom advertise where virtual site with vip"], "anchor": "section", "description": "This defines a reference to a customer site virtual site along with network type and IP where a load balancer could be advertised.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:virtual_site_with_vip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "virtual_site_with_vip"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement advertise custom advertise where vk8s service"], "anchor": "section", "description": "This defines a reference to a RE site or virtual site where a load balancer could be advertised in the vK8s service network.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where:vk8s_service", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where", "vk8s_service"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Where should this load balancer be available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom.advertise_where

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/)
- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/)
- proxy_advertisement.advertise_custom.advertise_where

<a id="section"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/): complete subsection reference.

- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/): complete subsection reference.

- [advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/): complete subsection reference.

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="schema-proxy_advertisement--advertise_custom--advertise_where--port_ranges"></a>

### port_ranges property

Type: `"string"`. Computed.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/): complete subsection reference.

- [use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/use_default_port/): complete subsection reference.

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/): complete subsection reference.

- [virtual_site_with_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/): complete subsection reference.

- [vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/)
- [proxy_advertisement.advertise_custom.advertise_where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/)
- [proxy_advertisement.advertise_custom.advertise_where.use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/use_default_port/)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/)
- [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
