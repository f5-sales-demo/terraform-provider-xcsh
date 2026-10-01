---
page_title: "vn_config.outside_static_routes"
subcategory: ""
description: "vn_config.outside_static_routes for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1157, "body_sha256": "sha256:5c61dc99a38df3bcfd1df62b3f756243d4f1b91d7d9f23368f096f70b0903ba0", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config:outside_static_routes", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vn_config", "path": "docs/guides/data-sources--aws_tgw_site--properties--vn_config--outside_static_routes.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "outside_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vn_config/outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.outside_static_routes for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- vn_config.outside_static_routes

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

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

- [static_route_list](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md): complete subsection reference.

## Next pages

- [vn_config.outside_static_routes.static_route_list](data-sources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md)
- [vn_config](data-sources--aws_tgw_site--properties--vn_config.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
