---
page_title: "api_specification"
subcategory: ""
description: "Settings for API specification (API definition, OpenAPI validation, etc.)."
xcsh_docs: {"aliases": ["api specification"], "body_bytes": 2413, "body_sha256": "sha256:12e5abf0a4c300ee9bf0930706564195245ca536229cadaf7e20b5fe5bfd39bc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:api_specification:api_definition", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list", "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_disabled"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:api_specification", "parent_id": "xcsh-docs:data-sources:third_party_application:reference", "path": "documentation/data-sources/third_party_application/properties/api_specification/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2122110330120311-1113110313011022-3302311001022320-0132212230220103-1320022230211320-3213132221313112-3232230332321003-3200002023102003", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification"], "schema_version": 1, "sections": [{"aliases": ["api definition"], "anchor": "section", "description": "Type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:api_definition", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "api_definition"], "syntax": "attribute", "type": "object"}, {"aliases": ["validation all spec endpoints"], "anchor": "section", "description": "API Inventory. Settings for API Inventory validation.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_all_spec_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["validation custom list"], "anchor": "section", "description": "Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to 'Fall Through Mode'.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_custom_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["validation disabled"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:third_party_application:properties:api_specification:validation_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_disabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/api_specification/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Settings for API specification (API definition, OpenAPI validation, etc.).", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- api_specification

<a id="section"></a>

Type: `"single"`. Computed.

Settings for API specification (API definition, OpenAPI validation, etc.).

## Direct properties

- [api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/api_definition/): complete subsection reference.

- [validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/): complete subsection reference.

- [validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/): complete subsection reference.

- [validation_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_disabled/): complete subsection reference.

## Next pages

- [api_specification.api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/api_definition/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_custom_list/)
- [api_specification.validation_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/api_specification/validation_disabled/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
