---
page_title: "vnet.existing_vnet.manual_routing"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["vnet existing vnet manual routing"], "body_bytes": 1406, "body_sha256": "sha256:e29a97412bb9fb6403d8378ec77a6a2692be7a94c657ac5fd226a31194bf051d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet:manual_routing", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:vnet:existing_vnet", "path": "documentation/resources/azure_vnet_site/properties/vnet/existing_vnet/manual_routing/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0220300112200000-0311221103023221-2113200132302320-1220030233023213-1302132110133030-3323220131113121-0302333020331100-2130320133332000", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vnet", "existing_vnet", "manual_routing"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/vnet/existing_vnet/manual_routing/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vnet.existing_vnet.manual_routing

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/)
- [vnet.existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/)
- vnet.existing_vnet.manual_routing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
manual_routing = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [vnet.existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/vnet/existing_vnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
