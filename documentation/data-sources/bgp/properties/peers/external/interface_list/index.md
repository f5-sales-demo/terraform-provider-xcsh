---
page_title: "peers.external.interface_list"
subcategory: ""
description: "List of network interfaces."
xcsh_docs: {"aliases": ["peers external interface list"], "body_bytes": 1513, "body_sha256": "sha256:4db9f656a7f4dfb2e104cead8c5ed4eb8d6b45653d0345247bd529abf8e447cb", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp:properties:peers:external:interface_list:interfaces"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp:properties:peers:external:interface_list", "parent_id": "xcsh-docs:data-sources:bgp:properties:peers:external", "path": "documentation/data-sources/bgp/properties/peers/external/interface_list/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0332112000321201-2012130032100313-3003023100112021-1221001212101031-3222213330302233-1032331013323132-1312222321313302-1202133211002223", "registry_path": "docs/guides/data-sources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "interface_list"], "schema_version": 1, "sections": [{"aliases": ["peers external interface list interfaces"], "anchor": "section", "description": "List of network interfaces.", "document_id": "xcsh-docs:data-sources:bgp:properties:peers:external:interface_list:interfaces", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["peers", "external", "interface_list", "interfaces"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp/properties/peers/external/interface_list/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of network interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.interface_list

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- peers.external.interface_list

<a id="section"></a>

Type: `"single"`. Computed.

Interface List. List of network interfaces.

Upstream description:

List of network interfaces.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface_list/interfaces/): complete subsection reference.

## Next pages

- [peers.external.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/interface_list/interfaces/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/properties/peers/external/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp/)
