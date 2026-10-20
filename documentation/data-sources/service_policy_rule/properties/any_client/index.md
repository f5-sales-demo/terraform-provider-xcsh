---
page_title: "any_client"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["any client"], "body_bytes": 1745, "body_sha256": "sha256:62d92a0abf674849ceb4f5dc108e00970df874ea00626cdd7458830df96257c2", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_rule:properties:any_client", "parent_id": "xcsh-docs:data-sources:service_policy_rule:reference", "path": "documentation/data-sources/service_policy_rule/properties/any_client/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2201201330130100-1123001021023020-3303233322313203-3331102111111233-1313213321012112-3332023310232102-0333301231202200-2113111111230220", "registry_path": "docs/guides/data-sources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["any_client"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_rule/properties/any_client/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_client

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/)
- any_client

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: any\_client, client\_name, client\_name\_matcher, client\_selector,
ip\_threat\_category\_list\] Enable this option

Additional upstream details:

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

- [any_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/any_client/#section)
- [client_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/#schema-client_name)
- [client_name_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/client_name_matcher/#section)
- [client_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/client_selector/#section)
- [ip_threat_category_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_rule/properties/ip_threat_category_list/#section)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.
