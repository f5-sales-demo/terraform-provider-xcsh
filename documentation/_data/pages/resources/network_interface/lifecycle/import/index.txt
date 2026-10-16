---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_network_interface."
xcsh_docs: {"aliases": ["network interface"], "body_bytes": 376, "body_sha256": "sha256:83b50a12a39c30c303ab3992a5bf66ea33cc4c97fb5043cf4646f6e2fecd6771", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:network_interface:fundamentals", "path": "documentation/resources/network_interface/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3202322311332023-0100223231033130-2213001230333101-2302222313120021-0323233203000131-3310112033131001-0021311030221121-1111303211301112", "registry_path": "docs/guides/resources--network_interface--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Import for xcsh_network_interface.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_network_interface.example system/example
```
