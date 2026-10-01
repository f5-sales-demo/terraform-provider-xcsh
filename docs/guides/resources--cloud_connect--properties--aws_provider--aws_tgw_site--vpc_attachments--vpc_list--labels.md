---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels"
subcategory: ""
description: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1519, "body_sha256": "sha256:e9e5929a00f7cae6510f0c0e261afde3a868682a3c845250b6bac6cd6d80159e", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:labels", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:labels", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "path": "docs/guides/resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--labels.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [aws_provider](resources--cloud_connect--properties--aws_provider.md)
- [aws_provider.aws_tgw_site](resources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Add labels for the VPC attachment. These labels can then be used in policies such as enhanced
firewall.

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

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
