---
page_title: "local_ip.ip_address.ip_address.dual_stack"
subcategory: ""
description: "local_ip.ip_address.ip_address.dual_stack for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1656, "body_sha256": "sha256:ea6a8801a4a0d62d6fa34332920a35a840b90ccb8f060ae1d16ed03dd1feae38", "canonical_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv4", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack:ipv6"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "path": "docs/guides/resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address", "dual_stack"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/ip_address/dual_stack/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.ip_address.dual_stack for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.ip_address.dual_stack

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [local_ip](resources--tunnel--properties--local_ip.md)
- [local_ip.ip_address](resources--tunnel--properties--local_ip--ip_address.md)
- [local_ip.ip_address.ip_address](resources--tunnel--properties--local_ip--ip_address--ip_address.md)
- local_ip.ip_address.ip_address.dual_stack

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

- [ipv4](resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv4.md): complete subsection reference.

- [ipv6](resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv6.md): complete subsection reference.

## Next pages

- [local_ip.ip_address.ip_address.dual_stack.ipv4](resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv4.md)
- [local_ip.ip_address.ip_address.dual_stack.ipv6](resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack--ipv6.md)
- [local_ip.ip_address.ip_address](resources--tunnel--properties--local_ip--ip_address--ip_address.md)
- [xcsh_tunnel](../resources/tunnel.md)
