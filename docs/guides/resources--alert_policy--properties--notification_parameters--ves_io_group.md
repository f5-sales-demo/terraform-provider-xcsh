---
page_title: "notification_parameters.ves_io_group"
subcategory: "Monitoring"
description: "notification_parameters.ves_io_group for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1066, "body_sha256": "sha256:fec9f5f7e637a36e23400f413e9e2026d996b60c61f8e76c3ed0a507ca6bd3a3", "canonical_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:ves_io_group", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:ves_io_group", "parent_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters", "path": "docs/guides/resources--alert_policy--properties--notification_parameters--ves_io_group.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["notification_parameters", "ves_io_group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/notification_parameters/ves_io_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "notification_parameters.ves_io_group for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# notification_parameters.ves_io_group

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md)
- [Property reference](resources--alert_policy--reference.md)
- [notification_parameters](resources--alert_policy--properties--notification_parameters.md)
- notification_parameters.ves_io_group

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ves io group.

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
ves_io_group = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [notification_parameters](resources--alert_policy--properties--notification_parameters.md)
- [xcsh_alert_policy](../resources/alert_policy.md)
