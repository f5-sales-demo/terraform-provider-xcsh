---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables"
subcategory: ""
description: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 5911, "body_sha256": "sha256:5fd6324ad2d483c3b06ea496ca1371ae47921d4e36233edbfe47b9182cdf22b3", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing:route_tables", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing", "path": "documentation/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/route_tables/index.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list", "custom_routing", "route_tables"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/route_tables/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/)
- [aws_provider.aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/)
- [aws_provider.aws_tgw_site.vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of route tables. Route Tables.

Upstream description:

Route Tables.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("static_routes")}
```

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

Terraform syntax:

```terraform
route_tables {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing--route_tables--route_table_id"></a>

### route_table_id property

Type: `"string"`. Optional.

Route table ID. Route table ID.

Upstream description:

Route table ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

Type: `["list", "string"]`. Optional.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 50),
}
```

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

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/custom_routing/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_connect/)
