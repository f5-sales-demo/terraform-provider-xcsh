---
page_title: "api_specification.validation_all_spec_endpoints"
subcategory: "Load Balancing"
description: "Settings for API Inventory validation."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints"], "body_bytes": 2714, "body_sha256": "sha256:14106d2dd9e171212be067734644173e11db6666ccaeec990c591c8a22f9516f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0220101021233310-1301213101111000-0022132033301311-1132223130021012-2232220303133123-0200303113033230-0010202001131210-3120323030011130", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints"], "schema_version": 1, "sections": [{"aliases": ["fall through mode"], "anchor": "section", "description": "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.fall_through_mode:ConflictingObjectAttributes:fall_through_mode_allow,fall_through_mode_custom", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.fall_through_mode:ConflictingObjectAttributes:fall_through_mode_allow,fall_through_mode_custom", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom", "type": "conflicts"}], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode"], "syntax": "block", "type": "object"}, {"aliases": ["settings"], "anchor": "section", "description": "OpenAPI specification validation settings relevant for \"API Inventory\" enforcement and for \"Custom list\" enforcement.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.settings:ConflictingObjectAttributes:oversized_body_fail_validation,oversized_body_skip_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_fail_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.settings:ConflictingObjectAttributes:oversized_body_fail_validation,oversized_body_skip_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_skip_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.settings:ConflictingObjectAttributes:property_validation_settings_custom,property_validation_settings_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.settings:ConflictingObjectAttributes:property_validation_settings_custom,property_validation_settings_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_default", "type": "conflicts"}], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings"], "syntax": "block", "type": "object"}, {"aliases": ["validation mode"], "anchor": "section", "description": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.validation_mode:ConflictingObjectAttributes:response_validation_mode_active,skip_response_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.validation_mode:ConflictingObjectAttributes:response_validation_mode_active,skip_response_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.validation_mode:ConflictingObjectAttributes:skip_validation,validation_mode_active", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.validation_mode:ConflictingObjectAttributes:skip_validation,validation_mode_active", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "type": "conflicts"}], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Settings for API Inventory validation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- api_specification.validation_all_spec_endpoints

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

API Inventory. Settings for API Inventory validation.

Upstream description:

Settings for API Inventory validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

Terraform syntax:

```terraform
validation_all_spec_endpoints {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/): complete subsection reference.

- [settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/): complete subsection reference.

- [validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/)
- [api_specification.validation_all_spec_endpoints.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
