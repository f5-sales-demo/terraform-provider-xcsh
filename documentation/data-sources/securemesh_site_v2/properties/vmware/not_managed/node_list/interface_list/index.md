---
page_title: "vmware.not_managed.node_list.interface_list"
subcategory: ""
description: "Manage interfaces belonging to this node."
xcsh_docs: {"aliases": ["vmware not managed node list interface list"], "body_bytes": 14162, "body_sha256": "sha256:ce72d8032eb7240e35b320ef524666983acf21b049b76a072f2116b62faecee5", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_client", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ethernet_interface", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:monitor", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:monitor_disabled", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:network_option", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv4_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv6_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:site_to_site_connectivity_interface_disabled", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:site_to_site_connectivity_interface_enabled", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ip", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ipv6_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:vlan_interface"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list", "path": "documentation/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["vmware", "not_managed", "node_list", "interface_list"], "schema_version": 1, "sections": [{"aliases": ["bond interface"], "anchor": "section", "description": "Bond devices configuration for fleet.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:bond_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "bond_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["description spec"], "anchor": "schema-vmware--not_managed--node_list--interface_list--description_spec", "description": "Interface Description. Description for this Interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["dhcp client"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_client", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "dhcp_client"], "syntax": "attribute", "type": "object"}, {"aliases": ["dhcp server"], "anchor": "section", "description": "DHCP server configuration for this interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:dhcp_server", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "dhcp_server"], "syntax": "attribute", "type": "object"}, {"aliases": ["ethernet interface"], "anchor": "section", "description": "Configuration parameter for ethernet interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ethernet_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "ethernet_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["ipv6 auto config"], "anchor": "section", "description": "IPV6AutoConfigType.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:ipv6_auto_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "ipv6_auto_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["is management"], "anchor": "schema-vmware--not_managed--node_list--interface_list--is_management", "description": "Configuration for is_management.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "is_management"], "syntax": "attribute", "type": "bool"}, {"aliases": ["is primary"], "anchor": "schema-vmware--not_managed--node_list--interface_list--is_primary", "description": "Configuration for is_primary.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "is_primary"], "syntax": "attribute", "type": "bool"}, {"aliases": ["labels"], "anchor": "schema-vmware--not_managed--node_list--interface_list--labels", "description": "Add Labels for this Interface, these labels can be used in firewall policy.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["monitor"], "anchor": "section", "description": "Link Quality Monitoring configuration for a network interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:monitor", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "monitor"], "syntax": "attribute", "type": "object"}, {"aliases": ["monitor disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:monitor_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "monitor_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["mtu"], "anchor": "schema-vmware--not_managed--node_list--interface_list--mtu", "description": "Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between 512 and 8000.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "mtu"], "syntax": "attribute", "type": "number"}, {"aliases": ["name"], "anchor": "schema-vmware--not_managed--node_list--interface_list--name", "description": "Name of this Interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["network option"], "anchor": "section", "description": "Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs, Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is optional. Global VRFs are", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:network_option", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "network_option"], "syntax": "attribute", "type": "object"}, {"aliases": ["no ipv4 address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv4_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "no_ipv4_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["no ipv6 address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:no_ipv6_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "no_ipv6_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["priority"], "anchor": "schema-vmware--not_managed--node_list--interface_list--priority", "description": "For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be used as active and interfaces with lower priority will be used as backup. If multiple interfaces have the same priority, ECMP will be used. Greater the value, higher the priority.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "priority"], "syntax": "attribute", "type": "number"}, {"aliases": ["site to site connectivity interface disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:site_to_site_connectivity_interface_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "site_to_site_connectivity_interface_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["site to site connectivity interface enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:site_to_site_connectivity_interface_enabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "site_to_site_connectivity_interface_enabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "static_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["static ipv6 address"], "anchor": "section", "description": "Configure Static IP parameters.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:static_ipv6_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "static_ipv6_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["vlan interface"], "anchor": "section", "description": "Configuration parameter for vlan interface.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:vmware:not_managed:node_list:interface_list:vlan_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vmware", "not_managed", "node_list", "interface_list", "vlan_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manage interfaces belonging to this node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vmware.not_managed.node_list.interface_list

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [vmware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/)
- [vmware.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/)
- [vmware.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/)
- vmware.not_managed.node_list.interface_list

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

- [bond_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/bond_interface/): complete subsection reference.

<a id="schema-vmware--not_managed--node_list--interface_list--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/dhcp_client/): complete subsection reference.

- [dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/dhcp_server/): complete subsection reference.

- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/ethernet_interface/): complete subsection reference.

- [ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/ipv6_auto_config/): complete subsection reference.

<a id="schema-vmware--not_managed--node_list--interface_list--is_management"></a>

### is_management property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="schema-vmware--not_managed--node_list--interface_list--is_primary"></a>

### is_primary property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="schema-vmware--not_managed--node_list--interface_list--labels"></a>

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

- [monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/monitor/): complete subsection reference.

- [monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/monitor_disabled/): complete subsection reference.

<a id="schema-vmware--not_managed--node_list--interface_list--mtu"></a>

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

<a id="schema-vmware--not_managed--node_list--interface_list--name"></a>

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

- [network_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/network_option/): complete subsection reference.

- [no_ipv4_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/no_ipv4_address/): complete subsection reference.

- [no_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/no_ipv6_address/): complete subsection reference.

<a id="schema-vmware--not_managed--node_list--interface_list--priority"></a>

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

- [site_to_site_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/site_to_site_connectivity_interface_disabled/): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/site_to_site_connectivity_interface_enabled/): complete subsection reference.

- [static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/static_ip/): complete subsection reference.

- [static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/static_ipv6_address/): complete subsection reference.

- [vlan_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/vlan_interface/): complete subsection reference.

## Next pages

- [vmware.not_managed.node_list.interface_list.bond_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/bond_interface/)
- [vmware.not_managed.node_list.interface_list.dhcp_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/dhcp_client/)
- [vmware.not_managed.node_list.interface_list.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/dhcp_server/)
- [vmware.not_managed.node_list.interface_list.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/ethernet_interface/)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/ipv6_auto_config/)
- [vmware.not_managed.node_list.interface_list.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/monitor/)
- [vmware.not_managed.node_list.interface_list.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/monitor_disabled/)
- [vmware.not_managed.node_list.interface_list.network_option](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/network_option/)
- [vmware.not_managed.node_list.interface_list.no_ipv4_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/no_ipv4_address/)
- [vmware.not_managed.node_list.interface_list.no_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/no_ipv6_address/)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/site_to_site_connectivity_interface_disabled/)
- [vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/site_to_site_connectivity_interface_enabled/)
- [vmware.not_managed.node_list.interface_list.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/static_ip/)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/static_ipv6_address/)
- [vmware.not_managed.node_list.interface_list.vlan_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/interface_list/vlan_interface/)
- [vmware.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/vmware/not_managed/node_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
