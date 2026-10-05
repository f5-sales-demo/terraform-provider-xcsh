---
page_title: "custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks"
subcategory: ""
description: "List of networks from which DHCP Server can allocate IPv4 Addresses."
xcsh_docs: {"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks"], "body_bytes": 10014, "body_sha256": "sha256:939254ee64f169f7f9c4a0a036a580ae099bc8565d3864c3de8b53c43eb6e330", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks:first_address", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks:last_address", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks:pools", "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks:same_as_dgw"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server", "path": "documentation/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks dgw address"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks--dgw_address", "description": "Exclusive with Enter a IPv4 address from the network prefix to be used as the default gateway.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks", "dgw_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks dns address"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks--dns_address", "description": "Exclusive with Enter a IPv4 address from the network prefix to be used as the DNS server.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks", "dns_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks first address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks:first_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks", "first_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks last address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks:last_address", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks", "last_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks network prefix"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks--network_prefix", "description": "Exclusive with Set the network prefix for the site. Ex: 192.0.2.0/24.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks", "network_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks pool settings"], "anchor": "schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks--pool_settings", "description": "Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks", "pool_settings"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks pools"], "anchor": "section", "description": "List of non overlapping IP address ranges.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks:pools", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks", "pools"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom storage config storage interface list storage interfaces storage interface dhcp server dhcp networks same as dgw"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_storage_config:storage_interface_list:storage_interfaces:storage_interface:dhcp_server:dhcp_networks:same_as_dgw", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_interface_list", "storage_interfaces", "storage_interface", "dhcp_server", "dhcp_networks", "same_as_dgw"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of networks from which DHCP Server can allocate IPv4 Addresses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/)
- [custom_storage_config.storage_interface_list.storage_interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks

<a id="section"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

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

## Direct properties

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks--dgw_address"></a>

### dgw_address property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks--dns_address"></a>

### dns_address property

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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

- [first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/first_address/): complete subsection reference.

- [last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/last_address/): complete subsection reference.

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks--network_prefix"></a>

### network_prefix property

Type: `"string"`. Computed.

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

<a id="schema-custom_storage_config--storage_interface_list--storage_interfaces--storage_interface--dhcp_server--dhcp_networks--pool_settings"></a>

### pool_settings property

Type: `"string"`. Computed.

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

- [pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/pools/): complete subsection reference.

- [same_as_dgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/same_as_dgw/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/first_address/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/last_address/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/pools/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.same_as_dgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/dhcp_networks/same_as_dgw/)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/custom_storage_config/storage_interface_list/storage_interfaces/storage_interface/dhcp_server/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
