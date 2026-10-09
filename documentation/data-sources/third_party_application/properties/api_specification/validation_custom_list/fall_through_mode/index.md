---
page_title: "api_specification.validation_custom_list.fall_through_mode"
subcategory: ""
description: "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)."
xcsh_docs: {"aliases": ["api specification validation custom list fall through mode"], "body_bytes": 1561, "body_sha256": "sha256:b5b9af9beef83f9b1b584ba650a30337e50cfb8f3eb00026754888ab16bf4884", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_allow", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:fall_through_mode", "parent_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list", "path": "documentation/data-sources/third_party_application/properties/api_specification/validation_custom_list/fall_through_mode/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1333310331131232-1101311110232110-3212300202103312-3321223001333302-3030213012112200-3312003011013122-3233313111100311-3003211220320122", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list fall through mode fall through mode allow"], "anchor": "section", "description": "Configuration parameter for fall through mode allow.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_allow", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list fall through mode fall through mode custom"], "anchor": "section", "description": "Configuration parameter for fall through mode custom.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/validation_custom_list/fall_through_mode/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.fall_through_mode

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/)
- api_specification.validation_custom_list.fall_through_mode

<a id="section"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

## Direct properties

- [fall_through_mode_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_allow/): complete subsection reference.

- [fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/): complete subsection reference.
