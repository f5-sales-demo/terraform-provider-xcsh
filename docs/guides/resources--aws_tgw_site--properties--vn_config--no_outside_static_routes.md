---
page_title: "vn_config.no_outside_static_routes"
subcategory: ""
description: "vn_config.no_outside_static_routes for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1030, "body_sha256": "sha256:f3eb66ec343f943b4635ec96f16085944d403770e26a7447cf2248dff7949c64", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_outside_static_routes", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:no_outside_static_routes", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--no_outside_static_routes.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "no_outside_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/no_outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.no_outside_static_routes for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.no_outside_static_routes

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- vn_config.no_outside_static_routes

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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
no_outside_static_routes = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
