---
page_title: "api_specification.validation_custom_list"
subcategory: "Load Balancing"
description: "Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to \"Fall Through Mode\"."
xcsh_docs: {"aliases": ["api specification validation custom list"], "body_bytes": 3063, "body_sha256": "sha256:f6f68ca12eb5c1a663307c840b3a851f33d47f19b20080c251c8302f8fcdf07e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification", "path": "documentation/resources/http_loadbalancer/properties/api_specification/validation_custom_list/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3100230121330010-3221133230132303-2201112120203123-0302210131222102-3011333230202231-2203133130103200-2032022122311022-2013122023303210", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list:RequiredObjectAttributes:open_api_validation_rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_custom_list"], "schema_version": 1, "sections": [{"aliases": ["fall through mode"], "anchor": "section", "description": "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.fall_through_mode:ConflictingObjectAttributes:fall_through_mode_allow,fall_through_mode_custom", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.fall_through_mode:ConflictingObjectAttributes:fall_through_mode_allow,fall_through_mode_custom", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:fall_through_mode:fall_through_mode_custom", "type": "conflicts"}], "schema_path": ["api_specification", "validation_custom_list", "fall_through_mode"], "syntax": "block", "type": "object"}, {"aliases": ["open api validation rules"], "anchor": "section", "description": "Rule or policy definition", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_specification", "validation_custom_list", "open_api_validation_rules"], "syntax": "block", "type": "object"}, {"aliases": ["settings"], "anchor": "section", "description": "OpenAPI specification validation settings relevant for \"API Inventory\" enforcement and for \"Custom list\" enforcement.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings:ConflictingObjectAttributes:oversized_body_fail_validation,oversized_body_skip_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_fail_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings:ConflictingObjectAttributes:oversized_body_fail_validation,oversized_body_skip_validation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:oversized_body_skip_validation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings:ConflictingObjectAttributes:property_validation_settings_custom,property_validation_settings_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_custom", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list.settings:ConflictingObjectAttributes:property_validation_settings_custom,property_validation_settings_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:settings:property_validation_settings_default", "type": "conflicts"}], "schema_path": ["api_specification", "validation_custom_list", "settings"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/validation_custom_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to \"Fall Through Mode\".", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_custom_list

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- api_specification.validation_custom_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

Terraform syntax:

```terraform
validation_custom_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/): complete subsection reference.

- [open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/): complete subsection reference.

- [settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/): complete subsection reference.

## Next pages

- [api_specification.validation_custom_list.fall_through_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/fall_through_mode/)
- [api_specification.validation_custom_list.open_api_validation_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/open_api_validation_rules/)
- [api_specification.validation_custom_list.settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/settings/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
