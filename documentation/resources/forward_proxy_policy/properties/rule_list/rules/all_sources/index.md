---
page_title: "rule_list.rules.all_sources"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["rule list rules all sources"], "body_bytes": 1451, "body_sha256": "sha256:e679a6ece921ad70c1a58283ab95b24e988327e93f30b166017948a74668f214", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules:all_sources", "parent_id": "xcsh-docs:resources:forward_proxy_policy:properties:rule_list:rules", "path": "documentation/resources/forward_proxy_policy/properties/rule_list/rules/all_sources/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2222002001020101-1111232232301313-0120210021110313-1122331022321320-1033330212102033-2100023110012212-2222020302330033-1303301022233201", "registry_path": "docs/guides/resources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "all_sources"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forward_proxy_policy/properties/rule_list/rules/all_sources/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.all_sources

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/)
- rule_list.rules.all_sources

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all sources.

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
all_sources = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/properties/rule_list/rules/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/forward_proxy_policy/)
