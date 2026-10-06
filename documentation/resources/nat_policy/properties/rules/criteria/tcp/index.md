---
page_title: "rules.criteria.tcp"
subcategory: ""
description: "Action to apply on the packet if the NAT rule is applied."
xcsh_docs: {"aliases": ["rules criteria tcp"], "body_bytes": 1367, "body_sha256": "sha256:9567465497b8403167bf7c8b66c371629893ddea9eb5715b59f98fe1846bc8c0", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "path": "documentation/resources/nat_policy/properties/rules/criteria/tcp/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria", "tcp"], "schema_version": 1, "sections": [{"aliases": ["rules criteria tcp destination port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--criteria--tcp--destination_port--port", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--destination_port--port", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--destination_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--destination_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port:no_port_match", "type": "conflicts"}], "schema_path": ["rules", "criteria", "tcp", "destination_port"], "syntax": "block", "type": "object"}, {"aliases": ["rules criteria tcp source port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--criteria--tcp--source_port--port", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--source_port--port", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--source_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--source_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port:no_port_match", "type": "conflicts"}], "schema_path": ["rules", "criteria", "tcp", "source_port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/tcp/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Action to apply on the packet if the NAT rule is applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.criteria.tcp

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/)
- rules.criteria.tcp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Action to apply on the packet if the NAT rule is applied.

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
tcp {
  # Configure direct properties listed below.
}
```

## Direct properties

- [destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/): complete subsection reference.

- [source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/): complete subsection reference.
