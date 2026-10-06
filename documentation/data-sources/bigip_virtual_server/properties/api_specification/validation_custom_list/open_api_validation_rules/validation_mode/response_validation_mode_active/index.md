---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active"
subcategory: ""
description: "Open API Validation Mode Active. Validation mode properties of response."
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active"], "body_bytes": 3064, "body_sha256": "sha256:d76d746bbf1911f51165af647410e12a438a5b0611f565984530dcf42b5eb1fc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "path": "documentation/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1210103103123102-1113132131001221-1111100312100011-2200112212020333-0310301320111103-1232230031033322-3310211212233233-3233122122001011", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active enforcement block"], "anchor": "section", "description": "Blocking validation: reject traffic that violates the selected OpenAPI validation properties. Invalid requests are returned as HTTP 403.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "enforcement_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active enforcement report"], "anchor": "section", "description": "Report-only validation: record OpenAPI violations while allowing the request or response to continue.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "enforcement_report"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active response validation properties"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties", "description": "List of properties of the response to validate according to the OpenAPI specification file (a.k.a. Swagger). Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`,", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "response_validation_properties"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Open API Validation Mode Active. Validation mode properties of response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="section"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

## Direct properties

- [enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_block/): complete subsection reference.

- [enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_report/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties"></a>

### response_validation_properties property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.
