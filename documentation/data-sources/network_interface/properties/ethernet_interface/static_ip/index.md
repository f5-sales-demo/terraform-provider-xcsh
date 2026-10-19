---
page_title: "ethernet_interface.static_ip"
subcategory: ""
description: "Configure Static IP parameters."
xcsh_docs: {"aliases": ["ethernet interface static ip"], "body_bytes": 1349, "body_sha256": "sha256:0f5d2dd66e974fac906b9f0bb237b73ce4da89aafacb2c243327e41712dddfc1", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ip:cluster_static_ip", "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ip:node_static_ip"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ip", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "path": "documentation/data-sources/network_interface/properties/ethernet_interface/static_ip/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3312221322021303-3120113113303322-0130201201103221-3232302303112021-3223032010200211-1122330011012312-2200002231303102-1310333113123231", "registry_path": "docs/guides/data-sources--network_interface--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "static_ip"], "schema_version": 1, "sections": [{"aliases": ["ethernet interface static ip cluster static ip"], "anchor": "section", "description": "Configure Static IP parameters for cluster.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ip:cluster_static_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ethernet_interface", "static_ip", "cluster_static_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["ethernet interface static ip node static ip"], "anchor": "section", "description": "Configure Static IP parameters for a node.", "document_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:static_ip:node_static_ip", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ethernet_interface", "static_ip", "node_static_ip"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/static_ip/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Configure Static IP parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.static_ip

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/)
- ethernet_interface.static_ip

<a id="section"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

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

- [cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/cluster_static_ip/): complete subsection reference.

- [node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/static_ip/node_static_ip/): complete subsection reference.
