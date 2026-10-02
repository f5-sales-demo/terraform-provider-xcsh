---
page_title: "query_params.check_not_present"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["query params check not present"], "body_bytes": 1326, "body_sha256": "sha256:526061cf18b5c7d702afd90e010f44f2ef9eccda0f9eccb649b80a19e6affdd6", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:query_params:check_not_present", "parent_id": "xcsh-docs:resources:service_policy_rule:properties:query_params", "path": "documentation/resources/service_policy_rule/properties/query_params/check_not_present/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2233100010311321-0013300020103330-2232123000232130-2313333300323232-2021121122103012-3321120200110032-2222013322013111-3111323320011102", "registry_path": "docs/guides/resources--service_policy_rule--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["query_params", "check_not_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/query_params/check_not_present/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# query_params.check_not_present

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/)
- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/)
- query_params.check_not_present

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

- [query_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/properties/query_params/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
