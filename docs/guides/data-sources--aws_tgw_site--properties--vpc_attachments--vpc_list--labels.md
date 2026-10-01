---
page_title: "vpc_attachments.vpc_list.labels"
subcategory: ""
description: "vpc_attachments.vpc_list.labels for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1076, "body_sha256": "sha256:29058d667dd9c9f78dfc8188217462b47a94f83712772ad96d7f2f8a55038f87", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels", "child_ids": [], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list", "path": "docs/guides/data-sources--aws_tgw_site--properties--vpc_attachments--vpc_list--labels.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vpc_attachments", "vpc_list", "labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vpc_attachments/vpc_list/labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vpc_attachments.vpc_list.labels for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc_attachments.vpc_list.labels

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [vpc_attachments](data-sources--aws_tgw_site--properties--vpc_attachments.md)
- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--properties--vpc_attachments--vpc_list.md)
- vpc_attachments.vpc_list.labels

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

- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--properties--vpc_attachments--vpc_list.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
