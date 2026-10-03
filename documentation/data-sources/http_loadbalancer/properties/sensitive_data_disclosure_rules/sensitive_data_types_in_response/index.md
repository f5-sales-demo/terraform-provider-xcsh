---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response"
subcategory: "Load Balancing"
description: "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses."
xcsh_docs: {"aliases": ["sensitive data disclosure rules sensitive data types in response"], "body_bytes": 3751, "body_sha256": "sha256:9957642857a21a2c1b7e4bb08c89bc0124cf7b2e29a362ad62d4f494664c30ec", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:report"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "path": "documentation/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2202311031302201-2300030202310013-0022020102200122-2033122123122002-1313131132112120-3221123310313023-2113100012233002-1130210321000321", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response"], "schema_version": 1, "sections": [{"aliases": ["sensitive data disclosure rules sensitive data types in response api endpoint"], "anchor": "section", "description": "This defines API endpoint.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "api_endpoint"], "syntax": "attribute", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response body"], "anchor": "section", "description": "OPTIONS for HTTP Body Masking.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "body"], "syntax": "attribute", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response mask"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "mask"], "syntax": "attribute", "type": "object"}, {"aliases": ["sensitive data disclosure rules sensitive data types in response report"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:report", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response", "report"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/api_endpoint/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.body](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/body/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/mask/)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.report](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/report/)
- [sensitive_data_disclosure_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
