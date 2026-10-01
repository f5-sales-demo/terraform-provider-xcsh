---
page_title: "eks_k8s.not_managed.node_list.interface_list"
subcategory: ""
description: "eks_k8s.not_managed.node_list.interface_list for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 12460, "body_sha256": "sha256:559ce338fa29d830287c5df2b54a4f74336fb15f0e111f926502d3633c29f34a", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:bond_interface", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:dhcp_client", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:dhcp_server", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ethernet_interface", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:ipv6_auto_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:monitor", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:monitor_disabled", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:network_option", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:no_ipv4_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:no_ipv6_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:site_to_site_connectivity_interface_disabled", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:site_to_site_connectivity_interface_enabled", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:static_ip", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:static_ipv6_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list:vlan_interface"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list:interface_list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:eks_k8s:not_managed:node_list", "path": "docs/guides/data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["eks_k8s", "not_managed", "node_list", "interface_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/eks_k8s/not_managed/node_list/interface_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "eks_k8s.not_managed.node_list.interface_list for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# eks_k8s.not_managed.node_list.interface_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [eks_k8s](data-sources--securemesh_site_v2--properties--eks_k8s.md)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed.md)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list.md)
- eks_k8s.not_managed.node_list.interface_list

<a id="section"></a>

Type: `"list"`. Computed.

Manage interfaces belonging to this node.

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

## Direct properties

- [bond_interface](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--bond_interface.md): complete subsection reference.

<a id="schema-eks_k8s--not_managed--node_list--interface_list--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--dhcp_client.md): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--dhcp_server.md): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ethernet_interface.md): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config.md): complete subsection reference.

<a id="schema-eks_k8s--not_managed--node_list--interface_list--is_management"></a>

### is_management property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="schema-eks_k8s--not_managed--node_list--interface_list--is_primary"></a>

### is_primary property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="schema-eks_k8s--not_managed--node_list--interface_list--labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

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

- [monitor](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--monitor.md): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--monitor_disabled.md): complete subsection reference.

<a id="schema-eks_k8s--not_managed--node_list--interface_list--mtu"></a>

### mtu property

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

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

<a id="schema-eks_k8s--not_managed--node_list--interface_list--name"></a>

### name property

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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

- [network_option](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--network_option.md): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--no_ipv4_address.md): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--no_ipv6_address.md): complete subsection reference.

<a id="schema-eks_k8s--not_managed--node_list--interface_list--priority"></a>

### priority property

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

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

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--site_to_site_connectivity_interface_disabled.md): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--site_to_site_connectivity_interface_enabled.md): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--static_ip.md): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--static_ipv6_address.md): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--vlan_interface.md): complete subsection reference.

## Next pages

- [eks_k8s.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--bond_interface.md)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--dhcp_client.md)
- [eks_k8s.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--dhcp_server.md)
- [eks_k8s.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ethernet_interface.md)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [eks_k8s.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--monitor.md)
- [eks_k8s.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--monitor_disabled.md)
- [eks_k8s.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--network_option.md)
- [eks_k8s.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--no_ipv4_address.md)
- [eks_k8s.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--no_ipv6_address.md)
- [eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--site_to_site_connectivity_interface_disabled.md)
- [eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--site_to_site_connectivity_interface_enabled.md)
- [eks_k8s.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--static_ip.md)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--static_ipv6_address.md)
- [eks_k8s.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list--interface_list--vlan_interface.md)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--properties--eks_k8s--not_managed--node_list.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
