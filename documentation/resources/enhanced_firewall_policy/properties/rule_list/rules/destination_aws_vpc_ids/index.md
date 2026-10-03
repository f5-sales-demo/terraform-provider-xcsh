---
page_title: "rule_list.rules.destination_aws_vpc_ids"
subcategory: ""
description: "List of VPC Identifiers in AWS."
xcsh_docs: {"aliases": ["rule list rules destination aws vpc ids"], "body_bytes": 3180, "body_sha256": "sha256:876bb4be595f52bed7f128778e076fdae3ecb821c9f47e1d3c07f3c0f9a23c67", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:destination_aws_vpc_ids", "parent_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules", "path": "documentation/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_aws_vpc_ids/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0203111133203023-1330113002232331-0110303203000213-1122030333032222-1003021330103020-1330113012310301-0203313022022110-1203121113122201", "registry_path": "docs/guides/resources--enhanced_firewall_policy--reference--group-001.md", "relationships": [{"anchor": "schema-rule_list--rules--destination_aws_vpc_ids--vpc_id", "enforcement": "provider-schema", "group": "rule_list.rules.destination_aws_vpc_ids:RequiredObjectAttributes:vpc_id", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:destination_aws_vpc_ids", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "destination_aws_vpc_ids"], "schema_version": 1, "sections": [{"aliases": ["rule list rules destination aws vpc ids vpc id"], "anchor": "schema-rule_list--rules--destination_aws_vpc_ids--vpc_id", "description": "List of VPC Identifiers in AWS.", "document_id": "xcsh-docs:resources:enhanced_firewall_policy:properties:rule_list:rules:destination_aws_vpc_ids", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "destination_aws_vpc_ids", "vpc_id"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/enhanced_firewall_policy/properties/rule_list/rules/destination_aws_vpc_ids/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of VPC Identifiers in AWS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.destination_aws_vpc_ids

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/)
- rule_list.rules.destination_aws_vpc_ids

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for destination aws vpc ids.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("vpc_id")}
```

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
destination_aws_vpc_ids {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-rule_list--rules--destination_aws_vpc_ids--vpc_id"></a>

### vpc_id property

Type: `["list", "string"]`. Optional.

AWS VPC List. List of VPC Identifiers in AWS.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/properties/rule_list/rules/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/enhanced_firewall_policy/)
