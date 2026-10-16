---
page_title: "Import"
subcategory: ""
description: "Import for xcsh_user_identification."
xcsh_docs: {"aliases": ["user identification"], "body_bytes": 382, "body_sha256": "sha256:9b7b22f1d8066fbd39e7e0910e778d23fbf8d61e940754b252ae124f547fd03c", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:import", "import_guidance": "Import using the `namespace/name` identifier format.", "parent_id": "xcsh-docs:resources:user_identification:fundamentals", "path": "documentation/resources/user_identification/lifecycle/import/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0100101021131001-3320300133032202-1010122102032110-1002122121202103-1310010102233133-3000200232000333-3110030122330012-2300300223331230", "registry_path": "docs/guides/resources--user_identification--lifecycle--group-001.md", "relationships": [], "retrieval_version": 1, "role": "import", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/lifecycle/import/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Import for xcsh_user_identification.", "tasks": ["import"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["user_identificationCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Import

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_user_identification.example system/example
```
