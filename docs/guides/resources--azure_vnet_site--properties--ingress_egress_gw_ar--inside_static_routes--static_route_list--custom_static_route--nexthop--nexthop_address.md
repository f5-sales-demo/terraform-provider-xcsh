---
page_title: "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3711, "body_sha256": "sha256:e4b203485594c33d2aa57a73e81e124f4ba388cfa353a83947ac9d1a9189d9e6", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:dual_stack", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv4", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address:ipv6"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop:nexthop_address", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "inside_static_routes", "static_route_list", "custom_static_route", "nexthop", "nexthop_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop.md)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dual_stack](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md): complete subsection reference.

- [ipv4](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md): complete subsection reference.

- [ipv6](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6.md)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--properties--ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
