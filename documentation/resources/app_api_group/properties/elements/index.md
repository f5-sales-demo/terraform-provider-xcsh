---
page_title: "elements"
subcategory: ""
description: "List of API group elements with methods and path regex for matching requests."
xcsh_docs: {"aliases": ["elements"], "body_bytes": 5342, "body_sha256": "sha256:2bf820d0a562d4f75f79fbba6f5de5792312790b3b8cc26ea56fcde93523922e", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:properties:elements", "parent_id": "xcsh-docs:resources:app_api_group:reference", "path": "documentation/resources/app_api_group/properties/elements/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1200112120223122-2201203333220320-0111101133003203-2030033111231322-2122100331112303-0132313203213213-0220302212111033-0103230323020001", "registry_path": "docs/guides/resources--app_api_group--reference--group-001.md", "relationships": [{"anchor": "schema-elements--methods", "enforcement": "provider-schema", "group": "elements:RequiredListObjectAttributes:methods,path_regex", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:app_api_group:properties:elements", "type": "requires"}, {"anchor": "schema-elements--path_regex", "enforcement": "provider-schema", "group": "elements:RequiredListObjectAttributes:methods,path_regex", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:app_api_group:properties:elements", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["elements"], "schema_version": 1, "sections": [{"aliases": ["elements methods"], "anchor": "schema-elements--methods", "description": "List of method values to match the input request API method against. The match is considered to succeed if the input request API method is a member of the list.", "document_id": "xcsh-docs:resources:app_api_group:properties:elements", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["elements", "methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["elements path regex"], "anchor": "schema-elements--path_regex", "description": "Regular expression to match the input request API path against. The match is considered to succeed if the input request API path matches the specified path regex.", "document_id": "xcsh-docs:resources:app_api_group:properties:elements", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["elements", "path_regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/properties/elements/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of API group elements with methods and path regex for matching requests.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# elements

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- elements

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of API group elements with methods and path regex for matching requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("methods",
    "path_regex")}
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5000"
  }
}
```

Terraform syntax:

```terraform
elements {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-elements--methods"></a>

### methods property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of method values to
match the input request API method against. The match is considered to succeed if the input request
API method is a member of the list. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`,
\`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of method values to match the input request API method against. The match is considered to
succeed if the input request API method is a member of the list.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "0",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-elements--path_regex"></a>

### path_regex property

Type: `"string"`. Optional.

Regular expression to match the input request API path against. The match is considered to succeed
if the input request API path matches the specified path regex.

Upstream description:

Regular expression to match the input request API path against. The match is considered to succeed
if the input request API path matches the specified path regex.

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
    "minLength": 1
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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/)
