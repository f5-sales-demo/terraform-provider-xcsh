---
page_title: "routes.notification_parameters.default"
subcategory: "Monitoring"
description: "routes.notification_parameters.default for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1131, "body_sha256": "sha256:7d2b68cc40b121753cc9938ea22b193b936806e81748d7b5fe1384931c7379cd", "canonical_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters:default", "parent_id": "xcsh-docs:resources:alert_policy:properties:routes:notification_parameters", "path": "docs/guides/resources--alert_policy--properties--routes--notification_parameters--default.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "notification_parameters", "default"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/routes/notification_parameters/default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.notification_parameters.default for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.notification_parameters.default

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md)
- [Property reference](resources--alert_policy--reference.md)
- [routes](resources--alert_policy--properties--routes.md)
- [routes.notification_parameters](resources--alert_policy--properties--routes--notification_parameters.md)
- routes.notification_parameters.default

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
default = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.notification_parameters](resources--alert_policy--properties--routes--notification_parameters.md)
- [xcsh_alert_policy](../resources/alert_policy.md)
