---
page_title: "aws_provider.aws_tgw_site.vpc_attachments"
subcategory: ""
description: "Configuration parameter for vpc attachments."
xcsh_docs: {"aliases": ["aws provider aws tgw site vpc attachments"], "body_bytes": 1672, "body_sha256": "sha256:1b0490c3992c1dbcfb42e6fc73af9565cd06a3ad07c7cad9a63edde6a3eaf06a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site", "path": "documentation/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3200311203003333-3212000013100213-1102221330110123-2132202003301122-3202310210102320-3120022021203120-3110111203122330-1233233203001313", "registry_path": "docs/guides/data-sources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments"], "schema_version": 1, "sections": [{"aliases": ["vpc list"], "anchor": "section", "description": "Collection of items or values", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration parameter for vpc attachments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/)
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

- [vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
