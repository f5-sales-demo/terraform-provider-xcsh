---
page_title: "local_ip.ip_address.ip_address"
subcategory: ""
description: "local_ip.ip_address.ip_address for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 2042, "body_sha256": "sha256:3eb886c09b9b26554c31dee595a6bcf8ba2bb566a60e878c3bde04c6908fd6fe", "canonical_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:dual_stack", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:ipv4", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address:ipv6"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "path": "docs/guides/resources--tunnel--properties--local_ip--ip_address--ip_address.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address", "ip_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/ip_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.ip_address for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.ip_address

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [local_ip](resources--tunnel--properties--local_ip.md)
- [local_ip.ip_address](resources--tunnel--properties--local_ip--ip_address.md)
- local_ip.ip_address.ip_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
```

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

Terraform syntax:

```terraform
ip_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dual_stack](resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md): complete subsection reference.

- [ipv4](resources--tunnel--properties--local_ip--ip_address--ip_address--ipv4.md): complete subsection reference.

- [ipv6](resources--tunnel--properties--local_ip--ip_address--ip_address--ipv6.md): complete subsection reference.

## Next pages

- [local_ip.ip_address.ip_address.dual_stack](resources--tunnel--properties--local_ip--ip_address--ip_address--dual_stack.md)
- [local_ip.ip_address.ip_address.ipv4](resources--tunnel--properties--local_ip--ip_address--ip_address--ipv4.md)
- [local_ip.ip_address.ip_address.ipv6](resources--tunnel--properties--local_ip--ip_address--ip_address--ipv6.md)
- [local_ip.ip_address](resources--tunnel--properties--local_ip--ip_address.md)
- [xcsh_tunnel](../resources/tunnel.md)
