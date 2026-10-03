---
page_title: "routes.match"
subcategory: ""
description: "Route match condition."
xcsh_docs: {"aliases": ["routes match"], "body_bytes": 3476, "body_sha256": "sha256:ea25c6c8099608a006c44547c3d6601cc4b0d7f8cbfa68d4480e10046c94276d", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:match:headers", "xcsh-docs:data-sources:route:properties:routes:match:incoming_port", "xcsh-docs:data-sources:route:properties:routes:match:path", "xcsh-docs:data-sources:route:properties:routes:match:query_params"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:match", "parent_id": "xcsh-docs:data-sources:route:properties:routes", "path": "documentation/data-sources/route/properties/routes/match/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2202021113211000-0321023232131332-0023133030000100-0113121112201031-1323001231200003-1323310313022102-0010110103313023-3023113002113312", "registry_path": "docs/guides/data-sources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "match"], "schema_version": 1, "sections": [{"aliases": ["routes match headers"], "anchor": "section", "description": "List of (key, value) headers.", "document_id": "xcsh-docs:data-sources:route:properties:routes:match:headers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "match", "headers"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes match http method"], "anchor": "schema-routes--match--http_method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:data-sources:route:properties:routes:match", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "match", "http_method"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes match incoming port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:data-sources:route:properties:routes:match:incoming_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "match", "incoming_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes match path"], "anchor": "section", "description": "Path match of the URI can be either be, Prefix match or exact match or regular expression match.", "document_id": "xcsh-docs:data-sources:route:properties:routes:match:path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "match", "path"], "syntax": "attribute", "type": "object"}, {"aliases": ["routes match query params"], "anchor": "section", "description": "List of (key, value) query parameters.", "document_id": "xcsh-docs:data-sources:route:properties:routes:match:query_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["routes", "match", "query_params"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/match/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Route match condition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["routeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.match

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- routes.match

<a id="section"></a>

Type: `"list"`. Computed.

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

## Direct properties

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/headers/): complete subsection reference.

<a id="schema-routes--match--http_method"></a>

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

- [incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/incoming_port/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/path/): complete subsection reference.

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/query_params/): complete subsection reference.

## Next pages

- [routes.match.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/headers/)
- [routes.match.incoming_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/incoming_port/)
- [routes.match.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/path/)
- [routes.match.query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/match/query_params/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
