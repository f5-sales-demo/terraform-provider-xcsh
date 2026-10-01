---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface"
subcategory: ""
description: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 14157, "body_sha256": "sha256:75d9bb83fcaee19d48b722f46e249dafd0a092ed3bc630a2a8e8bc7b09bc02d4", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:cluster", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_client", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:ipv6_auto_config", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:is_primary", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:monitor", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:monitor_disabled", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:no_ipv6_address", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:not_primary", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_inside_network", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:site_local_network", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ip", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:static_ipv6_address", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:storage_network", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:untagged"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces", "path": "docs/guides/data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [custom_storage_config](data-sources--voltstack_site--properties--custom_storage_config.md)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list.md)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for storage interface.

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

- [cluster](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--cluster.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--device"></a>

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

- [dhcp_client](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_client.md): complete subsection reference.

- [dhcp_server](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server.md): complete subsection reference.

- [ipv6_auto_config](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config.md): complete subsection reference.

- [is_primary](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--is_primary.md): complete subsection reference.

- [monitor](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--monitor.md): complete subsection reference.

- [monitor_disabled](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--monitor_disabled.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--mtu"></a>

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

- [no_ipv6_address](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--no_ipv6_address.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--node"></a>

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

- [not_primary](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--not_primary.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--priority"></a>

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

- [site_local_inside_network](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--site_local_inside_network.md): complete subsection reference.

- [site_local_network](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--site_local_network.md): complete subsection reference.

- [static_ip](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip.md): complete subsection reference.

- [static_ipv6_address](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ipv6_address.md): complete subsection reference.

- [storage_network](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--storage_network.md): complete subsection reference.

- [untagged](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--untagged.md): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--vlan_id"></a>

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

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--cluster.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_client.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--ipv6_auto_config.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--is_primary.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--monitor.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--monitor_disabled.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--no_ipv6_address.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--not_primary.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--site_local_inside_network.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--site_local_network.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ip.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--static_ipv6_address.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.storage_network](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--storage_network.md)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.untagged](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--untagged.md)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--properties--custom_storage_config--storage_interface_list--storage_interfaces.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
