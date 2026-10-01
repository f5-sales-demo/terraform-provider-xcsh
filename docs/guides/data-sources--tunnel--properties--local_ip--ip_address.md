---
page_title: "local_ip.ip_address"
subcategory: ""
description: "local_ip.ip_address for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1534, "body_sha256": "sha256:51759f2fe5d3adf0bf3b755671ed6fa753c3bf9b51433a7906bf5a23da93802e", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:auto", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:ip_address", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip", "path": "docs/guides/data-sources--tunnel--properties--local_ip--ip_address.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/ip_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [local_ip](data-sources--tunnel--properties--local_ip.md)
- local_ip.ip_address

<a id="section"></a>

Type: `"single"`. Computed.

Provides the configuration to pick up source IP and network for transporting encapsulated packet.

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

## Direct properties

- [auto](data-sources--tunnel--properties--local_ip--ip_address--auto.md): complete subsection reference.

- [ip_address](data-sources--tunnel--properties--local_ip--ip_address--ip_address.md): complete subsection reference.

- [virtual_network_type](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type.md): complete subsection reference.

## Next pages

- [local_ip.ip_address.auto](data-sources--tunnel--properties--local_ip--ip_address--auto.md)
- [local_ip.ip_address.ip_address](data-sources--tunnel--properties--local_ip--ip_address--ip_address.md)
- [local_ip.ip_address.virtual_network_type](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type.md)
- [local_ip](data-sources--tunnel--properties--local_ip.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
