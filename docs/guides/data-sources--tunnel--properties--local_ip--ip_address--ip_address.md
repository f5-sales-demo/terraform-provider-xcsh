---
page_title: "local_ip.ip_address.ip_address"
subcategory: ""
description: "local_ip.ip_address.ip_address for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1566, "body_sha256": "sha256:7ef00f3e748af09b24f739d2383fd8b6b6b1af1abfc7fe616941fb8fb3aa8d4b", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:ipv4", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address:ipv6"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "path": "docs/guides/data-sources--tunnel--properties--local_ip--ip_address--ip_address.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/ip_address/ip_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.ip_address for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_ip.ip_address.ip_address

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [local_ip](data-sources--tunnel--properties--local_ip.md)
- [local_ip.ip_address](data-sources--tunnel--properties--local_ip--ip_address.md)
- local_ip.ip_address.ip_address

<a id="section"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

## Direct properties

- [dual_stack](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md): complete subsection reference.

- [ipv4](data-sources--tunnel--properties--local_ip--ip_address--ip_address--ipv4.md): complete subsection reference.

- [ipv6](data-sources--tunnel--properties--local_ip--ip_address--ip_address--ipv6.md): complete subsection reference.

## Next pages

- [local_ip.ip_address.ip_address.dual_stack](data-sources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md)
- [local_ip.ip_address.ip_address.ipv4](data-sources--tunnel--properties--local_ip--ip_address--ip_address--ipv4.md)
- [local_ip.ip_address.ip_address.ipv6](data-sources--tunnel--properties--local_ip--ip_address--ip_address--ipv6.md)
- [local_ip.ip_address](data-sources--tunnel--properties--local_ip--ip_address.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
