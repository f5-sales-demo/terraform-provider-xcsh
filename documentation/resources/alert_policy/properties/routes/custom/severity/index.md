---
page_title: "routes.custom.severity"
subcategory: "Monitoring"
description: "Label Matcher."
xcsh_docs: {"aliases": ["routes custom severity"], "body_bytes": 2885, "body_sha256": "sha256:1ded53edec7a500cbb8214e6f9b7ea0c6b5179127e1628fa808b01a41ceb3463", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes:custom:severity", "parent_id": "xcsh-docs:resources:alert_policy:properties:routes:custom", "path": "documentation/resources/alert_policy/properties/routes/custom/severity/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3010333121102303-1101131221000001-0131322123023133-0130331133201303-2021210233033132-3200213201213312-0002312032003303-1232113323213003", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [{"anchor": "schema-routes--custom--severity--exact_match", "enforcement": "provider-schema", "group": "routes.custom.severity:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:severity", "type": "conflicts"}, {"anchor": "schema-routes--custom--severity--regex_match", "enforcement": "provider-schema", "group": "routes.custom.severity:ConflictingObjectAttributes:exact_match,regex_match", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:severity", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "custom", "severity"], "schema_version": 1, "sections": [{"aliases": ["routes custom severity exact match"], "anchor": "schema-routes--custom--severity--exact_match", "description": "Exclusive with Equality match value for the label.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:severity", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom", "severity", "exact_match"], "syntax": "attribute", "type": "string"}, {"aliases": ["routes custom severity regex match"], "anchor": "schema-routes--custom--severity--regex_match", "description": "Exclusive with Regular expression match value for the label.", "document_id": "xcsh-docs:resources:alert_policy:properties:routes:custom:severity", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["routes", "custom", "severity", "regex_match"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/custom/severity/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Label Matcher.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom.severity

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/)
- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/)
- routes.custom.severity

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
severity {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-routes--custom--severity--exact_match"></a>

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

- [routes.custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/routes/custom/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
