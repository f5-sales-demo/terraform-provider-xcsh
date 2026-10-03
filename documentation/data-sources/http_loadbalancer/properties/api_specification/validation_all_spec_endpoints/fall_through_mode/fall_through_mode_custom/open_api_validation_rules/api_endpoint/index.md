---
page_title: "api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint"
subcategory: "Load Balancing"
description: "This defines API endpoint."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints fall through mode fall through mode custom open api validation rules api endpoint"], "body_bytes": 5248, "body_sha256": "sha256:a5a7bcf38676a30b68691216628c660e357b7c3b927fbad30c2edbc79eb15b4e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:api_endpoint", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "path": "documentation/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/api_endpoint/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0012223332023113-1301002030011201-3211230101223131-0303322232110312-2323130313210031-3123013220200023-0323220203200223-2333333023201213", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "api_endpoint"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints fall through mode fall through mode custom open api validation rules api endpoint methods"], "anchor": "schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--methods", "description": "Methods to be matched.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "api_endpoint", "methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["api specification validation all spec endpoints fall through mode fall through mode custom open api validation rules api endpoint path"], "anchor": "schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--path", "description": "Path to be matched.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "api_endpoint", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/api_endpoint/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines API endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="section"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

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

## Direct properties

<a id="schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--methods"></a>

### methods property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--path"></a>

### path property

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
