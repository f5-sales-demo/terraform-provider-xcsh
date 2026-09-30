---
page_title: "equinix.not_managed.node_list.interface_list"
subcategory: ""
description: "equinix.not_managed.node_list.interface_list for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 14382, "body_sha256": "sha256:fe2fa80955ec57f30492d8dc2a29550459721e93e6e21d33f082d2956e355137", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:bond_interface", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:dhcp_client", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:dhcp_server", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ethernet_interface", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:ipv6_auto_config", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:monitor", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:monitor_disabled", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:network_option", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:no_ipv4_address", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:no_ipv6_address", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:site_to_site_connectivity_interface_disabled", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:site_to_site_connectivity_interface_enabled", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ip", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:vlan_interface"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list", "path": "docs/guides/resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["equinix", "not_managed", "node_list", "interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "equinix.not_managed.node_list.interface_list for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# equinix.not_managed.node_list.interface_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [equinix](resources--securemesh_site_v2--properties--equinix.md)
- [equinix.not_managed](resources--securemesh_site_v2--properties--equinix--not_managed.md)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--properties--equinix--not_managed--node_list.md)
- equinix.not_managed.node_list.interface_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bond_interface](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--bond_interface.md): complete subsection reference.

<a id="schema-equinix--not_managed--node_list--interface_list--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--dhcp_client.md): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--dhcp_server.md): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ethernet_interface.md): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config.md): complete subsection reference.

<a id="schema-equinix--not_managed--node_list--interface_list--is_management"></a>

### is_management property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="schema-equinix--not_managed--node_list--interface_list--is_primary"></a>

### is_primary property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="schema-equinix--not_managed--node_list--interface_list--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--monitor.md): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--monitor_disabled.md): complete subsection reference.

<a id="schema-equinix--not_managed--node_list--interface_list--mtu"></a>

### mtu property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
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
    "maximum": 8000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="schema-equinix--not_managed--node_list--interface_list--name"></a>

### name property

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--network_option.md): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--no_ipv4_address.md): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--no_ipv6_address.md): complete subsection reference.

<a id="schema-equinix--not_managed--node_list--interface_list--priority"></a>

### priority property

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--site_to_site_connectivity_interface_disabled.md): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--site_to_site_connectivity_interface_enabled.md): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--static_ip.md): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--static_ipv6_address.md): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--vlan_interface.md): complete subsection reference.

## Next pages

- [equinix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--bond_interface.md)
- [equinix.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--dhcp_client.md)
- [equinix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--dhcp_server.md)
- [equinix.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ethernet_interface.md)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [equinix.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--monitor.md)
- [equinix.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--monitor_disabled.md)
- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--network_option.md)
- [equinix.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--no_ipv4_address.md)
- [equinix.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--no_ipv6_address.md)
- [equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--site_to_site_connectivity_interface_disabled.md)
- [equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--site_to_site_connectivity_interface_enabled.md)
- [equinix.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--static_ip.md)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--static_ipv6_address.md)
- [equinix.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--properties--equinix--not_managed--node_list--interface_list--vlan_interface.md)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--properties--equinix--not_managed--node_list.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
