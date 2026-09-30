---
page_title: "ingress_egress_gw_ar"
subcategory: "Infrastructure"
description: "ingress_egress_gw_ar for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 13210, "body_sha256": "sha256:8dafae41f6247a441e8a1d6f557cef049df7ce4eadb3e172ebb04da2f475eb54", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:accelerated_networking", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:active_enhanced_firewall_policies", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:active_forward_proxy_policies", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:active_network_policies", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:dc_cluster_group_inside_vn", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:dc_cluster_group_outside_vn", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:forward_proxy_allow_all", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:global_network_list", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:hub", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:inside_static_routes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:no_dc_cluster_group", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:no_forward_proxy", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:no_global_network", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:no_inside_static_routes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:no_network_policy", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:no_outside_static_routes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:node", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:not_hub", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:outside_static_routes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:performance_enhancement_mode", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:sm_connection_public_ip", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw_ar", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "documentation/resources/azure_vnet_site/properties/ingress_egress_gw_ar/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["ingress_egress_gw_ar"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw_ar/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw_ar for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw_ar

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- ingress_egress_gw_ar

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Two interface Azure ingress/egress site on Alternate Region with no support for zones.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("azure_certified_hw"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("hub",
    "not_hub"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-hub_choice": "[\"hub\",\"not_hub\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
ingress_egress_gw_ar {
  # Configure direct properties listed below.
}
```

## Direct properties

- [accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/accelerated_networking/): complete subsection reference.

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/active_network_policies/): complete subsection reference.

<a id="schema-ingress_egress_gw_ar--azure_certified_hw"></a>

### azure_certified_hw property

Type: `"string"`. Optional.

\[Enum: azure-byol-multi-nic-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-multi-nic-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-multi-nic-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_inside_vn/): complete subsection reference.

- [dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_outside_vn/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/forward_proxy_allow_all/): complete subsection reference.

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/): complete subsection reference.

- [hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/): complete subsection reference.

- [inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_dc_cluster_group/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_forward_proxy/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_global_network/): complete subsection reference.

- [no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_inside_static_routes/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_network_policy/): complete subsection reference.

- [no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_outside_static_routes/): complete subsection reference.

- [node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/): complete subsection reference.

- [not_hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/not_hub/): complete subsection reference.

- [outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/): complete subsection reference.

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/sm_connection_pvt_ip/): complete subsection reference.

## Next pages

- [ingress_egress_gw_ar.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/accelerated_networking/)
- [ingress_egress_gw_ar.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/active_enhanced_firewall_policies/)
- [ingress_egress_gw_ar.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/)
- [ingress_egress_gw_ar.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/active_network_policies/)
- [ingress_egress_gw_ar.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_inside_vn/)
- [ingress_egress_gw_ar.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_outside_vn/)
- [ingress_egress_gw_ar.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/forward_proxy_allow_all/)
- [ingress_egress_gw_ar.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/)
- [ingress_egress_gw_ar.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/)
- [ingress_egress_gw_ar.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/)
- [ingress_egress_gw_ar.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_dc_cluster_group/)
- [ingress_egress_gw_ar.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_forward_proxy/)
- [ingress_egress_gw_ar.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_global_network/)
- [ingress_egress_gw_ar.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_inside_static_routes/)
- [ingress_egress_gw_ar.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_network_policy/)
- [ingress_egress_gw_ar.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/no_outside_static_routes/)
- [ingress_egress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/node/)
- [ingress_egress_gw_ar.not_hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/not_hub/)
- [ingress_egress_gw_ar.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/)
- [ingress_egress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/)
- [ingress_egress_gw_ar.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/sm_connection_public_ip/)
- [ingress_egress_gw_ar.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_egress_gw_ar/sm_connection_pvt_ip/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
