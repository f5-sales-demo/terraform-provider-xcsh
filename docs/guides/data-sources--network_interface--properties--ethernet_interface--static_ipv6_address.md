---
page_title: "ethernet_interface.static_ipv6_address"
subcategory: ""
description: "ethernet_interface.static_ipv6_address for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1597, "body_sha256": "sha256:53e188679cdbced740d45e8c9a419a91bcd51e0e43a1d926738f271ca453dfd9", "canonical_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address:cluster_static_ip", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address:node_static_ip"], "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "path": "docs/guides/data-sources--network_interface--properties--ethernet_interface--static_ipv6_address.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "static_ipv6_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.static_ipv6_address for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ethernet_interface.static_ipv6_address

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md)
- [Property reference](data-sources--network_interface--reference.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- ethernet_interface.static_ipv6_address

<a id="section"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

## Direct properties

- [cluster_static_ip](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--node_static_ip.md): complete subsection reference.

## Next pages

- [ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--cluster_static_ip.md)
- [ethernet_interface.static_ipv6_address.node_static_ip](data-sources--network_interface--properties--ethernet_interface--static_ipv6_address--node_static_ip.md)
- [ethernet_interface](data-sources--network_interface--properties--ethernet_interface.md)
- [xcsh_network_interface](../data-sources/network_interface.md)
