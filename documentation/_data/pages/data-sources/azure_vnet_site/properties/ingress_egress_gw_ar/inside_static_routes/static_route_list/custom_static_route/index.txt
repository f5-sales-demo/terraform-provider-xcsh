---
page_title: "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 4461, "body_sha256": "sha256:80d54b955caaad8c9f4450ffcc1ea3e601878e2edfa244a8a22db12af41f25ad", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:labels", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:nexthop", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route:subnets"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list:custom_static_route", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes:static_route_list", "path": "documentation/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["ingress_egress_gw_ar", "inside_static_routes", "static_route_list", "custom_static_route"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_egress_gw_ar.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route

<a id="section"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

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

## Direct properties

<a id="schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--attrs"></a>

### attrs property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/labels/): complete subsection reference.

- [nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/): complete subsection reference.

- [subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/labels/)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
