---
page_title: "ingress_egress_gw_ar.hub.spoke_vnets.labels"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub.spoke_vnets.labels for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1473, "body_sha256": "sha256:1a60b70dc9032f7722646fb85b0d6ef8289895ad79563e7f5146bda3ad1d115f", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets:labels", "child_ids": [], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets:labels", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:spoke_vnets", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets--labels.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "spoke_vnets", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub.spoke_vnets.labels for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.hub.spoke_vnets.labels

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets.md)
- ingress_egress_gw_ar.hub.spoke_vnets.labels

<a id="section"></a>

Type: `"single"`. Computed.

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Upstream description:

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

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

- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--spoke_vnets.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
