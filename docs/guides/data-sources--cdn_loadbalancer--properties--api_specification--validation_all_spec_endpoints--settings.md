---
page_title: "api_specification.validation_all_spec_endpoints.settings"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.settings for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3447, "body_sha256": "sha256:2fe3d3faa2fe71499dbbd6e689012f46f57e544bbaefa565ade3825f0844541c", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_fail_validation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:oversized_body_skip_validation", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_default"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.settings for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.settings

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [api_specification](data-sources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- api_specification.validation_all_spec_endpoints.settings

<a id="section"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

## Direct properties

- [oversized_body_fail_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--oversized_body_fail_validation.md): complete subsection reference.

- [oversized_body_skip_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--oversized_body_skip_validation.md): complete subsection reference.

- [property_validation_settings_custom](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom.md): complete subsection reference.

- [property_validation_settings_default](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_default.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--oversized_body_fail_validation.md)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--oversized_body_skip_validation.md)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom.md)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_default.md)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
