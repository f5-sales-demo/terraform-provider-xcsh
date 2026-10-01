---
page_title: "ingress_egress_gw.hub.spoke_vnets.manual"
subcategory: "Infrastructure"
description: "ingress_egress_gw.hub.spoke_vnets.manual for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1287, "body_sha256": "sha256:a63d7511e60a2a2a08a0a19f0e0beae2f114d486593ba6785a1002094301e4a0", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:manual", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:manual", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--manual.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets", "manual"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/manual/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.hub.spoke_vnets.manual for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.spoke_vnets.manual

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.hub](resources--azure_vnet_site--properties--ingress_egress_gw--hub.md)
- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets.md)
- ingress_egress_gw.hub.spoke_vnets.manual

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
manual = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw.hub.spoke_vnets](resources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
