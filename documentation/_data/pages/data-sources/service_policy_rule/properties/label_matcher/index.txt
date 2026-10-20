---
page_title: "label_matcher"
subcategory: ""
description: "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name."
xcsh_docs: {"aliases": ["label matcher"], "body_bytes": 2043, "body_sha256": "sha256:b979347337e961784a91b7da959638677a74df5a31f37baf4a5f4f46dcfa0adf", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:label_matcher", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "documentation/data-sources/service_policy_rule/properties/label_matcher/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3210311000320320-2030222121122100-0220110100321312-1121102131023230-3300133221311211-2200300323311103-1132001130323212-2010010230320023", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["label_matcher"], "schema_version": 1, "sections": [{"aliases": ["label matcher keys"], "anchor": "schema-label_matcher--keys", "description": "The list of label key names that have to match.", "document_id": "xcsh-docs:data-sources:service_policy_rule:properties:label_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["label_matcher", "keys"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/label_matcher/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# label_matcher

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- label_matcher

<a id="section"></a>

Type: `"single"`. Computed.

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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

<a id="schema-label_matcher--keys"></a>

### keys property

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
