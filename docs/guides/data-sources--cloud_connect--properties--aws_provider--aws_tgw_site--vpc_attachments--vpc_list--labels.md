---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels"
subcategory: ""
description: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1376, "body_sha256": "sha256:d35681c4661ea624af53db63b6d6e08bceaf82ce2faed01b7cb1a0e0aa9add69", "canonical_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:labels", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:labels", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "path": "docs/guides/data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--labels.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
- [Property reference](data-sources--cloud_connect--reference.md)
- [aws_provider](data-sources--cloud_connect--properties--aws_provider.md)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
