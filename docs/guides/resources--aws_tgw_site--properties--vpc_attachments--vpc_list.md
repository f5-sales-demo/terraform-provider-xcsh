---
page_title: "vpc_attachments.vpc_list"
subcategory: ""
description: "vpc_attachments.vpc_list for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2772, "body_sha256": "sha256:61a01a61878bba40c1fbf87b162225484bfa6ffe32f530b500334645a858ad31", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list:labels"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments:vpc_list", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vpc_attachments", "path": "docs/guides/resources--aws_tgw_site--properties--vpc_attachments--vpc_list.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vpc_attachments", "vpc_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vpc_attachments/vpc_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vpc_attachments.vpc_list for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vpc_attachments.vpc_list

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vpc_attachments](resources--aws_tgw_site--properties--vpc_attachments.md)
- vpc_attachments.vpc_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of VPC attachments to transit gateway.

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

- [labels](resources--aws_tgw_site--properties--vpc_attachments--vpc_list--labels.md): complete subsection reference.

<a id="schema-vpc_attachments--vpc_list--vpc_id"></a>

### vpc_id property

Type: `"string"`. Optional.

VPC ID. Information about existing VPC.

Upstream description:

Information about existing VPC.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

## Next pages

- [vpc_attachments.vpc_list.labels](resources--aws_tgw_site--properties--vpc_attachments--vpc_list--labels.md)
- [vpc_attachments](resources--aws_tgw_site--properties--vpc_attachments.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
