---
page_title: "non_api_endpoints"
subcategory: "API Management"
description: "List of Non-API Endpoints."
xcsh_docs: {"aliases": ["non api endpoints"], "body_bytes": 4480, "body_sha256": "sha256:dd0fed33a94dc20c2e04fc8dc7f9a9ef76d72968f05f0f34948a79912659402c", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:properties:non_api_endpoints", "parent_id": "xcsh-docs:resources:api_definition:reference", "path": "documentation/resources/api_definition/properties/non_api_endpoints/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3013111320133022-1120132002103332-1322121200101212-1332122003230001-3132113222020231-0022312331223301-2130330000300003-1313003013011323", "registry_path": "docs/guides/resources--api_definition--reference--group-001.md", "relationships": [{"anchor": "schema-non_api_endpoints--path", "enforcement": "provider-schema", "group": "non_api_endpoints:RequiredListObjectAttributes:path", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:api_definition:properties:non_api_endpoints", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["non_api_endpoints"], "schema_version": 1, "sections": [{"aliases": ["non api endpoints method"], "anchor": "schema-non_api_endpoints--method", "description": "Specifies the HTTP method used to access a resource. Any HTTP Method.", "document_id": "xcsh-docs:resources:api_definition:properties:non_api_endpoints", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ANY", "CONNECT", "COPY", "DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT", "TRACE"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["non_api_endpoints", "method"], "syntax": "attribute", "type": "string"}, {"aliases": ["non api endpoints path"], "anchor": "schema-non_api_endpoints--path", "description": "An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC 3986 and may have parameters according to OpenAPI specification.", "document_id": "xcsh-docs:resources:api_definition:properties:non_api_endpoints", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["non_api_endpoints", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/properties/non_api_endpoints/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of Non-API Endpoints.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["api_definitionCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# non_api_endpoints

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/properties/)
- non_api_endpoints

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

API Discovery Exclusion List. List of Non-API Endpoints. Defaults to \`\[\]\`. Server applies
default when omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
non_api_endpoints {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-non_api_endpoints--method"></a>

### method property

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

<a id="schema-non_api_endpoints--path"></a>

### path property

Type: `"string"`. Optional.

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
