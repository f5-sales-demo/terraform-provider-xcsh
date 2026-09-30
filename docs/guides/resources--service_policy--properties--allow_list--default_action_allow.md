---
page_title: "allow_list.default_action_allow"
subcategory: "Security"
description: "allow_list.default_action_allow for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 904, "body_sha256": "sha256:dc0790e8d26e81c2f1b02b3d2c71c93d729a254ebf83a49b4a63a5a94229681e", "canonical_id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_allow", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:allow_list:default_action_allow", "parent_id": "xcsh-docs:resources:service_policy:properties:allow_list", "path": "docs/guides/resources--service_policy--properties--allow_list--default_action_allow.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_list", "default_action_allow"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/allow_list/default_action_allow/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_list.default_action_allow for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# allow_list.default_action_allow

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [allow_list](resources--service_policy--properties--allow_list.md)
- allow_list.default_action_allow

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
default_action_allow = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [allow_list](resources--service_policy--properties--allow_list.md)
- [xcsh_service_policy](../resources/service_policy.md)
