---
page_title: "protocol_policer.protocol.udp"
subcategory: ""
description: "Match all UDP packets."
xcsh_docs: {"aliases": ["protocol policer protocol udp"], "body_bytes": 1081, "body_sha256": "sha256:7a11719c4000192021e9bca1d40fc0a00a0b7f3d1fcc1682e2b36abb812a279e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol:udp", "parent_id": "xcsh-docs:resources:protocol_policer:properties:protocol_policer:protocol", "path": "documentation/resources/protocol_policer/properties/protocol_policer/protocol/udp/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3232001013002031-3120311332330222-3021000303020022-2020032310230232-1030303131021111-2312112011011111-3212032322233022-3321311010112131", "registry_path": "docs/guides/resources--protocol_policer--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["protocol_policer", "protocol", "udp"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_policer/properties/protocol_policer/protocol/udp/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Match all UDP packets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# protocol_policer.protocol.udp

Breadcrumbs:

- [xcsh_protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/)
- [protocol_policer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/)
- [protocol_policer.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_policer/properties/protocol_policer/protocol/)
- protocol_policer.protocol.udp

<a id="section"></a>

Type: `["object", {}]`. Optional.

UDP Packets. Match all UDP packets.

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

This is an empty object or choice marker. It has no direct properties.
