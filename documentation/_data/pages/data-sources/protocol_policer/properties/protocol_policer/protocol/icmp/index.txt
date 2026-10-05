---
page_title: "protocol_policer.protocol.icmp"
subcategory: ""
description: "ICMP message type to match in packet."
xcsh_docs: {"aliases": ["protocol policer protocol icmp"], "body_bytes": 1920, "body_sha256": "sha256:b4701a38d86db1240c9dc24b6e3c480e7a81a0d4cc73456d2999f4268d6d4999", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:icmp", "parent_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "path": "documentation/data-sources/protocol_policer/properties/protocol_policer/protocol/icmp/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3010200132302232-3303103021120313-0321121120213302-2202132312133203-2223333213312221-0232313320330001-1032002102302323-3101123033311201", "registry_path": "docs/guides/data-sources--protocol_policer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer", "protocol", "icmp"], "schema_version": 1, "sections": [{"aliases": ["protocol policer protocol icmp type"], "anchor": "schema-protocol_policer--protocol--icmp--type", "description": "ICMP message type to be matched in packet.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:icmp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol_policer", "protocol", "icmp", "type"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/protocol_policer/protocol/icmp/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "ICMP message type to match in packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

Upstream description:

ICMP message type to match in packet.

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

Upstream description:

ICMP message type to be matched in packet.

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

## Next pages

- [protocol_policer.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/)
