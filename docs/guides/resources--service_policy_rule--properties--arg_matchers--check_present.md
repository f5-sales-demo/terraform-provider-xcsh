---
page_title: "arg_matchers.check_present"
subcategory: ""
description: "arg_matchers.check_present for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1053, "body_sha256": "sha256:c507d66b3831c1ff1a1c537f90d4a23ba88b9b2c3c6ff5ded956a68303dcd4fc", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:check_present", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:check_present", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers", "path": "docs/guides/resources--service_policy_rule--properties--arg_matchers--check_present.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["arg_matchers", "check_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/arg_matchers/check_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "arg_matchers.check_present for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# arg_matchers.check_present

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [arg_matchers](resources--service_policy_rule--properties--arg_matchers.md)
- arg_matchers.check_present

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

- [arg_matchers](resources--service_policy_rule--properties--arg_matchers.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
