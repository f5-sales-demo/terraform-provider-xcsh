---
page_title: "request_constraints.max_header_count_none"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["request constraints max header count none"], "body_bytes": 1384, "body_sha256": "sha256:373f803355dbe3a6b7b68ed533de5d54d62fa2c50f0c5d5cd984b036f6c7d62a", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints:max_header_count_none", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:request_constraints", "path": "documentation/resources/service_policy_rule/properties/request_constraints/max_header_count_none/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2031333020012231-3100132320333220-1321002111211212-1121322112320113-1033112301122111-3221032333001212-3332000110133331-1031013122313000", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["request_constraints", "max_header_count_none"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/request_constraints/max_header_count_none/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# request_constraints.max_header_count_none

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/)
- request_constraints.max_header_count_none

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for max header count none.

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
max_header_count_none = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [request_constraints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/request_constraints/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
