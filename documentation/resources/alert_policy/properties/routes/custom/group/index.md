---
page_title: "routes.custom.group"
subcategory: "Monitoring"
description: "Label Matcher."
xcsh_docs: {"aliases": ["routes custom group"], "body_bytes": 2870, "body_sha256": "sha256:78b053170d108f5aea898e30fc203846f09d8e287415b01d3e77d5b70b2507f4", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "parent_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "path": "documentation/resources/alert_policy/properties/routes/custom/group/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2011010113021203-2333000210101230-3032311201110303-1013211322323222-1113220300312112-0201232013320033-3012310221102211-0031001203003130", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [{"anchor": "schema-routes--custom--group--exact_match", "enforcement": "provider-schema", "group": "routes.custom.group:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "type": "conflicts"}, {"anchor": "schema-routes--custom--group--regex_match", "enforcement": "provider-schema", "group": "routes.custom.group:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom", "group"], "schema_version": 1, "sections": [{"aliases": ["exact match"], "anchor": "schema-routes--custom--group--exact_match", "description": "Exclusive with Equality match value for the label.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom", "group", "exact_match"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex match"], "anchor": "schema-routes--custom--group--regex_match", "description": "Exclusive with Regular expression match value for the label.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom", "group", "regex_match"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/custom/group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Label Matcher.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom.group

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/)
- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/)
- routes.custom.group

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Label Matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_match",
    "regex_match")}
```

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

Terraform syntax:

```terraform
group {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--custom--group--exact_match"></a>

### exact_match property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-routes--custom--group--regex_match"></a>

### regex_match property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
