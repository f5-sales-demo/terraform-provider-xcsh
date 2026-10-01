---
page_title: "local_ip"
subcategory: ""
description: "local_ip for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1849, "body_sha256": "sha256:61da229ec226a4ed1c5fc81348246a8aaed0517ff021655c5891b51b387d694d", "canonical_id": "xcsh-docs:resources:tunnel:properties:local_ip", "child_ids": ["xcsh-docs:resources:tunnel:properties:local_ip:intf", "xcsh-docs:resources:tunnel:properties:local_ip:ip_address"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip", "parent_id": "xcsh-docs:resources:tunnel:reference", "path": "docs/guides/resources--tunnel--properties--local_ip.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- local_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - 1. Local Interface - Network Interface from which IP address and network will be
selected 2. IP Address - IP address and network can be configured explicitly.

Upstream description:

Defines the OPTIONS to select local IP address and virtual network for tunnel object OPTIONS
available are - &#8203;1. Local Interface - Network Interface from which IP address and network will
be selected &#8203;2. IP Address - IP address and network can be configured explicitly.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("intf",
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
  "x-ves-oneof-field-type": "[\"intf\",\"ip_address\"]"
}
```

Terraform syntax:

```terraform
local_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [intf](resources--tunnel--properties--local_ip--intf.md): complete subsection reference.

- [ip_address](resources--tunnel--properties--local_ip--ip_address.md): complete subsection reference.

## Next pages

- [local_ip.intf](resources--tunnel--properties--local_ip--intf.md)
- [local_ip.ip_address](resources--tunnel--properties--local_ip--ip_address.md)
- [Property reference](resources--tunnel--reference.md)
- [xcsh_tunnel](../resources/tunnel.md)
