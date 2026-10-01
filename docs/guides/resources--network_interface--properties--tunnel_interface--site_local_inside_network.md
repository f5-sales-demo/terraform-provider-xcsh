---
page_title: "tunnel_interface.site_local_inside_network"
subcategory: ""
description: "tunnel_interface.site_local_inside_network for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1075, "body_sha256": "sha256:d422cd3605b2071f7a0ceef9d2d999a3c68aa0e29bacf8ccb1d7ddbe77c70667", "canonical_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_inside_network", "child_ids": [], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:tunnel_interface:site_local_inside_network", "parent_id": "xcsh-docs:resources:network_interface:properties:tunnel_interface", "path": "docs/guides/resources--network_interface--properties--tunnel_interface--site_local_inside_network.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tunnel_interface", "site_local_inside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/tunnel_interface/site_local_inside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tunnel_interface.site_local_inside_network for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface.site_local_inside_network

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [tunnel_interface](resources--network_interface--properties--tunnel_interface.md)
- tunnel_interface.site_local_inside_network

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
site_local_inside_network = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tunnel_interface](resources--network_interface--properties--tunnel_interface.md)
- [xcsh_network_interface](../resources/network_interface.md)
