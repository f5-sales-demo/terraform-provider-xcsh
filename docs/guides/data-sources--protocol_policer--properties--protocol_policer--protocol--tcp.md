---
page_title: "protocol_policer.protocol.tcp"
subcategory: ""
description: "protocol_policer.protocol.tcp for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1493, "body_sha256": "sha256:75b114ba2a34b9d1c99fe27e4c6625154ef1f55e73853e26d674502e69d57ec1", "canonical_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:tcp", "child_ids": [], "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:tcp", "parent_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "path": "docs/guides/data-sources--protocol_policer--properties--protocol_policer--protocol--tcp.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protocol_policer", "protocol", "tcp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/protocol_policer/protocol/tcp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protocol_policer.protocol.tcp for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# protocol_policer.protocol.tcp

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md)
- [Property reference](data-sources--protocol_policer--reference.md)
- [protocol_policer](data-sources--protocol_policer--properties--protocol_policer.md)
- [protocol_policer.protocol](data-sources--protocol_policer--properties--protocol_policer--protocol.md)
- protocol_policer.protocol.tcp

<a id="section"></a>

Type: `"single"`. Computed.

Specification of TCP flag to be matched in a TCP packet.

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

<a id="schema-protocol_policer--protocol--tcp--flags"></a>

### flags property

Type: `["list", "string"]`. Computed.

\[Enum: FIN|SYN|RST|PSH|ACK|URG|ALL\_TCP\_FLAGS|KEEPALIVE\] TCP flags. TCP flag to be matched in a
TCP packet. Possible values are \`FIN\`, \`SYN\`, \`RST\`, \`PSH\`, \`ACK\`, \`URG\`,
\`ALL\_TCP\_FLAGS\`, \`KEEPALIVE\`. Defaults to \`FIN\`.

Upstream description:

TCP flag to be matched in a TCP packet.

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

- [protocol_policer.protocol](data-sources--protocol_policer--properties--protocol_policer--protocol.md)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md)
