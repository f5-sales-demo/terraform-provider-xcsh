---
page_title: "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route"
subcategory: "Container"
description: "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL."
xcsh_docs: {"aliases": ["service advertise options advertise custom ports http loadbalancer specific routes routes redirect route"], "body_bytes": 6469, "body_sha256": "sha256:991148a2ba17c33232bee4045bbde2a303c3dc7ab40a3105b7a77ef56eadd44e", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:headers", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:path", "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route", "parent_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes", "path": "documentation/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113", "registry_path": "docs/guides/resources--workload--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "redirect_route"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "redirect_route", "headers"], "syntax": "block", "type": "object"}, {"aliases": ["http method"], "anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "redirect_route", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--incoming_port--port", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--incoming_port--port", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:incoming_port:no_port_match", "type": "conflicts"}], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "redirect_route", "incoming_port"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--path--path", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--path--path", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--path--prefix", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--path--prefix", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--path--regex", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--path--regex", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:path", "type": "conflicts"}], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "redirect_route", "path"], "syntax": "block", "type": "object"}, {"aliases": ["route redirect"], "anchor": "section", "description": "Route redirect parameters when match action is redirect.", "document_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--route_redirect--path_redirect", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect:ConflictingObjectAttributes:path_redirect,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--route_redirect--prefix_rewrite", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect:ConflictingObjectAttributes:path_redirect,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--route_redirect--replace_params", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect", "type": "conflicts"}, {"anchor": "schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--route_redirect--replace_params", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect:retain_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:advertise_options:advertise_custom:ports:http_loadbalancer:specific_routes:routes:redirect_route:route_redirect:retain_all_params", "type": "conflicts"}], "schema_path": ["service", "advertise_options", "advertise_custom", "ports", "http_loadbalancer", "specific_routes", "routes", "redirect_route", "route_redirect"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.advertise_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/)
- [service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/)
- [service.advertise_options.advertise_custom.ports](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/headers/): complete subsection reference.

<a id="schema-service--advertise_options--advertise_custom--ports--http_loadbalancer--specific_routes--routes--redirect_route--http_method"></a>

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

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/incoming_port/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/path/): complete subsection reference.

- [route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/route_redirect/): complete subsection reference.

## Next pages

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/headers/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/incoming_port/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/path/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/redirect_route/route_redirect/)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/advertise_options/advertise_custom/ports/http_loadbalancer/specific_routes/routes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
