---
page_title: "routes.direct_response_route"
subcategory: "Load Balancing"
description: "A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic."
xcsh_docs: {"aliases": ["routes direct response route"], "body_bytes": 4055, "body_sha256": "sha256:8adf4bdd1d7df79ad1f5f4b1f7e8cde36cd8214ef927eb8653e83ca8c471d865", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port", "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:path", "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:route_direct_response"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "path": "documentation/resources/http_loadbalancer/properties/routes/direct_response_route/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2022020131111323-3202323023313121-3232301211300210-3200302022310101-2311200321100013-0302020023130211-3333323200113311-0021321000303311", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "direct_response_route"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--direct_response_route--headers--exact", "enforcement": "provider-schema", "group": "routes.direct_response_route.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--headers--exact", "enforcement": "provider-schema", "group": "routes.direct_response_route.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--headers--presence", "enforcement": "provider-schema", "group": "routes.direct_response_route.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--headers--presence", "enforcement": "provider-schema", "group": "routes.direct_response_route.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--headers--regex", "enforcement": "provider-schema", "group": "routes.direct_response_route.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--headers--regex", "enforcement": "provider-schema", "group": "routes.direct_response_route.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--headers--name", "enforcement": "provider-schema", "group": "routes.direct_response_route.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:headers", "type": "requires"}], "schema_path": ["routes", "direct_response_route", "headers"], "syntax": "block", "type": "object"}, {"aliases": ["http method"], "anchor": "schema-routes--direct_response_route--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "direct_response_route", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--direct_response_route--incoming_port--port", "enforcement": "provider-schema", "group": "routes.direct_response_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--incoming_port--port", "enforcement": "provider-schema", "group": "routes.direct_response_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.direct_response_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.direct_response_route.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.direct_response_route.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.direct_response_route.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:incoming_port:no_port_match", "type": "conflicts"}], "schema_path": ["routes", "direct_response_route", "incoming_port"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--direct_response_route--path--path", "enforcement": "provider-schema", "group": "routes.direct_response_route.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:path", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--path--path", "enforcement": "provider-schema", "group": "routes.direct_response_route.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:path", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--path--prefix", "enforcement": "provider-schema", "group": "routes.direct_response_route.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:path", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--path--prefix", "enforcement": "provider-schema", "group": "routes.direct_response_route.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:path", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--path--regex", "enforcement": "provider-schema", "group": "routes.direct_response_route.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:path", "type": "conflicts"}, {"anchor": "schema-routes--direct_response_route--path--regex", "enforcement": "provider-schema", "group": "routes.direct_response_route.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:path", "type": "conflicts"}], "schema_path": ["routes", "direct_response_route", "path"], "syntax": "block", "type": "object"}, {"aliases": ["route direct response"], "anchor": "section", "description": "Send this direct response in case of route match action is direct response.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:route_direct_response", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--direct_response_route--route_direct_response--response_code", "enforcement": "provider-schema", "group": "routes.direct_response_route.route_direct_response:RequiredObjectAttributes:response_code", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:direct_response_route:route_direct_response", "type": "requires"}], "schema_path": ["routes", "direct_response_route", "route_direct_response"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/direct_response_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "A direct response route matches on path, incoming header, incoming port and/or HTTP method and responds directly to the matching traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.direct_response_route

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- routes.direct_response_route

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/headers/): complete subsection reference.

<a id="schema-routes--direct_response_route--http_method"></a>

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

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/incoming_port/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/path/): complete subsection reference.

- [route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/route_direct_response/): complete subsection reference.

## Next pages

- [routes.direct_response_route.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/headers/)
- [routes.direct_response_route.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/incoming_port/)
- [routes.direct_response_route.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/path/)
- [routes.direct_response_route.route_direct_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/direct_response_route/route_direct_response/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/routes/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
