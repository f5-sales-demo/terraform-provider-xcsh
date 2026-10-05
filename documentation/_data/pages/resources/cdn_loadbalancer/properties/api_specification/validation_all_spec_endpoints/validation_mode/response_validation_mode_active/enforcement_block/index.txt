---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block"
subcategory: "Load Balancing"
description: "Blocking validation: reject traffic that violates the selected OpenAPI validation properties. Invalid requests are returned as HTTP 403."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints validation mode response validation mode active enforcement block"], "body_bytes": 2371, "body_sha256": "sha256:15e6211d7e495ac5ecc006febee64bcaa3b4714072c639104855d6b7086db66e", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active:enforcement_block", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/enforcement_block/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0210302333310310-1231311212222012-3133320013101221-1121211230120001-1130311210323112-0033122313133313-2323310223321013-3033132303130320", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "response_validation_mode_active", "enforcement_block"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/enforcement_block/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Blocking validation: reject traffic that violates the selected OpenAPI validation properties. Invalid requests are returned as HTTP 403.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="section"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
enforcement_block = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
