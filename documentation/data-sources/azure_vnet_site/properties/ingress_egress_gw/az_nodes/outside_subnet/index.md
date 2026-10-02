---
page_title: "ingress_egress_gw.az_nodes.outside_subnet"
subcategory: "Infrastructure"
description: "Parameters for Azure subnet."
xcsh_docs: {"aliases": ["ingress egress gw az nodes outside subnet"], "body_bytes": 2228, "body_sha256": "sha256:8c5a7778e0549bac4f30f5b585ddcf714a953aa88bb2e5ddfca335543cb95a5b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet_param"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1103012331130120-1311013230120021-1030122101011311-0120100231223022-0113130123323333-2032210010302121-0230330223121323-3232133200102221", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "az_nodes", "outside_subnet"], "schema_version": 1, "sections": [{"aliases": ["subnet"], "anchor": "section", "description": "Parameters for Azure subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "outside_subnet", "subnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["subnet param"], "anchor": "section", "description": "Parameters for creating a new cloud subnet.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:az_nodes:outside_subnet:subnet_param", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "az_nodes", "outside_subnet", "subnet_param"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for Azure subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.az_nodes.outside_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/)
- ingress_egress_gw.az_nodes.outside_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for outside subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

## Direct properties

- [subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/): complete subsection reference.

- [subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet_param/): complete subsection reference.

## Next pages

- [ingress_egress_gw.az_nodes.outside_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/)
- [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet_param/)
- [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
