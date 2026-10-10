---
page_title: "rule_list.rules.spec.request_constraints.max_cookie_value_size_none"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rule list rules spec request constraints max cookie value size none"], "body_bytes": 1565, "body_sha256": "sha256:280b6b632961e412c6bc06a91b13124282f2fe8c3ccc463f4e061064221d4617", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints:max_cookie_value_size_none", "parent_id": "xcsh-docs:resources:service_policy:properties:rule_list:rules:spec:request_constraints", "path": "documentation/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_cookie_value_size_none/index.md", "product": "distributed-cloud", "provider_name": "service_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2003211303022323-3332223010332321-2301222030020233-1021330312213310-0011101303003300-0102320320312332-1102220311121300-0113023322332102", "registry_path": "docs/guides/resources--service_policy--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "spec", "request_constraints", "max_cookie_value_size_none"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/properties/rule_list/rules/spec/request_constraints/max_cookie_value_size_none/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["service_policyCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.spec.request_constraints.max_cookie_value_size_none

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/)
- [rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/)
- [rule_list.rules.spec.request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/properties/rule_list/rules/spec/request_constraints/)
- rule_list.rules.spec.request_constraints.max_cookie_value_size_none

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max cookie value size none.

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

Terraform syntax:

```terraform
max_cookie_value_size_none = {}
```

This is an empty object or choice marker. It has no direct properties.
