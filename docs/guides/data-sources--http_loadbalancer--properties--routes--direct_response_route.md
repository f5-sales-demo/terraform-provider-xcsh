---
page_title: "routes.direct_response_route"
subcategory: "Load Balancing"
description: "routes.direct_response_route for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2955, "body_sha256": "sha256:1fc6b5ce9ffa63c36ff16cda3ba465b7f27cb58630dc319c04e8a028ccdd903d", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:incoming_port", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:path", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route:route_direct_response"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:direct_response_route", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--direct_response_route.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "direct_response_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/direct_response_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.direct_response_route for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.direct_response_route

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- routes.direct_response_route

<a id="section"></a>

Type: `"single"`. Computed.

Direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Upstream description:

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

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

- [headers](data-sources--http_loadbalancer--properties--routes--direct_response_route--headers.md): complete subsection reference.

<a id="schema-routes--direct_response_route--http_method"></a>

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

- [incoming_port](data-sources--http_loadbalancer--properties--routes--direct_response_route--incoming_port.md): complete subsection reference.

- [path](data-sources--http_loadbalancer--properties--routes--direct_response_route--path.md): complete subsection reference.

- [route_direct_response](data-sources--http_loadbalancer--properties--routes--direct_response_route--route_direct_response.md): complete subsection reference.

## Next pages

- [routes.direct_response_route.headers](data-sources--http_loadbalancer--properties--routes--direct_response_route--headers.md)
- [routes.direct_response_route.incoming_port](data-sources--http_loadbalancer--properties--routes--direct_response_route--incoming_port.md)
- [routes.direct_response_route.path](data-sources--http_loadbalancer--properties--routes--direct_response_route--path.md)
- [routes.direct_response_route.route_direct_response](data-sources--http_loadbalancer--properties--routes--direct_response_route--route_direct_response.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
