---
page_title: "rule_list.rules.tls_list"
subcategory: "Security"
description: "DomainListType."
xcsh_docs: {"aliases": ["rule list rules tls list"], "body_bytes": 1582, "body_sha256": "sha256:934ebfec9ed822070a1fff6c7c16097e34ca2cb9a4ea46dc91622e0577a101b4", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list:tls_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules", "path": "documentation/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/index.md", "product": "distributed-cloud", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3130031113132122-1313310010332323-3012313301202331-1020032123210331-2033230333200121-0311201313100230-3313331323102333-1023320301301321", "registry_path": "docs/guides/data-sources--forward_proxy_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "tls_list"], "schema_version": 1, "sections": [{"aliases": ["tls list"], "anchor": "section", "description": "Domains in SNI for TLS connections.", "document_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:rule_list:rules:tls_list:tls_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["rule_list", "rules", "tls_list", "tls_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "DomainListType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.tls_list

Breadcrumbs:

- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- rule_list.rules.tls_list

<a id="section"></a>

Type: `"single"`. Computed.

DomainListType.

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

- [tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/tls_list/): complete subsection reference.

## Next pages

- [rule_list.rules.tls_list.tls_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/tls_list/tls_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/properties/rule_list/rules/)
- [xcsh_forward_proxy_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/forward_proxy_policy/)
