---
page_title: "api_specification.validation_all_spec_endpoints.settings"
subcategory: ""
description: "OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints settings"], "body_bytes": 3571, "body_sha256": "sha256:8fef8f75abc2fac6c0abe55060368db2ae3a34923d89cc41218db4621d6c8a72", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_fail_validation", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_skip_validation", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_default"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings", "parent_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints", "path": "documentation/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3311220020032133-3000201111210133-0321112310321021-2102221123011323-0103220011020332-2021023122320130-1223133102300110-2330100200223021", "registry_path": "docs/guides/data-sources--bigip_virtual_server--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings"], "schema_version": 1, "sections": [{"aliases": ["oversized body fail validation"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_fail_validation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "oversized_body_fail_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["oversized body skip validation"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_skip_validation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "oversized_body_skip_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["property validation settings custom"], "anchor": "section", "description": "Configuration parameter for property validation settings custom.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom"], "syntax": "attribute", "type": "object"}, {"aliases": ["property validation settings default"], "anchor": "section", "description": "Configuration parameter for property validation settings default.", "document_id": "xcsh-docs:data-sources:bigip_virtual_server:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_default", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_default"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom list' enforcement.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.settings

Breadcrumbs:

- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/)
- api_specification.validation_all_spec_endpoints.settings

<a id="section"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

## Direct properties

- [oversized_body_fail_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/oversized_body_fail_validation/): complete subsection reference.

- [oversized_body_skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/oversized_body_skip_validation/): complete subsection reference.

- [property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/): complete subsection reference.

- [property_validation_settings_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_default/): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/oversized_body_fail_validation/)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/oversized_body_skip_validation/)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_default/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/api_specification/validation_all_spec_endpoints/)
- [xcsh_bigip_virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/)
