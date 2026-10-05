---
page_title: "oci.not_managed.node_list.interface_list.static_ipv6_address"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["oci not managed node list interface list static ipv6 address"], "body_bytes": 3104, "body_sha256": "sha256:ecc7577aa638d45a167dd3b75593d0bc04bc86292f5dea47f77b037a2bbf668b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list", "path": "documentation/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/static_ipv6_address/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0022320311200201-2022123231133030-1003100132020003-2023121231320123-1130203301010300-0010101333111201-1013111001221323-3131232313332233", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-014.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "oci.not_managed.node_list.interface_list.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "oci.not_managed.node_list.interface_list.static_ipv6_address:ConflictingObjectAttributes:cluster_static_ip,node_static_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "static_ipv6_address"], "schema_version": 1, "sections": [{"aliases": ["oci not managed node list interface list static ipv6 address cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:cluster_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "static_ipv6_address", "cluster_static_ip"], "syntax": "block", "type": "object"}, {"aliases": ["oci not managed node list interface list static ipv6 address node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-oci--not_managed--node_list--interface_list--static_ipv6_address--node_static_ip--ip_address", "enforcement": "provider-schema", "group": "oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip:RequiredObjectAttributes:ip_address", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:static_ipv6_address:node_static_ip", "type": "requires"}], "schema_path": ["oci", "not_managed", "node_list", "interface_list", "static_ipv6_address", "node_static_ip"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# oci.not_managed.node_list.interface_list.static_ipv6_address

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [oci](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/)
- [oci.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/)
- [oci.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/)
- [oci.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/)
- oci.not_managed.node_list.interface_list.static_ipv6_address

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

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/): complete subsection reference.

## Next pages

- [oci.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/static_ipv6_address/cluster_static_ip/)
- [oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/static_ipv6_address/node_static_ip/)
- [oci.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
