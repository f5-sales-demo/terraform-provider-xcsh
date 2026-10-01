---
page_title: "ingress_egress_gw"
subcategory: "Infrastructure"
description: "ingress_egress_gw for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 11520, "body_sha256": "sha256:294b3c068370bcb0d4449b680f8c0f6f5f9be14e2d4a8a1b85cd577d502e83d1", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:accelerated_networking", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_forward_proxy_policies", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:active_network_policies", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:az_nodes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:forward_proxy_allow_all", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:global_network_list", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:hub", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:inside_static_routes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:no_dc_cluster_group", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:no_forward_proxy", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:no_global_network", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:no_inside_static_routes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:no_network_policy", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:no_outside_static_routes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:not_hub", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:outside_static_routes", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:sm_connection_public_ip", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw", "parent_id": "xcsh-docs:resources:azure_vnet_site:reference", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- ingress_egress_gw

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ingress\_egress\_gw, ingress\_egress\_gw\_ar, ingress\_gw, ingress\_gw\_ar,
voltstack\_cluster, voltstack\_cluster\_ar\] Two interface Azure ingress/egress site.

Upstream description:

Two interface Azure ingress/egress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("az_nodes",
    "azure_certified_hw"),
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

OneOf alternatives in this subsection:

- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md#section)
- [ingress_egress_gw_ar](resources--azure_vnet_site--properties--ingress_egress_gw_ar.md#section)
- [ingress_gw](resources--azure_vnet_site--properties--ingress_gw.md#section)
- [ingress_gw_ar](resources--azure_vnet_site--properties--ingress_gw_ar.md#section)
- [voltstack_cluster](resources--azure_vnet_site--properties--voltstack_cluster.md#section)
- [voltstack_cluster_ar](resources--azure_vnet_site--properties--voltstack_cluster_ar.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ingress_egress_gw {
  # Configure direct properties listed below.
}
```

## Direct properties

- [accelerated_networking](resources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking.md): complete subsection reference.

- [active_enhanced_firewall_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_enhanced_firewall_policies.md): complete subsection reference.

- [active_forward_proxy_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies.md): complete subsection reference.

- [active_network_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies.md): complete subsection reference.

- [az_nodes](resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes.md): complete subsection reference.

<a id="schema-ingress_egress_gw--azure_certified_hw"></a>

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

- [dc_cluster_group_inside_vn](resources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md): complete subsection reference.

- [forward_proxy_allow_all](resources--azure_vnet_site--properties--ingress_egress_gw--forward_proxy_allow_all.md): complete subsection reference.

- [global_network_list](resources--azure_vnet_site--properties--ingress_egress_gw--global_network_list.md): complete subsection reference.

- [hub](resources--azure_vnet_site--properties--ingress_egress_gw--hub.md): complete subsection reference.

- [inside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes.md): complete subsection reference.

- [no_dc_cluster_group](resources--azure_vnet_site--properties--ingress_egress_gw--no_dc_cluster_group.md): complete subsection reference.

- [no_forward_proxy](resources--azure_vnet_site--properties--ingress_egress_gw--no_forward_proxy.md): complete subsection reference.

- [no_global_network](resources--azure_vnet_site--properties--ingress_egress_gw--no_global_network.md): complete subsection reference.

- [no_inside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--no_inside_static_routes.md): complete subsection reference.

- [no_network_policy](resources--azure_vnet_site--properties--ingress_egress_gw--no_network_policy.md): complete subsection reference.

- [no_outside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--no_outside_static_routes.md): complete subsection reference.

- [not_hub](resources--azure_vnet_site--properties--ingress_egress_gw--not_hub.md): complete subsection reference.

- [outside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes.md): complete subsection reference.

- [performance_enhancement_mode](resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode.md): complete subsection reference.

- [sm_connection_public_ip](resources--azure_vnet_site--properties--ingress_egress_gw--sm_connection_public_ip.md): complete subsection reference.

- [sm_connection_pvt_ip](resources--azure_vnet_site--properties--ingress_egress_gw--sm_connection_pvt_ip.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.accelerated_networking](resources--azure_vnet_site--properties--ingress_egress_gw--accelerated_networking.md)
- [ingress_egress_gw.active_enhanced_firewall_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_enhanced_firewall_policies.md)
- [ingress_egress_gw.active_forward_proxy_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_forward_proxy_policies.md)
- [ingress_egress_gw.active_network_policies](resources--azure_vnet_site--properties--ingress_egress_gw--active_network_policies.md)
- [ingress_egress_gw.az_nodes](resources--azure_vnet_site--properties--ingress_egress_gw--az_nodes.md)
- [ingress_egress_gw.dc_cluster_group_inside_vn](resources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md)
- [ingress_egress_gw.dc_cluster_group_outside_vn](resources--azure_vnet_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md)
- [ingress_egress_gw.forward_proxy_allow_all](resources--azure_vnet_site--properties--ingress_egress_gw--forward_proxy_allow_all.md)
- [ingress_egress_gw.global_network_list](resources--azure_vnet_site--properties--ingress_egress_gw--global_network_list.md)
- [ingress_egress_gw.hub](resources--azure_vnet_site--properties--ingress_egress_gw--hub.md)
- [ingress_egress_gw.inside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--inside_static_routes.md)
- [ingress_egress_gw.no_dc_cluster_group](resources--azure_vnet_site--properties--ingress_egress_gw--no_dc_cluster_group.md)
- [ingress_egress_gw.no_forward_proxy](resources--azure_vnet_site--properties--ingress_egress_gw--no_forward_proxy.md)
- [ingress_egress_gw.no_global_network](resources--azure_vnet_site--properties--ingress_egress_gw--no_global_network.md)
- [ingress_egress_gw.no_inside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--no_inside_static_routes.md)
- [ingress_egress_gw.no_network_policy](resources--azure_vnet_site--properties--ingress_egress_gw--no_network_policy.md)
- [ingress_egress_gw.no_outside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--no_outside_static_routes.md)
- [ingress_egress_gw.not_hub](resources--azure_vnet_site--properties--ingress_egress_gw--not_hub.md)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--properties--ingress_egress_gw--outside_static_routes.md)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode.md)
- [ingress_egress_gw.sm_connection_public_ip](resources--azure_vnet_site--properties--ingress_egress_gw--sm_connection_public_ip.md)
- [ingress_egress_gw.sm_connection_pvt_ip](resources--azure_vnet_site--properties--ingress_egress_gw--sm_connection_pvt_ip.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
