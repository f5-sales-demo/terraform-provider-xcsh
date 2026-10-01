---
page_title: "local_ip.ip_address.virtual_network_type.public"
subcategory: ""
description: "local_ip.ip_address.virtual_network_type.public for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1226, "body_sha256": "sha256:977cc8a823457713f1d340631aac974750fbc0f8771e467b038f9a77974eba50", "canonical_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "child_ids": [], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type:public", "parent_id": "xcsh-docs:resources:tunnel:properties:local_ip:ip_address:virtual_network_type", "path": "docs/guides/resources--tunnel--properties--local_ip--ip_address--virtual_network_type--public.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["local_ip", "ip_address", "virtual_network_type", "public"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/local_ip/ip_address/virtual_network_type/public/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "local_ip.ip_address.virtual_network_type.public for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_ip.ip_address.virtual_network_type.public

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [local_ip](resources--tunnel--properties--local_ip.md)
- [local_ip.ip_address](resources--tunnel--properties--local_ip--ip_address.md)
- [local_ip.ip_address.virtual_network_type](resources--tunnel--properties--local_ip--ip_address--virtual_network_type.md)
- local_ip.ip_address.virtual_network_type.public

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
public = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [local_ip.ip_address.virtual_network_type](resources--tunnel--properties--local_ip--ip_address--virtual_network_type.md)
- [xcsh_tunnel](../resources/tunnel.md)
