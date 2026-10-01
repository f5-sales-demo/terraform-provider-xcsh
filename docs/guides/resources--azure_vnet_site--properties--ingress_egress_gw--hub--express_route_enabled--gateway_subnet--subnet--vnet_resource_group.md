---
page_title: "ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group"
subcategory: "Infrastructure"
description: "ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1859, "body_sha256": "sha256:16fa4fd18a9c4531b1c34fff3bdce55c819050512fd0cbc7bda482aac3cf3539", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:gateway_subnet:subnet:vnet_resource_group", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:gateway_subnet:subnet:vnet_resource_group", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:gateway_subnet:subnet", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet--vnet_resource_group.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "gateway_subnet", "subnet", "vnet_resource_group"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/gateway_subnet/subnet/vnet_resource_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.hub](resources--azure_vnet_site--properties--ingress_egress_gw--hub.md)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled.md)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet.md)
- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet.md)
- ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

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

- [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](resources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
