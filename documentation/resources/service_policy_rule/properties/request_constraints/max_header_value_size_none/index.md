---
page_title: "request_constraints.max_header_value_size_none"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["request constraints max header value size none"], "body_bytes": 1404, "body_sha256": "sha256:178613920262b0fb1a390eac9e90a049c6f1b3be9c03bca84e0b3700020e934d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_value_size_none", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "path": "documentation/resources/service_policy_rule/properties/request_constraints/max_header_value_size_none/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2102120230210231-2100101103110120-1123203011201001-3031333211322330-1210231111011330-2310102033001102-1320213011013311-0021120132333103", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["request_constraints", "max_header_value_size_none"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/request_constraints/max_header_value_size_none/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_constraints.max_header_value_size_none

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/)
- request_constraints.max_header_value_size_none

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header value size none.

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
max_header_value_size_none = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
