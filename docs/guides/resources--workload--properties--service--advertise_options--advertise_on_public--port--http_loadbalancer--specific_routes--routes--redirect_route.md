---
page_title: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route"
subcategory: "Container"
description: "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 5546, "body_sha256": "sha256:6d3393180c1be63ef5a02d493822f778590e37ebcd9b3f45dad273bc80451c2e", "canonical_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route:headers", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route:path", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes:redirect_route", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_on_public:port:http_loadbalancer:specific_routes:routes", "path": "docs/guides/resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route.md", "provider_name": "workload", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["service", "advertise_options", "advertise_on_public", "port", "http_loadbalancer", "specific_routes", "routes", "redirect_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_on_public/port/http_loadbalancer/specific_routes/routes/redirect_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

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
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

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

- [headers](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--headers.md): complete subsection reference.

<a id="schema-service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--http_method"></a>

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

- [incoming_port](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--incoming_port.md): complete subsection reference.

- [path](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--path.md): complete subsection reference.

- [route_redirect](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--route_redirect.md): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--headers.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--incoming_port.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--path.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes--redirect_route--route_redirect.md)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--properties--service--advertise_options--advertise_on_public--port--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../resources/workload.md)
