---
page_title: "protocol_policer.protocol"
subcategory: ""
description: "Protocol and protocol specific flags to be matched in packet."
xcsh_docs: {"aliases": ["protocol policer protocol"], "body_bytes": 3060, "body_sha256": "sha256:da6f5c65a141710404b25d80f725b4cd9764ca811dbca2413548f9e2c70f571a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "parent_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer", "path": "documentation/resources/protocol_policer/properties/protocol_policer/protocol/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-1100220312202001-1132123022333322-1131101032220112-2213013000211201-2032331131032232-0230102201211101-0223303131111000-0332211321020103", "registry_path": "docs/guides/resources--protocol_policer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,icmp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,icmp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:icmp,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:icmp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:icmp,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:tcp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:icmp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:tcp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer", "protocol"], "schema_version": 1, "sections": [{"aliases": ["dns"], "anchor": "section", "description": "Match all DNS packets including UDP and TCP.", "document_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol_policer", "protocol", "dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["icmp"], "anchor": "section", "description": "ICMP message type to match in packet.", "document_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["protocol_policer", "protocol", "icmp"], "syntax": "block", "type": "object"}, {"aliases": ["tcp"], "anchor": "section", "description": "Specification of TCP flag to be matched in a TCP packet.", "document_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["protocol_policer", "protocol", "tcp"], "syntax": "block", "type": "object"}, {"aliases": ["udp"], "anchor": "section", "description": "Match all UDP packets.", "document_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol_policer", "protocol", "udp"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/protocol/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Protocol and protocol specific flags to be matched in packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer.protocol

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/)
- protocol_policer.protocol

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Protocol and protocol specific flags to be matched in packet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dns",
    "icmp"),
  validators.ConflictingObjectAttributes("dns",
    "tcp"),
  validators.ConflictingObjectAttributes("dns",
    "udp"),
  validators.ConflictingObjectAttributes("icmp",
    "tcp"),
  validators.ConflictingObjectAttributes("icmp",
    "udp"),
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
  "x-ves-oneof-field-type": "[\"dns\",\"icmp\",\"tcp\",\"udp\"]"
}
```

Terraform syntax:

```terraform
protocol {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/dns/): complete subsection reference.

- [icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/icmp/): complete subsection reference.

- [tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/tcp/): complete subsection reference.

- [udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/udp/): complete subsection reference.

## Next pages

- [protocol_policer.protocol.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/dns/)
- [protocol_policer.protocol.icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/icmp/)
- [protocol_policer.protocol.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/tcp/)
- [protocol_policer.protocol.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/udp/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
