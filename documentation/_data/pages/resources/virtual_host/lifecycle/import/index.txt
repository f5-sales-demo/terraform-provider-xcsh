---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_virtual_host."
xcsh_docs: {"aliases": ["virtual host"], "body_bytes": 361, "body_sha256": "sha256:298f4647170688e9846e38ac118d82b86585c20abcd45a31f39f6be4452b0cb3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:virtual_host:fundamentals", "path": "documentation/resources/virtual_host/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3010122311030023-3301131312110032-1003122023332231-3031220211222311-0010113103131003-3110200330233031-1023022301030303-3300223123110301", "registry_path": "docs/guides/resources--virtual_host--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Import for xcsh_virtual_host.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_virtual_host.example system/example
```
