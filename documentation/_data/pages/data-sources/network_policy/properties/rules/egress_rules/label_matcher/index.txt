---
page_title: "rules.egress_rules.label_matcher"
subcategory: "Security"
description: "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name."
xcsh_docs: {"aliases": ["rules egress rules label matcher"], "body_bytes": 2339, "body_sha256": "sha256:131a31030c789d6deffd3b6527fe1c31926cdc7d2e70f66b642785b94624013b", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules:label_matcher", "parent_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules", "path": "documentation/data-sources/network_policy/properties/rules/egress_rules/label_matcher/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3103120202331220-3020222320021103-3231231302132201-0221111120010303-3201333021311002-2122002331032200-0321021001230303-3230301111333210", "registry_path": "docs/guides/data-sources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "egress_rules", "label_matcher"], "schema_version": 1, "sections": [{"aliases": ["rules egress rules label matcher keys"], "anchor": "schema-rules--egress_rules--label_matcher--keys", "description": "The list of label key names that have to match.", "document_id": "xcsh-docs:data-sources:network_policy:properties:rules:egress_rules:label_matcher", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "egress_rules", "label_matcher", "keys"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/rules/egress_rules/label_matcher/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "A label matcher specifies a list of label keys whose values need to match for source/client and destination/server. Note that the actual label values are not specified and do not matter. This allows an ability to scope grouping by the label key name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["network_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.egress_rules.label_matcher

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/)
- [rules.egress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/rules/egress_rules/)
- rules.egress_rules.label_matcher

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

<a id="schema-rules--egress_rules--label_matcher--keys"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
