---
page_title: "api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint"
subcategory: ""
description: "API Endpoint. This defines API endpoint."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints fall through mode fall through mode custom open api validation rules api endpoint"], "body_bytes": 3194, "body_sha256": "sha256:bfbfdbb45be326659edf142ec13e34631b61fbaed60e2f0abed3abf6fb9d0467", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:api_endpoint", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "path": "documentation/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/api_endpoint/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3200300222312333-2323221120220012-0320131322123012-1112021323230232-2123302110011031-2030332330121212-3211113122313202-1003030213201201", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "api_endpoint"], "schema_version": 1, "sections": [{"aliases": ["methods"], "anchor": "schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--methods", "description": "Methods. Methods to be matched. Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`. Defaults to `ANY`.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "api_endpoint", "methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["path"], "anchor": "schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--path", "description": "Path. Path to be matched.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom", "open_api_validation_rules", "api_endpoint", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/api_endpoint/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "API Endpoint. This defines API endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="section"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

## Direct properties

<a id="schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--methods"></a>

### methods property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

<a id="schema-api_specification--validation_all_spec_endpoints--fall_through_mode--fall_through_mode_custom--open_api_validation_rules--api_endpoint--path"></a>

### path property

Type: `"string"`. Computed.

Path. Path to be matched.

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/open_api_validation_rules/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
