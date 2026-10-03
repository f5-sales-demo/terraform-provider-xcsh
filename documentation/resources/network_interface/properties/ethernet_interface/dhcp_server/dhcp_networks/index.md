---
page_title: "ethernet_interface.dhcp_server.dhcp_networks"
subcategory: ""
description: "List of networks from which DHCP Server can allocate IPv4 Addresses."
xcsh_docs: {"aliases": ["ethernet interface dhcp server dhcp networks"], "body_bytes": 9021, "body_sha256": "sha256:e2a0a8980a799058935f944cfa801d7914239cafdced0caa0c0c59bbb4c0af69", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:first_address", "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:last_address", "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:pools", "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:same_as_dgw"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server", "path": "documentation/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1101231230331123-0120002001113102-0100312322303002-2001332201223231-2320301212302003-1331032013103200-0312202201100100-3322013131331322", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [{"anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--dgw_address", "enforcement": "provider-schema", "group": "ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,first_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--dgw_address", "enforcement": "provider-schema", "group": "ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--dns_address", "enforcement": "provider-schema", "group": "ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dns_address,same_as_dgw", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,first_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:first_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dgw_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:first_address,last_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:last_address", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ethernet_interface.dhcp_server.dhcp_networks:ConflictingListObjectAttributes:dns_address,same_as_dgw", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:same_as_dgw", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks"], "schema_version": 1, "sections": [{"aliases": ["ethernet interface dhcp server dhcp networks dgw address"], "anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--dgw_address", "description": "Exclusive with Enter a IPv4 address from the network prefix to be used as the default gateway.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "dgw_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["ethernet interface dhcp server dhcp networks dns address"], "anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--dns_address", "description": "Exclusive with Enter a IPv4 address from the network prefix to be used as the DNS server.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "dns_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["ethernet interface dhcp server dhcp networks first address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:first_address", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "first_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["ethernet interface dhcp server dhcp networks last address"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:last_address", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "last_address"], "syntax": "attribute", "type": "object"}, {"aliases": ["ethernet interface dhcp server dhcp networks network prefix"], "anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--network_prefix", "description": "Exclusive with Set the network prefix for the site. Ex: 192.0.2.0/24.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "network_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["ethernet interface dhcp server dhcp networks pool settings"], "anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--pool_settings", "description": "Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "pool_settings"], "syntax": "attribute", "type": "string"}, {"aliases": ["ethernet interface dhcp server dhcp networks pools"], "anchor": "section", "description": "List of non overlapping IP address ranges.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:pools", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "pools"], "syntax": "block", "type": "object"}, {"aliases": ["ethernet interface dhcp server dhcp networks same as dgw"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:same_as_dgw", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "same_as_dgw"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of networks from which DHCP Server can allocate IPv4 Addresses.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.dhcp_server.dhcp_networks

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- [ethernet_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/)
- ethernet_interface.dhcp_server.dhcp_networks

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-ethernet_interface--dhcp_server--dhcp_networks--dgw_address"></a>

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-ethernet_interface--dhcp_server--dhcp_networks--dns_address"></a>

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/first_address/): complete subsection reference.

- [last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/last_address/): complete subsection reference.

<a id="schema-ethernet_interface--dhcp_server--dhcp_networks--network_prefix"></a>

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="schema-ethernet_interface--dhcp_server--dhcp_networks--pool_settings"></a>

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

- [pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/pools/): complete subsection reference.

- [same_as_dgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/same_as_dgw/): complete subsection reference.

## Next pages

- [ethernet_interface.dhcp_server.dhcp_networks.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/first_address/)
- [ethernet_interface.dhcp_server.dhcp_networks.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/last_address/)
- [ethernet_interface.dhcp_server.dhcp_networks.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/pools/)
- [ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/same_as_dgw/)
- [ethernet_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
