---
page_title: "local_ip.ip_address.ip_address.dual_stack"
subcategory: ""
description: "local_ip.ip_address.ip_address.dual_stack for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1570, "body_sha256": "sha256:eeeb5f8bfc7ffa28276b2f628e386ac463f919578aebde9d446afbf0ff4d636a", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv4", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv6"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address", "path": "docs/guides/data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address", "dual_stack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.ip_address.dual_stack for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.ip_address.dual_stack

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [local_ip](data-sources--tunnel--properties--local_ip.md)
- [local_ip.ip_address](data-sources--tunnel--properties--local_ip--ip_address.md)
- [local_ip.ip_address.ip_address](data-sources--tunnel--properties--local_ip--ip_address--ip_address.md)
- local_ip.ip_address.ip_address.dual_stack

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

- [ipv4](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv4.md): complete subsection reference.

- [ipv6](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv6.md): complete subsection reference.

## Next pages

- [local_ip.ip_address.ip_address.dual_stack.ipv4](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv4.md)
- [local_ip.ip_address.ip_address.dual_stack.ipv6](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv6.md)
- [local_ip.ip_address.ip_address](data-sources--tunnel--properties--local_ip--ip_address--ip_address.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
