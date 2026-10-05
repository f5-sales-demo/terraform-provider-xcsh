---
page_title: "api_specification"
subcategory: ""
description: "Settings for API specification (API definition, OpenAPI validation, etc.)."
xcsh_docs: {"aliases": ["api specification"], "body_bytes": 2371, "body_sha256": "sha256:b5131394f449ccd3d28ecd98c8e146e6d32a770cfd636abdcbb92483db1dbbdb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:api_definition", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_disabled"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:reference", "path": "documentation/data-sources/bigip_virtual_server/properties/api_specification/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2213121213201103-1311331110220313-3223333313332233-2301330300100331-1322321323010131-2310201211203221-1033320121012323-1332031232323233", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification"], "schema_version": 1, "sections": [{"aliases": ["api specification api definition"], "anchor": "section", "description": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:api_definition", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "api_definition"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints"], "anchor": "section", "description": "API Inventory. Settings for API Inventory validation.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list"], "anchor": "section", "description": "Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to 'Fall Through Mode'.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_custom_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation disabled"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_disabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/api_specification/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Settings for API specification (API definition, OpenAPI validation, etc.).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- api_specification

<a id="section"></a>

Type: `"single"`. Computed.

Settings for API specification (API definition, OpenAPI validation, etc.).

## Direct properties

- [api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/api_definition/): complete subsection reference.

- [validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/): complete subsection reference.

- [validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/): complete subsection reference.

- [validation_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_disabled/): complete subsection reference.

## Next pages

- [api_specification.api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/api_definition/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_custom_list/)
- [api_specification.validation_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_disabled/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
