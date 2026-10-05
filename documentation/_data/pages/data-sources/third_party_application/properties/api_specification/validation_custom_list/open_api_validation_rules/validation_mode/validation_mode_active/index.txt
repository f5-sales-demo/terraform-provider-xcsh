---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active"
subcategory: ""
description: "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403."
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules validation mode validation mode active"], "body_bytes": 4274, "body_sha256": "sha256:1335e4e3ff1496a6f81c149f271b9400aec84e9c5f57f0e7fe1080c056f6fdac", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active:enforcement_block", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active:enforcement_report"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "path": "documentation/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2021331333103120-1211233012210110-1213010111300100-3320103120022323-1203203330302010-3001211202203220-1003232022001203-1221212312122030", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "validation_mode_active"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list open api validation rules validation mode validation mode active enforcement block"], "anchor": "section", "description": "Blocking validation: reject traffic that violates the selected OpenAPI validation properties. Invalid requests are returned as HTTP 403.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active:enforcement_block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "validation_mode_active", "enforcement_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode validation mode active enforcement report"], "anchor": "section", "description": "Report-only validation: record OpenAPI violations while allowing the request or response to continue.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active:enforcement_report", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "validation_mode_active", "enforcement_report"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode validation mode active request validation properties"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active--request_validation_properties", "description": "List of properties of the request to validate according to the OpenAPI specification file (a.k.a. Swagger). Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`,", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "validation_mode_active", "request_validation_properties"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="section"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

## Direct properties

- [enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/enforcement_block/): complete subsection reference.

- [enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/enforcement_report/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active--request_validation_properties"></a>

### request_validation_properties property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/enforcement_block/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/enforcement_report/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
