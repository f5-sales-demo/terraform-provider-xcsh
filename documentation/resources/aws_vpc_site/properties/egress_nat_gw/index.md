---
page_title: "egress_nat_gw"
subcategory: "Infrastructure"
description: "With this option, egress site traffic will be routed through an Network Address Translation(NAT) Gateway."
xcsh_docs: {"aliases": ["egress nat gw"], "body_bytes": 2285, "body_sha256": "sha256:b5c3739d7eaaaede26da11dd3f5b211769b5ada39734b87a726bdee7eaf2f177", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:egress_nat_gw", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "documentation/resources/aws_vpc_site/properties/egress_nat_gw/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2132333123021112-2121130110202202-3232330110112120-1023112123222122-2013332222002322-0212030223100002-1131133120313012-3111311320301310", "registry_path": "docs/guides/resources--aws_vpc_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["egress_nat_gw"], "schema_version": 1, "sections": [{"aliases": ["nat gw id"], "anchor": "schema-egress_nat_gw--nat_gw_id", "description": "Exclusive with", "document_id": "xcsh-docs:resources:aws_vpc_site:properties:egress_nat_gw", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["egress_nat_gw", "nat_gw_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/egress_nat_gw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With this option, egress site traffic will be routed through an Network Address Translation(NAT) Gateway.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_nat_gw

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_vpc_site/)
