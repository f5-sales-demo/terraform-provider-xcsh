---
page_title: "routes.redirect_route"
subcategory: "Load Balancing"
description: "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL."
xcsh_docs: {"aliases": ["routes redirect route"], "body_bytes": 3073, "body_sha256": "sha256:adf80d07192dd89e9e2c7a2e2995792de939bdd07bd9936e1bed2d9cdc12a6d5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "path": "documentation/resources/http_loadbalancer/properties/routes/redirect_route/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1211202033310002-2022303010112111-1203302130013122-0230232000103023-1331332320302032-1123010322232032-3033123123323331-0133321031111223", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-025.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "redirect_route"], "schema_version": 1, "sections": [{"aliases": ["routes redirect route headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--redirect_route--headers--exact", "enforcement": "provider-schema", "group": "routes.redirect_route.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--headers--exact", "enforcement": "provider-schema", "group": "routes.redirect_route.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--headers--presence", "enforcement": "provider-schema", "group": "routes.redirect_route.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--headers--presence", "enforcement": "provider-schema", "group": "routes.redirect_route.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--headers--regex", "enforcement": "provider-schema", "group": "routes.redirect_route.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--headers--regex", "enforcement": "provider-schema", "group": "routes.redirect_route.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--headers--name", "enforcement": "provider-schema", "group": "routes.redirect_route.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:headers", "type": "requires"}], "schema_path": ["routes", "redirect_route", "headers"], "syntax": "block", "type": "object"}, {"aliases": ["routes redirect route http method"], "anchor": "schema-routes--redirect_route--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ANY", "CONNECT", "COPY", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "redirect_route", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes redirect route incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--redirect_route--incoming_port--port", "enforcement": "provider-schema", "group": "routes.redirect_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--incoming_port--port", "enforcement": "provider-schema", "group": "routes.redirect_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.redirect_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.redirect_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.redirect_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.redirect_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:incoming_port:no_port_match", "type": "conflicts"}], "schema_path": ["routes", "redirect_route", "incoming_port"], "syntax": "block", "type": "object"}, {"aliases": ["routes redirect route path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--redirect_route--path--path", "enforcement": "provider-schema", "group": "routes.redirect_route.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--path--path", "enforcement": "provider-schema", "group": "routes.redirect_route.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--path--prefix", "enforcement": "provider-schema", "group": "routes.redirect_route.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--path--prefix", "enforcement": "provider-schema", "group": "routes.redirect_route.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--path--regex", "enforcement": "provider-schema", "group": "routes.redirect_route.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--path--regex", "enforcement": "provider-schema", "group": "routes.redirect_route.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:path", "type": "conflicts"}], "schema_path": ["routes", "redirect_route", "path"], "syntax": "block", "type": "object"}, {"aliases": ["routes redirect route route redirect"], "anchor": "section", "description": "Route redirect parameters when match action is redirect.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--redirect_route--route_redirect--path_redirect", "enforcement": "provider-schema", "group": "routes.redirect_route.route_redirect:ConflictingObjectAttributes:path_redirect,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--route_redirect--prefix_rewrite", "enforcement": "provider-schema", "group": "routes.redirect_route.route_redirect:ConflictingObjectAttributes:path_redirect,prefix_rewrite", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--route_redirect--replace_params", "enforcement": "provider-schema", "group": "routes.redirect_route.route_redirect:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect", "type": "conflicts"}, {"anchor": "schema-routes--redirect_route--route_redirect--replace_params", "enforcement": "provider-schema", "group": "routes.redirect_route.route_redirect:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.redirect_route.route_redirect:ConflictingObjectAttributes:remove_all_params,replace_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.redirect_route.route_redirect:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect:remove_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.redirect_route.route_redirect:ConflictingObjectAttributes:remove_all_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect:retain_all_params", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.redirect_route.route_redirect:ConflictingObjectAttributes:replace_params,retain_all_params", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect:retain_all_params", "type": "conflicts"}], "schema_path": ["routes", "redirect_route", "route_redirect"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/redirect_route/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the matching traffic to a different URL.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.redirect_route

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- routes.redirect_route

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/headers/): complete subsection reference.

<a id="schema-routes--redirect_route--http_method"></a>

### http_method property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/incoming_port/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/path/): complete subsection reference.

- [route_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/redirect_route/route_redirect/): complete subsection reference.
