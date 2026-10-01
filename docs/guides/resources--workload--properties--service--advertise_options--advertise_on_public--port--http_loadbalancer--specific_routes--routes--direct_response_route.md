---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5783, "body_sha256": "sha256:ee586f53d3b3ad759902cf6bebb8ad66f68611b3733420a9abd1d6154f356e2a", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:direct_response_route", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:direct_response_route:headers", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:direct_response_route:path", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:direct_response_route:route_direct_response"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:direct_response_route", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes", "routes", "direct_response_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/direct_response_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [service](resources--workload--properties--service.md)
- [service.advertise_options](resources--workload--properties--service--advertise_options.md)
- [service.advertise_options.advertise_on_public](resources--workload--properties--service--advertise_options--advertise_on_public.md)
- [service.advertise_options.advertise_on_public.port](resources--workload--properties--service--advertise_options--advertise_on_public--port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
direct_response_route {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--headers.md): complete subsection reference.

<a id="schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--http_method"></a>

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

- [incoming_port](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port.md): complete subsection reference.

- [path](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--path.md): complete subsection reference.

- [route_direct_response](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.headers](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--headers.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--path.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../resources/workload.md)
