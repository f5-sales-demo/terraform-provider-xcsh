---
page_title: "local_ip.intf"
subcategory: ""
description: "local_ip.intf for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 884, "body_sha256": "sha256:830c78592d75cbe4b9fcb7f0c0e03a7a5f93f39d4b19e0d44a894e7bb5006ad4", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:intf", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:intf:local_intf"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:intf", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip", "path": "docs/guides/data-sources--tunnel--properties--local_ip--intf.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "intf"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/intf/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.intf for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
