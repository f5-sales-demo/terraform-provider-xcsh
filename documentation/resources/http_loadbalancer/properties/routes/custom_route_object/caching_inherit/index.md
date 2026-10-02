---
page_title: "routes.custom_route_object.caching_inherit"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["routes custom route object caching inherit"], "body_bytes": 1503, "body_sha256": "sha256:4537b51fd3b09e2b6d563e939a2a1d371140d051e250065ca09e995be502eded", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "path": "documentation/resources/http_loadbalancer/properties/routes/custom_route_object/caching_inherit/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1212201113232021-1213023333100221-3121132020000323-0333203233000210-2112121231111212-0112211131311133-3323220002021033-3201103230031110", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom_route_object", "caching_inherit"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/custom_route_object/caching_inherit/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom_route_object.caching_inherit

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.custom_route_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/)
- routes.custom_route_object.caching_inherit

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

Upstream description:

This can be used for messages where no values are needed.

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
caching_inherit = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.custom_route_object](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/custom_route_object/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
