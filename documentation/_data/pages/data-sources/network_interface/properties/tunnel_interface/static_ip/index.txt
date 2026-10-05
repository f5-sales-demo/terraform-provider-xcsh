---
page_title: "tunnel_interface.static_ip"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["tunnel interface static ip"], "body_bytes": 2045, "body_sha256": "sha256:5c231970104aa3627ea709dfc2a28600ac84042be7bf584c29b71e3ef099ac20", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip:cluster_static_ip", "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip:node_static_ip"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip", "parent_id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface", "path": "documentation/data-sources/network_interface/properties/tunnel_interface/static_ip/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3121212033100033-0132200113233111-1130113303120101-0331313032310301-2312200021332010-0102333211323101-0333123111113031-2303203102211032", "registry_path": "docs/guides/data-sources--network_interface--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tunnel_interface", "static_ip"], "schema_version": 1, "sections": [{"aliases": ["tunnel interface static ip cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip:cluster_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tunnel_interface", "static_ip", "cluster_static_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["tunnel interface static ip node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:data-sources:network_interface:properties:tunnel_interface:static_ip:node_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["tunnel_interface", "static_ip", "node_static_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/tunnel_interface/static_ip/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tunnel_interface.static_ip

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/)
- tunnel_interface.static_ip

<a id="section"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

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

## Direct properties

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/): complete subsection reference.

## Next pages

- [tunnel_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/cluster_static_ip/)
- [tunnel_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/static_ip/node_static_ip/)
- [tunnel_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/tunnel_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
