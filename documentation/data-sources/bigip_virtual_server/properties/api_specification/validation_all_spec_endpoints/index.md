---
page_title: "api_specification.validation_all_spec_endpoints"
subcategory: ""
description: "API Inventory. Settings for API Inventory validation."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints"], "body_bytes": 2359, "body_sha256": "sha256:224f4965b5b4489f0092e8c5235e10cc2c233d7206fdea77d9f13c3d776f9410", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:validation_mode"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification", "path": "documentation/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3301110110012032-3220203300202303-0201103332211020-0212003103000302-3031313231231221-3211331322000212-2222001230232123-2120202203300330", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints fall through mode"], "anchor": "section", "description": "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules).", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints settings"], "anchor": "section", "description": "OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode"], "anchor": "section", "description": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:validation_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "API Inventory. Settings for API Inventory validation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/)
- api_specification.validation_all_spec_endpoints

<a id="section"></a>

Type: `"single"`. Computed.

API Inventory. Settings for API Inventory validation.

## Direct properties

- [fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/): complete subsection reference.

- [settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/): complete subsection reference.

- [validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/validation_mode/): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/)
- [api_specification.validation_all_spec_endpoints.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/)
- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
