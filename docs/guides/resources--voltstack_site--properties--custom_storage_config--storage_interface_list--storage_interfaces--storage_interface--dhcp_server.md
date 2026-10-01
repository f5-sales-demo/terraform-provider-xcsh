---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server"
subcategory: ""
description: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 5409, "body_sha256": "sha256:46630def36da65ac9e6ec4d47ccd52fa3e396fc1b272a54ee86192f11d0cff5b", "canonical_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:automatic_from_end", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:automatic_from_start", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:interface_ip_map"], "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "path": "docs/guides/resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md)
- [Property reference](resources--voltstack_site--reference.md)
- [custom_storage_config](resources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--properties--custom_storage_config--storage_interface_list.md)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dhcp server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

## Direct properties

- [automatic_from_end](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--automatic_from_end.md): complete subsection reference.

- [automatic_from_start](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--automatic_from_start.md): complete subsection reference.

- [dhcp_networks](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_option82_tag"></a>

### dhcp_option82_tag property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--fixed_ip_map"></a>

### fixed_ip_map property

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--interface_ip_map.md): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--automatic_from_end.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--automatic_from_start.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--interface_ip_map.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md)
- [xcsh_voltstack_site](../resources/voltstack_site.md)
