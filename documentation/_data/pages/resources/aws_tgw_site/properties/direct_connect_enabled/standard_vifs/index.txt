---
page_title: "direct_connect_enabled.standard_vifs"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["direct connect enabled standard vifs"], "body_bytes": 1321, "body_sha256": "sha256:74af67deff62a0b5c0d7112cdcb713c5aaca42bb7cfe8d1f401dec946dabdf35", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:standard_vifs", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "path": "documentation/resources/aws_tgw_site/properties/direct_connect_enabled/standard_vifs/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0203322303111331-2203130323333113-1223320210331012-3020332313111013-3323101013301101-3033212221200222-2013111020230323-2103231132202000", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled", "standard_vifs"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/standard_vifs/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.standard_vifs

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/)
- direct_connect_enabled.standard_vifs

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for standard vifs.

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
standard_vifs = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
