---
page_title: "api_specification.validation_custom_list.open_api_validation_rules"
subcategory: ""
description: "Validation List. Rule or policy definition"
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules"], "body_bytes": 2635, "body_sha256": "sha256:d471fec331332224df555a170b9d337494d662e7814b9d05c7ef3839ef7e8efe", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:any_domain", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:metadata", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list", "path": "documentation/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1203032213021212-2233221032301313-2011101010210023-0102233213323110-2231122103031132-1303103013210320-1021211221002020-3113231110321300", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list open api validation rules any domain"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:any_domain", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules api endpoint"], "anchor": "section", "description": "API Endpoint. This defines API endpoint.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:api_endpoint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules api group"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--api_group", "description": "Exclusive with The API group which this validation applies to.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "api_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["api specification validation custom list open api validation rules base path"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--base_path", "description": "Exclusive with The base path which this validation applies to.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "base_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["api specification validation custom list open api validation rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create..", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:metadata", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules specific domain"], "anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain", "description": "Exclusive with The rule will apply for a specific domain.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "specific_domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["api specification validation custom list open api validation rules validation mode"], "anchor": "section", "description": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Validation List. Rule or policy definition", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="section"></a>

Type: `"list"`. Computed.

Validation List. Rule or policy definition

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/any_domain/): complete subsection reference.

- [api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/api_endpoint/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--api_group"></a>

### api_group property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--base_path"></a>

### base_path property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/metadata/): complete subsection reference.

<a id="schema-api_specification--validation_custom_list--open_api_validation_rules--specific_domain"></a>

### specific_domain property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

- [validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/): complete subsection reference.
