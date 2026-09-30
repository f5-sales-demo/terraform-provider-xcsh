---
page_title: "local_ip.ip_address"
subcategory: ""
description: "local_ip.ip_address for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1685, "body_sha256": "sha256:63994a94b268fbbf8457ce6f525e1e1931f1a4b6db13842090d60ef46f89e076", "canonical_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:ip_address:auto", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:ip_address", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip", "path": "docs/guides/resources--tunnel--properties--local_ip--ip_address.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_ip.ip_address

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [local_ip](resources--tunnel--properties--local_ip.md)
- local_ip.ip_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Provides the configuration to pick up source IP and network for transporting encapsulated packet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto",
    "ip_address")}
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
  "x-ves-oneof-field-type": "[\"auto\",\"ip_address\"]"
}
```

Terraform syntax:

```terraform
ip_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto](resources--tunnel--properties--local_ip--ip_address--auto.md): complete subsection reference.

- [ip_address](resources--tunnel--properties--local_ip--ip_address--ip_address.md): complete subsection reference.

- [virtual_network_type](resources--tunnel--properties--local_ip--ip_address--virtual_network_type.md): complete subsection reference.

## Next pages

- [local_ip.ip_address.auto](resources--tunnel--properties--local_ip--ip_address--auto.md)
- [local_ip.ip_address.ip_address](resources--tunnel--properties--local_ip--ip_address--ip_address.md)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--properties--local_ip--ip_address--virtual_network_type.md)
- [local_ip](resources--tunnel--properties--local_ip.md)
- [xcsh_tunnel](../resources/tunnel.md)
