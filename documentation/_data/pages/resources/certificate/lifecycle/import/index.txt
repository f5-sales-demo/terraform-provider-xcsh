---
page_title: "Import"
subcategory: "Security"
description: "Import for xcsh_certificate."
xcsh_docs: {"aliases": ["certificate"], "body_bytes": 358, "body_sha256": "sha256:d363ac05107cf6e845192eb5e51a5ce6940c811e55e2f9ddd4720f65aac23333", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:resources:certificate:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:certificate:fundamentals", "path": "documentation/resources/certificate/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "certificate", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-1101233031133031-3130000233121232-3102010323131202-3101303002132031-0012011222011220-2110202303020302-1303110123110020-3321102223002210", "registry_path": "docs/guides/resources--certificate--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/certificate/lifecycle/import/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Import for xcsh_certificate.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["certificateCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/certificate/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_certificate.example system/example
```
