---
page_title: "routes.redirect_route"
subcategory: "Load Balancing"
description: "routes.redirect_route for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2923, "body_sha256": "sha256:b4bfccb300c65b7391a228d81809b20c2c6ac19dacc15a815a176b79b6202637", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:path", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:route_redirect"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--redirect_route.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "redirect_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/redirect_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.redirect_route for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.redirect_route

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
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

- [headers](data-sources--http_loadbalancer--properties--routes--redirect_route--headers.md): complete subsection reference.

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

- [incoming_port](data-sources--http_loadbalancer--properties--routes--redirect_route--incoming_port.md): complete subsection reference.

- [path](data-sources--http_loadbalancer--properties--routes--redirect_route--path.md): complete subsection reference.

- [route_redirect](data-sources--http_loadbalancer--properties--routes--redirect_route--route_redirect.md): complete subsection reference.

## Next pages

- [routes.redirect_route.headers](data-sources--http_loadbalancer--properties--routes--redirect_route--headers.md)
- [routes.redirect_route.incoming_port](data-sources--http_loadbalancer--properties--routes--redirect_route--incoming_port.md)
- [routes.redirect_route.path](data-sources--http_loadbalancer--properties--routes--redirect_route--path.md)
- [routes.redirect_route.route_redirect](data-sources--http_loadbalancer--properties--routes--redirect_route--route_redirect.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
