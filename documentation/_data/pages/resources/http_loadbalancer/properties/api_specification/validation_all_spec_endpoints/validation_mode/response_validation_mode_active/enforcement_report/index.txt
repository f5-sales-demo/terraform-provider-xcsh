---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report"
subcategory: "Load Balancing"
description: "Report-only validation: record OpenAPI violations while allowing the request or response to continue."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints validation mode response validation mode active enforcement report"], "body_bytes": 2349, "body_sha256": "sha256:979b6a7e423c5546a09a2e3f0bd17ba8ce0eda1eadb7ecd9c2b3843bd2230d6e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active:enforcement_report", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:response_validation_mode_active", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/enforcement_report/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1013012111113200-2213022201012030-2223033120201113-2311311012300130-3030103320001200-0313013100220103-2110200222311212-0220100302030222", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "response_validation_mode_active", "enforcement_report"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/enforcement_report/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Report-only validation: record OpenAPI violations while allowing the request or response to continue.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="section"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

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
enforcement_report = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/response_validation_mode_active/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
