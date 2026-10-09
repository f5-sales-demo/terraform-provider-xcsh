---
page_title: "api_specification.validation_all_spec_endpoints.settings"
subcategory: "Load Balancing"
description: "OpenAPI specification validation settings relevant for \"API Inventory\" enforcement and for \"Custom list\" enforcement."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints settings"], "body_bytes": 2779, "body_sha256": "sha256:a3dd98282a91974bd2d194f790e23be09011ded28ae6cf5c1cdb54f04729956a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_fail_validation", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_skip_validation", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_default"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2331331310030110-3011131313323200-3232121003112200-3101232023222123-0033102102300313-0323132303001233-0320103003121030-2022001010233201", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints settings oversized body fail validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_fail_validation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "oversized_body_fail_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints settings oversized body skip validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_skip_validation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "oversized_body_skip_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints settings property validation settings custom"], "anchor": "section", "description": "Custom property validation settings.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom"], "syntax": "block", "type": "object"}, {"aliases": ["api specification validation all spec endpoints settings property validation settings default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_default", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_default"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "OpenAPI specification validation settings relevant for \"API Inventory\" enforcement and for \"Custom list\" enforcement.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.settings

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- api_specification.validation_all_spec_endpoints.settings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Additional upstream details:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [oversized_body_fail_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/oversized_body_fail_validation/): complete subsection reference.

- [oversized_body_skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/oversized_body_skip_validation/): complete subsection reference.

- [property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/): complete subsection reference.

- [property_validation_settings_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_default/): complete subsection reference.
