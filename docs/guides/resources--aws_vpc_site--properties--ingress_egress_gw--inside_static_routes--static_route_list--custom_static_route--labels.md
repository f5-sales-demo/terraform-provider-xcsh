---
page_title: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels"
subcategory: "Infrastructure"
description: "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1705, "body_sha256": "sha256:944c258a68401a7b321458b9c9134ee57dc7009c49519727fe36dd78595b3432", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:labels", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route:labels", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:ingress_egress_gw:inside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--labels.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "inside_static_routes", "static_route_list", "custom_static_route", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [ingress_egress_gw](resources--aws_vpc_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes.md)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list.md)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

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

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--properties--ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
