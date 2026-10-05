---
page_title: "stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route"
subcategory: "Container"
description: "A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic."
xcsh_docs: {"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer specific routes routes direct response route"], "body_bytes": 7163, "body_sha256": "sha256:7487584347310f22afb4e289d327d9670df2caa5ad9578a654a793909fdbd58e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:headers", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:path", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:route_direct_response"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2101110022312210-0030310322121123-1033101310330332-1312310032311100-3100032320110000-0021230210123302-0001022012103121-3022213313133301", "registry_path": "docs/guides/data-sources--workload--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route"], "schema_version": 1, "sections": [{"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer specific routes routes direct response route headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer specific routes routes direct response route http method"], "anchor": "schema-stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer specific routes routes direct response route incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:incoming_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "incoming_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer specific routes routes direct response route path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "path"], "syntax": "attribute", "type": "object"}, {"aliases": ["stateful service advertise options advertise on public multi ports ports http loadbalancer specific routes routes direct response route route direct response"], "anchor": "section", "description": "Send this direct response in case of route match action is direct response.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public:multi_ports:ports:http_loadbalancer:specific_routes:routes:direct_response_route:route_direct_response", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public", "multi_ports", "ports", "http_loadbalancer", "specific_routes", "routes", "direct_response_route", "route_direct_response"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [stateful_service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/headers/): complete subsection reference.

<a id="schema-stateful_service--advertise_options--advertise_on_public--multi_ports--ports--http_loadbalancer--specific_routes--routes--direct_response_route--http_method"></a>

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

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/incoming_port/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/path/): complete subsection reference.

- [route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/route_direct_response/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/headers/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/incoming_port/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/path/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/direct_response_route/route_direct_response/)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/multi_ports/ports/http_loadbalancer/specific_routes/routes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
