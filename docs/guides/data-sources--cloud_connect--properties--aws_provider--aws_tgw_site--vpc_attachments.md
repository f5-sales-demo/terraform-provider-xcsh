---
page_title: "aws_provider.aws_tgw_site.vpc_attachments"
subcategory: ""
description: "aws_provider.aws_tgw_site.vpc_attachments for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 1270, "body_sha256": "sha256:5a10d89a8e3fb48e708cc24b24208d1a35849512d73b2ce9c02d9592defd9393", "canonical_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments", "child_ids": ["xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list"], "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site", "path": "docs/guides/data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site.vpc_attachments for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
- [Property reference](data-sources--cloud_connect--reference.md)
- [aws_provider](data-sources--cloud_connect--properties--aws_provider.md)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- aws_provider.aws_tgw_site.vpc_attachments

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for vpc attachments.

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

- [vpc_list](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
