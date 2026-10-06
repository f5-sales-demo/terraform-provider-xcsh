---
page_title: "api_specification.validation_custom_list.fall_through_mode"
subcategory: ""
description: "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)."
xcsh_docs: {"aliases": ["api specification validation custom list fall through mode"], "body_bytes": 1540, "body_sha256": "sha256:20f42486a34c4e07e39c35676ca4f0747d995c01a1c2591ff7bd17727fb57790", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_allow", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:fall_through_mode", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list", "path": "documentation/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/fall_through_mode/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1203203121111030-0010121322013103-2210101322101031-0001112231033223-1320002223232310-2013321323302002-1320111112102322-1130232323130213", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list fall through mode fall through mode allow"], "anchor": "section", "description": "Configuration parameter for fall through mode allow.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_allow", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list fall through mode fall through mode custom"], "anchor": "section", "description": "Configuration parameter for fall through mode custom.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode", "fall_through_mode_custom"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/fall_through_mode/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.fall_through_mode

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/)
- api_specification.validation_custom_list.fall_through_mode

<a id="section"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

## Direct properties

- [fall_through_mode_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_allow/): complete subsection reference.

- [fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/fall_through_mode/fall_through_mode_custom/): complete subsection reference.
