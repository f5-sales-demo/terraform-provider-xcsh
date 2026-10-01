---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5318, "body_sha256": "sha256:4f7a6fd00b3df2c29df9c78f6fbc56eb8c380daebb805d2c526e8fc34ef98223", "canonical_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route:headers", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route:path", "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect"], "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route", "parent_id": "xcsh-docs:data-sources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes", "path": "docs/guides/data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes", "routes", "redirect_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/redirect_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md)
- [Property reference](data-sources--workload--reference.md)
- [service](data-sources--workload--properties--service.md)
- [service.advertise_options](data-sources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](data-sources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

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

- [headers](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--headers.md): complete subsection reference.

<a id="schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--http_method"></a>

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

- [incoming_port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--incoming_port.md): complete subsection reference.

- [path](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--path.md): complete subsection reference.

- [route_redirect](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--route_redirect.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--headers.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--incoming_port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--path.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--route_redirect.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](data-sources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../data-sources/workload.md)
