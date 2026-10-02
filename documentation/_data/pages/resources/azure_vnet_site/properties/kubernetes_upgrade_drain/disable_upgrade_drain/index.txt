---
page_title: "kubernetes_upgrade_drain.disable_upgrade_drain"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["kubernetes upgrade drain disable upgrade drain"], "body_bytes": 1386, "body_sha256": "sha256:6e60df701b7f790d511cc71e5d59fe9f5fdddaf0389b7ca6914566b5aea0f65b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:kubernetes_upgrade_drain", "path": "documentation/resources/azure_vnet_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0002220310300102-0001032002222311-3303232132222222-3002013331013201-3121133202111113-2231102130012112-1133133321322220-0101200122310020", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "disable_upgrade_drain"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.disable_upgrade_drain

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/kubernetes_upgrade_drain/)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/kubernetes_upgrade_drain/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
