---
page_title: "layer2_interface"
subcategory: ""
description: "Layer2 Interface Configuration."
xcsh_docs: {"aliases": ["layer2 interface"], "body_bytes": 2217, "body_sha256": "sha256:fe3e42a85003c4bc4a2ebc31abb488ccd5fcb0414243f02f0d4dc375134b53c9", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2sriov_interface", "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_slo_interface"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "parent_id": "xcsh-docs:data-sources:network_interface:reference", "path": "documentation/data-sources/network_interface/properties/layer2_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030", "registry_path": "docs/guides/data-sources--network_interface--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["layer2_interface"], "schema_version": 1, "sections": [{"aliases": ["l2sriov interface"], "anchor": "section", "description": "Layer2 SR-IOV Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2sriov_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["layer2_interface", "l2sriov_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["l2vlan interface"], "anchor": "section", "description": "Layer2 VLAN Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["layer2_interface", "l2vlan_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["l2vlan slo interface"], "anchor": "section", "description": "Layer2 Site Local Outside VLAN Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["layer2_interface", "l2vlan_slo_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/layer2_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Layer2 Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# layer2_interface

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- layer2_interface

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for layer2 interface.

Upstream description:

Layer2 Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-layer2_interface_choice": "[\"l2sriov_interface\",\"l2vlan_interface\",\"l2vlan_slo_interface\"]"
}
```

## Direct properties

- [l2sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2sriov_interface/): complete subsection reference.

- [l2vlan_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/): complete subsection reference.

- [l2vlan_slo_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_slo_interface/): complete subsection reference.

## Next pages

- [layer2_interface.l2sriov_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2sriov_interface/)
- [layer2_interface.l2vlan_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_interface/)
- [layer2_interface.l2vlan_slo_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/layer2_interface/l2vlan_slo_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
