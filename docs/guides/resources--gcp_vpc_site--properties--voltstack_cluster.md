---
page_title: "voltstack_cluster"
subcategory: "Infrastructure"
description: "voltstack_cluster for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 11496, "body_sha256": "sha256:0edd1f95f5be979598d103104bc875816785101ae427ab8624aa8acee2a17677", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:active_network_policies", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:dc_cluster_group", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:default_storage", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:forward_proxy_allow_all", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:global_network_list", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:k8s_cluster", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_dc_cluster_group", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_forward_proxy", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_global_network", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_k8s_cluster", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_network_policy", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:no_outside_static_routes", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:outside_static_routes", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:sm_connection_public_ip", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:sm_connection_pvt_ip", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:storage_class_list"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "docs/guides/resources--gcp_vpc_site--properties--voltstack_cluster.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# voltstack_cluster

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- voltstack_cluster

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

App Stack cluster of single interface GCP site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("gcp_certified_hw",
    "gcp_zone_names"),
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
  validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("default_storage",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("k8s_cluster",
    "no_k8s_cluster"),
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

Terraform syntax:

```terraform
voltstack_cluster {
  # Configure direct properties listed below.
}
```

## Direct properties

- [active_enhanced_firewall_policies](resources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md): complete subsection reference.

- [active_forward_proxy_policies](resources--gcp_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies.md): complete subsection reference.

- [active_network_policies](resources--gcp_vpc_site--properties--voltstack_cluster--active_network_policies.md): complete subsection reference.

- [dc_cluster_group](resources--gcp_vpc_site--properties--voltstack_cluster--dc_cluster_group.md): complete subsection reference.

- [default_storage](resources--gcp_vpc_site--properties--voltstack_cluster--default_storage.md): complete subsection reference.

- [forward_proxy_allow_all](resources--gcp_vpc_site--properties--voltstack_cluster--forward_proxy_allow_all.md): complete subsection reference.

<a id="schema-voltstack_cluster--gcp_certified_hw"></a>

### gcp_certified_hw property

Type: `"string"`. Optional.

\[Enum: gcp-byol-voltstack-combo\] GCP Certified Hardware. Name for GCP certified hardware. The only
possible value is \`gcp-byol-voltstack-combo\`.

Upstream description:

Name for GCP certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("gcp-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-voltstack_cluster--gcp_zone_names"></a>

### gcp_zone_names property

Type: `["list", "string"]`. Optional.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(3),
}
```

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

- [global_network_list](resources--gcp_vpc_site--properties--voltstack_cluster--global_network_list.md): complete subsection reference.

- [k8s_cluster](resources--gcp_vpc_site--properties--voltstack_cluster--k8s_cluster.md): complete subsection reference.

- [no_dc_cluster_group](resources--gcp_vpc_site--properties--voltstack_cluster--no_dc_cluster_group.md): complete subsection reference.

- [no_forward_proxy](resources--gcp_vpc_site--properties--voltstack_cluster--no_forward_proxy.md): complete subsection reference.

- [no_global_network](resources--gcp_vpc_site--properties--voltstack_cluster--no_global_network.md): complete subsection reference.

- [no_k8s_cluster](resources--gcp_vpc_site--properties--voltstack_cluster--no_k8s_cluster.md): complete subsection reference.

- [no_network_policy](resources--gcp_vpc_site--properties--voltstack_cluster--no_network_policy.md): complete subsection reference.

- [no_outside_static_routes](resources--gcp_vpc_site--properties--voltstack_cluster--no_outside_static_routes.md): complete subsection reference.

<a id="schema-voltstack_cluster--node_number"></a>

### node_number property

Type: `"number"`. Optional.

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

- [outside_static_routes](resources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes.md): complete subsection reference.

- [site_local_network](resources--gcp_vpc_site--properties--voltstack_cluster--site_local_network.md): complete subsection reference.

- [site_local_subnet](resources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet.md): complete subsection reference.

- [sm_connection_public_ip](resources--gcp_vpc_site--properties--voltstack_cluster--sm_connection_public_ip.md): complete subsection reference.

- [sm_connection_pvt_ip](resources--gcp_vpc_site--properties--voltstack_cluster--sm_connection_pvt_ip.md): complete subsection reference.

- [storage_class_list](resources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list.md): complete subsection reference.

## Next pages

- [voltstack_cluster.active_enhanced_firewall_policies](resources--gcp_vpc_site--properties--voltstack_cluster--active_enhanced_firewall_policies.md)
- [voltstack_cluster.active_forward_proxy_policies](resources--gcp_vpc_site--properties--voltstack_cluster--active_forward_proxy_policies.md)
- [voltstack_cluster.active_network_policies](resources--gcp_vpc_site--properties--voltstack_cluster--active_network_policies.md)
- [voltstack_cluster.dc_cluster_group](resources--gcp_vpc_site--properties--voltstack_cluster--dc_cluster_group.md)
- [voltstack_cluster.default_storage](resources--gcp_vpc_site--properties--voltstack_cluster--default_storage.md)
- [voltstack_cluster.forward_proxy_allow_all](resources--gcp_vpc_site--properties--voltstack_cluster--forward_proxy_allow_all.md)
- [voltstack_cluster.global_network_list](resources--gcp_vpc_site--properties--voltstack_cluster--global_network_list.md)
- [voltstack_cluster.k8s_cluster](resources--gcp_vpc_site--properties--voltstack_cluster--k8s_cluster.md)
- [voltstack_cluster.no_dc_cluster_group](resources--gcp_vpc_site--properties--voltstack_cluster--no_dc_cluster_group.md)
- [voltstack_cluster.no_forward_proxy](resources--gcp_vpc_site--properties--voltstack_cluster--no_forward_proxy.md)
- [voltstack_cluster.no_global_network](resources--gcp_vpc_site--properties--voltstack_cluster--no_global_network.md)
- [voltstack_cluster.no_k8s_cluster](resources--gcp_vpc_site--properties--voltstack_cluster--no_k8s_cluster.md)
- [voltstack_cluster.no_network_policy](resources--gcp_vpc_site--properties--voltstack_cluster--no_network_policy.md)
- [voltstack_cluster.no_outside_static_routes](resources--gcp_vpc_site--properties--voltstack_cluster--no_outside_static_routes.md)
- [voltstack_cluster.outside_static_routes](resources--gcp_vpc_site--properties--voltstack_cluster--outside_static_routes.md)
- [voltstack_cluster.site_local_network](resources--gcp_vpc_site--properties--voltstack_cluster--site_local_network.md)
- [voltstack_cluster.site_local_subnet](resources--gcp_vpc_site--properties--voltstack_cluster--site_local_subnet.md)
- [voltstack_cluster.sm_connection_public_ip](resources--gcp_vpc_site--properties--voltstack_cluster--sm_connection_public_ip.md)
- [voltstack_cluster.sm_connection_pvt_ip](resources--gcp_vpc_site--properties--voltstack_cluster--sm_connection_pvt_ip.md)
- [voltstack_cluster.storage_class_list](resources--gcp_vpc_site--properties--voltstack_cluster--storage_class_list.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
