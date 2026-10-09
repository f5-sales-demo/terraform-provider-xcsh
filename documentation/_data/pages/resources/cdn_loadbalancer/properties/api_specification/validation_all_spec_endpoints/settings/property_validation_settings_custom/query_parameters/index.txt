---
page_title: "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters"
subcategory: "Load Balancing"
description: "Custom settings for query parameters validation."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints settings property validation settings custom query parameters"], "body_bytes": 2489, "body_sha256": "sha256:1c8391828b60bda5d7c9a38c7740f47d9d6a891c47297c1c14dfa1902b13fab5", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters:ConflictingObjectAttributes:allow_additional_parameters,disallow_additional_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters:ConflictingObjectAttributes:allow_additional_parameters,disallow_additional_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom", "query_parameters"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints settings property validation settings custom query parameters allow additional parameters"], "anchor": "section", "description": "Configuration parameter for allow additional parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom", "query_parameters", "allow_additional_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints settings property validation settings custom query parameters disallow additional parameters"], "anchor": "section", "description": "Configuration parameter for disallow additional parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom", "query_parameters", "disallow_additional_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Custom settings for query parameters validation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow_additional_parameters",
    "disallow_additional_parameters")}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/allow_additional_parameters/): complete subsection reference.

- [disallow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/disallow_additional_parameters/): complete subsection reference.
