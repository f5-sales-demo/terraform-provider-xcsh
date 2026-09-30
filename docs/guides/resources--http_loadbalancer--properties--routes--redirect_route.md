---
page_title: "routes.redirect_route"
subcategory: "Load Balancing"
description: "routes.redirect_route for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3169, "body_sha256": "sha256:b0b6c90e033c9397a146613cfe1d8d6aa822ef704f9759e455e466b3389962ac", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "path": "docs/guides/resources--http_loadbalancer--properties--routes--redirect_route.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "redirect_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/redirect_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.redirect_route for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.redirect_route

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- routes.redirect_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
redirect_route {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](resources--http_loadbalancer--properties--routes--redirect_route--headers.md): complete subsection reference.

<a id="schema-routes--redirect_route--http_method"></a>

### http_method property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

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

- [incoming_port](resources--http_loadbalancer--properties--routes--redirect_route--incoming_port.md): complete subsection reference.

- [path](resources--http_loadbalancer--properties--routes--redirect_route--path.md): complete subsection reference.

- [route_redirect](resources--http_loadbalancer--properties--routes--redirect_route--route_redirect.md): complete subsection reference.

## Next pages

- [routes.redirect_route.headers](resources--http_loadbalancer--properties--routes--redirect_route--headers.md)
- [routes.redirect_route.incoming_port](resources--http_loadbalancer--properties--routes--redirect_route--incoming_port.md)
- [routes.redirect_route.path](resources--http_loadbalancer--properties--routes--redirect_route--path.md)
- [routes.redirect_route.route_redirect](resources--http_loadbalancer--properties--routes--redirect_route--route_redirect.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
