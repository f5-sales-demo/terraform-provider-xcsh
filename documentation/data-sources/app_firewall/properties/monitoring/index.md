---
page_title: "monitoring"
subcategory: "Security"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["monitoring"], "body_bytes": 1095, "body_sha256": "sha256:07dda267ae04808354ee6e1e69473aab8990125e5a688467239b4a27517721d1", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:monitoring", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/monitoring/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0132332001110103-0010201300032310-0100330312110132-3301223331032312-2131202331131103-3320010120230110-0301101202322330-0001011321201113", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["monitoring"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/monitoring/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# monitoring

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- monitoring

<a id="section"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
