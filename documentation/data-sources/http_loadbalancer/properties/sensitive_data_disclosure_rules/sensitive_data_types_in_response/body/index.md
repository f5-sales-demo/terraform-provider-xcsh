---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response.body"
subcategory: "Load Balancing"
description: "OPTIONS for HTTP Body Masking."
xcsh_docs: {"aliases": ["sensitive data disclosure rules sensitive data types in response body"], "body_bytes": 3147, "body_sha256": "sha256:ec1079574e4ce2eeab9e74c990234839277fb0360bdfcf8ad6405e7f24673d46", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "path": "documentation/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3010010003013121-2203101333021322-3221113010222013-2201013301203101-1302201201223013-2332231003202311-1033201232210021-0101111031000100", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "body"], "schema_version": 1, "sections": [{"aliases": ["sensitive data disclosure rules sensitive data types in response body fields"], "anchor": "schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--body--fields", "description": "List of JSON Path field values. Use square brackets with an underscore to indicate array elements (e.g., person.emails). To reference JSON keys that contain spaces, enclose the entire path in double quotes. For example: \"person.first name\".", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "body", "fields"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "OPTIONS for HTTP Body Masking.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules.sensitive_data_types_in_response.body

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [sensitive_data_disclosure_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response.body

<a id="section"></a>

Type: `"single"`. Computed.

Body Section Masking OPTIONS. OPTIONS for HTTP Body Masking.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-sensitive_data_disclosure_rules--sensitive_data_types_in_response--body--fields"></a>

### fields property

Type: `["list", "string"]`. Computed.

List of JSON Path field values. Use square brackets with an underscore \[\_\] to indicate array
elements (e.g., person.emails\[\_\]). To reference JSON keys that contain spaces, enclose the entire
path in double quotes. For example: "person.first name".

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-16T10:31:02+00:00"
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
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.json_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
