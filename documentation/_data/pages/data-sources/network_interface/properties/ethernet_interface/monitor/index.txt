---
page_title: "ethernet_interface.monitor"
subcategory: ""
description: "Link Quality Monitoring configuration for a network interface."
xcsh_docs: {"aliases": ["ethernet interface monitor"], "body_bytes": 921, "body_sha256": "sha256:4c5889bb5199b801f0919f92b150beab4680166d39644a55a8c626163a6e443a", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface:monitor", "parent_id": "xcsh-docs:data-sources:network_interface:properties:ethernet_interface", "path": "documentation/data-sources/network_interface/properties/ethernet_interface/monitor/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1332112132010022-2002033222322231-0013003130323320-3003020332121213-2220110320101212-2120330123312013-1221021231031213-2211123321110201", "registry_path": "docs/guides/data-sources--network_interface--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "monitor"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/properties/ethernet_interface/monitor/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Link Quality Monitoring configuration for a network interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.monitor

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/ethernet_interface/)
- ethernet_interface.monitor

<a id="section"></a>

Type: `["object", {}]`. Computed.

Link Quality Monitoring configuration for a network interface.

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

This is an empty object or choice marker. It has no direct properties.
