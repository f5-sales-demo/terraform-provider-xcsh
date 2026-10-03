---
page_title: "api_specification.validation_custom_list.open_api_validation_rules.validation_mode"
subcategory: "Load Balancing"
description: "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)"
xcsh_docs: {"aliases": ["api specification validation custom list open api validation rules validation mode"], "body_bytes": 5015, "body_sha256": "sha256:01da2675383ae0a8b5deff04643f9f2988b45854ec0397a24ac56ce456c22993", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1300210222332112-3202232332012200-0000013303132120-2201221030000121-1230122120233211-1001120020312010-0300212200032021-0203310121102333", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:response_validation_mode_active,skip_response_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:response_validation_mode_active,skip_response_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:skip_validation,validation_mode_active", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode:ConflictingObjectAttributes:skip_validation,validation_mode_active", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list open api validation rules validation mode response validation mode active"], "anchor": "section", "description": "Validation mode properties of response.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active:ConflictingObjectAttributes:enforcement_block,enforcement_report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active:ConflictingObjectAttributes:enforcement_block,enforcement_report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active:enforcement_report", "type": "conflicts"}, {"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--response_validation_mode_active--response_validation_properties", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active:RequiredObjectAttributes:response_validation_properties", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:response_validation_mode_active", "type": "requires"}], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "response_validation_mode_active"], "syntax": "block", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode skip response validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_response_validation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "skip_response_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode skip validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:skip_validation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "skip_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation custom list open api validation rules validation mode validation mode active"], "anchor": "section", "description": "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active:ConflictingObjectAttributes:enforcement_block,enforcement_report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active:enforcement_block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active:ConflictingObjectAttributes:enforcement_block,enforcement_report", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active:enforcement_report", "type": "conflicts"}, {"anchor": "schema-api_specification--validation_custom_list--open_api_validation_rules--validation_mode--validation_mode_active--request_validation_properties", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active:RequiredObjectAttributes:request_validation_properties", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules:validation_mode:validation_mode_active", "type": "requires"}], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules", "validation_mode", "validation_mode_active"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.open_api_validation_rules.validation_mode

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/)
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

- [response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/): complete subsection reference.

- [skip_response_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_response_validation/): complete subsection reference.

- [skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_validation/): complete subsection reference.

- [validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/response_validation_mode_active/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_response_validation/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/skip_validation/)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/validation_mode/validation_mode_active/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
