---
page_title: "custom_network_config"
subcategory: ""
description: "SmsNetworkConfiguration."
xcsh_docs: {"aliases": ["custom network config"], "body_bytes": 10230, "body_sha256": "sha256:faf2bdb55429cdb3b1eb2240c088097d0131e5f21209847d98128152729d1823", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_enhanced_firewall_policies", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_forward_proxy_policies", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_network_policies", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:default_config", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:default_interface_config", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:default_sli_config", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:forward_proxy_allow_all", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:global_network_list", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:no_forward_proxy", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:no_global_network", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:no_network_policy", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:sli_config", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:sm_connection_public_ip", "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:sm_connection_pvt_ip"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config", "parent_id": "xcsh-docs:data-sources:securemesh_site:reference", "path": "documentation/data-sources/securemesh_site/properties/custom_network_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302", "registry_path": "docs/guides/data-sources--securemesh_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config"], "schema_version": 1, "sections": [{"aliases": ["custom network config active enhanced firewall policies"], "anchor": "section", "description": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_enhanced_firewall_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "active_enhanced_firewall_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config active forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_forward_proxy_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "active_forward_proxy_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config active network policies"], "anchor": "section", "description": "List of firewall policy views.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:active_network_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "active_network_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config default config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:default_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "default_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config default interface config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:default_interface_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "default_interface_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config default sli config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:default_sli_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "default_sli_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config forward proxy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:forward_proxy_allow_all", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "forward_proxy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config global network list"], "anchor": "section", "description": "List of global network connections.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:global_network_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "global_network_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config interface list"], "anchor": "section", "description": "Configure network interfaces for this Secure Mesh site.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:interface_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "interface_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config no forward proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:no_forward_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "no_forward_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config no global network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:no_global_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "no_global_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config no network policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:no_network_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "no_network_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config sli config"], "anchor": "section", "description": "Site local network configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:sli_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "sli_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config slo config"], "anchor": "section", "description": "Site local network configuration.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:slo_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "slo_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:sm_connection_public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config:sm_connection_pvt_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config tunnel dead timeout", "duration"], "anchor": "schema-custom_network_config--tunnel_dead_timeout", "description": "Time interval, in millisec, within which any IPsec / SSL connection from the site going down is detected. When not set (== 0), a default value of 10000 msec will be used.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "tunnel_dead_timeout"], "syntax": "attribute", "type": "number"}, {"aliases": ["custom network config vip vrrp mode"], "anchor": "schema-custom_network_config--vip_vrrp_mode", "description": "VRRP advertisement mode for VIP Invalid VRRP mode.", "document_id": "xcsh-docs:data-sources:securemesh_site:properties:custom_network_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "vip_vrrp_mode"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/custom_network_config/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "SmsNetworkConfiguration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- custom_network_config

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_network\_config, default\_network\_config; Default: default\_network\_config\]
SmsNetworkConfiguration.

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
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-sli_choice": "[\"default_sli_config\",\"sli_config\"]",
  "x-ves-oneof-field-slo_choice": "[\"default_config\",\"slo_config\"]"
}
```

OneOf alternatives in this subsection:

- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/#section)
- [default_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/default_network_config/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/active_network_policies/): complete subsection reference.

- [default_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/default_config/): complete subsection reference.

- [default_interface_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/default_interface_config/): complete subsection reference.

- [default_sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/default_sli_config/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/forward_proxy_allow_all/): complete subsection reference.

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/global_network_list/): complete subsection reference.

- [interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/no_forward_proxy/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/no_global_network/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/no_network_policy/): complete subsection reference.

- [sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/sli_config/): complete subsection reference.

- [slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/sm_connection_pvt_ip/): complete subsection reference.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [custom_network_config.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/active_enhanced_firewall_policies/)
- [custom_network_config.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/active_forward_proxy_policies/)
- [custom_network_config.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/active_network_policies/)
- [custom_network_config.default_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/default_config/)
- [custom_network_config.default_interface_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/default_interface_config/)
- [custom_network_config.default_sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/default_sli_config/)
- [custom_network_config.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/forward_proxy_allow_all/)
- [custom_network_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/global_network_list/)
- [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/interface_list/)
- [custom_network_config.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/no_forward_proxy/)
- [custom_network_config.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/no_global_network/)
- [custom_network_config.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/no_network_policy/)
- [custom_network_config.sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/sli_config/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/slo_config/)
- [custom_network_config.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/sm_connection_public_ip/)
- [custom_network_config.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/custom_network_config/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/)
