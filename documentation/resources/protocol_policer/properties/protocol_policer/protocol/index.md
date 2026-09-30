---
page_title: "protocol_policer.protocol"
subcategory: ""
description: "protocol_policer.protocol for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 2961, "body_sha256": "sha256:d30adeabfaf9a288315a3f22ab0107b99444e5aee8d507b502dbba0c03060711", "child_ids": ["xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:dns", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:icmp", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:tcp", "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp"], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "parent_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer", "path": "documentation/resources/protocol_policer/properties/protocol_policer/protocol/index.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["protocol_policer", "protocol"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/protocol/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protocol_policer.protocol for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# protocol_policer.protocol

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/)
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

- [dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/dns/): complete subsection reference.

- [icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/icmp/): complete subsection reference.

- [tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/tcp/): complete subsection reference.

- [udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/udp/): complete subsection reference.

## Next pages

- [protocol_policer.protocol.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/dns/)
- [protocol_policer.protocol.icmp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/icmp/)
- [protocol_policer.protocol.tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/tcp/)
- [protocol_policer.protocol.udp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/udp/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/)
- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
