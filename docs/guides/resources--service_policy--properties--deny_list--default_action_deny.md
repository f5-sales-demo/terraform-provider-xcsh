---
page_title: "deny_list.default_action_deny"
subcategory: "Security"
description: "deny_list.default_action_deny for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 994, "body_sha256": "sha256:d6b9e6e8b142a3a44ab822cba2dde64f380c9429af13b3f20d3c4c796ebb9a68", "canonical_id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_deny", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_deny", "parent_id": "xcsh-docs:resources:service_policy:properties:deny_list", "path": "docs/guides/resources--service_policy--properties--deny_list--default_action_deny.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["deny_list", "default_action_deny"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/deny_list/default_action_deny/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "deny_list.default_action_deny for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# deny_list.default_action_deny

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [deny_list](resources--service_policy--properties--deny_list.md)
- deny_list.default_action_deny

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
default_action_deny = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [deny_list](resources--service_policy--properties--deny_list.md)
- [xcsh_service_policy](../resources/service_policy.md)
