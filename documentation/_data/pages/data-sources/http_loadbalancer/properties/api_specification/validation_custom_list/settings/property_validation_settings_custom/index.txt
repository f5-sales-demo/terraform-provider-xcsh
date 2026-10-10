---
page_title: "api_specification.validation_custom_list.settings.property_validation_settings_custom"
subcategory: "Load Balancing"
description: "Custom property validation settings."
xcsh_docs: {"aliases": ["api specification validation custom list settings property validation settings custom"], "body_bytes": 1698, "body_sha256": "sha256:f791c0742a5089d7961192bfc45d916f9eb24312e1c1830ec6c073458a463750", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings", "path": "documentation/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1023103312220320-1013301300323132-0203031202102332-1313013201030010-1132110001213123-0100302223200011-1203033301223110-1223202222013033", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom"], "schema_version": 1, "sections": [{"aliases": ["api specification validation custom list settings property validation settings custom query parameters"], "anchor": "section", "description": "Custom settings for query parameters validation.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Custom property validation settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.settings.property_validation_settings_custom

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/)
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

- [query_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/): complete subsection reference.
