---
page_title: "arg_matchers.check_not_present"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["arg matchers check not present"], "body_bytes": 1326, "body_sha256": "sha256:097f52a96dd18d8e931450a92c0da7eaecdd82a6d0e0241cbaac1e183ace7e45", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers:check_not_present", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:arg_matchers", "path": "documentation/resources/service_policy_rule/properties/arg_matchers/check_not_present/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1221111202311312-0313312121110232-3032221101101202-1030222010000113-0223233211102003-0122112210311023-0033223223120233-2001311233123032", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["arg_matchers", "check_not_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/arg_matchers/check_not_present/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# arg_matchers.check_not_present

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/)
- arg_matchers.check_not_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [arg_matchers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/arg_matchers/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
