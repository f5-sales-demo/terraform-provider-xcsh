---
page_title: "custom_network_config"
subcategory: ""
description: "custom_network_config for xcsh_voltstack_site."
xcsh_docs: {"aliases": [], "body_bytes": 13325, "body_sha256": "sha256:73d56e9f952c947f560b316e23e4d835afbf32e300f1492f2173290075ec880f", "canonical_id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config", "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_enhanced_firewall_policies", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_forward_proxy_policies", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:active_network_policies", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:default_config", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:default_interface_config", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:default_sli_config", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:forward_proxy_allow_all", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:global_network_list", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:interface_list", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:no_forward_proxy", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:no_global_network", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:no_network_policy", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:sli_config", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:slo_config", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:sm_connection_public_ip", "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:custom_network_config", "parent_id": "xcsh-docs:data-sources:voltstack_site:reference", "path": "docs/guides/data-sources--voltstack_site--properties--custom_network_config.md", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/custom_network_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config for xcsh_voltstack_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# custom_network_config

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- custom_network_config

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_network\_config, default\_network\_config; Default: default\_network\_config\]
VssNetworkConfiguration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-interface_choice": "[\"default_interface_config\",\"interface_list\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"site_to_site_tunnel_ip\",\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

OneOf alternatives in this subsection:

- [custom_network_config](data-sources--voltstack_site--properties--custom_network_config.md#section)
- [default_network_config](data-sources--voltstack_site--properties--default_network_config.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [active_enhanced_firewall_policies](data-sources--voltstack_site--properties--custom_network_config--active_enhanced_firewall_policies.md): complete subsection reference.

- [active_forward_proxy_policies](data-sources--voltstack_site--properties--custom_network_config--active_forward_proxy_policies.md): complete subsection reference.

- [active_network_policies](data-sources--voltstack_site--properties--custom_network_config--active_network_policies.md): complete subsection reference.

<a id="schema-custom_network_config--bgp_peer_address"></a>

### bgp_peer_address property

Type: `"string"`. Computed.

Optional BGP peer address that can be used as parameter for BGP configuration when BGP is configured
to fetch BGP peer address from site Object. This can be used to change peer address per site in
fleet.

Upstream description:

Optional BGP peer address that can be used as parameter for BGP configuration when BGP is configured
to fetch BGP peer address from site Object. This can be used to change peer address per site in
fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-custom_network_config--bgp_router_id"></a>

### bgp_router_id property

Type: `"string"`. Computed.

Optional BGP router ID that can be used as parameter for BGP configuration when BGP is configured to
fetch BGP router ID from site object.

Upstream description:

Optional BGP router ID that can be used as parameter for BGP configuration when BGP is configured to
fetch BGP router ID from site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [default_config](data-sources--voltstack_site--properties--custom_network_config--default_config.md): complete subsection reference.

- [default_interface_config](data-sources--voltstack_site--properties--custom_network_config--default_interface_config.md): complete subsection reference.

- [default_sli_config](data-sources--voltstack_site--properties--custom_network_config--default_sli_config.md): complete subsection reference.

- [forward_proxy_allow_all](data-sources--voltstack_site--properties--custom_network_config--forward_proxy_allow_all.md): complete subsection reference.

- [global_network_list](data-sources--voltstack_site--properties--custom_network_config--global_network_list.md): complete subsection reference.

- [interface_list](data-sources--voltstack_site--properties--custom_network_config--interface_list.md): complete subsection reference.

- [no_forward_proxy](data-sources--voltstack_site--properties--custom_network_config--no_forward_proxy.md): complete subsection reference.

- [no_global_network](data-sources--voltstack_site--properties--custom_network_config--no_global_network.md): complete subsection reference.

- [no_network_policy](data-sources--voltstack_site--properties--custom_network_config--no_network_policy.md): complete subsection reference.

<a id="schema-custom_network_config--outside_nameserver"></a>

### outside_nameserver property

Type: `"string"`. Computed.

Optional DNS server V4 IP to be used for name resolution in local network.

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

<a id="schema-custom_network_config--outside_vip"></a>

### outside_vip property

Type: `"string"`. Computed.

Optional common virtual V4 IP across all nodes to be used as automatic VIP for site local network.

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

<a id="schema-custom_network_config--site_to_site_tunnel_ip"></a>

### site_to_site_tunnel_ip property

Type: `"string"`. Computed.

Exclusive with \[sm\_connection\_public\_ip sm\_connection\_pvt\_ip\] Site Mesh Group Connection Via
Virtual IP. This option will use the Virtual IP provided for creating IPsec between two sites which
are part of the site mesh group.

Upstream description:

Exclusive with \[sm\_connection\_public\_ip sm\_connection\_pvt\_ip\] Site Mesh Group Connection Via
Virtual IP. This option will use the Virtual IP provided for creating IPsec between two sites which
are part of the site mesh group.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [sli_config](data-sources--voltstack_site--properties--custom_network_config--sli_config.md): complete subsection reference.

- [slo_config](data-sources--voltstack_site--properties--custom_network_config--slo_config.md): complete subsection reference.

- [sm_connection_public_ip](data-sources--voltstack_site--properties--custom_network_config--sm_connection_public_ip.md): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--voltstack_site--properties--custom_network_config--sm_connection_pvt_ip.md): complete subsection reference.

<a id="schema-custom_network_config--tunnel_dead_timeout"></a>

### tunnel_dead_timeout property

Type: `"number"`. Computed.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Upstream description:

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  }
}
```

<a id="schema-custom_network_config--vip_vrrp_mode"></a>

### vip_vrrp_mode property

Type: `"string"`. Computed.

\[Enum: VIP\_VRRP\_INVALID|VIP\_VRRP\_ENABLE|VIP\_VRRP\_DISABLE\] VRRP advertisement mode for VIP
Invalid VRRP mode. Possible values are \`VIP\_VRRP\_INVALID\`, \`VIP\_VRRP\_ENABLE\`,
\`VIP\_VRRP\_DISABLE\`. Defaults to \`VIP\_VRRP\_INVALID\`.

Upstream description:

VRRP advertisement mode for VIP

Invalid VRRP mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIP_VRRP_INVALID",
  "enum": [
    "VIP_VRRP_INVALID",
    "VIP_VRRP_ENABLE",
    "VIP_VRRP_DISABLE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [custom_network_config.active_enhanced_firewall_policies](data-sources--voltstack_site--properties--custom_network_config--active_enhanced_firewall_policies.md)
- [custom_network_config.active_forward_proxy_policies](data-sources--voltstack_site--properties--custom_network_config--active_forward_proxy_policies.md)
- [custom_network_config.active_network_policies](data-sources--voltstack_site--properties--custom_network_config--active_network_policies.md)
- [custom_network_config.default_config](data-sources--voltstack_site--properties--custom_network_config--default_config.md)
- [custom_network_config.default_interface_config](data-sources--voltstack_site--properties--custom_network_config--default_interface_config.md)
- [custom_network_config.default_sli_config](data-sources--voltstack_site--properties--custom_network_config--default_sli_config.md)
- [custom_network_config.forward_proxy_allow_all](data-sources--voltstack_site--properties--custom_network_config--forward_proxy_allow_all.md)
- [custom_network_config.global_network_list](data-sources--voltstack_site--properties--custom_network_config--global_network_list.md)
- [custom_network_config.interface_list](data-sources--voltstack_site--properties--custom_network_config--interface_list.md)
- [custom_network_config.no_forward_proxy](data-sources--voltstack_site--properties--custom_network_config--no_forward_proxy.md)
- [custom_network_config.no_global_network](data-sources--voltstack_site--properties--custom_network_config--no_global_network.md)
- [custom_network_config.no_network_policy](data-sources--voltstack_site--properties--custom_network_config--no_network_policy.md)
- [custom_network_config.sli_config](data-sources--voltstack_site--properties--custom_network_config--sli_config.md)
- [custom_network_config.slo_config](data-sources--voltstack_site--properties--custom_network_config--slo_config.md)
- [custom_network_config.sm_connection_public_ip](data-sources--voltstack_site--properties--custom_network_config--sm_connection_public_ip.md)
- [custom_network_config.sm_connection_pvt_ip](data-sources--voltstack_site--properties--custom_network_config--sm_connection_pvt_ip.md)
- [Property reference](data-sources--voltstack_site--reference.md)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md)
