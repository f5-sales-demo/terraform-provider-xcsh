---
page_title: "rule_list.rules.spec.mum_action"
subcategory: "Security"
description: "rule_list.rules.spec.mum_action for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1869, "body_sha256": "sha256:572248ef47dbbe188022155150b4a4cb0e69cf0d559a2848eda42aa4980c8ae1", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action", "child_ids": ["xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:default", "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action:skip_processing"], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:mum_action", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--mum_action.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "mum_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/mum_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.mum_action for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules.spec.mum_action

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- rule_list.rules.spec.mum_action

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

- [default](resources--service_policy--properties--rule_list--rules--spec--mum_action--default.md): complete subsection reference.

- [skip_processing](resources--service_policy--properties--rule_list--rules--spec--mum_action--skip_processing.md): complete subsection reference.

## Next pages

- [rule_list.rules.spec.mum_action.default](resources--service_policy--properties--rule_list--rules--spec--mum_action--default.md)
- [rule_list.rules.spec.mum_action.skip_processing](resources--service_policy--properties--rule_list--rules--spec--mum_action--skip_processing.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [xcsh_service_policy](../resources/service_policy.md)
