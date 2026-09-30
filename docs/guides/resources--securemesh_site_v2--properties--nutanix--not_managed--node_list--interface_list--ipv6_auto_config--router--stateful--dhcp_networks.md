---
page_title: "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks"
subcategory: ""
description: "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 5723, "body_sha256": "sha256:b0376eae1e1c36ffe0960e6474809edc56fe2d9bf6820e6ab11a630880515a62", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful", "path": "docs/guides/resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [nutanix](resources--securemesh_site_v2--properties--nutanix.md)
- [nutanix.not_managed](resources--securemesh_site_v2--properties--nutanix--not_managed.md)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list.md)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

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

<a id="schema-nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--network_prefix"></a>

### network_prefix property

Type: `"string"`. Optional.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="schema-nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pool_settings"></a>

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

- [pools](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools.md): complete subsection reference.

## Next pages

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
