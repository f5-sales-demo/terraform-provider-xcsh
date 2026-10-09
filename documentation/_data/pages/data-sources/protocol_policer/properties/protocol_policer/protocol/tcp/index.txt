---
page_title: "protocol_policer.protocol.tcp"
subcategory: ""
description: "Specification of TCP flag to be matched in a TCP packet."
xcsh_docs: {"aliases": ["protocol policer protocol tcp"], "body_bytes": 1549, "body_sha256": "sha256:f96db58570fdc22f4f260324b12c8050a9c3457de2220c2dd49fe96f67d477d1", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:tcp", "parent_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "path": "documentation/data-sources/protocol_policer/properties/protocol_policer/protocol/tcp/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3001021011222003-2321001123332012-3102133112033010-3212213222211030-2121313332320032-0213311321231211-3312212131132220-1123101033122330", "registry_path": "docs/guides/data-sources--protocol_policer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer", "protocol", "tcp"], "schema_version": 1, "sections": [{"aliases": ["protocol policer protocol tcp flags"], "anchor": "schema-protocol_policer--protocol--tcp--flags", "description": "TCP flag to be matched in a TCP packet.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:tcp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol_policer", "protocol", "tcp", "flags"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/protocol_policer/protocol/tcp/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Specification of TCP flag to be matched in a TCP packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer.protocol.tcp

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/)
- [protocol_policer.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/)
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
