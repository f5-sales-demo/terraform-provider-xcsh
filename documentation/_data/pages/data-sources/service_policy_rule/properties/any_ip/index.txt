---
page_title: "any_ip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["any ip"], "body_bytes": 1630, "body_sha256": "sha256:9760e68953b93a29d88387d357205ceda04cfc8bfeddcbd5e28d8af671856371", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:any_ip", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "documentation/data-sources/service_policy_rule/properties/any_ip/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2010322331300203-1113233230123032-2023011223232330-0323222220231000-0113102103000012-1012103310023010-2111000200320203-1232223201332301", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["any_ip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/any_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_ip

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- any_ip

<a id="section"></a>

Type: `["object", {}]`. Computed.

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

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/any_ip/#section)
- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/ip_matcher/#section)
- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/ip_prefix_list/#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
