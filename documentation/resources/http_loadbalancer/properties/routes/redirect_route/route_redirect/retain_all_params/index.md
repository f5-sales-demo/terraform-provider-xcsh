---
page_title: "routes.redirect_route.route_redirect.retain_all_params"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["routes redirect route route redirect retain all params"], "body_bytes": 1716, "body_sha256": "sha256:d4509b4d58b14371d032fd7d8f4a43b8eecf3dcf516b47ae39b869592f07d27a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect:retain_all_params", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect", "path": "documentation/resources/http_loadbalancer/properties/routes/redirect_route/route_redirect/retain_all_params/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1101201133222023-1231130333213023-2122000132101301-0201202130021010-1011100000001301-1131232322101100-3020000300030123-0123332022113020", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "redirect_route", "route_redirect", "retain_all_params"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/redirect_route/route_redirect/retain_all_params/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.redirect_route.route_redirect.retain_all_params

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [routes.redirect_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/)
- [routes.redirect_route.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/route_redirect/)
- routes.redirect_route.route_redirect.retain_all_params

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.redirect_route.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/route_redirect/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
