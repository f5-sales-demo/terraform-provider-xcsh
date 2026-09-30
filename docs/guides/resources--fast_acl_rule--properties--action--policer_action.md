---
page_title: "action.policer_action"
subcategory: ""
description: "action.policer_action for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1054, "body_sha256": "sha256:412d01c67a983dd17ddd7e4fb557ffaa4a1a6280cc9a72cb522a7874a87b1398", "canonical_id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "child_ids": ["xcsh-docs:resources:fast_acl_rule:properties:action:policer_action:ref"], "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "parent_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "path": "docs/guides/resources--fast_acl_rule--properties--action--policer_action.md", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["action", "policer_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/properties/action/policer_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "action.policer_action for xcsh_fast_acl_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# action.policer_action

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md)
- [Property reference](resources--fast_acl_rule--reference.md)
- [action](resources--fast_acl_rule--properties--action.md)
- action.policer_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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
policer_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ref](resources--fast_acl_rule--properties--action--policer_action--ref.md): complete subsection reference.

## Next pages

- [action.policer_action.ref](resources--fast_acl_rule--properties--action--policer_action--ref.md)
- [action](resources--fast_acl_rule--properties--action.md)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md)
