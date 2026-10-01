---
page_title: "egress_nat_gw"
subcategory: "Infrastructure"
description: "egress_nat_gw for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2077, "body_sha256": "sha256:9117bb6d02d66c0727f60e6d89b46f4c7ef29556197c28622e4180b6f9bb2704", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:egress_nat_gw", "child_ids": [], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:egress_nat_gw", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "docs/guides/resources--aws_vpc_site--properties--egress_nat_gw.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["egress_nat_gw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/egress_nat_gw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_nat_gw for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_nat_gw

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- egress_nat_gw

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

With this option, egress site traffic will be routed through an Network Address Translation(NAT)
Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"nat_gw_id\"]"
}
```

Terraform syntax:

```terraform
egress_nat_gw {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-egress_nat_gw--nat_gw_id"></a>

### nat_gw_id property

Type: `"string"`. Optional.

Existing NAT Gateway ID. Exclusive with \[\]

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
    "pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

## Next pages

- [Property reference](resources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
