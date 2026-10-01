---
page_title: "headers.check_present"
subcategory: ""
description: "headers.check_present for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1023, "body_sha256": "sha256:c4547205b98ca50de765ff462a509332d4ca02c62c66a325167e3bd8250b5ba4", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:headers:check_present", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:headers:check_present", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:headers", "path": "docs/guides/resources--service_policy_rule--properties--headers--check_present.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["headers", "check_present"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/headers/check_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "headers.check_present for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# headers.check_present

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [headers](resources--service_policy_rule--properties--headers.md)
- headers.check_present

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

- [headers](resources--service_policy_rule--properties--headers.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
