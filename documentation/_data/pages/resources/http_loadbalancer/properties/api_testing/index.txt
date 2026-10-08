---
page_title: "api_testing"
subcategory: "Load Balancing"
description: "API Testing."
xcsh_docs: {"aliases": ["api testing"], "body_bytes": 3384, "body_sha256": "sha256:f010a8b9ca7e0700f2b68df0980d389fc6d5e5ccca10897c9243205e977fab1e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/api_testing/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_day,every_month", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_day,every_week", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_day,every_month", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_month,every_week", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_day,every_week", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_month,every_week", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_testing"], "schema_version": 1, "sections": [{"aliases": ["api testing custom header value"], "anchor": "schema-api_testing--custom_header_value", "description": "Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "custom_header_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["api testing domains"], "anchor": "section", "description": "Add and configure testing domains and credentials.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-api_testing--domains--domain", "enforcement": "provider-schema", "group": "api_testing.domains:RequiredListObjectAttributes:credentials,domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains:RequiredListObjectAttributes:credentials,domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "type": "requires"}], "schema_path": ["api_testing", "domains"], "syntax": "block", "type": "object"}, {"aliases": ["api testing every day"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_day"], "syntax": "attribute", "type": "object"}, {"aliases": ["api testing every month"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_month"], "syntax": "attribute", "type": "object"}, {"aliases": ["api testing every week"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_week"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "API Testing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- api_testing

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_testing, disable\_api\_testing; Default: disable\_api\_testing\] API Testing.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("every_day",
    "every_month"),
  validators.ConflictingObjectAttributes("every_day",
    "every_week"),
  validators.ConflictingObjectAttributes("every_month",
    "every_week")}
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
  "x-ves-oneof-field-frequency_choice": "[\"every_day\",\"every_month\",\"every_week\"]"
}
```

OneOf alternatives in this subsection:

- [api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/#section)
- [disable_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/disable_api_testing/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_testing {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-api_testing--custom_header_value"></a>

### custom_header_value property

Type: `"string"`. Optional.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/): complete subsection reference.

- [every_day](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/every_day/): complete subsection reference.

- [every_month](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/every_month/): complete subsection reference.

- [every_week](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/every_week/): complete subsection reference.
