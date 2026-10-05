---
page_title: "protocol_policer.protocol.tcp"
subcategory: ""
description: "Specification of TCP flag to be matched in a TCP packet."
xcsh_docs: {"aliases": ["protocol policer protocol tcp"], "body_bytes": 1992, "body_sha256": "sha256:8ef23bc11b3f8f597f93ea1c6c06e837944b02f04c6347af1893597ad6902251", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "parent_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "path": "documentation/resources/protocol_policer/properties/protocol_policer/protocol/tcp/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0122020023100021-0020003133211223-1133032331211030-2223201003111100-0330213031012130-2223022230231131-1122312030003211-1123003022032011", "registry_path": "docs/guides/resources--protocol_policer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer", "protocol", "tcp"], "schema_version": 1, "sections": [{"aliases": ["protocol policer protocol tcp flags"], "anchor": "schema-protocol_policer--protocol--tcp--flags", "description": "TCP flag to be matched in a TCP packet.", "document_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol_policer", "protocol", "tcp", "flags"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/protocol/tcp/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Specification of TCP flag to be matched in a TCP packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer.protocol.tcp

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/)
- [protocol_policer.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/)
- protocol_policer.protocol.tcp

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tcp {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-protocol_policer--protocol--tcp--flags"></a>

### flags property

Type: `["list", "string"]`. Optional.

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

- [protocol_policer.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
