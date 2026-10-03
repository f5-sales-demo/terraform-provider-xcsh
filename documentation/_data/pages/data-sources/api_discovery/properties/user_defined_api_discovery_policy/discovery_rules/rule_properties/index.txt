---
page_title: "user_defined_api_discovery_policy.discovery_rules.rule_properties"
subcategory: ""
description: "Determines whether matching endpoints are included in API Discovery or excluded."
xcsh_docs: {"aliases": ["user defined api discovery policy discovery rules rule properties"], "body_bytes": 4607, "body_sha256": "sha256:e0148e441235e3576aabffdd2a6b1efd3f06178aef4a890f9f47b555f4d2593b", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:inclusion"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "parent_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules", "path": "documentation/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/index.md", "product": "distributed-cloud", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3323303011223032-3320130223201120-1023021232331211-1310102121212133-1130012211020030-3020120001300323-2232230022103321-0131120302102312", "registry_path": "docs/guides/data-sources--api_discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties"], "schema_version": 1, "sections": [{"aliases": ["user defined api discovery policy discovery rules rule properties exclusion"], "anchor": "section", "description": "Configuration for exclusion action.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:exclusion", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "exclusion"], "syntax": "attribute", "type": "object"}, {"aliases": ["user defined api discovery policy discovery rules rule properties http header criteria"], "anchor": "section", "description": "Criteria for matching HTTP headers.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:http_header_criteria", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "http_header_criteria"], "syntax": "attribute", "type": "object"}, {"aliases": ["user defined api discovery policy discovery rules rule properties inclusion"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties:inclusion", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "inclusion"], "syntax": "attribute", "type": "object"}, {"aliases": ["user defined api discovery policy discovery rules rule properties pattern"], "anchor": "schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--pattern", "description": "Exclusive with Patterns are matched against the request path to identify endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern are affected by the rule.", "document_id": "xcsh-docs:data-sources:api_discovery:properties:user_defined_api_discovery_policy:discovery_rules:rule_properties", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_defined_api_discovery_policy", "discovery_rules", "rule_properties", "pattern"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Determines whether matching endpoints are included in API Discovery or excluded.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_defined_api_discovery_policy.discovery_rules.rule_properties

Breadcrumbs:

- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/)
- [user_defined_api_discovery_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/)
- [user_defined_api_discovery_policy.discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/)
- user_defined_api_discovery_policy.discovery_rules.rule_properties

<a id="section"></a>

Type: `"single"`. Computed.

Determines whether matching endpoints are included in API Discovery or excluded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-criteria": "[\"http_header_criteria\",\"pattern\"]",
  "x-ves-oneof-field-rule_type_choice": "[\"exclusion\",\"inclusion\"]"
}
```

## Direct properties

- [exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/): complete subsection reference.

- [http_header_criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/): complete subsection reference.

- [inclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/inclusion/): complete subsection reference.

<a id="schema-user_defined_api_discovery_policy--discovery_rules--rule_properties--pattern"></a>

### pattern property

Type: `"string"`. Computed.

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Upstream description:

Exclusive with \[http\_header\_criteria\] Patterns are matched against the request path to identify
endpoints by path structure, file extension, or version prefix. Endpoints that match this pattern
are affected by the rule.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

## Next pages

- [user_defined_api_discovery_policy.discovery_rules.rule_properties.exclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/exclusion/)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.http_header_criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/http_header_criteria/)
- [user_defined_api_discovery_policy.discovery_rules.rule_properties.inclusion](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/rule_properties/inclusion/)
- [user_defined_api_discovery_policy.discovery_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/properties/user_defined_api_discovery_policy/discovery_rules/)
- [xcsh_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_discovery/)
