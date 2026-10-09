---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode"
subcategory: "Load Balancing"
description: "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)"
xcsh_docs: {"aliases": ["api specification validation all spec endpoints validation mode"], "body_bytes": 2465, "body_sha256": "sha256:965e0860549d468b5e1b60e7d5a45547af40281f366219e9366dcbe36dfd36b0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_validation", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "path": "documentation/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints validation mode response validation mode active"], "anchor": "section", "description": "Validation mode properties of response.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "response_validation_mode_active"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode skip response validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "skip_response_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode skip validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_validation", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "skip_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode validation mode active"], "anchor": "section", "description": "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="section"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

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

## Direct properties

- [response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/): complete subsection reference.

- [skip_response_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_response_validation/): complete subsection reference.

- [skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_validation/): complete subsection reference.

- [validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/): complete subsection reference.
