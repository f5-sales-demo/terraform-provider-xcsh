---
page_title: "routes.custom"
subcategory: "Monitoring"
description: "routes.custom for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1700, "body_sha256": "sha256:bbece54bbcc4507f8ebd3f1e4bad79f1f0a08d89b239f2ea17b728d109fd25f2", "canonical_id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom", "child_ids": ["xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertlabel", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:alertname", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:group", "xcsh-docs:data-sources:alert_policy:properties:routes:custom:severity"], "collection_id": "xcsh-docs:data-sources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:alert_policy:properties:routes:custom", "parent_id": "xcsh-docs:data-sources:alert_policy:properties:routes", "path": "docs/guides/data-sources--alert_policy--properties--routes--custom.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/alert_policy/properties/routes/custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.custom for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.custom

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md)
- [Property reference](data-sources--alert_policy--reference.md)
- [routes](data-sources--alert_policy--properties--routes.md)
- routes.custom

<a id="section"></a>

Type: `"single"`. Computed.

Set of matchers an alert has to fulfill to match the route.

Upstream description:

A set of matchers an alert has to fulfill to match the route.

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

- [alertlabel](data-sources--alert_policy--properties--routes--custom--alertlabel.md): complete subsection reference.

- [alertname](data-sources--alert_policy--properties--routes--custom--alertname.md): complete subsection reference.

- [group](data-sources--alert_policy--properties--routes--custom--group.md): complete subsection reference.

- [severity](data-sources--alert_policy--properties--routes--custom--severity.md): complete subsection reference.

## Next pages

- [routes.custom.alertlabel](data-sources--alert_policy--properties--routes--custom--alertlabel.md)
- [routes.custom.alertname](data-sources--alert_policy--properties--routes--custom--alertname.md)
- [routes.custom.group](data-sources--alert_policy--properties--routes--custom--group.md)
- [routes.custom.severity](data-sources--alert_policy--properties--routes--custom--severity.md)
- [routes](data-sources--alert_policy--properties--routes.md)
- [xcsh_alert_policy](../data-sources/alert_policy.md)
