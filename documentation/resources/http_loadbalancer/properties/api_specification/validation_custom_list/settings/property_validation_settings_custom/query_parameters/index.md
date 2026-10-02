---
page_title: "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters"
subcategory: "Load Balancing"
description: "Custom settings for query parameters validation."
xcsh_docs: {"aliases": ["api specification validation custom list settings property validation settings custom query parameters"], "body_bytes": 3532, "body_sha256": "sha256:768451e151745258b2b38f109212e5a7781aa9b329491934ca256862eb135417", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3021030320203033-0223013023100133-2213202322202003-2203303231023101-3203013330020331-1030233301212013-3231230321302033-2333002103031122", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters:ConflictingObjectAttributes:allow_additional_parameters,disallow_additional_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters:ConflictingObjectAttributes:allow_additional_parameters,disallow_additional_parameters", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters"], "schema_version": 1, "sections": [{"aliases": ["allow additional parameters"], "anchor": "section", "description": "Configuration parameter for allow additional parameters.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:allow_additional_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters", "allow_additional_parameters"], "syntax": "attribute", "type": "object"}, {"aliases": ["disallow additional parameters"], "anchor": "section", "description": "Configuration parameter for disallow additional parameters.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom:query_parameters:disallow_additional_parameters", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom", "query_parameters", "disallow_additional_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/query_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Custom settings for query parameters validation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
