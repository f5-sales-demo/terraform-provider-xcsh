---
page_title: "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters"
subcategory: "Load Balancing"
description: "Custom settings for query parameters validation."
xcsh_docs: {"aliases": ["api specification validation custom list settings property validation settings custom query parameters"], "body_bytes": 3518, "body_sha256": "sha256:464d294e3254ac497eafd948d0dbc01d7229f9e8ab5a5b2b296073b609e6ba8c", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1000200113012312-1213233230310003-3331010123211031-2312011301332111-3202110021022322-3133012130322101-1133001121322212-3123212230310110", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters:ConflictingObjectAttributes:allow_additional_parameters,disallow_additional_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters:ConflictingObjectAttributes:allow_additional_parameters,disallow_additional_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters"], "schema_version": 1, "sections": [{"aliases": ["allow additional parameters"], "anchor": "section", "description": "Configuration parameter for allow additional parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters", "allow_additional_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["disallow additional parameters"], "anchor": "section", "description": "Configuration parameter for disallow additional parameters.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters", "disallow_additional_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Custom settings for query parameters validation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_custom_list.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/)
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

- [allow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/allow_additional_parameters/): complete subsection reference.

- [disallow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/disallow_additional_parameters/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/allow_additional_parameters/)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/disallow_additional_parameters/)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
