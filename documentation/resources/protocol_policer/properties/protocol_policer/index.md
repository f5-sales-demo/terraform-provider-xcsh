---
page_title: "protocol_policer"
subcategory: ""
description: "List of L4 protocol match condition and associated traffic rate limits."
xcsh_docs: {"aliases": ["protocol policer"], "body_bytes": 1917, "body_sha256": "sha256:1aecdf1f8b50413234e31946e390e55bb636248ca3db57add65aa5b9d1b52dcb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:protocol_policer:properties:protocol_policer:policer", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer", "parent_id": "xcsh-docs:resources:protocol_policer:reference", "path": "documentation/resources/protocol_policer/properties/protocol_policer/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1103203330131213-3232300131010200-1312133231130010-2111331121132203-3131111112032203-2132000223302321-0222101101022113-1011231101032132", "registry_path": "docs/guides/resources--protocol_policer--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer:RequiredListObjectAttributes:policer", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:policer", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer"], "schema_version": 1, "sections": [{"aliases": ["protocol policer policer"], "anchor": "section", "description": "Reference to policer object to apply traffic rate limits.", "document_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:policer", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["protocol_policer", "policer"], "syntax": "block", "type": "object"}, {"aliases": ["protocol policer protocol"], "anchor": "section", "description": "Protocol and protocol specific flags to be matched in packet.", "document_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,icmp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,icmp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:icmp,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:icmp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:icmp,tcp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:tcp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:dns,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:icmp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "protocol_policer.protocol:ConflictingObjectAttributes:tcp,udp", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "type": "conflicts"}], "schema_path": ["protocol_policer", "protocol"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "List of L4 protocol match condition and associated traffic rate limits.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/)
- protocol_policer

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of L4 protocol match condition and associated traffic rate limits.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("policer")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protocol_policer {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/policer/): complete subsection reference.

- [protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/): complete subsection reference.
