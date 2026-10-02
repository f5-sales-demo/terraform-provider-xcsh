---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active"
subcategory: ""
description: "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints validation mode validation mode active"], "body_bytes": 3796, "body_sha256": "sha256:30c45b7143169f2d98ae76158751090d790a912aca9007731d792f35a291b010", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_block", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_report"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:validation_mode", "path": "documentation/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0331133021331323-0210002121010133-1230120012333133-0021131012212013-1000030013322020-2123230123220330-1223220122311313-0023131331122311", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active"], "schema_version": 1, "sections": [{"aliases": ["enforcement block"], "anchor": "section", "description": "Blocking validation: reject traffic that violates the selected OpenAPI validation properties. Invalid requests are returned as HTTP 403.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_block", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active", "enforcement_block"], "syntax": "attribute", "type": "object"}, {"aliases": ["enforcement report"], "anchor": "section", "description": "Report-only validation: record OpenAPI violations while allowing the request or response to continue.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_report", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active", "enforcement_report"], "syntax": "attribute", "type": "object"}, {"aliases": ["request validation properties"], "anchor": "schema-api_specification--validation_all_spec_endpoints--validation_mode--validation_mode_active--request_validation_properties", "description": "List of properties of the request to validate according to the OpenAPI specification file (a.k.a. Swagger). Possible values are `PROPERTY_QUERY_PARAMETERS`, `PROPERTY_PATH_PARAMETERS`, `PROPERTY_CONTENT_TYPE`, `PROPERTY_COOKIE_PARAMETERS`, `PROPERTY_HTTP_HEADERS`, `PROPERTY_HTTP_BODY`, `PROPERTY_SECURITY_SCHEMA`,", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active", "request_validation_properties"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="section"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

## Direct properties

- [enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_block/): complete subsection reference.

- [enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_report/): complete subsection reference.

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

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_block/)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_report/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
