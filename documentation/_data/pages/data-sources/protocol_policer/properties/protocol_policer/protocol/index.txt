---
page_title: "protocol_policer.protocol"
subcategory: ""
description: "Protocol and protocol specific flags to be matched in packet."
xcsh_docs: {"aliases": ["protocol policer protocol"], "body_bytes": 2523, "body_sha256": "sha256:d5aeec9e14985a9783171c9ed8d021514841156e48d4db156918c7fe81d1fd8c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:dns", "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:icmp", "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:tcp", "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:udp"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "parent_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer", "path": "documentation/data-sources/protocol_policer/properties/protocol_policer/protocol/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1312022101001313-0132230313033120-0020110303200101-0211120303200322-1110030330001232-3020130333132232-2210120210310010-3122131200302131", "registry_path": "docs/guides/data-sources--protocol_policer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer", "protocol"], "schema_version": 1, "sections": [{"aliases": ["protocol policer protocol dns"], "anchor": "section", "description": "Match all DNS packets including UDP and TCP.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:dns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol_policer", "protocol", "dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol policer protocol icmp"], "anchor": "section", "description": "ICMP message type to match in packet.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:icmp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["protocol_policer", "protocol", "icmp"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol policer protocol tcp"], "anchor": "section", "description": "Specification of TCP flag to be matched in a TCP packet.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:tcp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["protocol_policer", "protocol", "tcp"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol policer protocol udp"], "anchor": "section", "description": "Match all UDP packets.", "document_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:udp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["protocol_policer", "protocol", "udp"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/protocol_policer/protocol/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Protocol and protocol specific flags to be matched in packet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer.protocol

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/)
- protocol_policer.protocol

<a id="section"></a>

Type: `"single"`. Computed.

Protocol and protocol specific flags to be matched in packet.

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

## Direct properties

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/dns/): complete subsection reference.

- [icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/icmp/): complete subsection reference.

- [tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/tcp/): complete subsection reference.

- [udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/udp/): complete subsection reference.

## Next pages

- [protocol_policer.protocol.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/dns/)
- [protocol_policer.protocol.icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/icmp/)
- [protocol_policer.protocol.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/tcp/)
- [protocol_policer.protocol.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/protocol/udp/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/protocol_policer/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/)
