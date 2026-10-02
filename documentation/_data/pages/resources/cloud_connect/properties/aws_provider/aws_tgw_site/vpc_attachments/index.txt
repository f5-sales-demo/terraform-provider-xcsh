---
page_title: "aws_provider.aws_tgw_site.vpc_attachments"
subcategory: ""
description: "Configuration parameter for vpc attachments."
xcsh_docs: {"aliases": ["aws provider aws tgw site vpc attachments"], "body_bytes": 1772, "body_sha256": "sha256:cc7e9509a8a418e198bbde7b33e5e69e1605c1597df4b3b88f955d6da51a55dd", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site", "path": "documentation/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments"], "schema_version": 1, "sections": [{"aliases": ["vpc list"], "anchor": "section", "description": "Collection of items or values", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list:ConflictingListObjectAttributes:custom_routing,default_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list:ConflictingListObjectAttributes:custom_routing,manual_routing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list:ConflictingListObjectAttributes:custom_routing,default_route", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list:ConflictingListObjectAttributes:default_route,manual_routing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list:ConflictingListObjectAttributes:custom_routing,manual_routing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:manual_routing", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list:ConflictingListObjectAttributes:default_route,manual_routing", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:manual_routing", "type": "conflicts"}, {"anchor": "schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--vpc_id", "enforcement": "provider-schema", "group": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list:RequiredListObjectAttributes:vpc_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "type": "requires"}], "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for vpc attachments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- aws_provider.aws_tgw_site.vpc_attachments

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
vpc_attachments {
  # Configure direct properties listed below.
}
```

## Direct properties

- [vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
