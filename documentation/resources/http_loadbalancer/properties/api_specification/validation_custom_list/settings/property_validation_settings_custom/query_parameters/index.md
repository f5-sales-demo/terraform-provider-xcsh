---
page_title: "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters"
subcategory: "Load Balancing"
description: "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3532, "body_sha256": "sha256:768451e151745258b2b38f109212e5a7781aa9b329491934ca256862eb135417", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/)
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

- [allow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/allow_additional_parameters/): complete subsection reference.

- [disallow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/disallow_additional_parameters/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/allow_additional_parameters/)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/disallow_additional_parameters/)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
