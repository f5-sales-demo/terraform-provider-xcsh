---
page_title: "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 1781, "body_sha256": "sha256:6ac0e13995824e2e308d5bfd635575710f023c60be3085a252d25470101edb30", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:labels", "child_ids": [], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:labels", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--labels.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "outside_static_routes", "static_route_list", "custom_static_route", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route.md)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

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
labels {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
