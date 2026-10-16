---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_namespace."
xcsh_docs: {"aliases": ["namespace"], "body_bytes": 383, "body_sha256": "sha256:cc77fab036d29a5660214b2f2d6b321192401dc52466f64e1fc7bd7c2bf41115", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:namespace:collection", "completeness": "complete", "id": "xcsh-docs:resources:namespace:import", "import_guidance": "This tenant-level resource uses its bare name. The `namespace` argument is omitted.", "parent_id": "xcsh-docs:resources:namespace:fundamentals", "path": "documentation/resources/namespace/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "namespace", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0320000300030320-2021002123203112-0303313311033303-1233033211003313-1002122203320111-2330321121202110-2211110330123011-3002101211000321", "registry_path": "docs/guides/resources--namespace--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/namespace/lifecycle/import/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Import for xcsh_namespace.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["namespaceCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/namespace/)
- Import

This tenant-level resource uses its bare name. The `namespace` argument is omitted.

```shell
terraform import xcsh_namespace.this example-namespace
```
