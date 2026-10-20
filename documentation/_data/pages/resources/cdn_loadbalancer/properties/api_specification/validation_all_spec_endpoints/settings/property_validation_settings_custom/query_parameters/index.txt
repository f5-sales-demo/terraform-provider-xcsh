---
page_title: "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters"
subcategory: "Load Balancing"
description: "Custom settings for query parameters validation."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints settings property validation settings custom query parameters"], "body_bytes": 2252, "body_sha256": "sha256:7b0bce4926c0e947d8a79335d2bb6e450dabc0c5dbc1a9e18297e8ce6161f025", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom", "query_parameters"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints settings property validation settings custom query parameters allow additional parameters"], "anchor": "section", "description": "Configuration parameter for allow additional parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom", "query_parameters", "allow_additional_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints settings property validation settings custom query parameters disallow additional parameters"], "anchor": "section", "description": "Configuration parameter for disallow additional parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom", "query_parameters", "disallow_additional_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Custom settings for query parameters validation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/allow_additional_parameters/): complete subsection reference.

- [disallow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/disallow_additional_parameters/): complete subsection reference.
