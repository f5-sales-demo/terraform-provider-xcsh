---
page_title: "ethernet_interface.static_ipv6_address"
subcategory: ""
description: "ethernet_interface.static_ipv6_address for xcsh_network_interface."
xcsh_docs: {"aliases": [], "body_bytes": 1978, "body_sha256": "sha256:b47ffa1b538044b7f40e08f45f4366c6a1c32cc13e8a8e2b00496f4d6f259c14", "canonical_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ipv6_address", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ipv6_address:cluster_static_ip", "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ipv6_address:node_static_ip"], "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:static_ipv6_address", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface", "path": "docs/guides/resources--network_interface--properties--ethernet_interface--static_ipv6_address.md", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ethernet_interface", "static_ipv6_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ethernet_interface.static_ipv6_address for xcsh_network_interface.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.static_ipv6_address

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md)
- [Property reference](resources--network_interface--reference.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- ethernet_interface.static_ipv6_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_static_ip](resources--network_interface--properties--ethernet_interface--static_ipv6_address--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](resources--network_interface--properties--ethernet_interface--static_ipv6_address--node_static_ip.md): complete subsection reference.

## Next pages

- [ethernet_interface.static_ipv6_address.cluster_static_ip](resources--network_interface--properties--ethernet_interface--static_ipv6_address--cluster_static_ip.md)
- [ethernet_interface.static_ipv6_address.node_static_ip](resources--network_interface--properties--ethernet_interface--static_ipv6_address--node_static_ip.md)
- [ethernet_interface](resources--network_interface--properties--ethernet_interface.md)
- [xcsh_network_interface](../resources/network_interface.md)
