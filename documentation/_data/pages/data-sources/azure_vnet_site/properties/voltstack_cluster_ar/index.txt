---
page_title: "voltstack_cluster_ar"
subcategory: "Infrastructure"
description: "voltstack_cluster_ar for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 10580, "body_sha256": "sha256:7a56590cb40748c94e7968cb7310d1ac2b60bae64d39495db1bd871f875ddd20", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:accelerated_networking", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:active_enhanced_firewall_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:active_forward_proxy_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:active_network_policies", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:dc_cluster_group", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:default_storage", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:forward_proxy_allow_all", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:global_network_list", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:k8s_cluster", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:no_dc_cluster_group", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:no_forward_proxy", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:no_global_network", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:no_k8s_cluster", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:no_network_policy", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:no_outside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:node", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:outside_static_routes", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:sm_connection_public_ip", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:sm_connection_pvt_ip", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar:storage_class_list"], "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "path": "documentation/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/index.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["voltstack_cluster_ar"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster_ar for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster_ar

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- voltstack_cluster_ar

<a id="section"></a>

Type: `"single"`. Computed.

App Stack Cluster of single interface Azure nodes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

## Direct properties

- [accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/): complete subsection reference.

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/): complete subsection reference.

<a id="schema-voltstack_cluster_ar--azure_certified_hw"></a>

### azure_certified_hw property

Type: `"string"`. Computed.

\[Enum: azure-byol-voltstack-combo\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-voltstack-combo\`.

Upstream description:

Name for Azure certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/dc_cluster_group/): complete subsection reference.

- [default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/default_storage/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/forward_proxy_allow_all/): complete subsection reference.

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/): complete subsection reference.

- [k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/k8s_cluster/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_dc_cluster_group/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_forward_proxy/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_global_network/): complete subsection reference.

- [no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_k8s_cluster/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_network_policy/): complete subsection reference.

- [no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_outside_static_routes/): complete subsection reference.

- [node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/): complete subsection reference.

- [outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/sm_connection_pvt_ip/): complete subsection reference.

- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/): complete subsection reference.

## Next pages

- [voltstack_cluster_ar.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/)
- [voltstack_cluster_ar.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_forward_proxy_policies/)
- [voltstack_cluster_ar.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/)
- [voltstack_cluster_ar.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/dc_cluster_group/)
- [voltstack_cluster_ar.default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/default_storage/)
- [voltstack_cluster_ar.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/forward_proxy_allow_all/)
- [voltstack_cluster_ar.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/)
- [voltstack_cluster_ar.k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/k8s_cluster/)
- [voltstack_cluster_ar.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_dc_cluster_group/)
- [voltstack_cluster_ar.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_forward_proxy/)
- [voltstack_cluster_ar.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_global_network/)
- [voltstack_cluster_ar.no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_k8s_cluster/)
- [voltstack_cluster_ar.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_network_policy/)
- [voltstack_cluster_ar.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_outside_static_routes/)
- [voltstack_cluster_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/)
- [voltstack_cluster_ar.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/)
- [voltstack_cluster_ar.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/sm_connection_public_ip/)
- [voltstack_cluster_ar.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/sm_connection_pvt_ip/)
- [voltstack_cluster_ar.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
