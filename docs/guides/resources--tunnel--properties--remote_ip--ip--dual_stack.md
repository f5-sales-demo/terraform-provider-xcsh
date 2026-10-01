---
page_title: "remote_ip.ip.dual_stack"
subcategory: ""
description: "remote_ip.ip.dual_stack for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1356, "body_sha256": "sha256:6932df13e8c478684fbce290cbe47506ecb44718a081cf5b8607a501ee25b51f", "canonical_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack:ipv4", "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack:ipv6"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "parent_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "path": "docs/guides/resources--tunnel--properties--remote_ip--ip--dual_stack.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip", "ip", "dual_stack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/ip/dual_stack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.ip.dual_stack for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.ip.dual_stack

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [remote_ip](resources--tunnel--properties--remote_ip.md)
- [remote_ip.ip](resources--tunnel--properties--remote_ip--ip.md)
- remote_ip.ip.dual_stack

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

## Direct properties

- [ipv4](resources--tunnel--properties--remote_ip--ip--dual_stack--ipv4.md): complete subsection reference.

- [ipv6](resources--tunnel--properties--remote_ip--ip--dual_stack--ipv6.md): complete subsection reference.

## Next pages

- [remote_ip.ip.dual_stack.ipv4](resources--tunnel--properties--remote_ip--ip--dual_stack--ipv4.md)
- [remote_ip.ip.dual_stack.ipv6](resources--tunnel--properties--remote_ip--ip--dual_stack--ipv6.md)
- [remote_ip.ip](resources--tunnel--properties--remote_ip--ip.md)
- [xcsh_tunnel](../resources/tunnel.md)
