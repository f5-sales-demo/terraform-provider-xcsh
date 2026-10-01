---
page_title: "allow_list.default_action_next_policy"
subcategory: "Security"
description: "allow_list.default_action_next_policy for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1041, "body_sha256": "sha256:4e3e2aa71722dcb407bb1c0c36b6893d9e0ca7a2634b5ca0a60cbf38bb84ca13", "canonical_id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_next_policy", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_next_policy", "parent_id": "xcsh-docs:resources:service_policy:properties:allow_list", "path": "docs/guides/resources--service_policy--properties--allow_list--default_action_next_policy.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_list", "default_action_next_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/allow_list/default_action_next_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_list.default_action_next_policy for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_list.default_action_next_policy

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [allow_list](resources--service_policy--properties--allow_list.md)
- allow_list.default_action_next_policy

<a id="section"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
default_action_next_policy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [allow_list](resources--service_policy--properties--allow_list.md)
- [xcsh_service_policy](../resources/service_policy.md)
