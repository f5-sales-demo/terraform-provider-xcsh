---
page_title: "routes.custom_route_object"
subcategory: "Load Balancing"
description: "A custom route uses a route object created outside of this view."
xcsh_docs: {"aliases": ["routes custom route object"], "body_bytes": 2390, "body_sha256": "sha256:a306a469d54eba720573ea0dda42efb69ee272c7c8282987a05afe40104a58f3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:caching_disable", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:route_ref"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes", "path": "documentation/data-sources/http_loadbalancer/properties/routes/custom_route_object/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2111330031230112-2100030222001233-0032133121122121-0332123032310212-0211320223220312-2120000231001112-0112012320313200-3022311310331230", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom_route_object"], "schema_version": 1, "sections": [{"aliases": ["routes custom route object caching disable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:caching_disable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom_route_object", "caching_disable"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes custom route object caching inherit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom_route_object", "caching_inherit"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes custom route object route ref"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:custom_route_object:route_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "custom_route_object", "route_ref"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/custom_route_object/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "A custom route uses a route object created outside of this view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom_route_object

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- routes.custom_route_object

<a id="section"></a>

Type: `"single"`. Computed.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

## Direct properties

- [caching_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/custom_route_object/caching_disable/): complete subsection reference.

- [caching_inherit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/custom_route_object/caching_inherit/): complete subsection reference.

- [route_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/custom_route_object/route_ref/): complete subsection reference.

## Next pages

- [routes.custom_route_object.caching_disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/custom_route_object/caching_disable/)
- [routes.custom_route_object.caching_inherit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/custom_route_object/caching_inherit/)
- [routes.custom_route_object.route_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/custom_route_object/route_ref/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
