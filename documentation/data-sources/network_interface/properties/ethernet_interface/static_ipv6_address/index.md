---
page_title: "ethernet_interface.static_ipv6_address"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["ethernet interface static ipv6 address"], "body_bytes": 2149, "body_sha256": "sha256:1ecc2d8758b4783acf06caa83253290d74ddaa75c522ad9650dcdf01f07f1551", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address:cluster_static_ip", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address:node_static_ip"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "path": "documentation/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2122131222310110-1332200023023011-0111031123111230-3110000330210323-2132201123233021-2222111300011033-3001000030212323-0221213333013310", "registry_path": "docs/guides/data-sources--network_interface--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "static_ipv6_address"], "schema_version": 1, "sections": [{"aliases": ["ethernet interface static ipv6 address cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address:cluster_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ethernet_interface", "static_ipv6_address", "cluster_static_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["ethernet interface static ipv6 address node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ipv6_address:node_static_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ethernet_interface", "static_ipv6_address", "node_static_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.static_ipv6_address

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/)
- ethernet_interface.static_ipv6_address

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

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/node_static_ip/): complete subsection reference.

## Next pages

- [ethernet_interface.static_ipv6_address.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/cluster_static_ip/)
- [ethernet_interface.static_ipv6_address.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ipv6_address/node_static_ip/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
