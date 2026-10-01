---
page_title: "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 3330, "body_sha256": "sha256:27b1487a124ab5a3550da36639c3a2e4c5b25511ca913b7c0cd4ff5137f6285b", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:subnets", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv4", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:subnets:ipv6"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route:subnets", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "outside_static_routes", "static_route_list", "custom_static_route", "subnets"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route.md)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipv4](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md): complete subsection reference.

- [ipv6](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6.md)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--properties--ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
