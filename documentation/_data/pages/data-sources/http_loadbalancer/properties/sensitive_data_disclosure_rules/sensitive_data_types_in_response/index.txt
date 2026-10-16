---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response"
subcategory: "Load Balancing"
description: "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses."
xcsh_docs: {"aliases": ["sensitive data disclosure rules sensitive data types in response"], "body_bytes": 2455, "body_sha256": "sha256:aeaa878cb2141c40d77ca4db78f8d0ae9da78e26e45d71aa346c5f2d9643a653", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:report"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "path": "documentation/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-2202311031302201-2300030202310013-0022020102200122-2033122123122002-1313131132112120-3221123310313023-2113100012233002-1130210321000321", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response"], "schema_version": 1, "sections": [{"aliases": ["sensitive data disclosure rules sensitive data types in response api endpoint"], "anchor": "section", "description": "This defines API endpoint.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "api_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response body"], "anchor": "section", "description": "OPTIONS for HTTP Body Masking.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "body"], "syntax": "attribute", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response mask"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "mask"], "syntax": "attribute", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response report"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:report", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "report"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules.sensitive_data_types_in_response

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [sensitive_data_disclosure_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response

<a id="section"></a>

Type: `"list"`. Computed.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-16T10:31:02+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/api_endpoint/): complete subsection reference.

- [body](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/): complete subsection reference.

- [mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/mask/): complete subsection reference.

- [report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/report/): complete subsection reference.
