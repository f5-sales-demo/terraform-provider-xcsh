---
page_title: "vn_config.outside_static_routes.static_route_list.custom_static_route.labels"
subcategory: ""
description: "vn_config.outside_static_routes.static_route_list.custom_static_route.labels for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1619, "body_sha256": "sha256:7799be26c50a390820c5e661fdd3be808959974ce763394fbdbb2ae3a837b121", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:labels", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route:labels", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list:custom_static_route", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route--labels.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "outside_static_routes", "static_route_list", "custom_static_route", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/outside_static_routes/static_route_list/custom_static_route/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.outside_static_routes.static_route_list.custom_static_route.labels for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes.static_route_list.custom_static_route.labels

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [vn_config.outside_static_routes](resources--aws_tgw_site--properties--vn_config--outside_static_routes.md)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md)
- vn_config.outside_static_routes.static_route_list.custom_static_route.labels

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

- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list--custom_static_route.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
