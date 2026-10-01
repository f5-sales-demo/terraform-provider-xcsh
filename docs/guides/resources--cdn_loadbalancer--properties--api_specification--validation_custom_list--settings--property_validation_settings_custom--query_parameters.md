---
page_title: "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2933, "body_sha256": "sha256:fe4ebd36994e847e8a46f1d817a7e90c1b1abbf1b277c7529884235dcea0fc38", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom--query_parameters.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_specification](resources--cdn_loadbalancer--properties--api_specification.md)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list.md)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings.md)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom.md)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow_additional_parameters",
    "disallow_additional_parameters")}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [allow_additional_parameters](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom--query_parameters--allow_additional_parameters.md): complete subsection reference.

- [disallow_additional_parameters](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom--query_parameters--disallow_additional_parameters.md): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom--query_parameters--allow_additional_parameters.md)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom--query_parameters--disallow_additional_parameters.md)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--properties--api_specification--validation_custom_list--settings--property_validation_settings_custom.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
