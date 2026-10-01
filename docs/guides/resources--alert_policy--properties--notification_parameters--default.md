---
page_title: "notification_parameters.default"
subcategory: "Monitoring"
description: "notification_parameters.default for xcsh_alert_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1028, "body_sha256": "sha256:5d56863322204dfa75c6707a40ca946c91024aa66ac7e11ce4c200059fe56149", "canonical_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:default", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_policy:properties:notification_parameters:default", "parent_id": "xcsh-docs:resources:alert_policy:properties:notification_parameters", "path": "docs/guides/resources--alert_policy--properties--notification_parameters--default.md", "provider_name": "alert_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["notification_parameters", "default"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_policy/properties/notification_parameters/default/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "notification_parameters.default for xcsh_alert_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# notification_parameters.default

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md)
- [Property reference](resources--alert_policy--reference.md)
- [notification_parameters](resources--alert_policy--properties--notification_parameters.md)
- notification_parameters.default

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

- [notification_parameters](resources--alert_policy--properties--notification_parameters.md)
- [xcsh_alert_policy](../resources/alert_policy.md)
