---
page_title: "api_specification"
subcategory: "Load Balancing"
description: "Settings for API specification (API definition, OpenAPI validation, etc.)"
xcsh_docs: {"aliases": ["api specification"], "body_bytes": 3725, "body_sha256": "sha256:50d4fe22225e1e75a799e819c60000b96feb8d5470bf6db6a7f3063ce21ce101", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_specification:api_definition", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_disabled"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/api_specification/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_all_spec_endpoints,validation_custom_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_all_spec_endpoints,validation_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_all_spec_endpoints,validation_custom_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_custom_list,validation_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_all_spec_endpoints,validation_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_specification:ConflictingObjectAttributes:validation_custom_list,validation_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_disabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_specification"], "schema_version": 1, "sections": [{"aliases": ["api definition"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:api_definition", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-api_specification--api_definition--name", "enforcement": "provider-schema", "group": "api_specification.api_definition:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:api_definition", "type": "requires"}], "schema_path": ["api_specification", "api_definition"], "syntax": "block", "type": "object"}, {"aliases": ["validation all spec endpoints"], "anchor": "section", "description": "Settings for API Inventory validation.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_all_spec_endpoints", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["api_specification", "validation_all_spec_endpoints"], "syntax": "block", "type": "object"}, {"aliases": ["validation custom list"], "anchor": "section", "description": "Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other API-endpoint not listed will act according to \"Fall Through Mode\".", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_specification.validation_custom_list:RequiredObjectAttributes:open_api_validation_rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_custom_list:open_api_validation_rules", "type": "requires"}], "schema_path": ["api_specification", "validation_custom_list"], "syntax": "block", "type": "object"}, {"aliases": ["validation disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_specification:validation_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_specification", "validation_disabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_specification/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Settings for API specification (API definition, OpenAPI validation, etc.)", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_specification

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- api_specification

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_specification, disable\_api\_definition; Default: disable\_api\_definition\] Settings
for API specification (API definition, OpenAPI validation, etc.).

Upstream description:

Settings for API specification (API definition, OpenAPI validation, etc.)

Provider validators and defaults (from schema source):

```go
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

- [api_specification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/#section)
- [disable_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/disable_api_definition/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_specification {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/api_definition/): complete subsection reference.

- [validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/): complete subsection reference.

- [validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/): complete subsection reference.

- [validation_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_disabled/): complete subsection reference.

## Next pages

- [api_specification.api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/api_definition/)
- [api_specification.validation_all_spec_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_all_spec_endpoints/)
- [api_specification.validation_custom_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_custom_list/)
- [api_specification.validation_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_specification/validation_disabled/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
