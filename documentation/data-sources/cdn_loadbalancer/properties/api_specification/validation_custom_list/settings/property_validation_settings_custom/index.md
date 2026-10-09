---
page_title: "api_specification.validation_custom_list.settings.property_validation_settings_custom"
subcategory: "Load Balancing"
description: "Custom property validation settings."
xcsh_docs: {"aliases": ["api specification validation custom list settings property validation settings custom"], "body_bytes": 1691, "body_sha256": "sha256:745d0cafbd943f07a5c432a2f7df2b3da76a49ed9bf6d608f49ba090c5174e91", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2201103103321303-1110332012011022-1322311310322203-0103201322210120-0110301003313200-0030310003002012-0203203113322212-3233031010220113", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list settings property validation settings custom query parameters"], "anchor": "section", "description": "Custom settings for query parameters validation.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Custom property validation settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.settings.property_validation_settings_custom

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/)
- api_specification.validation_custom_list.settings.property_validation_settings_custom

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

- [query_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/): complete subsection reference.
