---
page_title: "rules.criteria"
subcategory: ""
description: "Match criteria of the packet to apply the NAT Rule."
xcsh_docs: {"aliases": ["rules criteria"], "body_bytes": 4827, "body_sha256": "sha256:b4703985982afff172b63119567ea33a2580a6cb43da1e104454200d3c953a2a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_inside_network", "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_network", "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules", "path": "documentation/resources/nat_policy/properties/rules/criteria/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,icmp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,icmp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:icmp,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:icmp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:site_local_inside_network,site_local_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:icmp,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:tcp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:any,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:icmp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria:ConflictingObjectAttributes:tcp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria"], "schema_version": 1, "sections": [{"aliases": ["rules criteria any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:any", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria destination cidr"], "anchor": "schema-rules--criteria--destination_cidr", "description": "Destination IP of the packet to match.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "destination_cidr"], "syntax": "attribute", "type": "list"}, {"aliases": ["rules criteria icmp"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:icmp", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "icmp"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria site local inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "site_local_inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria site local network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:site_local_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "site_local_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["rules criteria source cidr"], "anchor": "schema-rules--criteria--source_cidr", "description": "Source IP of the packet to match.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["rules", "criteria", "source_cidr"], "syntax": "attribute", "type": "list"}, {"aliases": ["rules criteria tcp"], "anchor": "section", "description": "Action to apply on the packet if the NAT rule is applied.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "criteria", "tcp"], "syntax": "block", "type": "object"}, {"aliases": ["rules criteria udp"], "anchor": "section", "description": "Action to apply on the packet if the NAT rule is applied.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:udp", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rules", "criteria", "udp"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Match criteria of the packet to apply the NAT Rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- rules.criteria

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match criteria of the packet to apply the NAT Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "icmp"),
  validators.ConflictingObjectAttributes("any",
    "tcp"),
  validators.ConflictingObjectAttributes("any",
    "udp"),
  validators.ConflictingObjectAttributes("icmp",
    "tcp"),
  validators.ConflictingObjectAttributes("icmp",
    "udp"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-protocol_choice": "[\"any\",\"icmp\",\"tcp\",\"udp\"]"
}
```

Terraform syntax:

```terraform
criteria {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/any/): complete subsection reference.

<a id="schema-rules--criteria--destination_cidr"></a>

### destination_cidr property

Type: `["list", "string"]`. Optional.

Destination IP. Destination IP of the packet to match.

Upstream description:

Destination IP of the packet to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/icmp/): complete subsection reference.

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/site_local_network/): complete subsection reference.

<a id="schema-rules--criteria--source_cidr"></a>

### source_cidr property

Type: `["list", "string"]`. Optional.

Source IP. Source IP of the packet to match.

Upstream description:

Source IP of the packet to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/): complete subsection reference.

- [udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/): complete subsection reference.

## Next pages

- [rules.criteria.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/any/)
- [rules.criteria.icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/icmp/)
- [rules.criteria.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/site_local_inside_network/)
- [rules.criteria.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/site_local_network/)
- [rules.criteria.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/)
- [rules.criteria.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/udp/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
