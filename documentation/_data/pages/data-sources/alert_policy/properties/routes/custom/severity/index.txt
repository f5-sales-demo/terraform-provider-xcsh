---
page_title: "routes.custom.severity"
subcategory: "Monitoring"
description: "Label Matcher."
xcsh_docs: {"aliases": ["routes custom severity"], "body_bytes": 2614, "body_sha256": "sha256:6ed7b7e092e941f28848982c19914975be3bb228ab7f37b31d9e7c79c90f033d", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:severity", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom", "path": "documentation/data-sources/alert_policy/properties/routes/custom/severity/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2223332111322120-2212302211303331-1121033012003020-3012032003021332-3001232020323333-3302321132331233-0220230103100022-3233130331330020", "registry_path": "docs/guides/data-sources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom", "severity"], "schema_version": 1, "sections": [{"aliases": ["routes custom severity exact match"], "anchor": "schema-routes--custom--severity--exact_match", "description": "Exclusive with Equality match value for the label.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:severity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom", "severity", "exact_match"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes custom severity regex match"], "anchor": "schema-routes--custom--severity--regex_match", "description": "Exclusive with Regular expression match value for the label.", "document_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:severity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom", "severity", "regex_match"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/custom/severity/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Label Matcher.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["alert_policyCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom.severity

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/)
- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/)
- routes.custom.severity

<a id="section"></a>

Type: `"single"`. Computed.

Label Matcher.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

## Direct properties

<a id="schema-routes--custom--severity--exact_match"></a>

### exact_match property

Type: `"string"`. Computed.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-routes--custom--severity--regex_match"></a>

### regex_match property

Type: `"string"`. Computed.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/properties/routes/custom/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/alert_policy/)
