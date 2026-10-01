---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list"
subcategory: ""
description: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 4147, "body_sha256": "sha256:ef57aa6c20464556bc2a42a6f4b2c590342cbddec8b2f86965b0e8734eed8e74", "canonical_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "child_ids": ["xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing", "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route", "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:labels", "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:manual_routing"], "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments", "path": "docs/guides/data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
- [Property reference](data-sources--cloud_connect--reference.md)
- [aws_provider](data-sources--cloud_connect--properties--aws_provider.md)
- [aws_provider.aws_tgw_site](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list

<a id="section"></a>

Type: `"list"`. Computed.

VPC List. Collection of items or values

Upstream description:

Collection of items or values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

## Direct properties

- [custom_routing](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing.md): complete subsection reference.

- [default_route](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route.md): complete subsection reference.

- [labels](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--labels.md): complete subsection reference.

- [manual_routing](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--manual_routing.md): complete subsection reference.

<a id="schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--vpc_id"></a>

### vpc_id property

Type: `"string"`. Computed.

Enter the VPC ID of the VPC to be attached.

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
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

## Next pages

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--labels.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--manual_routing.md)
- [aws_provider.aws_tgw_site.vpc_attachments](data-sources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
