---
page_title: "notification_parameters.custom"
subcategory: "Monitoring"
description: "Specify list of custom labels to group/aggregate the alerts."
xcsh_docs: {"aliases": ["notification parameters custom"], "body_bytes": 2448, "body_sha256": "sha256:a00391d16064cd71c8a4e0a4cf18f76e9d9f03847b527afa0e113aedbb6842ea", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:custom", "parent_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters", "path": "documentation/resources/alert_policy/properties/notification_parameters/custom/index.md", "product": "distributed-cloud", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3113333100001232-1201212313303031-0121102133332311-2231102000302312-2033011323111100-3332002323232120-1021202100212120-3331103121103022", "registry_path": "docs/guides/resources--alert_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["notification_parameters", "custom"], "schema_version": 1, "sections": [{"aliases": ["notification parameters custom labels"], "anchor": "schema-notification_parameters--custom--labels", "description": "Name of labels to group/aggregate the alerts.", "document_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:custom", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["notification_parameters", "custom", "labels"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/notification_parameters/custom/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specify list of custom labels to group/aggregate the alerts.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["alert_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# notification_parameters.custom

Breadcrumbs:

- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/)
- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/)
- notification_parameters.custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-notification_parameters--custom--labels"></a>

### labels property

Type: `["list", "string"]`. Optional.

Name of labels to group/aggregate the alerts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(5),
}
```

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

## Next pages

- [notification_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/properties/notification_parameters/)
- [xcsh_alert_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_policy/)
