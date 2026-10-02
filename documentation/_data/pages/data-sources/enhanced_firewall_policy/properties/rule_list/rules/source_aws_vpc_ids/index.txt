---
page_title: "rule_list.rules.source_aws_vpc_ids"
subcategory: ""
description: "List of VPC Identifiers in AWS."
xcsh_docs: {"aliases": ["rule list rules source aws vpc ids"], "body_bytes": 2769, "body_sha256": "sha256:911ad9549c8c9301421248456c4229181bf2a31f734ad926db55e519f370200e", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_aws_vpc_ids", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "path": "documentation/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_aws_vpc_ids/index.md", "product": "distributed-cloud", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1303330131133222-2011312300131102-3103333300023021-3113302301021331-2022112113023002-1302120120322100-0312301222031211-3331302220131002", "registry_path": "docs/guides/data-sources--enhanced_firewall_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rule_list", "rules", "source_aws_vpc_ids"], "schema_version": 1, "sections": [{"aliases": ["vpc id"], "anchor": "schema-rule_list--rules--source_aws_vpc_ids--vpc_id", "description": "List of VPC Identifiers in AWS.", "document_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:source_aws_vpc_ids", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rule_list", "rules", "source_aws_vpc_ids", "vpc_id"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/rules/source_aws_vpc_ids/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of VPC Identifiers in AWS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.source_aws_vpc_ids

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/)
- rule_list.rules.source_aws_vpc_ids

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for source aws vpc ids.

Upstream description:

List of VPC Identifiers in AWS.

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

<a id="schema-rule_list--rules--source_aws_vpc_ids--vpc_id"></a>

### vpc_id property

Type: `["list", "string"]`. Computed.

AWS VPC List. List of VPC Identifiers in AWS.

Upstream description:

List of VPC Identifiers in AWS.

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

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/)
