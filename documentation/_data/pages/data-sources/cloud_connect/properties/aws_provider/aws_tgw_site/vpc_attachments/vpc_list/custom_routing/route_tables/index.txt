---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables"
subcategory: ""
description: "Route Tables."
xcsh_docs: {"aliases": ["aws provider aws tgw site vpc attachments vpc list custom routing route tables"], "body_bytes": 5391, "body_sha256": "sha256:bfcebdd199f66a95d7a26d1ad3d1601c58567d85ac438d40fe1c99bab962177c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing:route_tables", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing", "path": "documentation/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/route_tables/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0311100001212222-1230333310312121-3021122331021313-1021001230001000-2232211320020210-2112020110301122-2210102002220203-1131110313113001", "registry_path": "docs/guides/data-sources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "custom_routing", "route_tables"], "schema_version": 1, "sections": [{"aliases": ["route table id"], "anchor": "schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--route_table_id", "description": "Route table ID.", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing:route_tables", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "custom_routing", "route_tables", "route_table_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["static routes"], "anchor": "schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--static_routes", "description": "List of Static Routes.", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing:route_tables", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "custom_routing", "route_tables", "static_routes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/route_tables/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Route Tables.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- [aws_provider.aws_tgw_site.vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables

<a id="section"></a>

Type: `"list"`. Computed.

List of route tables. Route Tables.

Upstream description:

Route Tables.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 200,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--route_table_id"></a>

### route_table_id property

Type: `"string"`. Computed.

Route table ID. Route table ID.

Upstream description:

Route table ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--static_routes"></a>

### static_routes property

Type: `["list", "string"]`. Computed.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
