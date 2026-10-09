---
page_title: "api_specification.validation_all_spec_endpoints.fall_through_mode"
subcategory: "Load Balancing"
description: "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)"
xcsh_docs: {"aliases": ["api specification validation all spec endpoints fall through mode"], "body_bytes": 2174, "body_sha256": "sha256:13711bcb4bd90d4f1a830650fe72f63b50141be9e8c7e4b5ad32342d7eb5e4b0", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_allow", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.fall_through_mode:ConflictingObjectAttributes:fall_through_mode_allow,fall_through_mode_custom", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_allow", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.fall_through_mode:ConflictingObjectAttributes:fall_through_mode_allow,fall_through_mode_custom", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode"], "schema_version": 1, "sections": [{"aliases": ["api specification validation all spec endpoints fall through mode fall through mode allow"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_allow", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_allow"], "syntax": "attribute", "type": "object"}, {"aliases": ["api specification validation all spec endpoints fall through mode fall through mode custom"], "anchor": "section", "description": "Define the fall through settings.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom:RequiredObjectAttributes:open_api_validation_rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints:fall_through_mode:fall_through_mode_custom:open_api_validation_rules", "type": "requires"}], "schema_path": ["api_specification", "validation_all_spec_endpoints", "fall_through_mode", "fall_through_mode_custom"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a. Swagger) or doesn't have a specific rule in custom rules)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification.validation_all_spec_endpoints.fall_through_mode

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("fall_through_mode_allow",
    "fall_through_mode_custom")}
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
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [fall_through_mode_allow](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_allow/): complete subsection reference.

- [fall_through_mode_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/fall_through_mode/fall_through_mode_custom/): complete subsection reference.
