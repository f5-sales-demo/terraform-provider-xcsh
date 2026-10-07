---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.api_endpoint"
subcategory: ""
description: "API Endpoint. This defines API endpoint."
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules api endpoint"], "body_bytes": 1836, "body_sha256": "sha256:4df62848137c3cb225650175dae1c719c8093ed225bc6f797d1b962599818378", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules", "path": "documentation/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3212033202223010-1022002030122221-2021031203333123-1200213123322120-2102111110300020-1303213300022330-0003021012331022-3122302331203232", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list open api validation rules api endpoint methods"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--methods", "description": "Methods. Methods to be matched. Possible values are `ANY`, `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`, `COPY`. Defaults to `ANY`.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint", "methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["api specification validation custom list open api validation rules api endpoint path"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--path", "description": "Path. Path to be matched.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint", "path"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "API Endpoint. This defines API endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="section"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

## Direct properties

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--methods"></a>

### methods property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--api_endpoint--path"></a>

### path property

Type: `"string"`. Computed.

Path. Path to be matched.
