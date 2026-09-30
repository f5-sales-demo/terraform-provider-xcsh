---
page_title: "deny_list.default_action_deny"
subcategory: "Security"
description: "deny_list.default_action_deny for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 895, "body_sha256": "sha256:e506173caf67c41484d3584ee276a5d0db0e8d9c5cc2f60d047cb2976d9e0c0a", "canonical_id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_deny", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:deny_list:default_action_deny", "parent_id": "xcsh-docs:resources:service_policy:properties:deny_list", "path": "docs/guides/resources--service_policy--properties--deny_list--default_action_deny.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["deny_list", "default_action_deny"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/deny_list/default_action_deny/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "deny_list.default_action_deny for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
