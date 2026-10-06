---
page_title: "routes.notification_parameters.custom"
subcategory: "Monitoring"
description: "Specify list of custom labels to group/aggregate the alerts."
xcsh_docs: {"aliases": ["routes notification parameters custom"], "body_bytes": 2098, "body_sha256": "sha256:1e0e7c2ddc6c6f18fbad2438fa9ed8649b0b9abc331e91211f79008410699dea", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:custom", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters", "path": "documentation/data-sources/alert_policy/properties/routes/notification_parameters/custom/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1010231102310202-0032222133032113-2000321113021313-0331323220033212-3202133213233002-0322312220221230-1032322231022032-2201203232220130", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "notification_parameters", "custom"], "schema_version": 1, "sections": [{"aliases": ["routes notification parameters custom labels"], "anchor": "schema-routes--notification_parameters--custom--labels", "description": "Name of labels to group/aggregate the alerts.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:notification_parameters:custom", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "notification_parameters", "custom", "labels"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/notification_parameters/custom/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Specify list of custom labels to group/aggregate the alerts.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.notification_parameters.custom

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/)
- [routes.notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/notification_parameters/)
- routes.notification_parameters.custom

<a id="section"></a>

Type: `"single"`. Computed.

Specify list of custom labels to group/aggregate the alerts.

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

<a id="schema-routes--notification_parameters--custom--labels"></a>

### labels property

Type: `["list", "string"]`. Computed.

Name of labels to group/aggregate the alerts.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
