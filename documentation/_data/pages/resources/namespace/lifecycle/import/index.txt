---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_namespace."
xcsh_docs: {"aliases": ["namespace"], "body_bytes": 383, "body_sha256": "sha256:cc77fab036d29a5660214b2f2d6b321192401dc52466f64e1fc7bd7c2bf41115", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:namespace:collection", "completeness": "complete", "id": "xcsh-docs:resources:namespace:import", "import_guidance": "This tenant-level resource uses its bare name. The `namespace` argument is omitted.", "parent_id": "xcsh-docs:resources:namespace:fundamentals", "path": "documentation/resources/namespace/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "namespace", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "resources", "registry_anchor": "canonical-0320000300030320-2021002123203112-0303313311033303-1233033211003313-1002122203320111-2330321121202110-2211110330123011-3002101211000321", "registry_path": "docs/guides/resources--namespace--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/namespace/lifecycle/import/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Import for xcsh_namespace.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["namespaceCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
