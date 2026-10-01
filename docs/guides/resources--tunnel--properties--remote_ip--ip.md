---
page_title: "remote_ip.ip"
subcategory: ""
description: "remote_ip.ip for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1731, "body_sha256": "sha256:635395178272f0eb8558ad5a5e48ba69b5da16da6f71195c44130587f3c7711b", "canonical_id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "child_ids": ["xcsh-docs:resources:tunnel:properties:remote_ip:ip:dual_stack", "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv4", "xcsh-docs:resources:tunnel:properties:remote_ip:ip:ipv6"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:remote_ip:ip", "parent_id": "xcsh-docs:resources:tunnel:properties:remote_ip", "path": "docs/guides/resources--tunnel--properties--remote_ip--ip.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["remote_ip", "ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/remote_ip/ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "remote_ip.ip for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# remote_ip.ip

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [remote_ip](resources--tunnel--properties--remote_ip.md)
- remote_ip.ip

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
ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dual_stack](resources--tunnel--properties--remote_ip--ip--dual_stack.md): complete subsection reference.

- [ipv4](resources--tunnel--properties--remote_ip--ip--ipv4.md): complete subsection reference.

- [ipv6](resources--tunnel--properties--remote_ip--ip--ipv6.md): complete subsection reference.

## Next pages

- [remote_ip.ip.dual_stack](resources--tunnel--properties--remote_ip--ip--dual_stack.md)
- [remote_ip.ip.ipv4](resources--tunnel--properties--remote_ip--ip--ipv4.md)
- [remote_ip.ip.ipv6](resources--tunnel--properties--remote_ip--ip--ipv6.md)
- [remote_ip](resources--tunnel--properties--remote_ip.md)
- [xcsh_tunnel](../resources/tunnel.md)
