---
page_title: "remote_ip.ip.dual_stack"
subcategory: ""
description: "remote_ip.ip.dual_stack for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1267, "body_sha256": "sha256:09a80843ec2a4c7fcd193e8f5624103ae896f97916c0cebd47502992caa4ece8", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack:ipv4", "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack:ipv6"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip", "path": "docs/guides/data-sources--tunnel--properties--remote_ip--ip--dual_stack.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip", "ip", "dual_stack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/ip/dual_stack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.ip.dual_stack for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.ip.dual_stack

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [remote_ip](data-sources--tunnel--properties--remote_ip.md)
- [remote_ip.ip](data-sources--tunnel--properties--remote_ip--ip.md)
- remote_ip.ip.dual_stack

<a id="section"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

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

- [ipv4](data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv4.md): complete subsection reference.

- [ipv6](data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv6.md): complete subsection reference.

## Next pages

- [remote_ip.ip.dual_stack.ipv4](data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv4.md)
- [remote_ip.ip.dual_stack.ipv6](data-sources--tunnel--properties--remote_ip--ip--dual_stack--ipv6.md)
- [remote_ip.ip](data-sources--tunnel--properties--remote_ip--ip.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
