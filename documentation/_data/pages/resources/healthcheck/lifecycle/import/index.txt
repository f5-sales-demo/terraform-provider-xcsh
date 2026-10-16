---
page_title: "Import"
subcategory: "Monitoring"
description: "Import for xcsh_healthcheck."
xcsh_docs: {"aliases": ["healthcheck"], "body_bytes": 358, "body_sha256": "sha256:22a0b4475ae00e738d7b8599a0dddee814c26d6dd13de4e6714ca86815c94663", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:healthcheck:fundamentals", "path": "documentation/resources/healthcheck/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-2200022120320020-3330113001002213-0131113001023023-2302333211211011-2120001023103030-1231103003310122-0201010303231030-3100212123301230", "registry_path": "docs/guides/resources--healthcheck--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/lifecycle/import/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Import for xcsh_healthcheck.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["healthcheckCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_healthcheck.example system/example
```
