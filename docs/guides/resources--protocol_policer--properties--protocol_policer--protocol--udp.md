---
page_title: "protocol_policer.protocol.udp"
subcategory: ""
description: "protocol_policer.protocol.udp for xcsh_protocol_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1024, "body_sha256": "sha256:33a682227e1d2110393fda51daf8993582475352fdc313e20ae69e2719da457b", "canonical_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "child_ids": [], "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "parent_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "path": "docs/guides/resources--protocol_policer--properties--protocol_policer--protocol--udp.md", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["protocol_policer", "protocol", "udp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/protocol/udp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "protocol_policer.protocol.udp for xcsh_protocol_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# protocol_policer.protocol.udp

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md)
- [Property reference](resources--protocol_policer--reference.md)
- [protocol_policer](resources--protocol_policer--properties--protocol_policer.md)
- [protocol_policer.protocol](resources--protocol_policer--properties--protocol_policer--protocol.md)
- protocol_policer.protocol.udp

<a id="section"></a>

Type: `["object", {}]`. Optional.

UDP Packets. Match all UDP packets.

Upstream description:

Match all UDP packets.

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
udp = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [protocol_policer.protocol](resources--protocol_policer--properties--protocol_policer--protocol.md)
- [xcsh_protocol_policer](../resources/protocol_policer.md)
