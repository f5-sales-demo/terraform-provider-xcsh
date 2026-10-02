---
page_title: "routes.match"
subcategory: ""
description: "Route match condition."
xcsh_docs: {"aliases": ["routes match"], "body_bytes": 3812, "body_sha256": "sha256:4892f8185d4d08d984c7e9dbd82f0406c87cde117ae4d2a97dad19cdae01d565", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:match:headers", "xcsh-docs:resources:route:properties:routes:match:incoming_port", "xcsh-docs:resources:route:properties:routes:match:path", "xcsh-docs:resources:route:properties:routes:match:query_params"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:match", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "documentation/resources/route/properties/routes/match/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303", "registry_path": "docs/guides/resources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "match"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:resources:route:properties:routes:match:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--match--headers--exact", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--exact", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--presence", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--presence", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--regex", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--regex", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--name", "enforcement": "provider-schema", "group": "routes.match.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "requires"}], "schema_path": ["routes", "match", "headers"], "syntax": "block", "type": "object"}, {"aliases": ["http method"], "anchor": "schema-routes--match--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:resources:route:properties:routes:match", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "match", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--match--incoming_port--port", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--match--incoming_port--port", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--match--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--match--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "type": "conflicts"}], "schema_path": ["routes", "match", "incoming_port"], "syntax": "block", "type": "object"}, {"aliases": ["path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:route:properties:routes:match:path", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--match--path--path", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--path", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--prefix", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--prefix", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--regex", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--regex", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}], "schema_path": ["routes", "match", "path"], "syntax": "block", "type": "object"}, {"aliases": ["query params"], "anchor": "section", "description": "List of (key, value) query parameters.", "document_id": "xcsh-docs:resources:route:properties:routes:match:query_params", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--match--query_params--exact", "enforcement": "provider-schema", "group": "routes.match.query_params:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:query_params", "type": "conflicts"}, {"anchor": "schema-routes--match--query_params--regex", "enforcement": "provider-schema", "group": "routes.match.query_params:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:query_params", "type": "conflicts"}, {"anchor": "schema-routes--match--query_params--key", "enforcement": "provider-schema", "group": "routes.match.query_params:RequiredListObjectAttributes:key", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:query_params", "type": "requires"}], "schema_path": ["routes", "match", "query_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/match/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Route match condition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.match

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- routes.match

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Match. Route match condition.

Upstream description:

Route match condition.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
match {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/headers/): complete subsection reference.

<a id="schema-routes--match--http_method"></a>

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

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/path/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/query_params/): complete subsection reference.

## Next pages

- [routes.match.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/headers/)
- [routes.match.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/)
- [routes.match.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/path/)
- [routes.match.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/query_params/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/)
