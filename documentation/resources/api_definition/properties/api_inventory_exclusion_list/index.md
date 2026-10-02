---
page_title: "api_inventory_exclusion_list"
subcategory: "API Management"
description: "List of API Endpoints excluded from the API Inventory."
xcsh_docs: {"aliases": ["api inventory exclusion list"], "body_bytes": 4751, "body_sha256": "sha256:eceee10b8fcd7b49c43222d3e146d70d49c889fddb7c2abd9688d537d7ab0349", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:properties:api_inventory_exclusion_list", "parent_id": "xcsh-docs:resources:api_definition:reference", "path": "documentation/resources/api_definition/properties/api_inventory_exclusion_list/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0322111332303200-0102230312330331-2230231331102322-0111003030300213-3103033311100113-2312223222013222-0332133333003331-1230011112023110", "registry_path": "docs/guides/resources--api_definition--reference--group-001.md", "relationships": [{"anchor": "schema-api_inventory_exclusion_list--path", "enforcement": "provider-schema", "group": "api_inventory_exclusion_list:RequiredListObjectAttributes:path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_definition:properties:api_inventory_exclusion_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_inventory_exclusion_list"], "schema_version": 1, "sections": [{"aliases": ["method"], "anchor": "schema-api_inventory_exclusion_list--method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:resources:api_definition:properties:api_inventory_exclusion_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_inventory_exclusion_list", "method"], "syntax": "attribute", "type": "string"}, {"aliases": ["path"], "anchor": "schema-api_inventory_exclusion_list--path", "description": "An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986 and may have parameters according to OpenAPI specification.", "document_id": "xcsh-docs:resources:api_definition:properties:api_inventory_exclusion_list", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_inventory_exclusion_list", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/properties/api_inventory_exclusion_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of API Endpoints excluded from the API Inventory.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_definitionCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_inventory_exclusion_list

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/properties/)
- api_inventory_exclusion_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of API Endpoints excluded from the API Inventory. Defaults to \`\[\]\`. Server applies default
when omitted.

Upstream description:

List of API Endpoints excluded from the API Inventory.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
api_inventory_exclusion_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-api_inventory_exclusion_list--method"></a>

### method property

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

<a id="schema-api_inventory_exclusion_list--path"></a>

### path property

Type: `"string"`. Optional.

Endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986
and may have parameters according to OpenAPI specification.

Upstream description:

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/properties/)
- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/)
