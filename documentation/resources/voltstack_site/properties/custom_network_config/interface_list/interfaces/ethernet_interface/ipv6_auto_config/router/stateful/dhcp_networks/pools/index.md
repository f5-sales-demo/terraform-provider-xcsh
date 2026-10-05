---
page_title: "custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools"
subcategory: ""
description: "List of non overlapping IP address ranges."
xcsh_docs: {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful dhcp networks pools"], "body_bytes": 6557, "body_sha256": "sha256:029b60334c2ca9c9e7db5e3c42db001b20e69a13652a98ed15a2e2658e1b518a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:dhcp_networks:pools", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:dhcp_networks", "path": "documentation/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/pools/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3213301131232323-3303331132010020-0220111030122023-0012322333122110-3100322320300112-0323210121321033-0233130313323301-1103130103312122", "registry_path": "docs/guides/resources--voltstack_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools"], "schema_version": 1, "sections": [{"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful dhcp networks pools end ip"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip", "description": "Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on network prefix.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:dhcp_networks:pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools", "end_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom network config interface list interfaces ethernet interface ipv6 auto config router stateful dhcp networks pools start ip"], "anchor": "schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip", "description": "Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on network prefix. 2001::1 with prefix length of 64, start offset is 5.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:interface_list:interfaces:ethernet_interface:ipv6_auto_config:router:stateful:dhcp_networks:pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "interface_list", "interfaces", "ethernet_interface", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools", "start_ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/pools/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of non overlapping IP address ranges.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/)
- [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/)
- [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip"></a>

### end_ip property

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip"></a>

### start_ip property

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

## Next pages

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
