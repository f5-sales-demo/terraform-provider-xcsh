---
page_title: "ingress_egress_gw.az_nodes.outside_subnet.subnet.vnet_resource_group"
subcategory: "Infrastructure"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["ingress egress gw az nodes outside subnet subnet vnet resource group"], "body_bytes": 1993, "body_sha256": "sha256:0157a43d2f7d4102285ae7e041cda9ebf3f55627339a0865da73f8ad791be929", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet:vnet_resource_group", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/vnet_resource_group/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1301221000033223-2303001312132101-2322223231133220-2131031210022023-3331130102132131-0212210322001301-2231230312001011-2301220321120021", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "outside_subnet", "subnet", "vnet_resource_group"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/vnet_resource_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.outside_subnet.subnet.vnet_resource_group

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/)
- [ingress_egress_gw.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/)
- [ingress_egress_gw.az_nodes.outside_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/)
- ingress_egress_gw.az_nodes.outside_subnet.subnet.vnet_resource_group

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw.az_nodes.outside_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
