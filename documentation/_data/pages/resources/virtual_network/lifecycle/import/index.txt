---
page_title: "Import"
subcategory: "Networking"
description: "Import for xcsh_virtual_network."
xcsh_docs: {"aliases": ["virtual network"], "body_bytes": 370, "body_sha256": "sha256:7855e10f2249747a607dfd483f919b8aa3e75906915742d72f6aff5d2bf29ce4", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:virtual_network:fundamentals", "path": "documentation/resources/virtual_network/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-2313133021302110-3012232223033200-1222231230132232-2031032101122323-2131331201003110-1001022232300313-1230230122212021-0020333131003110", "registry_path": "docs/guides/resources--virtual_network--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/lifecycle/import/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Import for xcsh_virtual_network.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_virtual_network.example system/example
```
