---
page_title: "azure.not_managed.node_list.interface_list.bond_interface.lacp"
subcategory: ""
description: "LACP parameters for the bond device."
xcsh_docs: {"aliases": ["azure not managed node list interface list bond interface lacp"], "body_bytes": 2922, "body_sha256": "sha256:10065585d2c802e7a23abee800f9f4cb8f8787fd733c5c67d08a43787c46cfda", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:bond_interface:lacp", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:bond_interface", "path": "documentation/data-sources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/bond_interface/lacp/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2010213103201331-1213000030222012-3030020013211123-3323102333221231-2031112012123221-0101212230122012-2300010111203113-2202030233003113", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["azure", "not_managed", "node_list", "interface_list", "bond_interface", "lacp"], "schema_version": 1, "sections": [{"aliases": ["azure not managed node list interface list bond interface lacp rate"], "anchor": "schema-azure--not_managed--node_list--interface_list--bond_interface--lacp--rate", "description": "Interval in seconds to transmit LACP packets.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:azure:not_managed:node_list:interface_list:bond_interface:lacp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure", "not_managed", "node_list", "interface_list", "bond_interface", "lacp", "rate"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/bond_interface/lacp/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "LACP parameters for the bond device.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# azure.not_managed.node_list.interface_list.bond_interface.lacp

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [azure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/azure/)
- [azure.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/azure/not_managed/)
- [azure.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/azure/not_managed/node_list/)
- [azure.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/)
- [azure.not_managed.node_list.interface_list.bond_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/bond_interface/)
- azure.not_managed.node_list.interface_list.bond_interface.lacp

<a id="section"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

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

<a id="schema-azure--not_managed--node_list--interface_list--bond_interface--lacp--rate"></a>

### rate property

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

## Next pages

- [azure.not_managed.node_list.interface_list.bond_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/azure/not_managed/node_list/interface_list/bond_interface/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
