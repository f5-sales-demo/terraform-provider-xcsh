---
page_title: "tls_intercept.policy.interception_rules"
subcategory: ""
description: "List of ordered rules to enable or disable for TLS interception."
xcsh_docs: {"aliases": ["tls intercept policy interception rules"], "body_bytes": 2389, "body_sha256": "sha256:30cec9d9920e9013ba2bb60740f64c41065602cb7bca992ee4f2818bd173fbe2", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules:disable_interception", "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules:enable_interception"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules", "parent_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy", "path": "documentation/data-sources/proxy/properties/tls_intercept/policy/interception_rules/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112", "registry_path": "docs/guides/data-sources--proxy--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_intercept", "policy", "interception_rules"], "schema_version": 1, "sections": [{"aliases": ["tls intercept policy interception rules disable interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules:disable_interception", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "policy", "interception_rules", "disable_interception"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls intercept policy interception rules domain match"], "anchor": "section", "description": "Domains names.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules:domain_match", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tls_intercept", "policy", "interception_rules", "domain_match"], "syntax": "attribute", "type": "object"}, {"aliases": ["tls intercept policy interception rules enable interception"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:tls_intercept:policy:interception_rules:enable_interception", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tls_intercept", "policy", "interception_rules", "enable_interception"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/tls_intercept/policy/interception_rules/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "List of ordered rules to enable or disable for TLS interception.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["proxyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_intercept.policy.interception_rules

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [tls_intercept](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/)
- [tls_intercept.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/policy/)
- tls_intercept.policy.interception_rules

<a id="section"></a>

Type: `"list"`. Computed.

List of ordered rules to enable or disable for TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-21T11:00:09+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [disable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/policy/interception_rules/disable_interception/): complete subsection reference.

- [domain_match](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/policy/interception_rules/domain_match/): complete subsection reference.

- [enable_interception](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/tls_intercept/policy/interception_rules/enable_interception/): complete subsection reference.
