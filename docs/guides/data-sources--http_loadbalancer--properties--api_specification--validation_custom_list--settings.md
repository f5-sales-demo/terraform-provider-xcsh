---
page_title: "api_specification.validation_custom_list.settings"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.settings for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3238, "body_sha256": "sha256:5e4cd0c4a378c977929521cf372d194e9e983080725424f0ec6ba87d82007e5e", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_fail_validation", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_skip_validation", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_default"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list:settings", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_specification:validation_custom_list", "path": "docs/guides/data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.settings for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_custom_list.settings

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [api_specification](data-sources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- api_specification.validation_custom_list.settings

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

- [oversized_body_fail_validation](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--oversized_body_fail_validation.md): complete subsection reference.

- [oversized_body_skip_validation](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--oversized_body_skip_validation.md): complete subsection reference.

- [property_validation_settings_custom](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom.md): complete subsection reference.

- [property_validation_settings_default](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_default.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.settings.oversized_body_fail_validation](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--oversized_body_fail_validation.md)
- [api_specification.validation_custom_list.settings.oversized_body_skip_validation](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--oversized_body_skip_validation.md)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom.md)
- [api_specification.validation_custom_list.settings.property_validation_settings_default](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_default.md)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--properties--api_specification--validation_custom_list.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
