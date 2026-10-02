---
page_title: "api_testing"
subcategory: "Load Balancing"
description: "API Testing."
xcsh_docs: {"aliases": ["api testing"], "body_bytes": 4156, "body_sha256": "sha256:4aca0377e18ed2fe07479f85636c0323ca0c52c2cead5a7064a790781340d808", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/api_testing/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_day,every_month", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_day,every_week", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_day,every_month", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_month,every_week", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_day,every_week", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:ConflictingObjectAttributes:every_month,every_week", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing:RequiredObjectAttributes:domains", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_testing"], "schema_version": 1, "sections": [{"aliases": ["custom header value"], "anchor": "schema-api_testing--custom_header_value", "description": "Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "custom_header_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["domains"], "anchor": "section", "description": "Add and configure testing domains and credentials.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-api_testing--domains--domain", "enforcement": "provider-schema", "group": "api_testing.domains:RequiredListObjectAttributes:credentials,domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "api_testing.domains:RequiredListObjectAttributes:credentials,domain", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:domains:credentials", "type": "requires"}], "schema_path": ["api_testing", "domains"], "syntax": "block", "type": "object"}, {"aliases": ["every day"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_day", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_day"], "syntax": "attribute", "type": "object"}, {"aliases": ["every month"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_month", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_month"], "syntax": "attribute", "type": "object"}, {"aliases": ["every week"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_testing:every_week", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_week"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_testing/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "API Testing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [api_testing.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/domains/)
- [api_testing.every_day](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/every_day/)
- [api_testing.every_month](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/every_month/)
- [api_testing.every_week](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_testing/every_week/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
