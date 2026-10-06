---
page_title: "api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report"
subcategory: "Load Balancing"
description: "Report-only validation: record OpenAPI violations while allowing the request or response to continue."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints validation mode validation mode active enforcement report"], "body_bytes": 1819, "body_sha256": "sha256:a237973b3422b42a19d31e9699dd247b2c5409b9eae80acef9bce5629822cab2", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active:enforcement_report", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:validation_mode:validation_mode_active", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_report/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0313311012111011-3032013003230330-1212000122022311-0300000002032031-2030031302100122-0102032312303202-1213221222120201-1321010331013331", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "validation_mode", "validation_mode_active", "enforcement_report"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/enforcement_report/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Report-only validation: record OpenAPI violations while allowing the request or response to continue.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.validation_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/validation_mode/validation_mode_active/)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="section"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.
