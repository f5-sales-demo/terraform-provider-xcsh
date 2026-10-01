---
page_title: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list"
subcategory: ""
description: "aws_provider.aws_tgw_site.vpc_attachments.vpc_list for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 4759, "body_sha256": "sha256:f48acb033a5a21079c132c80befba1fe5118e386b96c3a1b3940d0ecf09a94bb", "canonical_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "child_ids": ["xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:custom_routing", "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:default_route", "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:labels", "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list:manual_routing"], "collection_id": "xcsh-docs:resources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments:vpc_list", "parent_id": "xcsh-docs:resources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments", "path": "docs/guides/resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments", "vpc_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/vpc_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_provider.aws_tgw_site.vpc_attachments.vpc_list for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site.vpc_attachments.vpc_list

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md)
- [Property reference](resources--cloud_connect--reference.md)
- [aws_provider](resources--cloud_connect--properties--aws_provider.md)
- [aws_provider.aws_tgw_site](resources--cloud_connect--properties--aws_provider--aws_tgw_site.md)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

VPC List. Collection of items or values

Upstream description:

Collection of items or values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("vpc_id"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "default_route"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "manual_routing"),
  validators.ConflictingListObjectAttributes("default_route",
    "manual_routing")}
```

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

Terraform syntax:

```terraform
vpc_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_routing](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing.md): complete subsection reference.

- [default_route](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route.md): complete subsection reference.

- [labels](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--labels.md): complete subsection reference.

- [manual_routing](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--manual_routing.md): complete subsection reference.

<a id="schema-aws_provider--aws_tgw_site--vpc_attachments--vpc_list--vpc_id"></a>

### vpc_id property

Type: `"string"`. Optional.

Enter the VPC ID of the VPC to be attached.

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

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--custom_routing.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--default_route.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--labels.md)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments--vpc_list--manual_routing.md)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--properties--aws_provider--aws_tgw_site--vpc_attachments.md)
- [xcsh_cloud_connect](../resources/cloud_connect.md)
