---
page_title: "request_constraints.max_parameter_count_none"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["request constraints max parameter count none"], "body_bytes": 1105, "body_sha256": "sha256:586cba6233ae229a2b6609ae3aa84c0f2c496c9084f1a232cf1a8a5403a6a928", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_parameter_count_none", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "path": "documentation/resources/service_policy_rule/properties/request_constraints/max_parameter_count_none/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3221300301123333-0233121001221021-1311302003000011-2010203103033011-3100131213301022-0011232323122220-0300231113313310-1313310111333232", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["request_constraints", "max_parameter_count_none"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/request_constraints/max_parameter_count_none/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_constraints.max_parameter_count_none

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/)
- request_constraints.max_parameter_count_none

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max parameter count none.

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
max_parameter_count_none = {}
```

This is an empty object or choice marker. It has no direct properties.
