---
page_title: "ethernet_interface.dhcp_server.dhcp_networks.first_address"
subcategory: ""
description: "ethernet_interface.dhcp_server.dhcp_networks.first_address for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1412, "body_sha256": "sha256:a8c041f520e9e95770417cc62ed5c6510bbd2ae1a0e4d696357a3e52be7c5f37", "canonical_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:first_address", "child_ids": [], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:first_address", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "path": "docs/guides/resources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks--first_address.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "first_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/first_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.dhcp_server.dhcp_networks.first_address for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.dhcp_server.dhcp_networks.first_address

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- [ethernet_interface.dhcp_server](resources--network_interface--properties--ethernet_interface--dhcp_server.md)
- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md)
- ethernet_interface.dhcp_server.dhcp_networks.first_address

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
first_address = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ethernet_interface.dhcp_server.dhcp_networks](resources--network_interface--properties--ethernet_interface--dhcp_server--dhcp_networks.md)
- [xcsh_network_interface](../resources/network_interface.md)
