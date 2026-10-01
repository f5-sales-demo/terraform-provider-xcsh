---
page_title: "rule_list.rules.spec.request_constraints.max_cookie_value_size_none"
subcategory: "Security"
description: "rule_list.rules.spec.request_constraints.max_cookie_value_size_none for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1482, "body_sha256": "sha256:3d1d0eb838efe9ca4013149d61169b105d240e43625ebee8879576db8e7b64b8", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_value_size_none", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_value_size_none", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--request_constraints--max_cookie_value_size_none.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "request_constraints", "max_cookie_value_size_none"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_cookie_value_size_none/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.request_constraints.max_cookie_value_size_none for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.request_constraints.max_cookie_value_size_none

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [rule_list.rules.spec.request_constraints](resources--service_policy--properties--rule_list--rules--spec--request_constraints.md)
- rule_list.rules.spec.request_constraints.max_cookie_value_size_none

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie value size none.

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
max_cookie_value_size_none = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rule_list.rules.spec.request_constraints](resources--service_policy--properties--rule_list--rules--spec--request_constraints.md)
- [xcsh_service_policy](../resources/service_policy.md)
