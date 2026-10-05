---
page_title: "static_routes.node_interface"
subcategory: "Networking"
description: "On multinode site, this type holds the information about per node interfaces."
xcsh_docs: {"aliases": ["static routes node interface"], "body_bytes": 1468, "body_sha256": "sha256:7bb57cf635b04abd74b6c9aa31301946b5ad303c52a9422ec5ca95dd1934693e", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:virtual_network:properties:static_routes:node_interface:list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_network:properties:static_routes:node_interface", "parent_id": "xcsh-docs:data-sources:virtual_network:properties:static_routes", "path": "documentation/data-sources/virtual_network/properties/static_routes/node_interface/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1203112303110222-2133131133220301-0313031022113033-0021122131133012-0022332222203001-3200320000113102-3132203122220232-2201200023310120", "registry_path": "docs/guides/data-sources--virtual_network--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["static_routes", "node_interface"], "schema_version": 1, "sections": [{"aliases": ["static routes node interface list"], "anchor": "section", "description": "On a multinode site, this list holds the nodes and corresponding networking_interface.", "document_id": "xcsh-docs:data-sources:virtual_network:properties:static_routes:node_interface:list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["static_routes", "node_interface", "list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_network/properties/static_routes/node_interface/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "On multinode site, this type holds the information about per node interfaces.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# static_routes.node_interface

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/)
- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/)
- static_routes.node_interface

<a id="section"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

- [list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/node_interface/list/): complete subsection reference.

## Next pages

- [static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/node_interface/list/)
- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/properties/static_routes/)
- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_network/)
