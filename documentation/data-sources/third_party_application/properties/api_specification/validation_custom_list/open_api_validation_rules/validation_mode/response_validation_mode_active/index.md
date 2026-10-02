---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active"
subcategory: ""
description: "Open API Validation Mode Active. Validation mode properties of response."
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active"], "body_bytes": 4265, "body_sha256": "sha256:c201497fa72d0ae05a403092200df7dc4ef65d6567a7b3bc0e85fd56eba78a1f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "path": "documentation/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0213130110300332-0132220202230230-1213013012221023-2021320130020120-1321030221222013-3213320112201333-3111021302230031-1303231133323101", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active"], "schema_version": 1, "sections": [{"aliases": ["enforcement block"], "anchor": "section", "description": "Blocking validation: reject traffic that violates the selected OpenAPI validation properties. Invalid requests are returned as HTTP 403.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "enforcement_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["enforcement report"], "anchor": "section", "description": "Report-only validation: record OpenAPI violations while allowing the request or response to continue.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "enforcement_report"], "syntax": "attribute", "type": "object"}, {"aliases": ["response validation properties"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties", "description": "List of properties of the response to validate according to the OpenAPI specification file (a.k.a. Swagger). Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`, `PRO", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active", "response_validation_properties"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Open API Validation Mode Active. Validation mode properties of response.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="section"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

## Direct properties

- [enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_block/): complete subsection reference.

- [enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_report/): complete subsection reference.

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

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_block/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/enforcement_report/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
