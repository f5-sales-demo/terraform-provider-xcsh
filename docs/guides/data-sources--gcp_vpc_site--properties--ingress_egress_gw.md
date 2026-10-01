---
page_title: "ingress_egress_gw"
subcategory: "Infrastructure"
description: "ingress_egress_gw for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 11104, "body_sha256": "sha256:5975aec41f75614a35da53a5f4029fc87d585305baf64d8aff671047526a1137", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:active_enhanced_firewall_policies", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:active_forward_proxy_policies", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:active_network_policies", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_inside_vn", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:dc_cluster_group_outside_vn", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:forward_proxy_allow_all", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:global_network_list", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_network", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_static_routes", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:no_dc_cluster_group", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:no_forward_proxy", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:no_global_network", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:no_inside_static_routes", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:no_network_policy", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:no_outside_static_routes", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_network", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_static_routes", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:outside_subnet", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:sm_connection_public_ip", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw:sm_connection_pvt_ip"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_egress_gw", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:reference", "path": "docs/guides/data-sources--gcp_vpc_site--properties--ingress_egress_gw.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_egress_gw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- ingress_egress_gw

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface GCP ingress/egress site.

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
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](data-sources--gcp_vpc_site--properties--ingress_egress_gw.md#section)
- [ingress_gw](data-sources--gcp_vpc_site--properties--ingress_gw.md#section)
- [voltstack_cluster](data-sources--gcp_vpc_site--properties--voltstack_cluster.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [active_enhanced_firewall_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies.md): complete subsection reference.

- [active_forward_proxy_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies.md): complete subsection reference.

- [active_network_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_network_policies.md): complete subsection reference.

- [dc_cluster_group_inside_vn](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md): complete subsection reference.

- [forward_proxy_allow_all](data-sources--gcp_vpc_site--properties--ingress_egress_gw--forward_proxy_allow_all.md): complete subsection reference.

<a id="schema-ingress_egress_gw--gcp_certified_hw"></a>

### gcp_certified_hw property

Type: `"string"`. Computed.

\[Enum: gcp-byol-multi-nic-voltmesh\] GCP Certified Hardware. Name for GCP certified hardware. The
only possible value is \`gcp-byol-multi-nic-voltmesh\`.

Upstream description:

Name for GCP certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-multi-nic-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-ingress_egress_gw--gcp_zone_names"></a>

### gcp_zone_names property

Type: `["list", "string"]`. Computed.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [global_network_list](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list.md): complete subsection reference.

- [inside_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network.md): complete subsection reference.

- [inside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes.md): complete subsection reference.

- [inside_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet.md): complete subsection reference.

- [no_dc_cluster_group](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_dc_cluster_group.md): complete subsection reference.

- [no_forward_proxy](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_forward_proxy.md): complete subsection reference.

- [no_global_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_global_network.md): complete subsection reference.

- [no_inside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_inside_static_routes.md): complete subsection reference.

- [no_network_policy](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_network_policy.md): complete subsection reference.

- [no_outside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_outside_static_routes.md): complete subsection reference.

<a id="schema-ingress_egress_gw--node_number"></a>

### node_number property

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_network.md): complete subsection reference.

- [outside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes.md): complete subsection reference.

- [outside_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_subnet.md): complete subsection reference.

- [performance_enhancement_mode](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode.md): complete subsection reference.

- [sm_connection_public_ip](data-sources--gcp_vpc_site--properties--ingress_egress_gw--sm_connection_public_ip.md): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--gcp_vpc_site--properties--ingress_egress_gw--sm_connection_pvt_ip.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_enhanced_firewall_policies.md)
- [ingress_egress_gw.active_forward_proxy_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_forward_proxy_policies.md)
- [ingress_egress_gw.active_network_policies](data-sources--gcp_vpc_site--properties--ingress_egress_gw--active_network_policies.md)
- [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_inside_vn.md)
- [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--gcp_vpc_site--properties--ingress_egress_gw--dc_cluster_group_outside_vn.md)
- [ingress_egress_gw.forward_proxy_allow_all](data-sources--gcp_vpc_site--properties--ingress_egress_gw--forward_proxy_allow_all.md)
- [ingress_egress_gw.global_network_list](data-sources--gcp_vpc_site--properties--ingress_egress_gw--global_network_list.md)
- [ingress_egress_gw.inside_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_network.md)
- [ingress_egress_gw.inside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_static_routes.md)
- [ingress_egress_gw.inside_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet.md)
- [ingress_egress_gw.no_dc_cluster_group](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_dc_cluster_group.md)
- [ingress_egress_gw.no_forward_proxy](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_forward_proxy.md)
- [ingress_egress_gw.no_global_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_global_network.md)
- [ingress_egress_gw.no_inside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_inside_static_routes.md)
- [ingress_egress_gw.no_network_policy](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_network_policy.md)
- [ingress_egress_gw.no_outside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--no_outside_static_routes.md)
- [ingress_egress_gw.outside_network](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_network.md)
- [ingress_egress_gw.outside_static_routes](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_static_routes.md)
- [ingress_egress_gw.outside_subnet](data-sources--gcp_vpc_site--properties--ingress_egress_gw--outside_subnet.md)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--properties--ingress_egress_gw--performance_enhancement_mode.md)
- [ingress_egress_gw.sm_connection_public_ip](data-sources--gcp_vpc_site--properties--ingress_egress_gw--sm_connection_public_ip.md)
- [ingress_egress_gw.sm_connection_pvt_ip](data-sources--gcp_vpc_site--properties--ingress_egress_gw--sm_connection_pvt_ip.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
