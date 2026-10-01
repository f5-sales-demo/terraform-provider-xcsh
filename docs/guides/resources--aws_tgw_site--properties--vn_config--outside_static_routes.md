---
page_title: "vn_config.outside_static_routes"
subcategory: ""
description: "vn_config.outside_static_routes for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1422, "body_sha256": "sha256:f38169e2ad6995804df38c1fad50bf574764d05c6031f80888eea8eb6bec0f1e", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes:static_route_list"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:outside_static_routes", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--outside_static_routes.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "outside_static_routes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/outside_static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.outside_static_routes for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.outside_static_routes

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- vn_config.outside_static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
```

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
outside_static_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [static_route_list](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md): complete subsection reference.

## Next pages

- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--properties--vn_config--outside_static_routes--static_route_list.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
