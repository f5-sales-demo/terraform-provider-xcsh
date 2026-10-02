---
page_title: "any_ip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["any ip"], "body_bytes": 1658, "body_sha256": "sha256:7f031a2096238c823c70b0d735bb1c259f7292be27265779d3a98693ba635f6d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:any_ip", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "documentation/resources/service_policy_rule/properties/any_ip/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2323033333212103-1131233331312011-2132333310220111-1000201311003232-3201032233230003-0111300321031231-0000200020031132-3212101202111300", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["any_ip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/any_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_ip

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
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

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/any_ip/#section)
- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_matcher/#section)
- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/ip_prefix_list/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_ip = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
