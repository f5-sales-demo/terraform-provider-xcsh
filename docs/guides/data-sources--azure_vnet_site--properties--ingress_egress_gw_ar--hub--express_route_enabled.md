---
page_title: "ingress_egress_gw_ar.hub.express_route_enabled"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.hub.express_route_enabled for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 7001, "body_sha256": "sha256:313b85d898dbb23d12c5878125f88252687615d663da6f6ecb7a8d5b286455f6", "canonical_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:advertise_to_route_server", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:auto_asn", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:connections", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:do_not_advertise_to_route_server", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:gateway_subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:route_server_subnet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:site_registration_over_express_route", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:site_registration_over_internet", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:sku_ergw1az", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:sku_ergw2az", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:sku_high_perf", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled:sku_standard"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub:express_route_enabled", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:hub", "path": "docs/guides/data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "hub", "express_route_enabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.hub.express_route_enabled for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw_ar.hub.express_route_enabled

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
- [Property reference](data-sources--azure_vnet_site--reference.md)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar.md)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- ingress_egress_gw_ar.hub.express_route_enabled

<a id="section"></a>

Type: `"single"`. Computed.

Express Route Configuration. Express Route Configuration.

Upstream description:

Express Route Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_express_route\",\"site_registration_over_internet\"]",
  "x-ves-oneof-field-sku_choice": "[\"sku_ergw1az\",\"sku_ergw2az\",\"sku_high_perf\",\"sku_standard\"]",
  "x-ves-oneof-field-spoke_vnet_routes": "[\"advertise_to_route_server\",\"do_not_advertise_to_route_server\"]"
}
```

## Direct properties

- [advertise_to_route_server](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--advertise_to_route_server.md): complete subsection reference.

- [auto_asn](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--auto_asn.md): complete subsection reference.

- [connections](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md): complete subsection reference.

<a id="schema-ingress_egress_gw_ar--hub--express_route_enabled--custom_asn"></a>

### custom_asn property

Type: `"number"`. Computed.

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

Upstream description:

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  }
}
```

- [do_not_advertise_to_route_server](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--do_not_advertise_to_route_server.md): complete subsection reference.

- [gateway_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet.md): complete subsection reference.

- [route_server_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet.md): complete subsection reference.

- [site_registration_over_express_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_express_route.md): complete subsection reference.

- [site_registration_over_internet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_internet.md): complete subsection reference.

- [sku_ergw1az](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_ergw1az.md): complete subsection reference.

- [sku_ergw2az](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_ergw2az.md): complete subsection reference.

- [sku_high_perf](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_high_perf.md): complete subsection reference.

- [sku_standard](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_standard.md): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--advertise_to_route_server.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.auto_asn](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--auto_asn.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--connections.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--do_not_advertise_to_route_server.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_express_route.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_internet.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_ergw1az.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_ergw2az.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_high_perf.md)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_standard](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub--express_route_enabled--sku_standard.md)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--properties--ingress_egress_gw_ar--hub.md)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md)
