---
page_title: "local_ip.ip_address.virtual_network_type"
subcategory: ""
description: "local_ip.ip_address.virtual_network_type for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1772, "body_sha256": "sha256:ff5e129abcf2a16dfe2b8abd18bb9af5553e504e69468f408a245febd494d136", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "path": "docs/guides/data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address", "virtual_network_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.virtual_network_type for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# local_ip.ip_address.virtual_network_type

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [local_ip](data-sources--tunnel--properties--local_ip.md)
- [local_ip.ip_address](data-sources--tunnel--properties--local_ip--ip_address.md)
- local_ip.ip_address.virtual_network_type

<a id="section"></a>

Type: `"single"`. Computed.

Different types of virtual networks understood by the system.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vn_type_choice": "[\"public\",\"site_local\",\"site_local_inside\"]"
}
```

## Direct properties

- [public](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--public.md): complete subsection reference.

- [site_local](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--site_local.md): complete subsection reference.

- [site_local_inside](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--site_local_inside.md): complete subsection reference.

## Next pages

- [local_ip.ip_address.virtual_network_type.public](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--public.md)
- [local_ip.ip_address.virtual_network_type.site_local](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--site_local.md)
- [local_ip.ip_address.virtual_network_type.site_local_inside](data-sources--tunnel--properties--local_ip--ip_address--virtual_network_type--site_local_inside.md)
- [local_ip.ip_address](data-sources--tunnel--properties--local_ip--ip_address.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
