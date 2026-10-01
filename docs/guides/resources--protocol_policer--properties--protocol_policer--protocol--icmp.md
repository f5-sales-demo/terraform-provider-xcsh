---
page_title: "protocol_policer.protocol.icmp"
subcategory: ""
description: "protocol_policer.protocol.icmp for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1709, "body_sha256": "sha256:5924b2ed1abb8f2348b3915704fec563f51dbe350eca5a571f054fa880faee9b", "canonical_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "child_ids": [], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "parent_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "path": "docs/guides/resources--protocol_policer--properties--protocol_policer--protocol--icmp.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protocol_policer", "protocol", "icmp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/protocol/icmp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protocol_policer.protocol.icmp for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer.protocol.icmp

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md)
- [Property reference](resources--protocol_policer--reference.md)
- [protocol_policer](resources--protocol_policer--properties--protocol_policer.md)
- [protocol_policer.protocol](resources--protocol_policer--properties--protocol_policer--protocol.md)
- protocol_policer.protocol.icmp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
icmp {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-protocol_policer--protocol--icmp--type"></a>

### type property

Type: `["list", "string"]`. Optional.

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

- [protocol_policer.protocol](resources--protocol_policer--properties--protocol_policer--protocol.md)
- [xcsh_protocol_policer](../resources/protocol_policer.md)
