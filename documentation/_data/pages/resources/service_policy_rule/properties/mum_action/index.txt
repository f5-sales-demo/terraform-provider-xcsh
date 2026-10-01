---
page_title: "mum_action"
subcategory: ""
description: "mum_action for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1981, "body_sha256": "sha256:ffc4f16649324533f14c736093df7a740f50ba867d59533a226ebbc24b8e2fdb", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:mum_action:default", "xcsh-docs:resources:service_policy_rule:properties:mum_action:skip_processing"], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:mum_action", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/mum_action/index.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["mum_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/mum_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "mum_action for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mum_action

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- mum_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify behavior for a matching request. The modification could be to entirely skip processing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default",
    "skip_processing")}
```

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

Terraform syntax:

```terraform
mum_action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/default/): complete subsection reference.

- [skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/skip_processing/): complete subsection reference.

## Next pages

- [mum_action.default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/default/)
- [mum_action.skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/mum_action/skip_processing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
