---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode"
subcategory: "Load Balancing"
description: "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)"
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules validation mode"], "body_bytes": 4998, "body_sha256": "sha256:1381d35a3d050b2febe858a6a3afa031d46ede4215d32ca92523019e21bcfd03", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:response_validation_mode_active,skip_response_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:response_validation_mode_active,skip_response_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:skip_validation,validation_mode_active", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:skip_validation,validation_mode_active", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode"], "schema_version": 1, "sections": [{"aliases": ["response validation mode active"], "anchor": "section", "description": "Validation mode properties of response.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active:ConflictingObjectAttributes:enforcement_block,enforcement_report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active:ConflictingObjectAttributes:enforcement_block,enforcement_report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report", "type": "conflicts"}, {"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active:RequiredObjectAttributes:response_validation_properties", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "type": "requires"}], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active"], "syntax": "block", "type": "object"}, {"aliases": ["skip response validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "skip_response_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["skip validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "skip_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["validation mode active"], "anchor": "section", "description": "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active:ConflictingObjectAttributes:enforcement_block,enforcement_report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active:enforcement_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active:ConflictingObjectAttributes:enforcement_block,enforcement_report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active:enforcement_report", "type": "conflicts"}, {"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active--request_validation_properties", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active:RequiredObjectAttributes:request_validation_properties", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "type": "requires"}], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "validation_mode_active"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("response_validation_mode_active",
    "skip_response_validation"),
  validators.ConflictingObjectAttributes("skip_validation",
    "validation_mode_active")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/): complete subsection reference.

- [skip_response_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_response_validation/): complete subsection reference.

- [skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_validation/): complete subsection reference.

- [validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_response_validation/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_validation/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
