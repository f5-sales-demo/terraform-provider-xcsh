---
page_title: "routes.redirect_route"
subcategory: "Load Balancing"
description: "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL."
xcsh_docs: {"aliases": ["routes redirect route"], "body_bytes": 3572, "body_sha256": "sha256:ae224ac1cba294d3128a59bf92e37249cfdafd10a7d15ac5a84f6c31a5dab86f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:path", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:route_redirect"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes", "path": "documentation/data-sources/http_loadbalancer/properties/routes/redirect_route/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3031130012330222-2232013233232232-2002103331033200-2211001300312033-3312120221121332-1323201212001023-1020112222113103-2303212300333332", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "redirect_route"], "schema_version": 1, "sections": [{"aliases": ["routes redirect route headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "redirect_route", "headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes redirect route http method"], "anchor": "schema-routes--redirect_route--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "redirect_route", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes redirect route incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "redirect_route", "incoming_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes redirect route path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "redirect_route", "path"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes redirect route route redirect"], "anchor": "section", "description": "Route redirect parameters when match action is redirect.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:route_redirect", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "redirect_route", "route_redirect"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/redirect_route/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.redirect_route

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- routes.redirect_route

<a id="section"></a>

Type: `"single"`. Computed.

Redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the
matching traffic to a different URL.

Upstream description:

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/headers/): complete subsection reference.

<a id="schema-routes--redirect_route--http_method"></a>

### http_method property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/incoming_port/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/path/): complete subsection reference.

- [route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/route_redirect/): complete subsection reference.

## Next pages

- [routes.redirect_route.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/headers/)
- [routes.redirect_route.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/incoming_port/)
- [routes.redirect_route.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/path/)
- [routes.redirect_route.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/redirect_route/route_redirect/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/routes/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
