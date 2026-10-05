---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode"
subcategory: "Load Balancing"
description: "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)"
xcsh_docs: {"aliases": ["api specification validation all spec endpoints validation mode"], "body_bytes": 4096, "body_sha256": "sha256:eef454d0e7932800db806fdb5ad6639c5d61b4a8f031a8ce15897db245bad144", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_validation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints validation mode response validation mode active"], "anchor": "section", "description": "Validation mode properties of response.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "response_validation_mode_active"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode skip response validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_response_validation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "skip_response_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode skip validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:skip_validation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "skip_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints validation mode validation mode active"], "anchor": "section", "description": "Enable OpenAPI validation and explicitly select enforcement_report to allow and log invalid traffic, or enforcement_block to reject invalid requests with HTTP 403.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="section"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

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

- [response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/): complete subsection reference.

- [skip_response_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_response_validation/): complete subsection reference.

- [skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_validation/): complete subsection reference.

- [validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_response_validation/)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/skip_validation/)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
