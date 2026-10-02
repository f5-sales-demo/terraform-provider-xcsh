---
page_title: "gcp.not_managed.node_list.interface_list.bond_interface.lacp"
subcategory: ""
description: "LACP parameters for the bond device."
xcsh_docs: {"aliases": ["gcp not managed node list interface list bond interface lacp"], "body_bytes": 2892, "body_sha256": "sha256:45db912bcd9dbda0dcb64ecdce150ff2a7c0a77b88321f342ee8e827df0a46b1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:bond_interface:lacp", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:bond_interface", "path": "documentation/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/bond_interface/lacp/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3122310201111311-3031030132132122-3303330320002122-0200302013132231-3020221110323301-2110111322012311-0001303131320203-2302131131211112", "registry_path": "docs/guides/data-sources--securemesh_site_v2--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "bond_interface", "lacp"], "schema_version": 1, "sections": [{"aliases": ["rate"], "anchor": "schema-gcp--not_managed--node_list--interface_list--bond_interface--lacp--rate", "description": "Interval in seconds to transmit LACP packets.", "document_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:gcp:not_managed:node_list:interface_list:bond_interface:lacp", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp", "not_managed", "node_list", "interface_list", "bond_interface", "lacp", "rate"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/bond_interface/lacp/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "LACP parameters for the bond device.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.not_managed.node_list.interface_list.bond_interface.lacp

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/)
- [gcp.not_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/)
- [gcp.not_managed.node_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/)
- [gcp.not_managed.node_list.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/)
- [gcp.not_managed.node_list.interface_list.bond_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/bond_interface/)
- gcp.not_managed.node_list.interface_list.bond_interface.lacp

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

<a id="schema-gcp--not_managed--node_list--interface_list--bond_interface--lacp--rate"></a>

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

- [gcp.not_managed.node_list.interface_list.bond_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/properties/gcp/not_managed/node_list/interface_list/bond_interface/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
