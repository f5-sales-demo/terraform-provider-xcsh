---
page_title: "rules.criteria.tcp"
subcategory: ""
description: "Action to apply on the packet if the NAT rule is applied."
xcsh_docs: {"aliases": ["rules criteria tcp"], "body_bytes": 1928, "body_sha256": "sha256:937886b24865feeaca2820751a1f4204e573754f385a793ec2ffb2e1908bcb96", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp", "parent_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria", "path": "documentation/resources/nat_policy/properties/rules/criteria/tcp/index.md", "product": "distributed-cloud", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033", "registry_path": "docs/guides/resources--nat_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rules", "criteria", "tcp"], "schema_version": 1, "sections": [{"aliases": ["destination port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--criteria--tcp--destination_port--port", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--destination_port--port", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--destination_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--destination_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.tcp.destination_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:destination_port:no_port_match", "type": "conflicts"}], "schema_path": ["rules", "criteria", "tcp", "destination_port"], "syntax": "block", "type": "object"}, {"aliases": ["source port"], "anchor": "section", "description": "Port match of the request can be a range or a specific port.", "document_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-rules--criteria--tcp--source_port--port", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--source_port--port", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--source_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "type": "conflicts"}, {"anchor": "schema-rules--criteria--tcp--source_port--port_ranges", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:port,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:no_port_match,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port:no_port_match", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "rules.criteria.tcp.source_port:ConflictingObjectAttributes:no_port_match,port_ranges", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:nat_policy:properties:rules:criteria:tcp:source_port:no_port_match", "type": "conflicts"}], "schema_path": ["rules", "criteria", "tcp", "source_port"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nat_policy/properties/rules/criteria/tcp/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action to apply on the packet if the NAT rule is applied.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nat_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [rules.criteria.tcp.destination_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/destination_port/)
- [rules.criteria.tcp.source_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/tcp/source_port/)
- [rules.criteria](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/properties/rules/criteria/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nat_policy/)
