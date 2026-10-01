---
page_title: "egress_virtual_private_gateway"
subcategory: "Infrastructure"
description: "egress_virtual_private_gateway for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2131, "body_sha256": "sha256:8ddcb102717ff9ab0e60a6b501bd03d085b93f8af9c6fe038166d7c9a9f2c332", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:egress_virtual_private_gateway", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:egress_virtual_private_gateway", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "docs/guides/resources--aws_vpc_site--properties--egress_virtual_private_gateway.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["egress_virtual_private_gateway"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/egress_virtual_private_gateway/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_virtual_private_gateway for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_virtual_private_gateway

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- egress_virtual_private_gateway

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

With this option, egress site traffic will be routed through an Virtual Private Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"vgw_id\"]"
}
```

Terraform syntax:

```terraform
egress_virtual_private_gateway {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-egress_virtual_private_gateway--vgw_id"></a>

### vgw_id property

Type: `"string"`. Optional.

Existing Virtual Private Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 21,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

## Next pages

- [Property reference](resources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
