---
page_title: "rule_list.rules.http_list"
subcategory: "Security"
description: "URLListType."
xcsh_docs: {"aliases": ["rule list rules http list"], "body_bytes": 1588, "body_sha256": "sha256:90a17a06904da069f184c35d5f1a3d382ce6fc06abf3f56df40d6aabfe6152a4", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:http_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules", "path": "documentation/data-sources/forward_proxy_policy/properties/rule_list/rules/http_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1103300030210331-2011320120302321-0221032310121221-0022100210211113-0020301030212012-3001201322312321-1200221130112110-0003030210222211", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "http_list"], "schema_version": 1, "sections": [{"aliases": ["http list"], "anchor": "section", "description": "URLs for HTTP connections.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:http_list:http_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "http_list", "http_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/rule_list/rules/http_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "URLListType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.http_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- rule_list.rules.http_list

<a id="section"></a>

Type: `"single"`. Computed.

URLListType.

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

## Direct properties

- [http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/): complete subsection reference.

## Next pages

- [rule_list.rules.http_list.http_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/http_list/http_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
