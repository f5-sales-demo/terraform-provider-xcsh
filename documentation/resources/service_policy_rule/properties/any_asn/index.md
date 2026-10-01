---
page_title: "any_asn"
subcategory: ""
description: "any_asn for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1648, "body_sha256": "sha256:339f7228fea7ada687ec04daf94896ba7e9517a9e0e94907c62e24cbe371cc2e", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:any_asn", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/any_asn/index.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["any_asn"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/any_asn/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "any_asn for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_asn

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- any_asn

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: any\_asn, asn\_list, asn\_matcher\] Enable this option

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

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/any_asn/#section)
- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_list/#section)
- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/asn_matcher/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_asn = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
