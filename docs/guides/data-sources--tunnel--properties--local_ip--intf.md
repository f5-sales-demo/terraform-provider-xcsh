---
page_title: "local_ip.intf"
subcategory: ""
description: "local_ip.intf for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 983, "body_sha256": "sha256:cae1cd65606ff44e69afcbc0b8a917a37aca91affa787d8aeb535c1cc8ac3edf", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:intf", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:intf:local_intf"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:intf", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip", "path": "docs/guides/data-sources--tunnel--properties--local_ip--intf.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "intf"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/intf/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.intf for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.intf

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [local_ip](data-sources--tunnel--properties--local_ip.md)
- local_ip.intf

<a id="section"></a>

Type: `"single"`. Computed.

Provides the local interface to pick up source IP and network for transporting encapsulated packet.

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

- [local_intf](data-sources--tunnel--properties--local_ip--intf--local_intf.md): complete subsection reference.

## Next pages

- [local_ip.intf.local_intf](data-sources--tunnel--properties--local_ip--intf--local_intf.md)
- [local_ip](data-sources--tunnel--properties--local_ip.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
