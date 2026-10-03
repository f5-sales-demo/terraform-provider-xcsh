---
page_title: "proxy_advertisement.advertise_custom"
subcategory: ""
description: "This defines a way to advertise a VIP on specific sites."
xcsh_docs: {"aliases": ["proxy advertisement advertise custom"], "body_bytes": 1640, "body_sha256": "sha256:f195f66be4f7299fc2725a4fb38d896d28a9721b2e7b60ba2d6e700e95f3ac96", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement", "path": "documentation/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["proxy_advertisement", "advertise_custom"], "schema_version": 1, "sections": [{"aliases": ["proxy advertisement advertise custom advertise where"], "anchor": "section", "description": "Where should this load balancer be available.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement:advertise_custom:advertise_where", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["proxy_advertisement", "advertise_custom", "advertise_where"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a way to advertise a VIP on specific sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# proxy_advertisement.advertise_custom

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/)
- proxy_advertisement.advertise_custom

<a id="section"></a>

Type: `"single"`. Computed.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

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

## Direct properties

- [advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/): complete subsection reference.

## Next pages

- [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
