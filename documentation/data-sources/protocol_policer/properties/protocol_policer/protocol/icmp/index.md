---
page_title: "protocol_policer.protocol.icmp"
subcategory: ""
description: "ICMP message type to match in packet."
xcsh_docs: {"aliases": ["protocol policer protocol icmp"], "body_bytes": 1506, "body_sha256": "sha256:81e0843770dd0ceb745d02fa7b0555a9731780f740321187d07ab92e4dfe19c4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:icmp", "parent_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "path": "documentation/data-sources/protocol_policer/properties/protocol_policer/protocol/icmp/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3010200132302232-3303103021120313-0321121120213302-2202132312133203-2223333213312221-0232313320330001-1032002102302323-3101123033311201", "registry_path": "docs/guides/data-sources--protocol_policer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer", "protocol", "icmp"], "schema_version": 1, "sections": [{"aliases": ["protocol policer protocol icmp type"], "anchor": "schema-protocol_policer--protocol--icmp--type", "description": "ICMP message type to be matched in packet.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:icmp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol_policer", "protocol", "icmp", "type"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/protocol_policer/protocol/icmp/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "ICMP message type to match in packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer.protocol.icmp

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/)
- [protocol_policer.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/)
- protocol_policer.protocol.icmp

<a id="section"></a>

Type: `"single"`. Computed.

ICMP Packet Type. ICMP message type to match in packet.

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

<a id="schema-protocol_policer--protocol--icmp--type"></a>

### type property

Type: `["list", "string"]`. Computed.

\[Enum: ECHO\_REPLY|ECHO\_REQUEST|ALL\_ICMP\_MSG\] ICMP message type to be matched in packet.
Possible values are \`ECHO\_REPLY\`, \`ECHO\_REQUEST\`, \`ALL\_ICMP\_MSG\`. Defaults to
\`ECHO\_REPLY\`.

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
