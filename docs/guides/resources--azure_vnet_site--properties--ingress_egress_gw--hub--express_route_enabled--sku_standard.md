---
page_title: "ingress_egress_gw.hub.express_route_enabled.sku_standard"
subcategory: "Infrastructure"
description: "ingress_egress_gw.hub.express_route_enabled.sku_standard for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1388, "body_sha256": "sha256:d93b3c1373e28eec87773acd934640a27bc2c4a12c35d25d79cc02359fbd6866", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:sku_standard", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled:sku_standard", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub:express_route_enabled", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled--sku_standard.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "hub", "express_route_enabled", "sku_standard"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/sku_standard/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.hub.express_route_enabled.sku_standard for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.hub.express_route_enabled.sku_standard

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.hub](resources--azure_vnet_site--properties--ingress_egress_gw--hub.md)
- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled.md)
- ingress_egress_gw.hub.express_route_enabled.sku_standard

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for sku standard.

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
sku_standard = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw.hub.express_route_enabled](resources--azure_vnet_site--properties--ingress_egress_gw--hub--express_route_enabled.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
