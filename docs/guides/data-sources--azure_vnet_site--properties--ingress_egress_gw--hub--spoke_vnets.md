---
page_title: "ingress_egress_gw.hub.spoke_vnets"
subcategory: "Infrastructure"
description: "ingress_egress_gw.hub.spoke_vnets for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2563, "body_sha256": "sha256:2d11f34561261d6c4ad015f217114bb65daf6e77fdd27f658298626eb05a19a7", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:auto", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:labels", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:manual", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets:vnet"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "spoke_vnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.hub.spoke_vnets for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.hub.spoke_vnets

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub.md)
- ingress_egress_gw.hub.spoke_vnets

<a id="section"></a>

Type: `"list"`. Computed.

Spoke VNet Peering (Legacy). Spoke VNet Peering.

Upstream description:

Spoke VNet Peering.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [auto](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--auto.md): complete subsection reference.

- [labels](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--labels.md): complete subsection reference.

- [manual](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--manual.md): complete subsection reference.

- [vnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--vnet.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.hub.spoke_vnets.auto](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--auto.md)
- [ingress_egress_gw.hub.spoke_vnets.labels](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--labels.md)
- [ingress_egress_gw.hub.spoke_vnets.manual](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--manual.md)
- [ingress_egress_gw.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets--vnet.md)
- [ingress_egress_gw.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
