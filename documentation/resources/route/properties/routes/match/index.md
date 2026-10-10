---
page_title: "routes.match"
subcategory: ""
description: "Route match condition."
xcsh_docs: {"aliases": ["routes match"], "body_bytes": 3242, "body_sha256": "sha256:1bbc576cba70a2898414aa7e43bbfedc371459039f5a33f5eab34532a5dc7558", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:route:properties:routes:match:headers", "xcsh-docs:resources:route:properties:routes:match:incoming_port", "xcsh-docs:resources:route:properties:routes:match:path", "xcsh-docs:resources:route:properties:routes:match:query_params"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "id": "xcsh-docs:resources:route:properties:routes:match", "parent_id": "xcsh-docs:resources:route:properties:routes", "path": "documentation/resources/route/properties/routes/match/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303", "registry_path": "docs/guides/resources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "match"], "schema_version": 1, "sections": [{"aliases": ["routes match headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:resources:route:properties:routes:match:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--match--headers--exact", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--exact", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--presence", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:exact,presence", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--presence", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--regex", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--regex", "enforcement": "provider-schema", "group": "routes.match.headers:ConflictingListObjectAttributes:presence,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "conflicts"}, {"anchor": "schema-routes--match--headers--name", "enforcement": "provider-schema", "group": "routes.match.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:headers", "type": "requires"}], "schema_path": ["routes", "match", "headers"], "syntax": "block", "type": "object"}, {"aliases": ["routes match http method"], "anchor": "schema-routes--match--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:resources:route:properties:routes:match", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ANY", "CONNECT", "COPY", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "match", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes match incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--match--incoming_port--port", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--match--incoming_port--port", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--match--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "type": "conflicts"}, {"anchor": "schema-routes--match--incoming_port--port_ranges", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "routes.match.incoming_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:incoming_port:no_port_match", "type": "conflicts"}], "schema_path": ["routes", "match", "incoming_port"], "syntax": "block", "type": "object"}, {"aliases": ["routes match path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:resources:route:properties:routes:match:path", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-routes--match--path--path", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--path", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--prefix", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,prefix", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--prefix", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--regex", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:path,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}, {"anchor": "schema-routes--match--path--regex", "enforcement": "provider-schema", "group": "routes.match.path:ConflictingObjectAttributes:prefix,regex", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:path", "type": "conflicts"}], "schema_path": ["routes", "match", "path"], "syntax": "block", "type": "object"}, {"aliases": ["routes match query params"], "anchor": "section", "description": "List of (key, value) query parameters.", "document_id": "xcsh-docs:resources:route:properties:routes:match:query_params", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-routes--match--query_params--exact", "enforcement": "provider-schema", "group": "routes.match.query_params:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:query_params", "type": "conflicts"}, {"anchor": "schema-routes--match--query_params--regex", "enforcement": "provider-schema", "group": "routes.match.query_params:ConflictingListObjectAttributes:exact,regex", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:query_params", "type": "conflicts"}, {"anchor": "schema-routes--match--query_params--key", "enforcement": "provider-schema", "group": "routes.match.query_params:RequiredListObjectAttributes:key", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:route:properties:routes:match:query_params", "type": "requires"}], "schema_path": ["routes", "match", "query_params"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/properties/routes/match/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Route match condition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["routeCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/incoming_port/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/path/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/route/properties/routes/match/query_params/): complete subsection reference.
