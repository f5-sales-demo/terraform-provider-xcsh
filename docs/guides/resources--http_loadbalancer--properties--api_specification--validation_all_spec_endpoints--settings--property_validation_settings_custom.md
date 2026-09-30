---
page_title: "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1984, "body_sha256": "sha256:2baa6b3ba25a295a270a87d469073d61b3652e7ae6c0ccaf4d4cc782fba34e3e", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for property validation settings custom.

Upstream description:

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

Terraform syntax:

```terraform
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

- [query_parameters](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters.md)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
