---
page_title: "any_ip"
subcategory: ""
description: "any_ip for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1297, "body_sha256": "sha256:20976047b7932185268483bbeccf5e4c45ecba25507f4ea6c81899ee723338d2", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:any_ip", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:any_ip", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "docs/guides/resources--service_policy_rule--properties--any_ip.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["any_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/any_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "any_ip for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_ip

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- any_ip

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: any\_ip, ip\_matcher, ip\_prefix\_list\] Enable this option

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

OneOf alternatives in this subsection:

- [any_ip](resources--service_policy_rule--properties--any_ip.md#section)
- [ip_matcher](resources--service_policy_rule--properties--ip_matcher.md#section)
- [ip_prefix_list](resources--service_policy_rule--properties--ip_prefix_list.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_ip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
