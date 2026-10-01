---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables"
subcategory: ""
description: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1797, "body_sha256": "sha256:32e7c1142fd19abee04b9c88bfba1479348c08a21593b84e87ec86a2347952b8", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:all_route_tables", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:all_route_tables", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route", "path": "docs/guides/resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route--all_route_tables.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "default_route", "all_route_tables"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/all_route_tables/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [aws_provider](resources--cloud_connect--properties--aws_provider.md)
- [aws_provider.aws_tgw_site](resources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route.md)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all route tables.

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
all_route_tables = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
