---
page_title: "notification_parameters.custom"
subcategory: "Monitoring"
description: "notification_parameters.custom for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1961, "body_sha256": "sha256:b817a926d76bb0c77861d0fb50d86efe4adee8678c4f273114dd26427cc62359", "canonical_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:custom", "child_ids": [], "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters:custom", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:notification_parameters", "path": "docs/guides/data-sources--alert_policy--properties--notification_parameters--custom.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["notification_parameters", "custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/notification_parameters/custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "notification_parameters.custom for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# notification_parameters.custom

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md)
- [Property reference](data-sources--alert_policy--reference.md)
- [notification_parameters](data-sources--alert_policy--properties--notification_parameters.md)
- notification_parameters.custom

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

<a id="schema-notification_parameters--custom--labels"></a>

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

- [notification_parameters](data-sources--alert_policy--properties--notification_parameters.md)
- [xcsh_alert_policy](../data-sources/alert_policy.md)
