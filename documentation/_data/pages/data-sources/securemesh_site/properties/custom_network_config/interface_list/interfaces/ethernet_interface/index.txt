---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.ethernet_interface for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 15002, "body_sha256": "sha256:4edd459a2fc3fd211e57527fdb095b49b052d0441ad3125d7ff90f3dea010dc9", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:cluster", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_client", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:is_primary", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:monitor", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:monitor_disabled", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:no_ipv6_address", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:not_primary", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:site_local_inside_network", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:site_local_network", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:storage_network", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:untagged"], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "parent_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list:interfaces", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/index.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.ethernet_interface for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.interface_list.interfaces.ethernet_interface

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- custom_network_config.interface_list.interfaces.ethernet_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

Upstream description:

Ethernet Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

## Direct properties

- [cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/cluster/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--device"></a>

### device property

Type: `"string"`. Computed.

Interface configuration for the ethernet device.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [dhcp_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_client/): complete subsection reference.

- [dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/): complete subsection reference.

- [ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/): complete subsection reference.

- [is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/is_primary/): complete subsection reference.

- [monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/monitor/): complete subsection reference.

- [monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/monitor_disabled/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--mtu"></a>

### mtu property

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/no_ipv6_address/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--node"></a>

### node property

Type: `"string"`. Computed.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/not_primary/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--priority"></a>

### priority property

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/site_local_inside_network/): complete subsection reference.

- [site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/site_local_network/): complete subsection reference.

- [static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/): complete subsection reference.

- [static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/): complete subsection reference.

- [storage_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/storage_network/): complete subsection reference.

- [untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/untagged/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Computed.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/cluster/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_client/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/is_primary/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/monitor/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/monitor_disabled/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/no_ipv6_address/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/not_primary/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/site_local_inside_network/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/site_local_network/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/storage_network/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/untagged/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/interfaces/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
