---
page_title: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route"
subcategory: "Container"
description: "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route for xcsh_workload."
xcsh_docs: {"aliases": [], "body_bytes": 6006, "body_sha256": "sha256:ed2e4f5886190452c59f3f725e360c5535197eeb83b384e02e3e1840c25108d4", "canonical_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "child_ids": ["xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:headers", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:path", "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route:route_direct_response"], "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "parent_id": "xcsh-docs:resources:workload:properties:stateful_service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes", "path": "docs/guides/resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route.md", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/stateful_service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/direct_response_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route for xcsh_workload.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route

Breadcrumbs:

- [xcsh_workload](../resources/workload.md)
- [Property reference](resources--workload--reference.md)
- [stateful_service](resources--workload--properties--stateful_service.md)
- [stateful_service.advertise_options](resources--workload--properties--stateful_service--advertise_options.md)
- [stateful_service.advertise_options.advertise_custom](resources--workload--properties--stateful_service--advertise_options--advertise_custom.md)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route

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

- [headers](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--headers.md): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--http_method"></a>

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

- [incoming_port](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port.md): complete subsection reference.

- [path](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--path.md): complete subsection reference.

- [route_direct_response](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response.md): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--headers.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--incoming_port.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--path.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--direct_response_route--route_direct_response.md)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--properties--stateful_service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes.md)
- [xcsh_workload](../resources/workload.md)
