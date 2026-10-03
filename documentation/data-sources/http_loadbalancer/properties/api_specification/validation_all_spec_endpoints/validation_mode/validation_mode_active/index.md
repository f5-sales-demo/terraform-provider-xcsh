---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active"
subcategory: "Load Balancing"
description: "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints validation mode validation mode active"], "body_bytes": 5235, "body_sha256": "sha256:11b9575fe937b35f9531828aa0a15c57c7ee5a605654653388213bc06f766b92", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_block", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_report"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "path": "documentation/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints validation mode validation mode active enforcement block"], "anchor": "section", "description": "Blocking validation: reject traffic that violates the selected OpenAPI validation properties. Invalid requests are returned as HTTP 403.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active", "enforcement_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode validation mode active enforcement report"], "anchor": "section", "description": "Report-only validation: record OpenAPI violations while allowing the request or response to continue.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_report", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active", "enforcement_report"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode validation mode active request validation properties"], "anchor": "schema-api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--request_validation_properties", "description": "List of properties of the request to validate according to the OpenAPI specification file (a.k.a. Swagger)", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active", "request_validation_properties"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="section"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

## Direct properties

- [enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_block/): complete subsection reference.

- [enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_report/): complete subsection reference.

<a id="schema-api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--request_validation_properties"></a>

### request_validation_properties property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_block/)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_report/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
