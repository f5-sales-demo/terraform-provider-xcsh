---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2502, "body_sha256": "sha256:bd17b6470aeada2ab4f52b75b02de5c9a93b6f35254323e7eca23448f69caec2", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:route_server_subnet", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:route_server_subnet:auto", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:route_server_subnet:subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:route_server_subnet:subnet_param"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:route_server_subnet", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled", "route_server_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/route_server_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for route server subnet.

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
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

## Direct properties

- [auto](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--auto.md): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet.md): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet_param.md): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--auto.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet_param.md)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
