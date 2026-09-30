---
page_title: "vpc_attachments"
subcategory: ""
description: "vpc_attachments for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 818, "body_sha256": "sha256:448ea29b23b03af151fbdff3054e10ff4a21c9f877bc29882abb78fd02e020b8", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments:vpc_list"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:vpc_attachments", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "docs/guides/data-sources--aws_tgw_site--properties--vpc_attachments.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vpc_attachments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/vpc_attachments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vpc_attachments for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vpc_attachments

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- vpc_attachments

<a id="section"></a>

Type: `"single"`. Computed.

Spoke VPCs to be attached to the AWS TGW Site.

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

- [vpc_list](data-sources--aws_tgw_site--properties--vpc_attachments--vpc_list.md): complete subsection reference.

## Next pages

- [vpc_attachments.vpc_list](data-sources--aws_tgw_site--properties--vpc_attachments--vpc_list.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
