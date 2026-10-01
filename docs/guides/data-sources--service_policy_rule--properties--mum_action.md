---
page_title: "mum_action"
subcategory: ""
description: "mum_action for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1306, "body_sha256": "sha256:995d5e268b5d373e6050e6af1b6933500a9bcbfa1a733bae121c4069a965305a", "canonical_id": "xcsh-docs:data-sources:service_policy_rule:properties:mum_action", "child_ids": ["xcsh-docs:data-sources:service_policy_rule:properties:mum_action:default", "xcsh-docs:data-sources:service_policy_rule:properties:mum_action:skip_processing"], "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:mum_action", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "docs/guides/data-sources--service_policy_rule--properties--mum_action.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["mum_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/mum_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mum_action for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mum_action

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- mum_action

<a id="section"></a>

Type: `"single"`. Computed.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"default\",\"skip_processing\"]"
}
```

## Direct properties

- [default](data-sources--service_policy_rule--properties--mum_action--default.md): complete subsection reference.

- [skip_processing](data-sources--service_policy_rule--properties--mum_action--skip_processing.md): complete subsection reference.

## Next pages

- [mum_action.default](data-sources--service_policy_rule--properties--mum_action--default.md)
- [mum_action.skip_processing](data-sources--service_policy_rule--properties--mum_action--skip_processing.md)
- [Property reference](data-sources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md)
