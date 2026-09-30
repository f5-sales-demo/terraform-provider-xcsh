---
page_title: "protocol_policer.protocol"
subcategory: ""
description: "protocol_policer.protocol for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1775, "body_sha256": "sha256:d87a3c864245a60d1d102662ed7abc48044513089be6c159b7b224359d579609", "canonical_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "child_ids": ["xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:dns", "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:icmp", "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:tcp", "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol:udp"], "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer:protocol", "parent_id": "xcsh-docs:data-sources:protocol_policer:properties:protocol_policer", "path": "docs/guides/data-sources--protocol_policer--properties--protocol_policer--protocol.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protocol_policer", "protocol"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/properties/protocol_policer/protocol/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protocol_policer.protocol for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# protocol_policer.protocol

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md)
- [Property reference](data-sources--protocol_policer--reference.md)
- [protocol_policer](data-sources--protocol_policer--properties--protocol_policer.md)
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

- [dns](data-sources--protocol_policer--properties--protocol_policer--protocol--dns.md): complete subsection reference.

- [icmp](data-sources--protocol_policer--properties--protocol_policer--protocol--icmp.md): complete subsection reference.

- [tcp](data-sources--protocol_policer--properties--protocol_policer--protocol--tcp.md): complete subsection reference.

- [udp](data-sources--protocol_policer--properties--protocol_policer--protocol--udp.md): complete subsection reference.

## Next pages

- [protocol_policer.protocol.dns](data-sources--protocol_policer--properties--protocol_policer--protocol--dns.md)
- [protocol_policer.protocol.icmp](data-sources--protocol_policer--properties--protocol_policer--protocol--icmp.md)
- [protocol_policer.protocol.tcp](data-sources--protocol_policer--properties--protocol_policer--protocol--tcp.md)
- [protocol_policer.protocol.udp](data-sources--protocol_policer--properties--protocol_policer--protocol--udp.md)
- [protocol_policer](data-sources--protocol_policer--properties--protocol_policer.md)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md)
