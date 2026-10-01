---
page_title: "protocol_policer.protocol"
subcategory: ""
description: "protocol_policer.protocol for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 2411, "body_sha256": "sha256:a372cb2037280a8a341920c8b7ffdb97e00d844ece89d1e55bdd0088c3f7eebb", "canonical_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "child_ids": ["xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp"], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "parent_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer", "path": "docs/guides/resources--protocol_policer--properties--protocol_policer--protocol.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protocol_policer", "protocol"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/protocol/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protocol_policer.protocol for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer.protocol

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md)
- [Property reference](resources--protocol_policer--reference.md)
- [protocol_policer](resources--protocol_policer--properties--protocol_policer.md)
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

- [dns](resources--protocol_policer--properties--protocol_policer--protocol--dns.md): complete subsection reference.

- [icmp](resources--protocol_policer--properties--protocol_policer--protocol--icmp.md): complete subsection reference.

- [tcp](resources--protocol_policer--properties--protocol_policer--protocol--tcp.md): complete subsection reference.

- [udp](resources--protocol_policer--properties--protocol_policer--protocol--udp.md): complete subsection reference.

## Next pages

- [protocol_policer.protocol.dns](resources--protocol_policer--properties--protocol_policer--protocol--dns.md)
- [protocol_policer.protocol.icmp](resources--protocol_policer--properties--protocol_policer--protocol--icmp.md)
- [protocol_policer.protocol.tcp](resources--protocol_policer--properties--protocol_policer--protocol--tcp.md)
- [protocol_policer.protocol.udp](resources--protocol_policer--properties--protocol_policer--protocol--udp.md)
- [protocol_policer](resources--protocol_policer--properties--protocol_policer.md)
- [xcsh_protocol_policer](../resources/protocol_policer.md)
