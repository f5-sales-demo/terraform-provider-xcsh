---
page_title: "api_specification.validation_custom_list.settings"
subcategory: "Load Balancing"
description: "OpenAPI specification validation settings relevant for \"API Inventory\" enforcement and for \"Custom list\" enforcement."
xcsh_docs: {"aliases": ["api specification validation custom list settings"], "body_bytes": 4437, "body_sha256": "sha256:e1b5bccdd9c15d19c8c66e20745ef874e9e7a2a90a216d4254a43f7e8253f8bd", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_fail_validation", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_skip_validation", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_default"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2120033000331323-2023121101030022-1120013320021212-2032223212012232-2012300302330330-2321131313020133-0332322003002322-2231033320320223", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings:ConflictingObjectAttributes:oversized_body_fail_validation,oversized_body_skip_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_fail_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings:ConflictingObjectAttributes:oversized_body_fail_validation,oversized_body_skip_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_skip_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings:ConflictingObjectAttributes:property_validation_settings_custom,property_validation_settings_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings:ConflictingObjectAttributes:property_validation_settings_custom,property_validation_settings_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_default", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list", "settings"], "schema_version": 1, "sections": [{"aliases": ["oversized body fail validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_fail_validation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "oversized_body_fail_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["oversized body skip validation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_skip_validation", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "oversized_body_skip_validation"], "syntax": "attribute", "type": "object"}, {"aliases": ["property validation settings custom"], "anchor": "section", "description": "Custom property validation settings.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_custom"], "syntax": "block", "type": "object"}, {"aliases": ["property validation settings default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_default", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "settings", "property_validation_settings_default"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "OpenAPI specification validation settings relevant for \"API Inventory\" enforcement and for \"Custom list\" enforcement.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list.settings

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/)
- api_specification.validation_custom_list.settings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("oversized_body_fail_validation",
    "oversized_body_skip_validation"),
  validators.ConflictingObjectAttributes("property_validation_settings_custom",
    "property_validation_settings_default")}
```

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

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [oversized_body_fail_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/oversized_body_fail_validation/): complete subsection reference.

- [oversized_body_skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/oversized_body_skip_validation/): complete subsection reference.

- [property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/): complete subsection reference.

- [property_validation_settings_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_default/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.settings.oversized_body_fail_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/oversized_body_fail_validation/)
- [api_specification.validation_custom_list.settings.oversized_body_skip_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/oversized_body_skip_validation/)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_custom/)
- [api_specification.validation_custom_list.settings.property_validation_settings_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/property_validation_settings_default/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
