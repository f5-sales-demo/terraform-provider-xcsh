---
page_title: "dedicated_interface.monitor"
subcategory: ""
description: "Link Quality Monitoring configuration for a network interface."
xcsh_docs: {"aliases": ["dedicated interface monitor"], "body_bytes": 966, "body_sha256": "sha256:4a1fc589f503e1a3c6c4f34d88bf19cea872365e6774c9a9d862bee013beaa83", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:dedicated_interface:monitor", "parent_id": "xcsh-docs:resources:network_interface:properties:dedicated_interface", "path": "documentation/resources/network_interface/properties/dedicated_interface/monitor/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3230321221323212-3020031231321210-2321212001020302-3332321332202330-0202313030101032-3113030003102211-1301123123303333-1031203002101112", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["dedicated_interface", "monitor"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/dedicated_interface/monitor/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Link Quality Monitoring configuration for a network interface.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dedicated_interface.monitor

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/dedicated_interface/)
- dedicated_interface.monitor

<a id="section"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
monitor = {}
```

This is an empty object or choice marker. It has no direct properties.
