---
page_title: "api_testing"
subcategory: "Load Balancing"
description: "API Testing."
xcsh_docs: {"aliases": ["api testing"], "body_bytes": 3572, "body_sha256": "sha256:480ce964372dde42d86d4ca9e55a36e87e70b0840be4687f7d20db4c05aa5510", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:every_day", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:every_month", "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:every_week"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "documentation/data-sources/http_loadbalancer/properties/api_testing/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3122013303110033-3021021311102313-3323021211213102-2100213133211231-2011301013301202-0233332230322311-0332101101001110-1223303130231312", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-010.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_testing"], "schema_version": 1, "sections": [{"aliases": ["custom header value"], "anchor": "schema-api_testing--custom_header_value", "description": "Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "custom_header_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication", "credential setup", "credentials", "domains"], "anchor": "section", "description": "Add and configure testing domains and credentials.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:domains", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_testing", "domains"], "syntax": "attribute", "type": "object"}, {"aliases": ["every day"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:every_day", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_day"], "syntax": "attribute", "type": "object"}, {"aliases": ["every month"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:every_month", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_month"], "syntax": "attribute", "type": "object"}, {"aliases": ["every week"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:api_testing:every_week", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_testing", "every_week"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/api_testing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "API Testing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_testing

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- api_testing

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: api\_testing, disable\_api\_testing; Default: disable\_api\_testing\] API Testing.

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

- [api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/#section)
- [disable_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/disable_api_testing/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-api_testing--custom_header_value"></a>

### custom_header_value property

Type: `"string"`. Computed.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

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

- [domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/): complete subsection reference.

- [every_day](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/every_day/): complete subsection reference.

- [every_month](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/every_month/): complete subsection reference.

- [every_week](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/every_week/): complete subsection reference.

## Next pages

- [api_testing.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/domains/)
- [api_testing.every_day](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/every_day/)
- [api_testing.every_month](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/every_month/)
- [api_testing.every_week](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/api_testing/every_week/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
