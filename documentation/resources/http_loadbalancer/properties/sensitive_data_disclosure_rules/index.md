---
page_title: "sensitive_data_disclosure_rules"
subcategory: "Load Balancing"
description: "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses."
xcsh_docs: {"aliases": ["sensitive data disclosure rules"], "body_bytes": 1138, "body_sha256": "sha256:8ba4c14e8c7c70340630145476aeefaeef50436d68aaa385f439a008406c3528", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "documentation/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3010130000211103-2132323303233130-3123203102231323-2302320310112320-0023100231121123-3223331112120111-3300221200232213-1133321320223332", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-027.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["sensitive_data_disclosure_rules"], "schema_version": 1, "sections": [{"aliases": ["sensitive data disclosure rules sensitive data types in response"], "anchor": "section", "description": "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:sensitive_data_disclosure_rules:sensitive_data_types_in_response", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["sensitive_data_disclosure_rules", "sensitive_data_types_in_response"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API responses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sensitive_data_disclosure_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- sensitive_data_disclosure_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Sensitive Data Exposure Rules allows specifying rules to mask sensitive data fields in API
responses.

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

Terraform syntax:

```terraform
sensitive_data_disclosure_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sensitive_data_types_in_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/sensitive_data_disclosure_rules/sensitive_data_types_in_response/): complete subsection reference.
