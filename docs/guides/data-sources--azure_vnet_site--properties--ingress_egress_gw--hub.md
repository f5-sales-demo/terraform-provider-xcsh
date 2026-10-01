---
page_title: "ingress_egress_gw.hub"
subcategory: "Infrastructure"
description: "ingress_egress_gw.hub for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1805, "body_sha256": "sha256:cf9a1406eaab10cd46b31ae389cae64121ae5ea6ee810da204c2202412502ee5", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_disabled", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub:spoke_vnets"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw:hub", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw--hub.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "hub"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.hub for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.hub

<a id="section"></a>

Type: `"single"`. Computed.

Hub VNet type. Hub VNet type.

Upstream description:

Hub VNet type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-express_route_choice": "[\"express_route_disabled\",\"express_route_enabled\"]"
}
```

## Direct properties

- [express_route_disabled](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_disabled.md): complete subsection reference.

- [express_route_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled.md): complete subsection reference.

- [spoke_vnets](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.hub.express_route_disabled](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_disabled.md)
- [ingress_egress_gw.hub.express_route_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled.md)
- [ingress_egress_gw.hub.spoke_vnets](data-sources--azure_vnet_site--properties--ingress_egress_gw--hub--spoke_vnets.md)
- [ingress_egress_gw](data-sources--azure_vnet_site--properties--ingress_egress_gw.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
