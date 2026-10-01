---
page_title: "routes.custom.alertlabel"
subcategory: "Monitoring"
description: "routes.custom.alertlabel for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1483, "body_sha256": "sha256:4a258577629b399d54b38b26dd938b7a6fa79d68206936d3b6c78cf1cc360c60", "canonical_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertlabel", "child_ids": [], "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertlabel", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom", "path": "docs/guides/data-sources--alert_policy--properties--routes--custom--alertlabel.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "custom", "alertlabel"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/custom/alertlabel/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.custom.alertlabel for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom.alertlabel

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md)
- [Property reference](data-sources--alert_policy--reference.md)
- [routes](data-sources--alert_policy--properties--routes.md)
- [routes.custom](data-sources--alert_policy--properties--routes--custom.md)
- routes.custom.alertlabel

<a id="section"></a>

Type: `"single"`. Computed.

AlertLabel to configure the alert policy rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  }
}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.custom](data-sources--alert_policy--properties--routes--custom.md)
- [xcsh_alert_policy](../data-sources/alert_policy.md)
