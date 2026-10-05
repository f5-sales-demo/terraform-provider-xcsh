---
page_title: "equinix.not_managed.node_list.interface_list.static_ipv6_address"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["equinix not managed node list interface list static ipv6 address"], "body_bytes": 3176, "body_sha256": "sha256:6b0ad3bce68644ac2b8a343c19817d7279a3ed5ca6c70af570dcecf9d3f996b1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/static_ipv6_address/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2100222011001113-1310320111302103-1201113123233031-0120121030312023-2301313031101303-2011123303232030-0310123322310000-2232110232312330", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "static_ipv6_address"], "schema_version": 1, "sections": [{"aliases": ["equinix not managed node list interface list static ipv6 address cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "static_ipv6_address", "cluster_static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["equinix not managed node list interface list static ipv6 address node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-equinix--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:equinix:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "requires"}], "schema_path": ["equinix", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# equinix.not_managed.node_list.interface_list.static_ipv6_address

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [equinix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/)
- [equinix.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/)
- [equinix.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/)
- [equinix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/)
- equinix.not_managed.node_list.interface_list.static_ipv6_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/): complete subsection reference.

## Next pages

- [equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/)
- [equinix.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/equinix/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
