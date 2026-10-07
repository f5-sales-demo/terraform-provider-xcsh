---
page_title: "layer2_interface"
subcategory: ""
description: "Layer2 Interface Configuration."
xcsh_docs: {"aliases": ["layer2 interface"], "body_bytes": 1443, "body_sha256": "sha256:7fda7a42d4ec79b66a907e8b650fbe7c7c89b3ac9934fd77904fdc736e05a466", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2sriov_interface", "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_slo_interface"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface", "parent_id": "xcsh-docs:data-sources:network_interface:reference", "path": "documentation/data-sources/network_interface/properties/layer2_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0203332310121320-0320201223332130-3003222223202313-0112031223031223-1020213311213033-3313000032021132-3232210333100223-0230211123130030", "registry_path": "docs/guides/data-sources--network_interface--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["layer2_interface"], "schema_version": 1, "sections": [{"aliases": ["layer2 interface l2sriov interface"], "anchor": "section", "description": "Layer2 SR-IOV Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2sriov_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["layer2_interface", "l2sriov_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["layer2 interface l2vlan interface"], "anchor": "section", "description": "Layer2 VLAN Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["layer2_interface", "l2vlan_interface"], "syntax": "attribute", "type": "object"}, {"aliases": ["layer2 interface l2vlan slo interface"], "anchor": "section", "description": "Layer2 Site Local Outside VLAN Interface Configuration.", "document_id": "xcsh-docs:data-sources:network_interface:properties:layer2_interface:l2vlan_slo_interface", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["layer2_interface", "l2vlan_slo_interface"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/layer2_interface/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Layer2 Interface Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
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

Additional upstream details:

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
