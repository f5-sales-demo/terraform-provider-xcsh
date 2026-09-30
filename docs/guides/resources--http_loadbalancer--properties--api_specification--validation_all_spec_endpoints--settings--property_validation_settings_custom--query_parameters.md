---
page_title: "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters"
subcategory: "Load Balancing"
description: "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2960, "body_sha256": "sha256:a9485f45ff535ae6bea655796b53a2625684cc8c3b3320c6d2b4967fd03814c2", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom:query_parameters", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints:settings:property_validation_settings_custom", "path": "docs/guides/resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "settings", "property_validation_settings_custom", "query_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/settings/property_validation_settings_custom/query_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [api_specification](resources--http_loadbalancer--properties--api_specification.md)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints.md)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings.md)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom.md)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

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

- [allow_additional_parameters](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters--allow_additional_parameters.md): complete subsection reference.

- [disallow_additional_parameters](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters--disallow_additional_parameters.md): complete subsection reference.

## Next pages

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters--allow_additional_parameters.md)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom--query_parameters--disallow_additional_parameters.md)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--properties--api_specification--validation_all_spec_endpoints--settings--property_validation_settings_custom.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
