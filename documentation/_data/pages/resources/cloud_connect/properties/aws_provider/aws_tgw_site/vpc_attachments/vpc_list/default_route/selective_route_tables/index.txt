---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables"
subcategory: ""
description: "AWS Route Table."
xcsh_docs: {"aliases": ["aws provider aws tgw site vpc attachments vpc list default route selective route tables"], "body_bytes": 3414, "body_sha256": "sha256:d8b18cd2b659f7dd8158c90e2a9525ad6054ec50df9d71bd0c3bd3665cec67bc", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:selective_route_tables", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route", "path": "documentation/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/selective_route_tables/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2221223103020031-3233303133133123-1311301323021132-0222313210010320-0121213023132100-2030020101221211-3222121030210213-1112030230300030", "registry_path": "docs/guides/resources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "default_route", "selective_route_tables"], "schema_version": 1, "sections": [{"aliases": ["route table id"], "anchor": "schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route--selective_route_tables--route_table_id", "description": "Route table ID.", "document_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route:selective_route_tables", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "default_route", "selective_route_tables", "route_table_id"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/selective_route_tables/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "AWS Route Table.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- [aws_provider.aws_tgw_site.vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for selective route tables.

Upstream description:

AWS Route Table.

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
selective_route_tables {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route--selective_route_tables--route_table_id"></a>

### route_table_id property

Type: `["list", "string"]`. Optional.

Route table ID. Route table ID.

Upstream description:

Route table ID.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/default_route/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
