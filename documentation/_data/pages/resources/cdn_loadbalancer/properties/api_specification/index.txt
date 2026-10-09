---
page_title: "api_specification"
subcategory: "Load Balancing"
description: "Settings for API specification (API definition, OpenAPI validation, etc.)"
xcsh_docs: {"aliases": ["api specification"], "body_bytes": 2675, "body_sha256": "sha256:2dc804c3dd6f1a82a7ee51bba05ebc4454ee58dc3851ff014a2671ef29458285", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:api_definition", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_disabled"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/api_specification/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_all_spec_endpoints,validation_custom_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_all_spec_endpoints,validation_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_all_spec_endpoints,validation_custom_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_custom_list,validation_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_all_spec_endpoints,validation_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_custom_list,validation_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_disabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification"], "schema_version": 1, "sections": [{"aliases": ["api specification api definition"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:api_definition", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_specification--api_definition--name", "enforcement": "provider-schema", "group": "api_specification.api_definition:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:api_definition", "type": "requires"}], "schema_path": ["api_specification", "api_definition"], "syntax": "block", "type": "object"}, {"aliases": ["api specification validation all spec endpoints"], "anchor": "section", "description": "Settings for API Inventory validation.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints"], "syntax": "block", "type": "object"}, {"aliases": ["api specification validation custom list"], "anchor": "section", "description": "Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to \"Fall Through Mode\".", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list:RequiredObjectAttributes:open_api_validation_rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "type": "requires"}], "schema_path": ["api_specification", "validation_custom_list"], "syntax": "block", "type": "object"}, {"aliases": ["api specification validation disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_specification:validation_disabled", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_disabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_specification/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Settings for API specification (API definition, OpenAPI validation, etc.)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- api_specification

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_custom_list"),
  validators.ConflictingObjectAttributes("validation_all_spec_endpoints",
    "validation_disabled"),
  validators.ConflictingObjectAttributes("validation_custom_list",
    "validation_disabled")}
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
  "x-ves-oneof-field-validation_target_choice": "[\"validation_all_spec_endpoints\",\"validation_custom_list\",\"validation_disabled\"]"
}
```

OneOf alternatives in this subsection:

- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/#section)
- [disable_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/disable_api_definition/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/api_definition/): complete subsection reference.

- [validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_all_spec_endpoints/): complete subsection reference.

- [validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_custom_list/): complete subsection reference.

- [validation_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_specification/validation_disabled/): complete subsection reference.
