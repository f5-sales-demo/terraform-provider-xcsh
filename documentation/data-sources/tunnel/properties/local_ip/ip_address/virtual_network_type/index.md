---
page_title: "local_ip.ip_address.virtual_network_type"
subcategory: ""
description: "local_ip.ip_address.virtual_network_type for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 2465, "body_sha256": "sha256:e036584528ce5adbfda79038419a3c85a48af73f55343222da19cc5ddf632a0d", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local", "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type:site_local_inside"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address:virtual_network_type", "parent_id": "xcsh-docs:data-sources:tunnel:properties:local_ip:ip_address", "path": "documentation/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/index.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["local_ip", "ip_address", "virtual_network_type"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.virtual_network_type for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.virtual_network_type

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [local_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/)
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

- [public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/public/): complete subsection reference.

- [site_local](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/): complete subsection reference.

- [site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local_inside/): complete subsection reference.

## Next pages

- [local_ip.ip_address.virtual_network_type.public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/public/)
- [local_ip.ip_address.virtual_network_type.site_local](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local/)
- [local_ip.ip_address.virtual_network_type.site_local_inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/virtual_network_type/site_local_inside/)
- [local_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/local_ip/ip_address/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
