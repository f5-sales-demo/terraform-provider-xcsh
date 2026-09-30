---
page_title: "voltstack_cluster"
subcategory: "Infrastructure"
description: "voltstack_cluster for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 8015, "body_sha256": "sha256:ec5393fd6760c0f77e259126227a1a3637c80455f0572736de9ef11f4e6f7f6e", "canonical_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_network_policies", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:az_nodes", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:dc_cluster_group", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:default_storage", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:forward_proxy_allow_all", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:global_network_list", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:k8s_cluster", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_dc_cluster_group", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_forward_proxy", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_global_network", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_k8s_cluster", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_network_policy", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_outside_static_routes", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:sm_connection_public_ip", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:sm_connection_pvt_ip", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:storage_class_list"], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:reference", "path": "docs/guides/data-sources--aws_vpc_site--properties--voltstack_cluster.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- voltstack_cluster

<a id="section"></a>

Type: `"single"`. Computed.

App Stack cluster of single interface AWS nodes.

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

- [active_enhanced_firewall_policies](data-sources--aws_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md): complete subsection reference.

- [active_forward_proxy_policies](data-sources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies.md): complete subsection reference.

- [active_network_policies](data-sources--aws_vpc_site--properties--voltstack_cluster--active_network_policies.md): complete subsection reference.

- [allowed_vip_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port.md): complete subsection reference.

<a id="schema-voltstack_cluster--aws_certified_hw"></a>

### aws_certified_hw property

Type: `"string"`. Computed.

\[Enum: aws-byol-voltstack-combo\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltstack-combo\`.

Upstream description:

Name for AWS certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](data-sources--aws_vpc_site--properties--voltstack_cluster--az_nodes.md): complete subsection reference.

- [dc_cluster_group](data-sources--aws_vpc_site--properties--voltstack_cluster--dc_cluster_group.md): complete subsection reference.

- [default_storage](data-sources--aws_vpc_site--properties--voltstack_cluster--default_storage.md): complete subsection reference.

- [forward_proxy_allow_all](data-sources--aws_vpc_site--properties--voltstack_cluster--forward_proxy_allow_all.md): complete subsection reference.

- [global_network_list](data-sources--aws_vpc_site--properties--voltstack_cluster--global_network_list.md): complete subsection reference.

- [k8s_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster--k8s_cluster.md): complete subsection reference.

- [no_dc_cluster_group](data-sources--aws_vpc_site--properties--voltstack_cluster--no_dc_cluster_group.md): complete subsection reference.

- [no_forward_proxy](data-sources--aws_vpc_site--properties--voltstack_cluster--no_forward_proxy.md): complete subsection reference.

- [no_global_network](data-sources--aws_vpc_site--properties--voltstack_cluster--no_global_network.md): complete subsection reference.

- [no_k8s_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster--no_k8s_cluster.md): complete subsection reference.

- [no_network_policy](data-sources--aws_vpc_site--properties--voltstack_cluster--no_network_policy.md): complete subsection reference.

- [no_outside_static_routes](data-sources--aws_vpc_site--properties--voltstack_cluster--no_outside_static_routes.md): complete subsection reference.

- [outside_static_routes](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes.md): complete subsection reference.

- [sm_connection_public_ip](data-sources--aws_vpc_site--properties--voltstack_cluster--sm_connection_public_ip.md): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--aws_vpc_site--properties--voltstack_cluster--sm_connection_pvt_ip.md): complete subsection reference.

- [storage_class_list](data-sources--aws_vpc_site--properties--voltstack_cluster--storage_class_list.md): complete subsection reference.

## Next pages

- [voltstack_cluster.active_enhanced_firewall_policies](data-sources--aws_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md)
- [voltstack_cluster.active_forward_proxy_policies](data-sources--aws_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies.md)
- [voltstack_cluster.active_network_policies](data-sources--aws_vpc_site--properties--voltstack_cluster--active_network_policies.md)
- [voltstack_cluster.allowed_vip_port](data-sources--aws_vpc_site--properties--voltstack_cluster--allowed_vip_port.md)
- [voltstack_cluster.az_nodes](data-sources--aws_vpc_site--properties--voltstack_cluster--az_nodes.md)
- [voltstack_cluster.dc_cluster_group](data-sources--aws_vpc_site--properties--voltstack_cluster--dc_cluster_group.md)
- [voltstack_cluster.default_storage](data-sources--aws_vpc_site--properties--voltstack_cluster--default_storage.md)
- [voltstack_cluster.forward_proxy_allow_all](data-sources--aws_vpc_site--properties--voltstack_cluster--forward_proxy_allow_all.md)
- [voltstack_cluster.global_network_list](data-sources--aws_vpc_site--properties--voltstack_cluster--global_network_list.md)
- [voltstack_cluster.k8s_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster--k8s_cluster.md)
- [voltstack_cluster.no_dc_cluster_group](data-sources--aws_vpc_site--properties--voltstack_cluster--no_dc_cluster_group.md)
- [voltstack_cluster.no_forward_proxy](data-sources--aws_vpc_site--properties--voltstack_cluster--no_forward_proxy.md)
- [voltstack_cluster.no_global_network](data-sources--aws_vpc_site--properties--voltstack_cluster--no_global_network.md)
- [voltstack_cluster.no_k8s_cluster](data-sources--aws_vpc_site--properties--voltstack_cluster--no_k8s_cluster.md)
- [voltstack_cluster.no_network_policy](data-sources--aws_vpc_site--properties--voltstack_cluster--no_network_policy.md)
- [voltstack_cluster.no_outside_static_routes](data-sources--aws_vpc_site--properties--voltstack_cluster--no_outside_static_routes.md)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--properties--voltstack_cluster--outside_static_routes.md)
- [voltstack_cluster.sm_connection_public_ip](data-sources--aws_vpc_site--properties--voltstack_cluster--sm_connection_public_ip.md)
- [voltstack_cluster.sm_connection_pvt_ip](data-sources--aws_vpc_site--properties--voltstack_cluster--sm_connection_pvt_ip.md)
- [voltstack_cluster.storage_class_list](data-sources--aws_vpc_site--properties--voltstack_cluster--storage_class_list.md)
- [Property reference](data-sources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md)
