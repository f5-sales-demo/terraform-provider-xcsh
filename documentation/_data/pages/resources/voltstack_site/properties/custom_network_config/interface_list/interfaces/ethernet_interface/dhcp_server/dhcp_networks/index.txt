---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks"
subcategory: ""
description: "List of networks from which DHCP Server can allocate IPv4 Addresses."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks"], "body_bytes": 10645, "body_sha256": "sha256:0ca07b61fad273ca5f0c9e088f733ebca16273277e1d4d97c36614c3c52ddf13", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:first_address", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:last_address", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:pools", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:same_as_dgw"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server", "path": "documentation/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2031300102020223-1122120102122320-1211133102230221-2312030221330301-3223233011133131-0313133222220331-3221222231211022-2311331122123132", "registry_path": "docs/guides/resources--voltstack_site--reference--group-004.md", "relationships": [{"anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,first_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dns_address", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dns_address,same_as_dgw", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,first_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dns_address,same_as_dgw", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:same_as_dgw", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks"], "schema_version": 1, "sections": [{"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks dgw address"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address", "description": "Exclusive with Enter a IPv4 address from the network prefix to be used as the default gateway.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks", "dgw_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks dns address"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dns_address", "description": "Exclusive with Enter a IPv4 address from the network prefix to be used as the DNS server.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks", "dns_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks first address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:first_address", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks", "first_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks last address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:last_address", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks", "last_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks network prefix"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--network_prefix", "description": "Exclusive with Set the network prefix for the site. Ex: 192.0.2.0/24.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks", "network_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks pool settings"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pool_settings", "description": "Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks", "pool_settings"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks pools"], "anchor": "section", "description": "List of non overlapping IP address ranges.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:pools", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks", "pools"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config interface list interfaces ethernet interface dhcp server dhcp networks same as dgw"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:dhcp_server:dhcp_networks:same_as_dgw", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "dhcp_server", "dhcp_networks", "same_as_dgw"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of networks from which DHCP Server can allocate IPv4 Addresses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address"></a>

### dgw_address property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dns_address"></a>

### dns_address property

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/first_address/): complete subsection reference.

- [last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/last_address/): complete subsection reference.

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--network_prefix"></a>

### network_prefix property

Type: `"string"`. Optional.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pool_settings"></a>

### pool_settings property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/pools/): complete subsection reference.

- [same_as_dgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/same_as_dgw/): complete subsection reference.

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/first_address/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/last_address/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/pools/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/same_as_dgw/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
