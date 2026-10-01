---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 2897, "body_sha256": "sha256:434aa6f07e8cf350a7743a6dd3f748d2af29753b18a8bd17aa19ee53eecb26d9", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address:cluster_static_ip", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address:node_static_ip"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "path": "docs/guides/resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "static_ipv6_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_network_config](resources--voltstack_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](resources--voltstack_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.interface_list.interfaces](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address

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

- [cluster_static_ip](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip.md): complete subsection reference.

- [node_static_ip](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface](resources--voltstack_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
