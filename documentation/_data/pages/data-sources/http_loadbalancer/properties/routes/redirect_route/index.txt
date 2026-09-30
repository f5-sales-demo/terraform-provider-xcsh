---
page_title: "routes.redirect_route"
subcategory: "Load Balancing"
description: "routes.redirect_route for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3473, "body_sha256": "sha256:1abe9f3b6adeb78213763f4256144a2baa1ec0537c3cc5abf9822fa7011bcc9f", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:path", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route:route_redirect"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:redirect_route", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes", "path": "documentation/data-sources/http_loadbalancer/properties/routes/redirect_route/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["routes", "redirect_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/redirect_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.redirect_route for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
