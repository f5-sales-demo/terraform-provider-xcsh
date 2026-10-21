---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response"
subcategory: "Load Balancing"
description: "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses."
xcsh_docs: {"aliases": ["sensitive data disclosure rules sensitive data types in response"], "body_bytes": 2575, "body_sha256": "sha256:a339ebb10330a009834ccdd62a87d5c51f9abe4866bfe6183910965fd1cf722d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:report"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "path": "documentation/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2023120330103011-1030003123103102-1102223030033323-3132213302112330-2131203003333230-1012102100122000-3211320013000233-2230021332122030", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response"], "schema_version": 1, "sections": [{"aliases": ["sensitive data disclosure rules sensitive data types in response api endpoint"], "anchor": "section", "description": "This defines API endpoint.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "api_endpoint"], "syntax": "block", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response body"], "anchor": "section", "description": "OPTIONS for HTTP Body Masking.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "body"], "syntax": "block", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response mask"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "mask"], "syntax": "attribute", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response report"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:report", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "report"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules.sensitive_data_types_in_response

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [sensitive_data_disclosure_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/)
- sensitive_data_disclosure_rules.sensitive_data_types_in_response

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-21T11:00:09+00:00"
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

Terraform syntax:

```terraform
sensitive_data_types_in_response {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/api_endpoint/): complete subsection reference.

- [body](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/): complete subsection reference.

- [mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/mask/): complete subsection reference.

- [report](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/report/): complete subsection reference.
