---
page_title: "remote_ip.ip"
subcategory: ""
description: "remote_ip.ip for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1359, "body_sha256": "sha256:7c36344746d98785728b18cb27877499824faf3eb6406e4bb6324934aaa8036e", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:dual_stack", "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:ipv4", "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip:ipv6"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:remote_ip:ip", "parent_id": "xcsh-docs:data-sources:tunnel:properties:remote_ip", "path": "docs/guides/data-sources--tunnel--properties--remote_ip--ip.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip", "ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/remote_ip/ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.ip for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.ip

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [remote_ip](data-sources--tunnel--properties--remote_ip.md)
- remote_ip.ip

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

- [dual_stack](data-sources--tunnel--properties--remote_ip--ip--dual_stack.md): complete subsection reference.

- [ipv4](data-sources--tunnel--properties--remote_ip--ip--ipv4.md): complete subsection reference.

- [ipv6](data-sources--tunnel--properties--remote_ip--ip--ipv6.md): complete subsection reference.

## Next pages

- [remote_ip.ip.dual_stack](data-sources--tunnel--properties--remote_ip--ip--dual_stack.md)
- [remote_ip.ip.ipv4](data-sources--tunnel--properties--remote_ip--ip--ipv4.md)
- [remote_ip.ip.ipv6](data-sources--tunnel--properties--remote_ip--ip--ipv6.md)
- [remote_ip](data-sources--tunnel--properties--remote_ip.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
