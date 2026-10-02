---
page_title: "vnet"
subcategory: "Infrastructure"
description: "This defines choice about Azure VNet for a view."
xcsh_docs: {"aliases": ["vnet"], "body_bytes": 1638, "body_sha256": "sha256:feb6bd2e9ce0b4896a1b29927e31d482c4696ba9edababe545642182b29ff3b9", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet", "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "documentation/data-sources/azure_vnet_site/properties/vnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1211302120213333-3002011311101310-0133223332111330-1313013330132133-2022210131232102-0210120323332022-0023122121132121-3100103210221122", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vnet"], "schema_version": 1, "sections": [{"aliases": ["existing vnet"], "anchor": "section", "description": "Resource group and name of existing Azure VNet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:existing_vnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vnet", "existing_vnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["new vnet"], "anchor": "section", "description": "Parameters to create a new Azure VNet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet:new_vnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vnet", "new_vnet"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/vnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines choice about Azure VNet for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- vnet

<a id="section"></a>

Type: `"single"`. Computed.

Defines choice about Azure VNet for a view.

Upstream description:

This defines choice about Azure VNet for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_vnet\",\"new_vnet\"]"
}
```

## Direct properties

- [existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/): complete subsection reference.

- [new_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/new_vnet/): complete subsection reference.

## Next pages

- [vnet.existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/)
- [vnet.new_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/new_vnet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
