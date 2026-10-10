---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active"
subcategory: "Load Balancing"
description: "Validation mode properties of response."
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active"], "body_bytes": 4534, "body_sha256": "sha256:0555499f3c4557e5f5a2a2a024c5a8bf2516442c2595f9a8731c8c15f274e6c7", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active enforcement block"], "anchor": "section", "description": "Blocking validation: reject traffic that violates the selected OpenAPI validation properties. Invalid requests are returned as HTTP 403.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "enforcement_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active enforcement report"], "anchor": "section", "description": "Report-only validation: record OpenAPI violations while allowing the request or response to continue.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "enforcement_report"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active response validation properties"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties", "description": "List of properties of the response to validate according to the OpenAPI specification file (a.k.a. Swagger)", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "response_validation_properties"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Validation mode properties of response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

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

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_block/): complete subsection reference.

- [enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_report/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties"></a>

### response_validation_properties property

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
