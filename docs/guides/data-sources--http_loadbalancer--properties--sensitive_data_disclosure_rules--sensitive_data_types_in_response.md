---
page_title: "sensitive_data_disclosure_rules.sensitive_data_types_in_response"
subcategory: "Load Balancing"
description: "sensitive_data_disclosure_rules.sensitive_data_types_in_response for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3102, "body_sha256": "sha256:6c0707b3f1ede7ffb4a86cbcb0f1e6b906299bddb60d65489e475522a00cf289", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:api_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:body", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:mask", "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response:report"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sensitive_data_disclosure_rules.sensitive_data_types_in_response for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules.sensitive_data_types_in_response

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules.md)
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

- [api_endpoint](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--api_endpoint.md): complete subsection reference.

- [body](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--body.md): complete subsection reference.

- [mask](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--mask.md): complete subsection reference.

- [report](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--report.md): complete subsection reference.

## Next pages

- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.api_endpoint](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--api_endpoint.md)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.body](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--body.md)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.mask](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--mask.md)
- [sensitive_data_disclosure_rules.sensitive_data_types_in_response.report](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules--sensitive_data_types_in_response--report.md)
- [sensitive_data_disclosure_rules](data-sources--http_loadbalancer--properties--sensitive_data_disclosure_rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
