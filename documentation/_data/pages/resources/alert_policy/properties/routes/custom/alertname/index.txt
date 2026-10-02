---
page_title: "routes.custom.alertname"
subcategory: "Monitoring"
description: "Label Matcher."
xcsh_docs: {"aliases": ["routes custom alertname"], "body_bytes": 2890, "body_sha256": "sha256:513d4fd648ecabcc40418b58401287170092e8408178816d7870ad5af5c0f4d5", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "parent_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "path": "documentation/resources/alert_policy/properties/routes/custom/alertname/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2020012332312022-1230112312111312-3111233110312222-2330002312011311-1201320303221211-1321231121320122-1301033130033020-3233100021113303", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [{"anchor": "schema-routes--custom--alertname--exact_match", "enforcement": "provider-schema", "group": "routes.custom.alertname:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "type": "conflicts"}, {"anchor": "schema-routes--custom--alertname--regex_match", "enforcement": "provider-schema", "group": "routes.custom.alertname:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom", "alertname"], "schema_version": 1, "sections": [{"aliases": ["exact match"], "anchor": "schema-routes--custom--alertname--exact_match", "description": "Exclusive with Equality match value for the label.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom", "alertname", "exact_match"], "syntax": "attribute", "type": "string"}, {"aliases": ["regex match"], "anchor": "schema-routes--custom--alertname--regex_match", "description": "Exclusive with Regular expression match value for the label.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:alertname", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom", "alertname", "regex_match"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/custom/alertname/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Label Matcher.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom.alertname

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/)
- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/)
- routes.custom.alertname

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
alertname {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--custom--alertname--exact_match"></a>

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

<a id="schema-routes--custom--alertname--regex_match"></a>

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
