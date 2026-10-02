---
page_title: "enable_api_discovery"
subcategory: ""
description: "Specifies the settings used for API discovery."
xcsh_docs: {"aliases": ["enable api discovery"], "body_bytes": 3779, "body_sha256": "sha256:58124706cc49ccfa489004cf1975a45294e0545fc60df39c620f0f0215c522bf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler", "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_discovery_from_code_scan", "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:custom_api_auth_discovery", "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:default_api_auth_discovery", "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:discovered_api_settings", "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:enable_learn_from_redirect_traffic"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery", "parent_id": "xcsh-docs:data-sources:third_party_application:reference", "path": "documentation/data-sources/third_party_application/properties/enable_api_discovery/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2212031120030331-0231323202213321-1332200003223030-1332331123331123-2331300130313130-0123201310233213-0002330203220331-3323111033133032", "registry_path": "docs/guides/data-sources--third_party_application--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery"], "schema_version": 1, "sections": [{"aliases": ["api crawler"], "anchor": "section", "description": "API Crawling. API Crawler message.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_crawler", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_crawler"], "syntax": "attribute", "type": "object"}, {"aliases": ["api discovery from code scan"], "anchor": "section", "description": "Select Code Base and Repositories.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:api_discovery_from_code_scan", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom api auth discovery"], "anchor": "section", "description": "API Discovery Advanced Settings. API Discovery Advanced settings.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:custom_api_auth_discovery", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "custom_api_auth_discovery"], "syntax": "attribute", "type": "object"}, {"aliases": ["default api auth discovery"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:default_api_auth_discovery", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "default_api_auth_discovery"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable learn from redirect traffic"], "anchor": "section", "description": "Configuration parameter for disable learn from redirect traffic.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:disable_learn_from_redirect_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "disable_learn_from_redirect_traffic"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovered api settings"], "anchor": "section", "description": "Discovered API Settings. Configure Discovered API Settings.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:discovered_api_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_api_discovery", "discovered_api_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable learn from redirect traffic"], "anchor": "section", "description": "Configuration parameter for enable learn from redirect traffic.", "document_id": "xcsh-docs:data-sources:third_party_application:properties:enable_api_discovery:enable_learn_from_redirect_traffic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enable_api_discovery", "enable_learn_from_redirect_traffic"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/properties/enable_api_discovery/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies the settings used for API discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- enable_api_discovery

<a id="section"></a>

Type: `"single"`. Computed.

Specifies the settings used for API discovery.

## Direct properties

- [api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/): complete subsection reference.

- [api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_discovery_from_code_scan/): complete subsection reference.

- [custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/custom_api_auth_discovery/): complete subsection reference.

- [default_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/default_api_auth_discovery/): complete subsection reference.

- [disable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/disable_learn_from_redirect_traffic/): complete subsection reference.

- [discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/discovered_api_settings/): complete subsection reference.

- [enable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/enable_learn_from_redirect_traffic/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_crawler](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_crawler/)
- [enable_api_discovery.api_discovery_from_code_scan](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/api_discovery_from_code_scan/)
- [enable_api_discovery.custom_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/custom_api_auth_discovery/)
- [enable_api_discovery.default_api_auth_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/default_api_auth_discovery/)
- [enable_api_discovery.disable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/disable_learn_from_redirect_traffic/)
- [enable_api_discovery.discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/discovered_api_settings/)
- [enable_api_discovery.enable_learn_from_redirect_traffic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/enable_api_discovery/enable_learn_from_redirect_traffic/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
