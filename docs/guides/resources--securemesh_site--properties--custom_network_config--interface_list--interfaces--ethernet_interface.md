---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface"
subcategory: ""
description: "custom_network_config.interface_list.interfaces.ethernet_interface for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 15249, "body_sha256": "sha256:ec353ea2dafb60d36bfe6c22bd794798b32c04c9256fe91b6c5d32ba7f428e2b", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:cluster", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_client", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:is_primary", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:monitor", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:monitor_disabled", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:no_ipv6_address", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:not_primary", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:site_local_inside_network", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:site_local_network", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ip", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:static_ipv6_address", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:storage_network", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:untagged"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list:interfaces", "path": "docs/guides/resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.interface_list.interfaces.ethernet_interface for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config.interface_list.interfaces.ethernet_interface

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [custom_network_config.interface_list](resources--securemesh_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md)
- custom_network_config.interface_list.interfaces.ethernet_interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Upstream description:

Ethernet Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("site_local_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("untagged",
    "vlan_id")}
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
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
ethernet_interface {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--cluster.md): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--device"></a>

### device property

Type: `"string"`. Optional.

Interface configuration for the ethernet device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

- [dhcp_client](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_client.md): complete subsection reference.

- [dhcp_server](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config.md): complete subsection reference.

- [is_primary](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--is_primary.md): complete subsection reference.

- [monitor](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--monitor.md): complete subsection reference.

- [monitor_disabled](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--monitor_disabled.md): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--mtu"></a>

### mtu property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
  ),
}
```

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

- [no_ipv6_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--no_ipv6_address.md): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--node"></a>

### node property

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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

- [not_primary](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--not_primary.md): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--priority"></a>

### priority property

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

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

- [site_local_inside_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--site_local_inside_network.md): complete subsection reference.

- [site_local_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--site_local_network.md): complete subsection reference.

- [static_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip.md): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address.md): complete subsection reference.

- [storage_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--storage_network.md): complete subsection reference.

- [untagged](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--untagged.md): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--vlan_id"></a>

### vlan_id property

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

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

- [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--cluster.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_client.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--is_primary.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--monitor.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--monitor_disabled.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--no_ipv6_address.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--not_primary.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--site_local_inside_network.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--site_local_network.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ip.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--storage_network.md)
- [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces--ethernet_interface--untagged.md)
- [custom_network_config.interface_list.interfaces](resources--securemesh_site--properties--custom_network_config--interface_list--interfaces.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
