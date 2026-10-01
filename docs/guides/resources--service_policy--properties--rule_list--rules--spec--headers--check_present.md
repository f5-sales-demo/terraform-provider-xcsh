---
page_title: "rule_list.rules.spec.headers.check_present"
subcategory: "Security"
description: "rule_list.rules.spec.headers.check_present for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1358, "body_sha256": "sha256:790c510eb14fe1446ef075c1b1dc4754b606559840ff201a8d2b326ced8abdb5", "canonical_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:headers:check_present", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:headers:check_present", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:headers", "path": "docs/guides/resources--service_policy--properties--rule_list--rules--spec--headers--check_present.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "spec", "headers", "check_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/headers/check_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.spec.headers.check_present for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.headers.check_present

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Property reference](resources--service_policy--reference.md)
- [rule_list](resources--service_policy--properties--rule_list.md)
- [rule_list.rules](resources--service_policy--properties--rule_list--rules.md)
- [rule_list.rules.spec](resources--service_policy--properties--rule_list--rules--spec.md)
- [rule_list.rules.spec.headers](resources--service_policy--properties--rule_list--rules--spec--headers.md)
- rule_list.rules.spec.headers.check_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rule_list.rules.spec.headers](resources--service_policy--properties--rule_list--rules--spec--headers.md)
- [xcsh_service_policy](../resources/service_policy.md)
