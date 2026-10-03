---
page_title: "aws_parameters.custom_security_group"
subcategory: ""
description: "Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface. Supported only for sites deployed on existing VPC."
xcsh_docs: {"aliases": ["aws parameters custom security group"], "body_bytes": 3329, "body_sha256": "sha256:894fd949dfaa67c7518489d34f5b738bea68da61d96f5e83422334143ac0aaf4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:custom_security_group", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters", "path": "documentation/data-sources/aws_tgw_site/properties/aws_parameters/custom_security_group/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1200330320023230-0322020333133231-1333202000222201-1300210113301131-3022222013031320-2323003101133112-0221332102112021-1220220320212322", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_parameters", "custom_security_group"], "schema_version": 1, "sections": [{"aliases": ["aws parameters custom security group inside security group id"], "anchor": "schema-aws_parameters--custom_security_group--inside_security_group_id", "description": "Security Group ID to be attached to SLI(Site Local Inside) Interface.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:custom_security_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "custom_security_group", "inside_security_group_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws parameters custom security group outside security group id"], "anchor": "schema-aws_parameters--custom_security_group--outside_security_group_id", "description": "Security Group ID to be attached to SLO(Site Local Outside) Interface.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:aws_parameters:custom_security_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_parameters", "custom_security_group", "outside_security_group_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/aws_parameters/custom_security_group/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface. Supported only for sites deployed on existing VPC.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_parameters.custom_security_group

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/)
- aws_parameters.custom_security_group

<a id="section"></a>

Type: `"single"`. Computed.

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

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

## Direct properties

<a id="schema-aws_parameters--custom_security_group--inside_security_group_id"></a>

### inside_security_group_id property

Type: `"string"`. Computed.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="schema-aws_parameters--custom_security_group--outside_security_group_id"></a>

### outside_security_group_id property

Type: `"string"`. Computed.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

## Next pages

- [aws_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/aws_parameters/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
