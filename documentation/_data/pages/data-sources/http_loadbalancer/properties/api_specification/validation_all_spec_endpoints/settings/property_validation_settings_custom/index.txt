---
page_title: "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom"
subcategory: "Load Balancing"
description: "Custom property validation settings."
xcsh_docs: {"aliases": ["api specification validation all spec endpoints settings property validation settings custom"], "body_bytes": 1747, "body_sha256": "sha256:423bdb3113bca57a1cd9e986e1e302aae775a93d22a3f29cdf08bbd2cb46f4ef", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "path": "documentation/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints settings property validation settings custom query parameters"], "anchor": "section", "description": "Custom settings for query parameters validation.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom", "query_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Custom property validation settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_all_spec_endpoints.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for property validation settings custom.

Additional upstream details:

Custom property validation settings.

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

## Direct properties

- [query_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/): complete subsection reference.
